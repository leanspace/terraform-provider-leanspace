package provider

import (
	"context"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

func (dataSourceType DataSourceType[T, PT]) toDataSource() *schema.Resource {
	var schemaType map[string]*schema.Schema
	if dataSourceType.IsUnique {
		schemaType = dataSourceType.FilterSchema
	} else {
		schemaType = general_objects.PaginatedListSchema(dataSourceType.Schema, dataSourceType.FilterSchema)
	}
	return &schema.Resource{
		ReadContext: dataSourceType.read,
		Schema:      schemaType,
	}
}

func (dataSourceType DataSourceType[T, PT]) read(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	client := m.(*Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	var filters map[string]any = nil
	if f, hasFilters := d.Get("filters").([]any); hasFilters && len(f) > 0 {
		filters = f[0].(map[string]any)
	}
	var genericClient GenericClient[T, PT] = dataSourceType.convert(client)
	var err error
	if genericClient.IsUnique == true {
		diags = dataSourceType.setUniqueFilter(genericClient, d, diags)
	} else {
		values, err := genericClient.GetAll(filters)
		err = dataSourceType.setData(values, d)
		if err != nil {
			return diag.FromErr(err)
		}
		d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	}
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func (dataSourceType DataSourceType[T, PT]) setUniqueFilter(genericClient GenericClient[T, PT], d *schema.ResourceData, diags diag.Diagnostics) diag.Diagnostics {
	value, err := genericClient.GetUnique()
	if value != nil {
		storedData := value.ToMap()
		for _, key := range dataSourceType.getFilterSchemaKeys() {
			err = d.Set(key, storedData[key])
			if err != nil {
				diags = append(diags, diag.FromErr(err)...)
			}
		}
		d.SetId(value.GetID())
	} else { // Object was not found (404)
		for _, key := range dataSourceType.getFilterSchemaKeys() {
			err = d.Set(key, nil)
			if err != nil {
				diags = append(diags, diag.FromErr(err)...)
			}
		}
	}
	return diags
}

func (dataSourceType DataSourceType[T, PT]) setData(paginatedList *general_objects.PaginatedList[T, PT], d *schema.ResourceData) error {
	paginatedListMap := paginatedList.ToMap()

	for key, value := range paginatedListMap {
		err := d.Set(key, value)
		if err != nil {
			return err
		}
	}
	return nil
}
