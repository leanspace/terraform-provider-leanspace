package types

import (
	"fmt"
	"strconv"
	"strings"
)

// formatFloat mirrors helper.ParseFloat: it parses a float to a string,
// kept local so this package has no dependency on the helper package
// (which pulls in terraform-plugin-sdk).
func formatFloat(num float64) string {
	return strconv.FormatFloat(num, 'g', -1, 64)
}

func (keyValue *KeyValue) ToMap() map[string]any {
	keyValueMap := make(map[string]any)
	keyValueMap["key"] = keyValue.Key
	keyValueMap["value"] = keyValue.Value
	return keyValueMap
}

func (keyValue *KeyValue) FromMap(keyValueMap map[string]any) error {
	keyValue.Key = keyValueMap["key"].(string)
	keyValue.Value = keyValueMap["value"].(string)
	return nil
}

func (attribute *ValueAttribute[T]) ToMap() map[string]any {
	attributeMap := make(map[string]any)
	attributeMap["type"] = attribute.Type
	attributeMap["data_type"] = attribute.DataType
	switch attribute.Type {
	case "NUMERIC":
		attributeMap["value"] = formatFloat(any(attribute.Value).(float64))
		attributeMap["unit_id"] = attribute.UnitId
	case "TEXT", "TIMESTAMP", "DATE", "TIME", "BINARY":
		attributeMap["value"] = attribute.Value
	case "BOOLEAN":
		attributeMap["value"] = strconv.FormatBool(any(attribute.Value).(bool))
	case "GEOPOINT":
		if attribute.Fields != nil {
			fieldList := make([]map[string]any, 1)
			fieldMap := make(map[string]any)
			elevationList := make([]map[string]any, 1)
			elevationList[0] = (&attribute.Fields.Elevation).ToMap()
			fieldMap["elevation"] = elevationList
			latitudeList := make([]map[string]any, 1)
			latitudeList[0] = (&attribute.Fields.Latitude).ToMap()
			fieldMap["latitude"] = latitudeList
			longitudeList := make([]map[string]any, 1)
			longitudeList[0] = (&attribute.Fields.Longitude).ToMap()
			fieldMap["longitude"] = longitudeList
			fieldList[0] = fieldMap
			attributeMap["fields"] = fieldList
		}
	case "ARRAY":
		if any(attribute.Value) != nil {
			var value string
			var interfaceArrayValues []interface{} = any(attribute.Value).([]interface{})
			for _, arrayValue := range interfaceArrayValues {
				value = value + "," + fmt.Sprint(arrayValue)
			}
			attributeMap["value"] = strings.TrimPrefix(value, ",")
		}

	}
	return attributeMap
}

func (attribute *DefinitionAttribute[T]) FromMap(attributeMap map[string]any) error {
	attribute.Type = attributeMap["type"].(string)
	if value, ok := attributeMap["required"]; ok {
		b := value.(bool)
		attribute.Required = &b
	}
	switch attribute.Type {
	case "NUMERIC":
		attribute.Min = attributeMap["min"].(float64)
		attribute.Max = attributeMap["max"].(float64)
		attribute.Scale = attributeMap["scale"].(int)
		attribute.Precision = attributeMap["precision"].(int)
		attribute.UnitId = attributeMap["unit_id"].(string)
	case "ENUM":
		if attributeMap["options"] != nil {
			option := attributeMap["options"].(map[string]any)
			attribute.Options = &option
		}
	case "TEXT":
		attribute.MinLength = attributeMap["min_length"].(int)
		attribute.MaxLength = attributeMap["max_length"].(int)
		attribute.Pattern = attributeMap["pattern"].(string)
	case "BINARY":
		attribute.MinLength = attributeMap["min_length"].(int)
		attribute.MaxLength = attributeMap["max_length"].(int)
	case "TIMESTAMP", "DATE", "TIME":
		attribute.Before = attributeMap["before"].(string)
		attribute.After = attributeMap["after"].(string)
	case "BOOLEAN":
		// no extra field
	case "GEOPOINT":
		if attributeMap["fields"] != nil {
			attribute.Fields = &FieldsDef{}
			if len(attributeMap["fields"].([]any)) > 0 {
				fields := attributeMap["fields"].([]any)[0].(map[string]any)
				attribute.Fields.Elevation.FromMap(fields["elevation"].([]any)[0].(map[string]any))
				attribute.Fields.Latitude.FromMap(fields["latitude"].([]any)[0].(map[string]any))
				attribute.Fields.Longitude.FromMap(fields["longitude"].([]any)[0].(map[string]any))
			}
		}
	case "ARRAY":
		attribute.MinSize = attributeMap["min_size"].(int)
		attribute.MaxSize = attributeMap["max_size"].(int)
		attribute.Unique = attributeMap["unique"].(bool)
		if len(attributeMap["constraint"].([]any)) > 0 {
			err := attribute.Constraint.FromMap(attributeMap["constraint"].([]any)[0].(map[string]any))
			if err != nil {
				return err
			}
		}

	}
	if defaultValue, ok := attributeMap["default_value"]; ok {
		if attribute.Type == "ARRAY" {
			var stringDefaultValues []string = strings.Split(defaultValue.(string), ",")
			var interfaceOfDefaultValues []interface{}
			for _, str := range stringDefaultValues {
				var stringValue = strings.TrimSpace(str)
				switch attribute.Constraint.Type {
				case "NUMERIC":
					if numericValue, err := strconv.ParseFloat(stringValue, 64); err == nil {
						interfaceOfDefaultValues = append(interfaceOfDefaultValues, numericValue)
					}
				case "ENUM":
					if enumValue, err := strconv.ParseInt(stringValue, 10, 16); err == nil {
						interfaceOfDefaultValues = append(interfaceOfDefaultValues, enumValue)
					}
				case "BOOLEAN":
					if booleanValue, err := strconv.ParseBool(stringValue); err == nil {
						interfaceOfDefaultValues = append(interfaceOfDefaultValues, booleanValue)
					}
				case "TEXT", "TIMESTAMP", "DATE", "TIME", "BINARY":
					interfaceOfDefaultValues = append(interfaceOfDefaultValues, stringValue)
				}
			}
			attribute.DefaultValue = any(interfaceOfDefaultValues).(T)
		} else if attribute.Type != "GEOPOINT" {
			attribute.DefaultValue = defaultValue.(T)
		}
	}
	return nil
}

func (constraint *ArrayConstraint[T]) FromMap(constraintMap map[string]any) error {
	constraint.Type = constraintMap["type"].(string)
	if value, ok := constraintMap["required"]; ok {
		b := value.(bool)
		constraint.Required = &b
	}
	switch constraint.Type {
	case "NUMERIC":
		constraint.Min = constraintMap["min"].(float64)
		constraint.Max = constraintMap["max"].(float64)
		constraint.Scale = constraintMap["scale"].(int)
		constraint.Precision = constraintMap["precision"].(int)
		constraint.UnitId = constraintMap["unit_id"].(string)
	case "ENUM":
		if constraintMap["options"] != nil {
			option := constraintMap["options"].(map[string]any)
			constraint.Options = &option
		}
	case "TEXT":
		constraint.MinLength = constraintMap["min_length"].(int)
		constraint.MaxLength = constraintMap["max_length"].(int)
		constraint.Pattern = constraintMap["pattern"].(string)
	case "TIMESTAMP", "DATE", "TIME":
		constraint.Before = constraintMap["before"].(string)
		constraint.After = constraintMap["after"].(string)
	case "BOOLEAN":
		// no extra field
	case "BINARY":
		constraint.MinLength = constraintMap["min_length"].(int)
		constraint.MaxLength = constraintMap["max_length"].(int)
	}
	return nil
}

func (attribute *ValueAttribute[T]) FromMap(attributeMap map[string]any) error {

	attribute.Type = attributeMap["type"].(string)
	attribute.DataType = attributeMap["data_type"].(string)
	if attribute.Type == "NUMERIC" {
		attribute.UnitId = attributeMap["unit_id"].(string)
		attribute.Value = attributeMap["value"].(T)
	} else if attribute.Type == "ARRAY" {
		var stringValues []string = strings.Split(attributeMap["value"].(string), ",")
		var interfaceOfValues []interface{}
		for _, str := range stringValues {
			var stringValue = strings.TrimSpace(str)
			switch attribute.DataType {
			case "NUMERIC":
				if numericValue, err := strconv.ParseFloat(stringValue, 64); err == nil {
					interfaceOfValues = append(interfaceOfValues, numericValue)
				}
			case "ENUM":
				if enumValue, err := strconv.ParseInt(stringValue, 10, 16); err == nil {
					interfaceOfValues = append(interfaceOfValues, enumValue)
				}
			case "BOOLEAN":
				if booleanValue, err := strconv.ParseBool(stringValue); err == nil {
					interfaceOfValues = append(interfaceOfValues, booleanValue)
				}
			case "TEXT", "TIMESTAMP", "DATE", "TIME", "BINARY":
				interfaceOfValues = append(interfaceOfValues, stringValue)
			}
		}
		attribute.Value = any(interfaceOfValues).(T)
	} else if attribute.Type == "GEOPOINT" {
		if attributeMap["fields"] != nil {
			fields := attributeMap["fields"].([]any)[0].(map[string]any)
			attribute.Fields = &Fields{}
			attribute.Fields.Elevation.FromMap(fields["elevation"].([]any)[0].(map[string]any))
			attribute.Fields.Latitude.FromMap(fields["latitude"].([]any)[0].(map[string]any))
			attribute.Fields.Longitude.FromMap(fields["longitude"].([]any)[0].(map[string]any))
		}
	} else {
		attribute.Value = attributeMap["value"].(T)
	}
	return nil
}

func (attribute *DefinitionAttribute[T]) ToMap() map[string]any {
	attributeMap := make(map[string]any)

	attributeMap["type"] = attribute.Type

	if attribute.Required != nil {
		attributeMap["required"] = attribute.Required
	}

	switch attribute.Type {
	case "TEXT":
		if any(attribute.DefaultValue) != nil {
			attributeMap["default_value"] = attribute.DefaultValue
		}
		attributeMap["min_length"] = attribute.MinLength
		attributeMap["max_length"] = attribute.MaxLength
		attributeMap["pattern"] = attribute.Pattern
	case "BINARY":
		if any(attribute.DefaultValue) != nil {
			attributeMap["default_value"] = attribute.DefaultValue
		}
		attributeMap["min_length"] = attribute.MinLength
		attributeMap["max_length"] = attribute.MaxLength
	case "NUMERIC":
		if any(attribute.DefaultValue) != nil {
			attributeMap["default_value"] = formatFloat(any(attribute.DefaultValue).(float64))
		}
		attributeMap["min"] = attribute.Min
		attributeMap["max"] = attribute.Max
		attributeMap["scale"] = attribute.Scale
		attributeMap["precision"] = attribute.Precision
		attributeMap["unit_id"] = attribute.UnitId
	case "BOOLEAN":
		if any(attribute.DefaultValue) != nil {
			attributeMap["default_value"] = strconv.FormatBool(any(attribute.DefaultValue).(bool))
		}
	case "TIMESTAMP", "DATE", "TIME":
		if any(attribute.DefaultValue) != nil {
			attributeMap["default_value"] = attribute.DefaultValue
		}
		attributeMap["before"] = attribute.Before
		attributeMap["after"] = attribute.After
	case "ENUM":
		if any(attribute.DefaultValue) != nil {
			attributeMap["default_value"] = formatFloat(any(attribute.DefaultValue).(float64))
		}
		if attribute.Options != nil {
			attributeMap["options"] = *attribute.Options
		}
	case "GEOPOINT":
		if attribute.Fields != nil {
			fieldList := make([]map[string]any, 1)
			fieldMap := make(map[string]any)
			elevationList := make([]map[string]any, 1)
			elevationList[0] = (&attribute.Fields.Elevation).ToMap()
			fieldMap["elevation"] = elevationList
			latitudeList := make([]map[string]any, 1)
			latitudeList[0] = (&attribute.Fields.Latitude).ToMap()
			fieldMap["latitude"] = latitudeList
			longitudeList := make([]map[string]any, 1)
			longitudeList[0] = (&attribute.Fields.Longitude).ToMap()
			fieldMap["longitude"] = longitudeList
			fieldList[0] = fieldMap
			attributeMap["fields"] = fieldList
		}
	case "ARRAY":
		attributeMap["min_size"] = attribute.MinSize
		attributeMap["max_size"] = attribute.MaxSize
		attributeMap["unique"] = attribute.Unique
		attributeMap["constraint"] = []any{attribute.Constraint.ToMap()}
		if any(attribute.DefaultValue) != nil {
			var defaultValue string
			var interfaceArrayDefaultValues []interface{} = any(attribute.DefaultValue).([]interface{})
			for _, value := range interfaceArrayDefaultValues {
				defaultValue = defaultValue + "," + fmt.Sprint(value)
			}
			attributeMap["default_value"] = strings.TrimPrefix(defaultValue, ",")
		}
	}
	return attributeMap
}

func (constraint *ArrayConstraint[T]) ToMap() map[string]any {
	constraintMap := make(map[string]any)

	constraintMap["type"] = constraint.Type

	if constraint.Required != nil {
		constraintMap["required"] = constraint.Required
	}

	switch constraint.Type {
	case "TEXT":
		constraintMap["min_length"] = constraint.MinLength
		constraintMap["max_length"] = constraint.MaxLength
		constraintMap["pattern"] = constraint.Pattern
	case "BINARY":
		constraintMap["min_length"] = constraint.MinLength
		constraintMap["max_length"] = constraint.MaxLength
	case "NUMERIC":
		constraintMap["min"] = constraint.Min
		constraintMap["max"] = constraint.Max
		constraintMap["scale"] = constraint.Scale
		constraintMap["precision"] = constraint.Precision
		constraintMap["unit_id"] = constraint.UnitId
	case "BOOLEAN":
		//nothing
	case "TIMESTAMP", "DATE", "TIME":
		constraintMap["before"] = constraint.Before
		constraintMap["after"] = constraint.After
	case "ENUM":
		if constraint.Options != nil {
			constraintMap["options"] = *constraint.Options
		}
	}
	return constraintMap
}

func (field *FieldDef[T]) ToMap() map[string]any {
	fieldMap := make(map[string]any)
	// This might be unsafe in the future - for now fields are only used for numbers
	// in the geopoint type so it's alright.
	if any(field.DefaultValue) != nil {
		fieldMap["default_value"] = formatFloat(any(field.DefaultValue).(float64))
	}
	fieldMap["min"] = field.Min
	fieldMap["max"] = field.Max
	fieldMap["scale"] = field.Scale
	fieldMap["precision"] = field.Precision
	fieldMap["unit_id"] = field.UnitId
	return fieldMap
}

func (field *FieldDef[T]) FromMap(fieldMap map[string]any) error {
	field.DefaultValue = fieldMap["default_value"].(T)
	field.Min = fieldMap["min"].(float64)
	field.Max = fieldMap["max"].(float64)
	field.Scale = fieldMap["scale"].(int)
	field.Precision = fieldMap["precision"].(int)
	field.UnitId = fieldMap["unit_id"].(string)
	return nil
}

func (field *Field[T]) ToMap() map[string]any {
	fieldMap := make(map[string]any)
	// This might be unsafe in the future - for now fields are only used for numbers
	// in the geopoint type so it's alright.
	if any(field.Value) != nil {
		fieldMap["value"] = formatFloat(any(field.Value).(float64))
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
