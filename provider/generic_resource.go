package provider

import (
	"context"
	"leanspace-terraform-provider/helper"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func (dataSource DataSourceType[T, PT]) toResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: dataSource.create,
		ReadContext:   dataSource.get,
		UpdateContext: dataSource.update,
		DeleteContext: dataSource.delete,
		Schema:        dataSource.Schema,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		CustomizeDiff: func(ctx context.Context, rd *schema.ResourceDiff, i any) error {
			if helper.IsInstance[T, ValidationModel]() {
				_, value, err := dataSource.getData(rd)
				if err != nil {
					return err
				}
				err = any(value).(ValidationModel).Validate()
				if err != nil {
					return err
				}
			}
			return nil
		},
	}
}

type ResourceData interface {
	Id() string
	Get(key string) any
}

func (dataSource DataSourceType[T, PT]) getSchemaKeys() []string {
	keys := []string{}
	for key := range dataSource.Schema {
		keys = append(keys, key)
	}
	return keys
}

func (dataSource DataSourceType[T, PT]) getData(d ResourceData) (string, PT, error) {
	valueId := d.Id()
	onlyNil := true
	valueRaw := make(map[string]any)
	for _, key := range dataSource.getSchemaKeys() {
		valueRaw[key] = d.Get(key)
		if valueRaw[key] != nil {
			onlyNil = false
		}
	}
	if onlyNil || len(valueRaw) == 0 {
		return valueId, nil, nil
	}
	var value PT = new(T)
	error := value.FromMap(valueRaw)
	return valueId, value, error
}

func (dataSource DataSourceType[T, PT]) create(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	client := m.(*Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	_, value, err := dataSource.getData(d)
	if err != nil {
		return diag.FromErr(err)
	}
	createdValue, err := dataSource.convert(client).Create(value)
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

	valueId, value, err := dataSource.getData(d)
	if err != nil {
		return diag.FromErr(err)
	}
	value, err = dataSource.convert(client).Get(valueId, value)
	if err != nil {
		return diag.FromErr(err)
	}

	if value != nil {
		storedData := value.ToMap()
		for _, key := range dataSource.getSchemaKeys() {
			err = d.Set(key, storedData[key])
			if err != nil {
				diags = append(diags, diag.FromErr(err)...)
			}
		}
	} else { // Object was not found (404)
		for _, key := range dataSource.getSchemaKeys() {
			err = d.Set(key, nil)
			if err != nil {
				diags = append(diags, diag.FromErr(err)...)
			}
		}
	}
	return diags
}

func (dataSource DataSourceType[T, PT]) update(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	client := m.(*Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	containsChange := false

	for _, key := range dataSource.getSchemaKeys() {
		if d.HasChange(key) {
			containsChange = true
			break
		}
	}

	if containsChange {
		valueId, value, err := dataSource.getData(d)
		if err != nil {
			return diag.FromErr(err)
		}
		_, err = dataSource.convert(client).Update(valueId, value)
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

	valueId, value, err := dataSource.getData(d)
	if err != nil {
		return diag.FromErr(err)
	}
	err = dataSource.convert(client).Delete(valueId, value)
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}
