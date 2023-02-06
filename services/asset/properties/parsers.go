package properties

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func (property *Property[T]) ToMap() map[string]any {
	propertyMap := make(map[string]any)
	propertyMap["id"] = property.ID
	propertyMap["name"] = property.Name
	propertyMap["description"] = property.Description
	propertyMap["built_in"] = property.IsBuiltIn
	propertyMap["node_id"] = property.NodeId
	propertyMap["created_at"] = property.CreatedAt
	propertyMap["created_by"] = property.CreatedBy
	propertyMap["last_modified_at"] = property.LastModifiedAt
	propertyMap["last_modified_by"] = property.LastModifiedBy
	propertyMap["tags"] = helper.ParseToMaps(property.Tags)
	propertyMap["type"] = property.Attributes.Type
	switch property.Attributes.Type {
	case "NUMERIC":
		if any(property.Attributes.Value) != nil {
			propertyMap["value"] = helper.ParseFloat(any(property.Attributes.Value).(float64))
		}
		propertyMap["min"] = property.Attributes.Min
		propertyMap["max"] = property.Attributes.Max
		propertyMap["scale"] = property.Attributes.Scale
		propertyMap["precision"] = property.Attributes.Precision
		propertyMap["unit_id"] = property.Attributes.UnitId
	case "ENUM":
		if any(property.Attributes.Value) != nil {
			propertyMap["value"] = helper.ParseFloat(any(property.Attributes.Value).(float64))
		}
		if property.Attributes.Options != nil {
			propertyMap["options"] = *property.Attributes.Options
		}
	case "TEXT":
		if any(property.Attributes.Value) != nil {
			propertyMap["value"] = property.Attributes.Value
		}
		propertyMap["min_length"] = property.Attributes.MinLength
		propertyMap["max_length"] = property.Attributes.MaxLength
		propertyMap["pattern"] = property.Attributes.Pattern
	case "TIMESTAMP", "DATE", "TIME":
		if any(property.Attributes.Value) != nil {
			propertyMap["value"] = property.Attributes.Value
		}
		propertyMap["before"] = property.Attributes.Before
		propertyMap["after"] = property.Attributes.After
	case "BOOLEAN":
		if any(property.Attributes.Value) != nil {
			propertyMap["value"] = strconv.FormatBool(any(property.Attributes.Value).(bool))
		}
	case "GEOPOINT":
		if property.Attributes.Fields != nil {
			fieldList := make([]map[string]any, 1)
			fieldMap := make(map[string]any)
			elevationList := make([]map[string]any, 1)
			elevationList[0] = (&property.Attributes.Fields.Elevation).ToMap()
			fieldMap["elevation"] = elevationList
			latitudeList := make([]map[string]any, 1)
			latitudeList[0] = (&property.Attributes.Fields.Latitude).ToMap()
			fieldMap["latitude"] = latitudeList
			longitudeList := make([]map[string]any, 1)
			longitudeList[0] = (&property.Attributes.Fields.Longitude).ToMap()
			fieldMap["longitude"] = longitudeList
			fieldList[0] = fieldMap
			propertyMap["fields"] = fieldList
		}
	case "TLE":
		if any(property.Attributes.Value) != nil {
			var tleValue string
			var tleValues []interface{} = any(property.Attributes.Value).([]interface{})
			for _, value := range tleValues {
				tleValue = tleValue + "," + fmt.Sprint(value)
			}
			propertyMap["value"] = strings.TrimPrefix(tleValue, ",")
		}
	}
	return propertyMap
}

func (property *Property[T]) FromMap(propertyMap map[string]any) error {
	property.ID = propertyMap["id"].(string)
	property.Name = propertyMap["name"].(string)
	property.LastModifiedAt = propertyMap["last_modified_at"].(string)
	property.LastModifiedBy = propertyMap["last_modified_by"].(string)
	property.IsBuiltIn = propertyMap["built_in"].(bool)
	property.NodeId = propertyMap["node_id"].(string)
	property.CreatedAt = propertyMap["created_at"].(string)
	property.CreatedBy = propertyMap["created_by"].(string)
	property.Description = propertyMap["description"].(string)
	if tags, err := helper.ParseFromMaps[general_objects.Tag](propertyMap["tags"].(*schema.Set).List()); err != nil {
		return err
	} else {
		property.Tags = tags
	}
	property.Attributes.Type = propertyMap["type"].(string)
	switch property.Attributes.Type {
	case "NUMERIC":
		if propertyMap["value"] != nil {
			property.Attributes.Value = any(propertyMap["value"]).(T)
		}
		property.Attributes.Min = propertyMap["min"].(float64)
		property.Attributes.Max = propertyMap["max"].(float64)
		property.Attributes.Scale = propertyMap["scale"].(int)
		property.Attributes.Precision = propertyMap["precision"].(int)
		property.Attributes.UnitId = propertyMap["unit_id"].(string)
	case "ENUM":
		if propertyMap["value"] != nil {
			property.Attributes.Value = any(propertyMap["value"]).(T)
		}
		if propertyMap["options"] != nil {
			option := propertyMap["options"].(map[string]any)
			property.Attributes.Options = &option
		}
	case "TEXT":
		if propertyMap["value"] != nil {
			property.Attributes.Value = any(propertyMap["value"]).(T)
		}
		property.Attributes.MinLength = propertyMap["min_length"].(int)
		property.Attributes.MaxLength = propertyMap["max_length"].(int)
		property.Attributes.Pattern = propertyMap["pattern"].(string)
	case "TIMESTAMP", "DATE", "TIME":
		if propertyMap["value"] != nil {
			property.Attributes.Value = any(propertyMap["value"]).(T)
		}
		property.Attributes.Before = propertyMap["before"].(string)
		property.Attributes.After = propertyMap["after"].(string)
	case "BOOLEAN":
		if propertyMap["value"] != nil {
			property.Attributes.Value = any(propertyMap["value"]).(T)
		}
	case "TLE":
		if tleValue, ok := propertyMap["value"]; ok {
			var stringTleValues []string = strings.Split(tleValue.(string), ",")
			if len(stringTleValues) == 2 {
				var interfaceOfTleValues []interface{}
				for _, str := range stringTleValues {
					var stringValue = strings.TrimSpace(str)
					interfaceOfTleValues = append(interfaceOfTleValues, stringValue)
				}
				property.Attributes.Value = any(interfaceOfTleValues).(T)
			}
		}
	case "GEOPOINT":
		if propertyMap["fields"] != nil {
			fields := propertyMap["fields"].([]any)[0].(map[string]any)
			property.Attributes.Fields = &Fields{}
			property.Attributes.Fields.Elevation.FromMap(fields["elevation"].([]any)[0].(map[string]any))
			property.Attributes.Fields.Latitude.FromMap(fields["latitude"].([]any)[0].(map[string]any))
			property.Attributes.Fields.Longitude.FromMap(fields["longitude"].([]any)[0].(map[string]any))
		}
	}

	return nil
}

func (field *Field[T]) ToMap() map[string]any {
	fieldMap := make(map[string]any)
	// This might be unsafe in the future - for now fields are only used for numbers
	// in the geopoint type so it's alright.
	if any(field.Value) != nil {
		fieldMap["value"] = helper.ParseFloat(any(field.Value).(float64))
	}
	fieldMap["min"] = field.Min
	fieldMap["max"] = field.Max
	fieldMap["scale"] = field.Scale
	fieldMap["precision"] = field.Precision
	fieldMap["unit_id"] = field.UnitId
	return fieldMap
}

func (field *Field[T]) FromMap(fieldMap map[string]any) error {
	field.Value = fieldMap["value"].(T)
	field.Min = fieldMap["min"].(float64)
	field.Max = fieldMap["max"].(float64)
	field.Scale = fieldMap["scale"].(int)
	field.Precision = fieldMap["precision"].(int)
	field.UnitId = fieldMap["unit_id"].(string)
	return nil
}
