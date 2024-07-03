package general_objects

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/leanspace/terraform-provider-leanspace/helper"
)

func (paginatedList *PaginatedList[T, PT]) ToMap() map[string]any {
	paginatedListMap := make(map[string]any)
	paginatedListMap["content"] = helper.ParseToMaps[T, PT](paginatedList.Content)
	paginatedListMap["total_elements"] = paginatedList.TotalElements
	paginatedListMap["total_pages"] = paginatedList.TotalPages
	paginatedListMap["number_of_elements"] = paginatedList.NumberOfElements
	paginatedListMap["number"] = paginatedList.Number
	paginatedListMap["size"] = paginatedList.Size
	paginatedListMap["sort"] = helper.ParseToMaps(paginatedList.Sort)
	paginatedListMap["first"] = paginatedList.First
	paginatedListMap["last"] = paginatedList.Last
	paginatedListMap["empty"] = paginatedList.Empty
	paginatedListMap["pageable"] = []any{paginatedList.Pageable.ToMap()}
	return paginatedListMap
}

func (keyValue *KeyValue) ToMap() map[string]any {
	keyValueMap := make(map[string]any)
	keyValueMap["key"] = keyValue.Key
	keyValueMap["value"] = keyValue.Value
	return keyValueMap
}

func (sort *Sort) ToMap() map[string]any {
	sortMap := make(map[string]any)
	sortMap["direction"] = sort.Direction
	sortMap["property"] = sort.Property
	sortMap["ignore_case"] = sort.IgnoreCase
	sortMap["null_handling"] = sort.NullHandling
	sortMap["ascending"] = sort.Ascending
	sortMap["descending"] = sort.Descending
	return sortMap
}

func (pageable *Pageable) ToMap() map[string]any {
	pageableMap := make(map[string]any)
	pageableMap["sort"] = helper.ParseToMaps(pageable.Sort)
	pageableMap["offset"] = pageable.Offset
	pageableMap["page_number"] = pageable.PageNumber
	pageableMap["page_size"] = pageable.PageSize
	pageableMap["paged"] = pageable.Paged
	pageableMap["unpaged"] = pageable.Unpaged
	return pageableMap
}

func (attribute *ValueAttribute[T]) ToMap() map[string]any {
	attributeMap := make(map[string]any)
	attributeMap["type"] = attribute.Type
	attributeMap["data_type"] = attribute.DataType
	switch attribute.Type {
	case "NUMERIC":
		attributeMap["value"] = helper.ParseFloat(any(attribute.Value).(float64))
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
			var defaultValue string
			var interfaceArrayValues []interface{} = any(attribute.Value).([]interface{})
			for _, value := range interfaceArrayValues {
				defaultValue = defaultValue + "," + fmt.Sprint(value)
			}
			attributeMap["value"] = strings.TrimPrefix(defaultValue, ",")
		}

	}
	return attributeMap
}

func (paginatedList *PaginatedList[T, PT]) FromMap(paginatedListMap map[string]any) error {
	if content, err := helper.ParseFromMaps[T, PT](paginatedListMap["content"].([]any)); err != nil {
		return err
	} else {
		paginatedList.Content = content
	}
	paginatedList.TotalElements = paginatedListMap["total_elements"].(int)
	paginatedList.TotalPages = paginatedListMap["total_pages"].(int)
	paginatedList.NumberOfElements = paginatedListMap["number_of_elements"].(int)
	paginatedList.Number = paginatedListMap["number"].(int)
	paginatedList.Size = paginatedListMap["size"].(int)
	if sort, err := helper.ParseFromMaps[Sort](paginatedListMap["sort"].([]any)); err != nil {
		return err
	} else {
		paginatedList.Sort = sort
	}
	paginatedList.First = paginatedListMap["first"].(bool)
	paginatedList.Last = paginatedListMap["last"].(bool)
	paginatedList.Empty = paginatedListMap["empty"].(bool)
	if len(paginatedListMap["pageable"].([]any)) > 0 {
		if err := paginatedList.Pageable.FromMap(paginatedListMap["pageable"].([]any)[0].(map[string]any)); err != nil {
			return err
		}
	}
	return nil
}

func (keyValue *KeyValue) FromMap(keyValueMap map[string]any) error {
	keyValue.Key = keyValueMap["key"].(string)
	keyValue.Value = keyValueMap["value"].(string)
	return nil
}

func (sort *Sort) FromMap(sortMap map[string]any) error {
	sort.Direction = sortMap["direction"].(string)
	sort.Property = sortMap["property"].(string)
	sort.NullHandling = sortMap["null_handling"].(string)
	sort.IgnoreCase = sortMap["ignore_case"].(bool)
	sort.Ascending = sortMap["ascending"].(bool)
	sort.Descending = sortMap["descending"].(bool)
	return nil
}

func (pageable *Pageable) FromMap(pageableMap map[string]any) error {
	if sorts, err := helper.ParseFromMaps[Sort](pageableMap["sorts"].([]any)); err != nil {
		return err
	} else {
		pageableMap["sorts"] = sorts
	}
	pageable.Offset = pageableMap["offset"].(int)
	pageable.PageNumber = pageableMap["page_number"].(int)
	pageable.PageSize = pageableMap["page_size"].(int)
	pageable.Paged = pageableMap["paged"].(bool)
	pageable.Unpaged = pageableMap["unpaged"].(bool)
	return nil
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
            fields := attributeMap["fields"].([]any)[0].(map[string]any)
            attribute.Fields = &FieldsDef{}
            attribute.Fields.Elevation.FromMap(fields["elevation"].([]any)[0].(map[string]any))
            attribute.Fields.Latitude.FromMap(fields["latitude"].([]any)[0].(map[string]any))
            attribute.Fields.Longitude.FromMap(fields["longitude"].([]any)[0].(map[string]any))
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
		} else {
			attribute.DefaultValue = defaultValue.(T)
		}
	}
	return nil
}

func (constraint *ArrayConstraint[T]) FromMap(constraintMap map[string]any) error {
	constraint.Type = constraintMap["type"].(string)
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

	attribute.Value = attributeMap["value"].(T)
	attribute.Type = attributeMap["type"].(string)
	attribute.DataType = attributeMap["data_type"].(string)
	if attributeMap["type"] == "NUMERIC" {
		attribute.UnitId = attributeMap["unit_id"].(string)
	}
	if attributeMap["type"] == "ARRAY" {
		var stringValues []string = strings.Split(attributeMap["value"].(string), ",")
		var interfaceOfValues []interface{}
		for _, str := range stringValues {
			var stringValue = strings.TrimSpace(str)
			interfaceOfValues = append(interfaceOfValues, stringValue)

		}
		attribute.Value = any(interfaceOfValues).(T)
	}
    if attributeMap["type"] == "GEOPOINT" {
        if attributeMap["fields"] != nil {
            fields := attributeMap["fields"].([]any)[0].(map[string]any)
            attribute.Fields = &Fields{}
            attribute.Fields.Elevation.FromMap(fields["elevation"].([]any)[0].(map[string]any))
            attribute.Fields.Latitude.FromMap(fields["latitude"].([]any)[0].(map[string]any))
            attribute.Fields.Longitude.FromMap(fields["longitude"].([]any)[0].(map[string]any))
        }
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
			attributeMap["default_value"] = helper.ParseFloat(any(attribute.DefaultValue).(float64))
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
			attributeMap["default_value"] = helper.ParseFloat(any(attribute.DefaultValue).(float64))
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
		fieldMap["default_value"] = helper.ParseFloat(any(field.DefaultValue).(float64))
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
