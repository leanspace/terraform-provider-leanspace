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

func propertyStructToInterface(property Property[any]) map[string]any {
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
	propertyMap["tags"] = general_objects.TagsStructToInterface(property.Tags)
	switch property.Type {
	case "NUMERIC":
		if property.Value != nil {
			propertyMap["value"] = strconv.FormatFloat(property.Value.(float64), 'g', -1, 64)
		}
		propertyMap["min"] = property.Min
		propertyMap["max"] = property.Max
		propertyMap["scale"] = property.Scale
		propertyMap["precision"] = property.Precision
		propertyMap["unit_id"] = property.UnitId
	case "ENUM":
		if property.Value != nil {
			propertyMap["value"] = strconv.FormatFloat(property.Value.(float64), 'g', -1, 64)
		}
		if property.Options != nil {
			propertyMap["options"] = *property.Options
		}
	case "TEXT":
		if property.Value != nil {
			propertyMap["value"] = property.Value.(string)
		}
		propertyMap["min_length"] = property.MinLength
		propertyMap["max_length"] = property.MaxLength
		propertyMap["pattern"] = property.Pattern
	case "TIMESTAMP", "DATE", "TIME":
		if property.Value != nil {
			propertyMap["value"] = property.Value.(string)
		}
		propertyMap["before"] = property.Before
		propertyMap["after"] = property.After
	case "BOOLEAN":
		if property.Value != nil {
			propertyMap["value"] = strconv.FormatBool(property.Value.(bool))
		}
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

func getPropertyData(property map[string]any) (Property[any], error) {
	propertyMap := Property[any]{}

	propertyMap.Name = property["name"].(string)
	propertyMap.Description = property["description"].(string)
	propertyMap.NodeId = property["node_id"].(string)
	propertyMap.CreatedAt = property["created_at"].(string)
	propertyMap.CreatedBy = property["created_by"].(string)
	propertyMap.LastModifiedAt = property["last_modified_at"].(string)
	propertyMap.LastModifiedBy = property["last_modified_by"].(string)
	propertyMap.Value = property["value"]
	propertyMap.Type = property["type"].(string)
	propertyMap.Tags = general_objects.TagsInterfaceToStruct(property["tags"])
	switch propertyMap.Type {
	case "NUMERIC":
		propertyMap.Min = property["min"].(float64)
		propertyMap.Max = property["max"].(float64)
		propertyMap.Scale = property["scale"].(int)
		propertyMap.Precision = property["precision"].(int)
		propertyMap.UnitId = property["unit_id"].(string)
	case "ENUM":
		if property["options"] != nil {
			option := property["options"].(map[string]any)
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
		// no extra property for booleans
	case "GEOPOINT":
		if property["fields"] != nil {
			propertyMap.Fields = &Fields{}
			propertyMap.Fields.Elevation = fieldInterfaceToStruct(property["fields"].([]any)[0].(map[string]any)["elevation"].([]any))
			propertyMap.Fields.Latitude = fieldInterfaceToStruct(property["fields"].([]any)[0].(map[string]any)["latitude"].([]any))
			propertyMap.Fields.Longitude = fieldInterfaceToStruct(property["fields"].([]any)[0].(map[string]any)["longitude"].([]any))
		}
	}

	return propertyMap, nil
}
