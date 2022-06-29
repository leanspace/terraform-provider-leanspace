package asset

import (
	"context"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mitchellh/mapstructure"

	"terraform-provider-asset/asset/general_objects"
)

func (dataSourceType DataSourceType[T]) toDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceType.read,
		Schema:      general_objects.PaginatedListSchema(dataSourceType.Schema),
	}
}

func (dataSourceType DataSourceType[T]) read(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	client := m.(*Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	values, err := dataSourceType.convert(client).GetAll()
	if err != nil {
		return diag.FromErr(err)
	}
	dataSourceType.setData(values, d)
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func (dataSourceType DataSourceType[T]) setData(paginatedList *general_objects.PaginatedList[T], d *schema.ResourceData) {
	data_as_map := make(map[string]any)
	decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		TagName: "terra",
		Result:  &data_as_map,
	})
	decoder.Decode(paginatedList)

	for key, value := range data_as_map {
		d.Set(key, value)
	}
}
