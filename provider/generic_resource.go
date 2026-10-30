package provider

import (
	"context"

	"github.com/leanspace/terraform-provider-leanspace/helper"

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
	}
}

func (dataSource DataSourceType[T, PT]) getSchemaKeys() []string {
	keys := []string{}
	for key := range dataSource.Schema {
		keys = append(keys, key)
	}
	return keys
}

func (dataSource DataSourceType[T, PT]) hasSecurityTags() bool {
	_, exists := dataSource.Schema["security_tags"]
	return exists
}

func (dataSource DataSourceType[T, PT]) getFilterSchemaKeys() []string {
	keys := []string{}
	for key := range dataSource.FilterSchema {
		keys = append(keys, key)
	}
	return keys
}

func (dataSource DataSourceType[T, PT]) getData(d *schema.ResourceData, checkValidity bool) (string, PT, error) {
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

	if checkValidity && helper.Implements[T, ValidationModel]() {
		err := any(value).(ValidationModel).Validate(valueRaw)
		if err != nil {
			return valueId, nil, err
		}
	}

	err := value.FromMap(valueRaw)
	return valueId, value, err
}

func (dataSource DataSourceType[T, PT]) create(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	client := m.(*Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	_, value, err := dataSource.getData(d, true)
	if err != nil {
		return diag.FromErr(err)
	}
	createdValue, err := dataSource.convert(client).Create(value)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(createdValue.GetID())
	if dataSource.hasSecurityTags() {
		configuration := d.GetRawConfig()
		if !configuration.IsNull() && configuration.IsKnown() && !configuration.GetAttr("security_tags").IsNull() {
			if err := dataSource.writeSecurityTags(ctx, client, d.Id(), value); err != nil {
				return diag.FromErr(err)
			}
		}
	}
	diags = append(diags, dataSource.get(ctx, d, m)...)

	return diags
}

func (dataSource DataSourceType[T, PT]) get(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	client := m.(*Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	valueId, value, err := dataSource.getData(d, false)
	if err != nil {
		return diag.FromErr(err)
	}
	value, err = dataSource.convert(client).Get(valueId, value)
	if err != nil {
		return diag.FromErr(err)
	}

	if value != nil {
		if dataSource.hasSecurityTags() {
			if err := dataSource.loadSecurityTagsAfterRead(ctx, client, valueId, value); err != nil {
				return diag.FromErr(err)
			}
		}
		storedData := value.ToMap()
		for _, key := range dataSource.getSchemaKeys() {
			err = d.Set(key, storedData[key])
			if err != nil {
				diags = append(diags, diag.FromErr(err)...)
			}
		}
	} else {
		// The remote object no longer exists, so clear its Terraform state
		// and stored attributes to let Terraform reconcile it on the next plan.
		d.SetId("")
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

	if dataSource.SkipHasChangeCheck {
		containsChange = true
	} else {
		for _, key := range dataSource.getSchemaKeys() {
			if key != "security_tags" && d.HasChange(key) {
				containsChange = true
				break
			}
		}
	}

	if containsChange || dataSource.hasSecurityTags() {
		valueId, value, err := dataSource.getData(d, true)
		if err != nil {
			return diag.FromErr(err)
		}
		if containsChange {
			_, err = dataSource.convert(client).Update(valueId, value)
			if err != nil {
				return diag.FromErr(err)
			}
		}
		if dataSource.hasSecurityTags() && d.HasChange("security_tags") {
			if err := dataSource.writeSecurityTags(ctx, client, valueId, value); err != nil {
				return diag.FromErr(err)
			}
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

	valueId, value, err := dataSource.getData(d, false)
	if err != nil {
		return diag.FromErr(err)
	}
	err = dataSource.convert(client).Delete(valueId, value)
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}
