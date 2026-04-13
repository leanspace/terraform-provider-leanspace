package provider

import (
	"context"
	"fmt"

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

func (r *GenericResource[T, PT]) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tfModel := r.dataType.NewTFModel()
	resp.Diagnostics.Append(req.Plan.Get(ctx, tfModel)...)
	if resp.Diagnostics.HasError() {
		return
	}

	apiPtr := tfModel.(TFToAPI).ToAPI().(PT)

	if v, ok := any(apiPtr).(ValidationModel); ok {
		if err := v.Validate(); err != nil {
			resp.Diagnostics.AddError("Validation error", err.Error())
			return
		}
	}

	createdValue, err := r.dataType.convert(r.client).Create(apiPtr)
	if err != nil {
		resp.Diagnostics.AddError("Error creating resource", err.Error())
		return
	}

	readValue, err := r.dataType.convert(r.client).Get(createdValue.GetID(), apiPtr)
	if err != nil {
		resp.Diagnostics.AddError("Error reading resource after create", err.Error())
		return
	}
	if readValue == nil {
		resp.Diagnostics.AddError("Resource not found after creation", "")
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, any(readValue).(APIToTF).ToTF())...)
}

func (r *GenericResource[T, PT]) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var id types.String
	diags := req.State.GetAttribute(ctx, path.Root("id"), &id)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readElement PT
	tfModel := r.dataType.NewTFModel()
	if d := req.State.Get(ctx, tfModel); !d.HasError() {
		if conv, ok := tfModel.(TFToAPI); ok {
			readElement = conv.ToAPI().(PT)
		}
	}

	value, err := r.dataType.convert(r.client).Get(id.ValueString(), readElement)
	if err != nil {
		resp.Diagnostics.AddError("Error reading resource", err.Error())
		return
	}

	if value == nil {
		resp.State.RemoveResource(ctx)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, any(value).(APIToTF).ToTF())...)
}

func (r *GenericResource[T, PT]) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tfModel := r.dataType.NewTFModel()
	resp.Diagnostics.Append(req.Plan.Get(ctx, tfModel)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var id types.String
	diags := req.State.GetAttribute(ctx, path.Root("id"), &id)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	apiPtr := tfModel.(TFToAPI).ToAPI().(PT)

	if v, ok := any(apiPtr).(ValidationModel); ok {
		if err := v.Validate(); err != nil {
			resp.Diagnostics.AddError("Validation error", err.Error())
			return
		}
	}

	_, err := r.dataType.convert(r.client).Update(id.ValueString(), apiPtr)
	if err != nil {
		resp.Diagnostics.AddError("Error updating resource", err.Error())
		return
	}

	readValue, err := r.dataType.convert(r.client).Get(id.ValueString(), apiPtr)
	if err != nil {
		resp.Diagnostics.AddError("Error reading resource after update", err.Error())
		return
	}
	if readValue == nil {
		resp.Diagnostics.AddError("Resource not found after update", "")
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, any(readValue).(APIToTF).ToTF())...)
}

func (r *GenericResource[T, PT]) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var id types.String
	diags := req.State.GetAttribute(ctx, path.Root("id"), &id)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var element PT
	tfModel := r.dataType.NewTFModel()
	if d := req.State.Get(ctx, tfModel); !d.HasError() {
		if conv, ok := tfModel.(TFToAPI); ok {
			element = conv.ToAPI().(PT)
		}
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
