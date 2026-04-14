package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	datasourceschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
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
	dsSchema := d.dataType.DataSourceSchema
	if dsSchema == nil {
		dsSchema = ResourceSchemaToDataSource(d.dataType.Schema)
	}
	if d.dataType.IsUnique {
		dsAttrs := make(map[string]datasourceschema.Attribute)
		for k, v := range d.dataType.FilterSchema {
			dsAttrs[k] = v
		}
		// Ensure unique data sources always expose an "id" attribute.
		if _, hasID := dsAttrs["id"]; !hasID {
			dsAttrs["id"] = datasourceschema.StringAttribute{Computed: true}
		}
		attrs, blocks := SplitDatasourceSchemaBlocks(dsAttrs)
		resp.Schema = datasourceschema.Schema{
			Attributes: attrs,
			Blocks:     blocks,
		}
	} else {
		allAttrs := general_objects.PaginatedListSchemaDS(dsSchema, d.dataType.FilterSchema)
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
	if value == nil {
		return
	}

	// Use the APIToDSTF interface to convert the API model to a data-source TF model.
	if dstf, ok := any(value).(APIToDSTF); ok {
		resp.Diagnostics.Append(resp.State.Set(ctx, dstf.ToDSTF())...)
	} else {
		// Fallback: set only the ID.
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(value.GetID()))...)
	}
}

func (d *GenericDataSource[T, PT]) readPaginated(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse, genericClient GenericClient[T, PT]) {
	// Extract filters from config
	var filtersObj types.Object
	diags := req.Config.GetAttribute(ctx, path.Root("filters"), &filtersObj)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var filters map[string]any
	if !filtersObj.IsNull() && !filtersObj.IsUnknown() {
		filters = general_objects.FilterObjectToMap(filtersObj)
	}

	values, err := genericClient.GetAll(filters)
	if err != nil {
		resp.Diagnostics.AddError("Error reading data source", err.Error())
		return
	}

	// Build content items by calling ToTF() on each API model.
	tfModels := make([]any, len(values.Content))
	for i := range values.Content {
		if apiToTF, ok := any(&values.Content[i]).(APIToTF); ok {
			tfModels[i] = apiToTF.ToTF()
		}
	}

	// Set all non-content state fields from the paginated metadata.
	result := values.ToDataSourceTF(filtersObj)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), result.ID)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("total_elements"), result.TotalElements)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("total_pages"), result.TotalPages)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("number_of_elements"), result.NumberOfElements)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("number"), result.Number)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("size"), result.Size)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("sort"), result.Sort)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("first"), result.First)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("last"), result.Last)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("empty"), result.Empty)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("pageable"), result.Pageable)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("filters"), result.Filters)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set content separately so the framework uses the schema's nested-attribute type
	// to reflect each ToTF() struct into a properly typed object.
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("content"), tfModels)...)
}
