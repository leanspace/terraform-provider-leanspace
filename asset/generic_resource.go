package asset

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func (dataSource DataSourceType[T, PT]) resourceGenericDefinition() *schema.Resource {
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

func (dataSource DataSourceType[T, PT]) create(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	client := m.(*Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	value := d.Get(dataSource.Name).([]any)
	valueData, err := dataSource.getValueData(value)
	if err != nil {
		return diag.FromErr(err)
	}
	createdValue, err := dataSource.convert(client).Create(valueData)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(createdValue.GetID())
	diags = append(diags, dataSource.get(ctx, d, m)...)

	return diags
}

func (dataSource DataSourceType[T, PT]) get(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	client := m.(*Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	valueId := d.Id()
	value, err := dataSource.convert(client).Get(valueId)
	if err != nil {
		return diag.FromErr(err)
	}

	var storedData any = nil
	if value != nil {
		storedData = []map[string]any{value.ToMap()}
	}
	err = d.Set(dataSource.Name, storedData)
	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func (dataSource DataSourceType[T, PT]) getValueData(valueList []any) (PT, error) {
	var value PT = new(T)
	err := value.FromMap(valueList[0].(map[string]any))
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (dataSource DataSourceType[T, PT]) update(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	client := m.(*Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if d.HasChange(dataSource.Name) {
		valueId := d.Id()
		value := d.Get(dataSource.Name).([]any)
		valueData, err := dataSource.getValueData(value)
		if err != nil {
			return diag.FromErr(err)
		}
		_, err = dataSource.convert(client).Update(valueId, valueData)
		if err != nil {
			return diag.FromErr(err)
		}
		diags = append(diags, dataSource.get(ctx, d, m)...)

		return diags
	}

	return dataSource.create(ctx, d, m)

}

func (dataSource DataSourceType[T, PT]) delete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
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
