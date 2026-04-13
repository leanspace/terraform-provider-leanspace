package provider

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	datasourceschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

// GenericDataSource implements datasource.DataSource for any DataSourceType.
type GenericDataSource[T any, PT ParseableModel[T]] struct {
	dataType *DataSourceType[T, PT]
	client   *Client
}

func NewGenericDataSource[T any, PT ParseableModel[T]](dt *DataSourceType[T, PT]) func() datasource.DataSource {
	return func() datasource.DataSource {
		return &GenericDataSource[T, PT]{dataType: dt}
	}
}

func (d *GenericDataSource[T, PT]) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = d.dataType.ResourceIdentifier
}

func (d *GenericDataSource[T, PT]) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	if d.dataType.IsUnique {
		attrs, blocks := SplitDatasourceSchemaBlocks(d.dataType.FilterSchema)
		resp.Schema = datasourceschema.Schema{
			Attributes: attrs,
			Blocks:     blocks,
		}
	} else {
		allAttrs := general_objects.PaginatedListSchemaDS(d.dataType.DataSourceSchema, d.dataType.FilterSchema)
		attrs, blocks := SplitDatasourceSchemaBlocks(allAttrs)
		resp.Schema = datasourceschema.Schema{
			Attributes: attrs,
			Blocks:     blocks,
		}
	}
}

func (d *GenericDataSource[T, PT]) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client, ok := req.ProviderData.(*Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Provider Data", fmt.Sprintf("Expected *Client, got %T", req.ProviderData))
		return
	}
	d.client = client
}

func (d *GenericDataSource[T, PT]) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	genericClient := d.dataType.convert(d.client)

	if genericClient.IsUnique {
		d.readUnique(ctx, req, resp, genericClient)
	} else {
		d.readPaginated(ctx, req, resp, genericClient)
	}
}

func (d *GenericDataSource[T, PT]) readUnique(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse, genericClient GenericClient[T, PT]) {
	value, err := genericClient.GetUnique()
	if err != nil {
		resp.Diagnostics.AddError("Error reading data source", err.Error())
		return
	}

	if value != nil {
		storedData := value.ToMap()
		attrVals, diags := MapToAttrValuesDatasource(ctx, d.dataType.FilterSchema, storedData)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
		for key, val := range attrVals {
			resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root(key), val)...)
		}
		// Set ID
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(value.GetID()))...)
	}
}

func (d *GenericDataSource[T, PT]) readPaginated(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse, genericClient GenericClient[T, PT]) {
	// Extract filters from config
	var filters map[string]any
	var filtersObj types.Object
	diags := req.Config.GetAttribute(ctx, path.Root("filters"), &filtersObj)
	resp.Diagnostics.Append(diags...)

	if !filtersObj.IsNull() && !filtersObj.IsUnknown() {
		// Build the filters schema attributes map to convert
		filterSchemaAttrs := general_objects.FilterSchemaDS(d.dataType.FilterSchema)
		filterMap, d := attrValuesToMapDS(ctx, filterSchemaAttrs, filtersObj.Attributes())
		resp.Diagnostics.Append(d...)
		if resp.Diagnostics.HasError() {
			return
		}
		filters = filterMap
	}

	values, err := genericClient.GetAll(filters)
	if err != nil {
		resp.Diagnostics.AddError("Error reading data source", err.Error())
		return
	}

	paginatedListMap := values.ToMap()

	// Build the full DS schema for conversion
	dsSchema := general_objects.PaginatedListSchemaDS(d.dataType.DataSourceSchema, d.dataType.FilterSchema)
	attrVals, d2 := MapToAttrValuesDatasource(ctx, dsSchema, paginatedListMap)
	resp.Diagnostics.Append(d2...)
	if resp.Diagnostics.HasError() {
		return
	}

	for key, val := range attrVals {
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root(key), val)...)
	}

	// Set a synthetic ID
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(strconv.FormatInt(time.Now().Unix(), 10)))...)
}

// attrValuesToMapDS converts datasource attr values to a raw map.
func attrValuesToMapDS(ctx context.Context, schemaAttrs map[string]datasourceschema.Attribute, values map[string]attr.Value) (map[string]any, diag.Diagnostics) {
	return AttrValuesToMapDatasource(ctx, schemaAttrs, values)
}

// AttrValuesToMapDatasource converts framework attr.Values back into map[string]any for datasource schemas.
func AttrValuesToMapDatasource(ctx context.Context, schemaAttrs map[string]datasourceschema.Attribute, values map[string]attr.Value) (map[string]any, diag.Diagnostics) {
	result := make(map[string]any)
	var diags diag.Diagnostics
	for key, schemaAttr := range schemaAttrs {
		attrType := schemaAttr.GetType()
		val, d := attrValueToNative(ctx, attrType, values[key])
		diags.Append(d...)
		result[key] = val
	}
	return result, diags
}
