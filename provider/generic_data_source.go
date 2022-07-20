package provider

import (
	"context"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"leanspace-terraform-provider/helper/general_objects"
)

func (dataSourceType DataSourceType[T, PT]) toDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceType.read,
		Schema:      general_objects.PaginatedListSchema(dataSourceType.Schema),
	}
}

func (dataSourceType DataSourceType[T, PT]) read(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	client := m.(*Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	values, err := dataSourceType.convert(client).GetAll()
	if err != nil {
		return diag.FromErr(err)
	}
	err = dataSourceType.setData(values, d)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

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