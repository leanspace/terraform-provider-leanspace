package asset

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func (dataSource DataSourceType[T]) resourceGenericDefinition() *schema.Resource {
	return &schema.Resource{
		CreateContext: dataSource.create,
		ReadContext:   dataSource.get,
		UpdateContext: dataSource.update,
		DeleteContext: dataSource.delete,
		Schema: map[string]*schema.Schema{
			dataSource.Name: {
				Type:     schema.TypeList,
				MaxItems: 1,
				Required: true,
				Elem: &schema.Resource{
					Schema: dataSource.Schema,
				},
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func (dataSource DataSourceType[T]) create(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	client := m.(*Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	value := d.Get(dataSource.Name).([]any)
	valueData := dataSource.getValueData(value)
	createdValue, err := dataSource.convert(client).Create(valueData)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(dataSource.GetID(createdValue))
	diags = append(diags, dataSource.get(ctx, d, m)...)

	return diags
}

func (dataSource DataSourceType[T]) get(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	client := m.(*Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	valueId := d.Id()
	value, err := dataSource.convert(client).Get(valueId)
	if err != nil {
		return diag.FromErr(err)
	}
	dataSource.setValueData(value, d)

	return diags
}

func (dataSource DataSourceType[T]) setValueData(value *T, d *schema.ResourceData) {
	valueList := make([]map[string]any, 1)

	valueList[0] = dataSource.StructToMap(*value)
	d.Set(dataSource.Name, valueList)
}

func (dataSource DataSourceType[T]) getValueData(valueList []any) T {
	value, _ := dataSource.MapToStruct(valueList[0].(map[string]any))
	return value
}

func (dataSource DataSourceType[T]) update(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	client := m.(*Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if d.HasChange(dataSource.Name) {
		valueId := d.Id()
		value := d.Get(dataSource.Name).([]any)
		_, err := dataSource.convert(client).Update(valueId, dataSource.getValueData(value))
		if err != nil {
			return diag.FromErr(err)
		}
		diags = append(diags, dataSource.get(ctx, d, m)...)

		return diags
	}

	return dataSource.create(ctx, d, m)

}

func (dataSource DataSourceType[T]) delete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	client := m.(*Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	valueId := d.Id()
	err := dataSource.convert(client).Delete(valueId)
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}
