package provider

import (
	"context"
	"fmt"

	"github.com/leanspace/terraform-provider-leanspace/helper"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// GenericResource implements resource.Resource for any DataSourceType.
type GenericResource[T any, PT ParseableModel[T]] struct {
	dataType *DataSourceType[T, PT]
	client   *Client
}

func NewGenericResource[T any, PT ParseableModel[T]](dt *DataSourceType[T, PT]) func() resource.Resource {
	return func() resource.Resource {
		return &GenericResource[T, PT]{dataType: dt}
	}
}

func (r *GenericResource[T, PT]) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = r.dataType.ResourceIdentifier
}

func (r *GenericResource[T, PT]) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := SplitResourceSchemaBlocks(r.dataType.Schema)
	resp.Schema = resourceschema.Schema{
		Attributes: attrs,
		Blocks:     blocks,
	}
}

func (r *GenericResource[T, PT]) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client, ok := req.ProviderData.(*Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Provider Data", fmt.Sprintf("Expected *Client, got %T", req.ProviderData))
		return
	}
	r.client = client
}

// getData reads all attributes from the framework plan/state into a raw map, validates, then converts to model.
func (r *GenericResource[T, PT]) getData(ctx context.Context, attrValues map[string]attr.Value, checkValidity bool) (PT, error) {
	rawMap, diags := AttrValuesToMap(ctx, r.dataType.Schema, attrValues)
	if diags.HasError() {
		return nil, fmt.Errorf("error converting plan to map: %s", diags.Errors()[0].Detail())
	}
	if rawMap == nil {
		return nil, nil
	}

	var value PT = new(T)

	if checkValidity && helper.Implements[T, ValidationModel]() {
		err := any(value).(ValidationModel).Validate(rawMap)
		if err != nil {
			return nil, err
		}
	}

	err := value.FromMap(rawMap)
	return value, err
}

func (r *GenericResource[T, PT]) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Get plan values
	planValues := make(map[string]attr.Value)
	for key := range r.dataType.Schema {
		var val attr.Value
		diags := req.Plan.GetAttribute(ctx, path.Root(key), &val)
		resp.Diagnostics.Append(diags...)
		planValues[key] = val
	}
	if resp.Diagnostics.HasError() {
		return
	}

	value, err := r.getData(ctx, planValues, true)
	if err != nil {
		resp.Diagnostics.AddError("Error parsing plan", err.Error())
		return
	}

	createdValue, err := r.dataType.convert(r.client).Create(value)
	if err != nil {
		resp.Diagnostics.AddError("Error creating resource", err.Error())
		return
	}

	// Re-read the resource to get the full state
	readValue, err := r.dataType.convert(r.client).Get(createdValue.GetID(), value)
	if err != nil {
		resp.Diagnostics.AddError("Error reading resource after create", err.Error())
		return
	}
	if readValue == nil {
		resp.Diagnostics.AddError("Resource not found after creation", "")
		return
	}

	storedData := readValue.ToMap()
	storedData["id"] = createdValue.GetID()
	attrVals, d := MapToAttrValues(ctx, r.dataType.Schema, storedData)
	resp.Diagnostics.Append(d...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set state
	for key, val := range attrVals {
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root(key), val)...)
	}
}

func (r *GenericResource[T, PT]) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get the ID from state
	var id types.String
	diags := req.State.GetAttribute(ctx, path.Root("id"), &id)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Read current state to pass to Get (for PostReadProcess)
	stateValues := make(map[string]attr.Value)
	for key := range r.dataType.Schema {
		var val attr.Value
		d := req.State.GetAttribute(ctx, path.Root(key), &val)
		resp.Diagnostics.Append(d...)
		stateValues[key] = val
	}

	var readElement PT
	rawMap, _ := AttrValuesToMap(ctx, r.dataType.Schema, stateValues)
	if rawMap != nil {
		readElement = new(T)
		readElement.FromMap(rawMap)
	}

	value, err := r.dataType.convert(r.client).Get(id.ValueString(), readElement)
	if err != nil {
		resp.Diagnostics.AddError("Error reading resource", err.Error())
		return
	}

	if value == nil {
		// Resource was deleted outside of Terraform
		resp.State.RemoveResource(ctx)
		return
	}

	storedData := value.ToMap()
	storedData["id"] = id.ValueString()
	attrVals, d := MapToAttrValues(ctx, r.dataType.Schema, storedData)
	resp.Diagnostics.Append(d...)
	if resp.Diagnostics.HasError() {
		return
	}

	for key, val := range attrVals {
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root(key), val)...)
	}
}

func (r *GenericResource[T, PT]) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Get plan values
	planValues := make(map[string]attr.Value)
	for key := range r.dataType.Schema {
		var val attr.Value
		diags := req.Plan.GetAttribute(ctx, path.Root(key), &val)
		resp.Diagnostics.Append(diags...)
		planValues[key] = val
	}
	if resp.Diagnostics.HasError() {
		return
	}

	// Get the ID from state
	var id types.String
	diags := req.State.GetAttribute(ctx, path.Root("id"), &id)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	value, err := r.getData(ctx, planValues, true)
	if err != nil {
		resp.Diagnostics.AddError("Error parsing plan", err.Error())
		return
	}

	_, err = r.dataType.convert(r.client).Update(id.ValueString(), value)
	if err != nil {
		resp.Diagnostics.AddError("Error updating resource", err.Error())
		return
	}

	// Re-read the resource
	readValue, err := r.dataType.convert(r.client).Get(id.ValueString(), value)
	if err != nil {
		resp.Diagnostics.AddError("Error reading resource after update", err.Error())
		return
	}
	if readValue == nil {
		resp.Diagnostics.AddError("Resource not found after update", "")
		return
	}

	storedData := readValue.ToMap()
	storedData["id"] = id.ValueString()
	attrVals, d := MapToAttrValues(ctx, r.dataType.Schema, storedData)
	resp.Diagnostics.Append(d...)
	if resp.Diagnostics.HasError() {
		return
	}

	for key, val := range attrVals {
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root(key), val)...)
	}
}

func (r *GenericResource[T, PT]) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Get the ID from state
	var id types.String
	diags := req.State.GetAttribute(ctx, path.Root("id"), &id)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Read current state for the model (needed for PreDeleteProcess etc.)
	stateValues := make(map[string]attr.Value)
	for key := range r.dataType.Schema {
		var val attr.Value
		d := req.State.GetAttribute(ctx, path.Root(key), &val)
		resp.Diagnostics.Append(d...)
		stateValues[key] = val
	}

	var element PT
	rawMap, _ := AttrValuesToMap(ctx, r.dataType.Schema, stateValues)
	if rawMap != nil {
		element = new(T)
		element.FromMap(rawMap)
	}

	err := r.dataType.convert(r.client).Delete(id.ValueString(), element)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting resource", err.Error())
		return
	}
}

func (r *GenericResource[T, PT]) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
