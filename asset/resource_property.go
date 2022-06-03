package asset

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceProperty() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePropertyCreate,
		ReadContext:   resourcePropertyRead,
		UpdateContext: resourcePropertyUpdate,
		DeleteContext: resourcePropertyDelete,
		Schema: map[string]*schema.Schema{
			"property": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: propertySchema,
				},
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourcePropertyCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	property := d.Get("property").([]interface{})
	propertyData := getPropertyData(property)
	createProperty, err := client.CreateProperty(propertyData.NodeId, propertyData)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(createProperty.ID)
	for _, diag := range resourcePropertyRead(ctx, d, m) {
		diags = append(diags, diag)
	}

	return diags
}

func resourcePropertyRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	propertyId := d.Id()
	property, err := client.GetProperty(propertyId)
	if err != nil {
		return diag.FromErr(err)
	}
	setPropertyData(property, d)

	return diags
}

func setPropertyData(property *Property[interface{}], d *schema.ResourceData) {
	propertyList := make([]map[string]interface{}, 1)

	propertyList[0] = propertyStructToInterface(*property)
	d.Set("property", propertyList)
}

func fieldInterfaceToStruct(fieldList []interface{}) Field[interface{}] {
	field := fieldList[0].(map[string]interface{})
	fieldStruct := Field[interface{}]{}
	fieldStruct.ID = field["id"].(string)
	fieldStruct.Name = field["name"].(string)
	fieldStruct.Description = field["description"].(string)
	fieldStruct.CreatedAt = field["created_at"].(string)
	fieldStruct.CreatedBy = field["created_by"].(string)
	fieldStruct.LastModifiedAt = field["last_modified_at"].(string)
	fieldStruct.LastModifiedBy = field["last_modified_by"].(string)
	fieldStruct.Type = field["type"].(string)
	fieldStruct.Value = field["value"]

	return fieldStruct
}

func getPropertyData(propertyList []interface{}) Property[interface{}] {
	property := propertyList[0].(map[string]interface{})
	propertyMap := Property[interface{}]{}

	propertyMap.Name = property["name"].(string)
	propertyMap.Description = property["description"].(string)
	propertyMap.NodeId = property["node_id"].(string)
	propertyMap.CreatedAt = property["created_at"].(string)
	propertyMap.CreatedBy = property["created_by"].(string)
	propertyMap.LastModifiedAt = property["last_modified_at"].(string)
	propertyMap.LastModifiedBy = property["last_modified_by"].(string)
	propertyMap.Value = property["value"]
	propertyMap.Type = property["type"].(string)
	propertyMap.Tags = tagsInterfaceToStruct(property["tags"])
	switch propertyMap.Type {
	case "NUMERIC":
		propertyMap.Min = property["min"].(float64)
		propertyMap.Max = property["max"].(float64)
		propertyMap.Scale = property["scale"].(int)
		propertyMap.Precision = property["precision"].(int)
		propertyMap.UnitId = property["unit_id"].(string)
	case "ENUM":
		if property["options"] != nil {
			option := property["options"].(map[string]interface{})
			propertyMap.Options = &option
		}
	case "TEXT":
		propertyMap.MinLength = property["min_length"].(int)
		propertyMap.MaxLength = property["max_length"].(int)
		propertyMap.Pattern = property["pattern"].(string)
	case "TIMESTAMP", "DATE", "TIME":
		propertyMap.Before = property["before"].(string)
		propertyMap.After = property["after"].(string)
	case "BOOLEAN":
	case "GEOPOINT":
		if property["fields"] != nil {
			propertyMap.Fields = &Fields{}
			propertyMap.Fields.Elevation = fieldInterfaceToStruct(property["fields"].([]interface{})[0].(map[string]interface{})["elevation"].([]interface{}))
			propertyMap.Fields.Latitude = fieldInterfaceToStruct(property["fields"].([]interface{})[0].(map[string]interface{})["latitude"].([]interface{}))
			propertyMap.Fields.Longitude = fieldInterfaceToStruct(property["fields"].([]interface{})[0].(map[string]interface{})["longitude"].([]interface{}))
		}
	}

	return propertyMap
}

func resourcePropertyUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if d.HasChange("property") {
		propertyId := d.Id()
		property := d.Get("property").([]interface{})
		_, err := client.UpdateProperty(propertyId, getPropertyData(property))
		if err != nil {
			return diag.FromErr(err)
		}
		for _, diag := range resourcePropertyRead(ctx, d, m) {
			diags = append(diags, diag)
		}

		return diags
	}

	return resourcePropertyCreate(ctx, d, m)

}

func resourcePropertyDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	propertyId := d.Id()
	err := client.DeleteProperty(propertyId)
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}
