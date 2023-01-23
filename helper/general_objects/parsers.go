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

func (tag *Tag) ToMap() map[string]any {
	tagMap := make(map[string]any)
	tagMap["key"] = tag.Key
	tagMap["value"] = tag.Value
	return tagMap
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
	switch attribute.Type {
	case "NUMERIC":
		attributeMap["value"] = helper.ParseFloat(any(attribute.Value).(float64))
		attributeMap["unit_id"] = attribute.UnitId
	case "TEXT", "TIMESTAMP", "DATE", "TIME":
		attributeMap["value"] = attribute.Value
	case "BOOLEAN":
		attributeMap["value"] = strconv.FormatBool(any(attribute.Value).(bool))
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

func (tag *Tag) FromMap(tagMap map[string]any) error {
	tag.Key = tagMap["key"].(string)
	tag.Value = tagMap["value"].(string)
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
	case "TIMESTAMP", "DATE", "TIME":
		attribute.Before = attributeMap["before"].(string)
		attribute.After = attributeMap["after"].(string)
	case "BOOLEAN":
		// no extra field
	case "ARRAY":
		attribute.MinSize = attributeMap["min_size"].(int)
		attribute.MaxSize = attributeMap["max_size"].(int)
		attribute.Unique = attributeMap["unique"].(bool)
		err := attribute.Constraint.FromMap(attributeMap["constraint"].([]any)[0].(map[string]any))
		if err != nil {
			return err
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
					if numericValue, err := strconv.ParseInt(stringValue, 10, 16); err == nil {
						interfaceOfDefaultValues = append(interfaceOfDefaultValues, numericValue)
					}
				case "BOOLEAN":
					if booleanValue, err := strconv.ParseBool(stringValue); err == nil {
						interfaceOfDefaultValues = append(interfaceOfDefaultValues, booleanValue)
					}
				case "TEXT", "TIMESTAMP", "DATE", "TIME":
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
	}
	return nil
}

func (attribute *ValueAttribute[T]) FromMap(attributeMap map[string]any) error {
	attribute.Value = attributeMap["value"].(T)
	attribute.Type = attributeMap["type"].(string)
	if attributeMap["type"] == "NUMERIC" {
		attribute.UnitId = attributeMap["unit_id"].(string)
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

func (attribute *PropertyAttribute[T]) ToMap() map[string]any {
	attributeMap := make(map[string]any)
	if attribute.AdditionalProperties != nil {
		attributeMap["additional_properties"] = any(attribute.AdditionalProperties)
	}
	attributeMap["type"] = attribute.Type
	switch attribute.Type {
	case "NUMERIC":
		if any(attribute.Value) != nil {
			attributeMap["value"] = helper.ParseFloat(any(attribute.Value).(float64))
		}
		attributeMap["min"] = attribute.Min
		attributeMap["max"] = attribute.Max
		attributeMap["scale"] = attribute.Scale
		attributeMap["precision"] = attribute.Precision
		attributeMap["unit_id"] = attribute.UnitId
	case "ENUM":
		if any(attribute.Value) != nil {
			attributeMap["value"] = helper.ParseFloat(any(attribute.Value).(float64))
		}
		if attribute.Options != nil {
			attributeMap["options"] = *attribute.Options
		}
	case "TEXT":
		if any(attribute.Value) != nil {
			attributeMap["value"] = attribute.Value
		}
		attributeMap["min_length"] = attribute.MinLength
		attributeMap["max_length"] = attribute.MaxLength
		attributeMap["pattern"] = attribute.Pattern
	case "TIMESTAMP", "DATE", "TIME":
		if any(attribute.Value) != nil {
			attributeMap["value"] = attribute.Value
		}
		attributeMap["before"] = attribute.Before
		attributeMap["after"] = attribute.After
	case "BOOLEAN":
		if any(attribute.Value) != nil {
			attributeMap["value"] = strconv.FormatBool(any(attribute.Value).(bool))
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
	case "TLE":
		if any(attribute.Value) != nil {
			var tleValue string
			var tleValues []interface{} = any(attribute.Value).([]interface{})
			for _, value := range tleValues {
				tleValue = tleValue + "," + fmt.Sprint(value)
			}
			attributeMap["value"] = strings.TrimPrefix(tleValue, ",")
		}
	}
	return attributeMap
}

func (attribute *PropertyAttribute[T]) FromMap(attributeMap map[string]any) error {
	attribute.AdditionalProperties = attributeMap["additional_properties"].(map[string]any)
	attribute.Type = attributeMap["type"].(string)
	switch attribute.Type {
	case "NUMERIC":
		if attributeMap["value"] != nil {
			attribute.Value = any(attributeMap["value"]).(T)
		}
		attribute.Min = attributeMap["min"].(float64)
		attribute.Max = attributeMap["max"].(float64)
		attribute.Scale = attributeMap["scale"].(int)
		attribute.Precision = attributeMap["precision"].(int)
		attribute.UnitId = attributeMap["unit_id"].(string)
	case "ENUM":
		if attributeMap["value"] != nil {
			attribute.Value = any(attributeMap["value"]).(T)
		}
		if attributeMap["options"] != nil {
			option := attributeMap["options"].(map[string]any)
			attribute.Options = &option
		}
	case "TEXT":
		if attributeMap["value"] != nil {
			attribute.Value = any(attributeMap["value"]).(T)
		}
		attribute.MinLength = attributeMap["min_length"].(int)
		attribute.MaxLength = attributeMap["max_length"].(int)
		attribute.Pattern = attributeMap["pattern"].(string)
	case "TIMESTAMP", "DATE", "TIME":
		if attributeMap["value"] != nil {
			attribute.Value = any(attributeMap["value"]).(T)
		}
		attribute.Before = attributeMap["before"].(string)
		attribute.After = attributeMap["after"].(string)
	case "BOOLEAN", "STRUCTURE":
		if attributeMap["value"] != nil {
			attribute.Value = any(attributeMap["value"]).(T)
		}
	case "GEOPOINT":
		if attributeMap["value"] != nil && len(any(attributeMap["value"]).(string)) != 0 {
			attribute.Value = any(attributeMap["value"]).(T)
		}
		if attributeMap["fields"] != nil {
			fields := attributeMap["fields"].([]any)[0].(map[string]any)
			attribute.Fields = &Fields{}
			attribute.Fields.Elevation.FromMap(fields["elevation"].([]any)[0].(map[string]any))
			attribute.Fields.Latitude.FromMap(fields["latitude"].([]any)[0].(map[string]any))
			attribute.Fields.Longitude.FromMap(fields["longitude"].([]any)[0].(map[string]any))
		}
	case "TLE":
		if tleValue, ok := attributeMap["value"]; ok {
			var stringTleValues []string = strings.Split(tleValue.(string), ",")
			var interfaceOfTleValues []interface{}
			for _, str := range stringTleValues {
				var stringValue = strings.TrimSpace(str)
				interfaceOfTleValues = append(interfaceOfTleValues, stringValue)
			}
			attribute.Value = any(interfaceOfTleValues).(T)
		}
	}

	return nil
}

func (field *Field[T]) ToMap() map[string]any {
	fieldMap := make(map[string]any)
	if field.AdditionalProperties != nil {
		fieldMap["additional_properties"] = any(field.AdditionalProperties)
	}
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
	if fieldMap["additional_properties"] != nil {
		//property := fieldMap["additional_properties"].(map[string]any)
		field.AdditionalProperties = fieldMap["additional_properties"].(map[string]any)
	}
	field.Value = fieldMap["value"].(T)
	field.Min = fieldMap["min"].(float64)
	field.Max = fieldMap["max"].(float64)
	field.Scale = fieldMap["scale"].(int)
	field.Precision = fieldMap["precision"].(int)
	field.UnitId = fieldMap["unit_id"].(string)
	return nil
}
