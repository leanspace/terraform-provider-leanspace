package properties

import (
	"strconv"
	"terraform-provider-asset/asset/general_objects"
)

func fieldStructToInterface(field *Field[any]) map[string]any {
	fieldMap := make(map[string]any)
	fieldMap["id"] = field.ID
	fieldMap["name"] = field.Name
	fieldMap["description"] = field.Description
	fieldMap["created_at"] = field.CreatedAt
	fieldMap["created_by"] = field.CreatedBy
	fieldMap["last_modified_at"] = field.LastModifiedAt
	fieldMap["last_modified_by"] = field.LastModifiedBy
	fieldMap["type"] = field.Type
	fieldMap["value"] = strconv.FormatFloat(field.Value.(float64), 'g', -1, 64)

	return fieldMap
}

func (property *Property[T]) ToMap() map[string]any {
	propertyMap := make(map[string]any)
	propertyMap["id"] = property.ID
	propertyMap["name"] = property.Name
	propertyMap["description"] = property.Description
	propertyMap["node_id"] = property.NodeId
	propertyMap["created_at"] = property.CreatedAt
	propertyMap["created_by"] = property.CreatedBy
	propertyMap["last_modified_at"] = property.LastModifiedAt
	propertyMap["last_modified_by"] = property.LastModifiedBy
	propertyMap["type"] = property.Type
	propertyMap["tags"] = general_objects.TagsStructToMap(property.Tags)
	switch property.Type {
	case "NUMERIC":
		propertyMap["value"] = strconv.FormatFloat(any(property.Value).(float64), 'g', -1, 64)
		propertyMap["min"] = property.Min
		propertyMap["max"] = property.Max
		propertyMap["scale"] = property.Scale
		propertyMap["precision"] = property.Precision
		propertyMap["unit_id"] = property.UnitId
	case "ENUM":
		propertyMap["value"] = strconv.FormatFloat(any(property.Value).(float64), 'g', -1, 64)
		if property.Options != nil {
			propertyMap["options"] = *property.Options
		}
	case "TEXT":
		propertyMap["value"] = property.Value
		propertyMap["min_length"] = property.MinLength
		propertyMap["max_length"] = property.MaxLength
		propertyMap["pattern"] = property.Pattern
	case "TIMESTAMP", "DATE", "TIME":
		propertyMap["value"] = property.Value
		propertyMap["before"] = property.Before
		propertyMap["after"] = property.After
	case "BOOLEAN":
		propertyMap["value"] = strconv.FormatBool(any(property.Value).(bool))
	case "GEOPOINT":
		if property.Fields != nil {
			fieldList := make([]map[string]any, 1)
			fieldMap := make(map[string]any)
			elevationList := make([]map[string]any, 1)
			elevationList[0] = fieldStructToInterface(&property.Fields.Elevation)
			fieldMap["elevation"] = elevationList
			latitudeList := make([]map[string]any, 1)
			latitudeList[0] = fieldStructToInterface(&property.Fields.Latitude)
			fieldMap["latitude"] = latitudeList
			longitudeList := make([]map[string]any, 1)
			longitudeList[0] = fieldStructToInterface(&property.Fields.Longitude)
			fieldMap["longitude"] = longitudeList
			fieldList[0] = fieldMap
			propertyMap["fields"] = fieldList
		}
	}
	return propertyMap
}

func fieldInterfaceToStruct(fieldList []any) Field[any] {
	field := fieldList[0].(map[string]any)
	fieldStruct := Field[any]{}
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

func (property *Property[T]) FromMap(propertyMap map[string]any) error {
	property.Name = propertyMap["name"].(string)
	property.Description = propertyMap["description"].(string)
	property.NodeId = propertyMap["node_id"].(string)
	property.CreatedAt = propertyMap["created_at"].(string)
	property.CreatedBy = propertyMap["created_by"].(string)
	property.LastModifiedAt = propertyMap["last_modified_at"].(string)
	property.LastModifiedBy = propertyMap["last_modified_by"].(string)
	property.Value = propertyMap["value"].(T)
	property.Type = propertyMap["type"].(string)
	property.Tags = general_objects.TagsInterfaceToStruct(propertyMap["tags"])
	switch property.Type {
	case "NUMERIC":
		property.Min = propertyMap["min"].(float64)
		property.Max = propertyMap["max"].(float64)
		property.Scale = propertyMap["scale"].(int)
		property.Precision = propertyMap["precision"].(int)
		property.UnitId = propertyMap["unit_id"].(string)
	case "ENUM":
		if propertyMap["options"] != nil {
			option := propertyMap["options"].(map[string]any)
			property.Options = &option
		}
	case "TEXT":
		property.MinLength = propertyMap["min_length"].(int)
		property.MaxLength = propertyMap["max_length"].(int)
		property.Pattern = propertyMap["pattern"].(string)
	case "TIMESTAMP", "DATE", "TIME":
		property.Before = propertyMap["before"].(string)
		property.After = propertyMap["after"].(string)
	case "BOOLEAN":
		// no extra property for booleans
	case "GEOPOINT":
		if propertyMap["fields"] != nil {
			property.Fields = &Fields{}
			property.Fields.Elevation = fieldInterfaceToStruct(propertyMap["fields"].([]any)[0].(map[string]any)["elevation"].([]any))
			property.Fields.Latitude = fieldInterfaceToStruct(propertyMap["fields"].([]any)[0].(map[string]any)["latitude"].([]any))
			property.Fields.Longitude = fieldInterfaceToStruct(propertyMap["fields"].([]any)[0].(map[string]any)["longitude"].([]any))
		}
	}

	return nil
}
