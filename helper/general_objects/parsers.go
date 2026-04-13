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
	paginatedListMap["total_elements"] = helper.NilIfEmpty(paginatedList.TotalElements)
	paginatedListMap["total_pages"] = helper.NilIfEmpty(paginatedList.TotalPages)
	paginatedListMap["number_of_elements"] = helper.NilIfEmpty(paginatedList.NumberOfElements)
	paginatedListMap["number"] = helper.NilIfEmpty(paginatedList.Number)
	paginatedListMap["size"] = helper.NilIfEmpty(paginatedList.Size)
	paginatedListMap["sort"] = helper.ParseToMaps(paginatedList.Sort)
	paginatedListMap["first"] = helper.NilIfEmpty(paginatedList.First)
	paginatedListMap["last"] = helper.NilIfEmpty(paginatedList.Last)
	paginatedListMap["empty"] = helper.NilIfEmpty(paginatedList.Empty)
	paginatedListMap["pageable"] = []any{paginatedList.Pageable.ToMap()}
	return paginatedListMap
}

func (keyValue *KeyValue) ToMap() map[string]any {
	keyValueMap := make(map[string]any)
	keyValueMap["key"] = helper.NilIfEmpty(keyValue.Key)
	keyValueMap["value"] = helper.NilIfEmpty(keyValue.Value)
	return keyValueMap
}

func (sort *Sort) ToMap() map[string]any {
	sortMap := make(map[string]any)
	sortMap["direction"] = helper.NilIfEmpty(sort.Direction)
	sortMap["property"] = helper.NilIfEmpty(sort.Property)
	sortMap["ignore_case"] = helper.NilIfEmpty(sort.IgnoreCase)
	sortMap["null_handling"] = helper.NilIfEmpty(sort.NullHandling)
	sortMap["ascending"] = helper.NilIfEmpty(sort.Ascending)
	sortMap["descending"] = helper.NilIfEmpty(sort.Descending)
	return sortMap
}

func (pageable *Pageable) ToMap() map[string]any {
	pageableMap := make(map[string]any)
	pageableMap["sort"] = helper.ParseToMaps(pageable.Sort)
	pageableMap["offset"] = helper.NilIfEmpty(pageable.Offset)
	pageableMap["page_number"] = helper.NilIfEmpty(pageable.PageNumber)
	pageableMap["page_size"] = helper.NilIfEmpty(pageable.PageSize)
	pageableMap["paged"] = helper.NilIfEmpty(pageable.Paged)
	pageableMap["unpaged"] = helper.NilIfEmpty(pageable.Unpaged)
	return pageableMap
}

func (attribute *ValueAttribute[T]) ToMap() map[string]any {
	attributeMap := make(map[string]any)
	attributeMap["type"] = helper.NilIfEmpty(attribute.Type)
	attributeMap["data_type"] = helper.NilIfEmpty(attribute.DataType)
	switch attribute.Type {
	case "NUMERIC":
		attributeMap["value"] = helper.ParseFloat(any(attribute.Value).(float64))
		attributeMap["unit_id"] = helper.NilIfEmpty(attribute.UnitId)
	case "TEXT", "TIMESTAMP", "DATE", "TIME", "BINARY":
		attributeMap["value"] = helper.NilIfEmpty(attribute.Value)
	case "BOOLEAN":
		attributeMap["value"] = strconv.FormatBool(any(attribute.Value).(bool))
	case "GEOPOINT":
		if attribute.Fields != nil {
			fieldMap := make(map[string]any)
			elevationMap := (&attribute.Fields.Elevation).ToMap()
			fieldMap["elevation"] = elevationMap
			latitudeMap := (&attribute.Fields.Latitude).ToMap()
			fieldMap["latitude"] = latitudeMap
			longitudeMap := (&attribute.Fields.Longitude).ToMap()
			fieldMap["longitude"] = longitudeMap
			attributeMap["fields"] = fieldMap
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
	if content, err := helper.ParseFromMaps[T, PT](helper.CastSlice(paginatedListMap, "content")); err != nil {
		return err
	} else {
		paginatedList.Content = content
	}
	paginatedList.TotalElements = helper.CastInt(paginatedListMap, "total_elements")
	paginatedList.TotalPages = helper.CastInt(paginatedListMap, "total_pages")
	paginatedList.NumberOfElements = helper.CastInt(paginatedListMap, "number_of_elements")
	paginatedList.Number = helper.CastInt(paginatedListMap, "number")
	paginatedList.Size = helper.CastInt(paginatedListMap, "size")
	if sort, err := helper.ParseFromMaps[Sort](helper.CastSlice(paginatedListMap, "sort")); err != nil {
		return err
	} else {
		paginatedList.Sort = sort
	}
	paginatedList.First = helper.CastBool(paginatedListMap, "first")
	paginatedList.Last = helper.CastBool(paginatedListMap, "last")
	paginatedList.Empty = helper.CastBool(paginatedListMap, "empty")
	if err := paginatedList.Pageable.FromMap(helper.CastMapAny(paginatedListMap, "pageable")); err != nil {
		return err
	}
	return nil
}

func (keyValue *KeyValue) FromMap(keyValueMap map[string]any) error {
	keyValue.Key = helper.CastString(keyValueMap, "key")
	keyValue.Value = helper.CastString(keyValueMap, "value")
	return nil
}

func (sort *Sort) FromMap(sortMap map[string]any) error {
	sort.Direction = helper.CastString(sortMap, "direction")
	sort.Property = helper.CastString(sortMap, "property")
	sort.NullHandling = helper.CastString(sortMap, "null_handling")
	sort.IgnoreCase = helper.CastBool(sortMap, "ignore_case")
	sort.Ascending = helper.CastBool(sortMap, "ascending")
	sort.Descending = helper.CastBool(sortMap, "descending")
	return nil
}

func (pageable *Pageable) FromMap(pageableMap map[string]any) error {
	if sorts, err := helper.ParseFromMaps[Sort](helper.CastSlice(pageableMap, "sorts")); err != nil {
		return err
	} else {
		pageableMap["sorts"] = sorts
	}
	pageable.Offset = helper.CastInt(pageableMap, "offset")
	pageable.PageNumber = helper.CastInt(pageableMap, "page_number")
	pageable.PageSize = helper.CastInt(pageableMap, "page_size")
	pageable.Paged = helper.CastBool(pageableMap, "paged")
	pageable.Unpaged = helper.CastBool(pageableMap, "unpaged")
	return nil
}

func (attribute *DefinitionAttribute[T]) FromMap(attributeMap map[string]any) error {
	attribute.Type = helper.CastString(attributeMap, "type")
	attribute.Required = helper.CastBoolPtr(attributeMap, "required")
	switch attribute.Type {
	case "NUMERIC":
		attribute.Min = helper.CastFloat64Ptr(attributeMap, "min")
		attribute.Max = helper.CastFloat64Ptr(attributeMap, "max")
		attribute.Scale = helper.CastIntPtr(attributeMap, "scale")
		attribute.Precision = helper.CastIntPtr(attributeMap, "precision")
		attribute.UnitId = helper.CastString(attributeMap, "unit_id")
	case "ENUM":
		if attributeMap["options"] != nil {
			option := helper.CastMapAny(attributeMap, "options")
			attribute.Options = &option
		}
	case "TEXT":
		attribute.MinLength = helper.CastIntPtr(attributeMap, "min_length")
		attribute.MaxLength = helper.CastIntPtr(attributeMap, "max_length")
		attribute.Pattern = helper.CastString(attributeMap, "pattern")
	case "BINARY":
		attribute.MinLength = helper.CastIntPtr(attributeMap, "min_length")
		attribute.MaxLength = helper.CastIntPtr(attributeMap, "max_length")
	case "TIMESTAMP", "DATE", "TIME":
		attribute.Before = helper.CastString(attributeMap, "before")
		attribute.After = helper.CastString(attributeMap, "after")
	case "BOOLEAN":
		// no extra field
	case "GEOPOINT":
		if attributeMap["fields"] != nil {
			attribute.Fields = &FieldsDef{}
			fields := helper.CastMapAny(attributeMap, "fields")
			attribute.Fields.Elevation.FromMap(helper.CastMapAny(fields, "elevation"))
			attribute.Fields.Latitude.FromMap(helper.CastMapAny(fields, "latitude"))
			attribute.Fields.Longitude.FromMap(helper.CastMapAny(fields, "longitude"))
		}
	case "ARRAY":
		attribute.MinSize = helper.CastIntPtr(attributeMap, "min_size")
		attribute.MaxSize = helper.CastIntPtr(attributeMap, "max_size")
		attribute.Unique = helper.CastBool(attributeMap, "unique")
		err := attribute.Constraint.FromMap(helper.CastMapAny(attributeMap, "constraint"))
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
			if defaultValue != nil {
				attribute.DefaultValue = defaultValue.(T)
			}
		}
	}
	return nil
}

func (constraint *ArrayConstraint[T]) FromMap(constraintMap map[string]any) error {
	constraint.Type = helper.CastString(constraintMap, "type")
	constraint.Required = helper.CastBoolPtr(constraintMap, "required")
	switch constraint.Type {
	case "NUMERIC":
		constraint.Min = helper.CastFloat64Ptr(constraintMap, "min")
		constraint.Max = helper.CastFloat64Ptr(constraintMap, "max")
		constraint.Scale = helper.CastIntPtr(constraintMap, "scale")
		constraint.Precision = helper.CastIntPtr(constraintMap, "precision")
		constraint.UnitId = helper.CastString(constraintMap, "unit_id")
	case "ENUM":
		if constraintMap["options"] != nil {
			option := helper.CastMapAny(constraintMap, "options")
			constraint.Options = &option
		}
	case "TEXT":
		constraint.MinLength = helper.CastIntPtr(constraintMap, "min_length")
		constraint.MaxLength = helper.CastIntPtr(constraintMap, "max_length")
		constraint.Pattern = helper.CastString(constraintMap, "pattern")
	case "TIMESTAMP", "DATE", "TIME":
		constraint.Before = helper.CastString(constraintMap, "before")
		constraint.After = helper.CastString(constraintMap, "after")
	case "BOOLEAN":
		// no extra field
	case "BINARY":
		constraint.MinLength = helper.CastIntPtr(constraintMap, "min_length")
		constraint.MaxLength = helper.CastIntPtr(constraintMap, "max_length")
	}
	return nil
}

func (attribute *ValueAttribute[T]) FromMap(attributeMap map[string]any) error {
	if attributeMap["value"] != nil {
		attribute.Value = attributeMap["value"].(T)
	}
	attribute.Type = helper.CastString(attributeMap, "type")
	attribute.DataType = helper.CastString(attributeMap, "data_type")
	if attributeMap["type"] == "NUMERIC" {
		attribute.UnitId = helper.CastString(attributeMap, "unit_id")
	}
	if attributeMap["type"] == "ARRAY" {
		var stringValues []string = strings.Split(helper.CastString(attributeMap, "value"), ",")
		var interfaceOfValues []interface{}
		for _, str := range stringValues {
			var stringValue = strings.TrimSpace(str)
			interfaceOfValues = append(interfaceOfValues, stringValue)

		}
		attribute.Value = any(interfaceOfValues).(T)
	}
	if attributeMap["type"] == "GEOPOINT" {
		if attributeMap["fields"] != nil {
			fields := helper.CastMapAny(attributeMap, "fields")
			attribute.Fields = &Fields{}
			attribute.Fields.Elevation.FromMap(helper.CastMapAny(fields, "elevation"))
			attribute.Fields.Latitude.FromMap(helper.CastMapAny(fields, "latitude"))
			attribute.Fields.Longitude.FromMap(helper.CastMapAny(fields, "longitude"))
		}
	}
	return nil
}

func (attribute *DefinitionAttribute[T]) ToMap() map[string]any {
	attributeMap := make(map[string]any)

	attributeMap["type"] = helper.NilIfEmpty(attribute.Type)

	attributeMap["required"] = helper.BoolPtrToAny(attribute.Required)

	switch attribute.Type {
	case "TEXT":
		if any(attribute.DefaultValue) != nil {
			attributeMap["default_value"] = helper.NilIfEmpty(attribute.DefaultValue)
		}
		attributeMap["min_length"] = helper.IntPtrToAny(attribute.MinLength)
		attributeMap["max_length"] = helper.IntPtrToAny(attribute.MaxLength)
		attributeMap["pattern"] = helper.NilIfEmpty(attribute.Pattern)
	case "BINARY":
		if any(attribute.DefaultValue) != nil {
			attributeMap["default_value"] = helper.NilIfEmpty(attribute.DefaultValue)
		}
		attributeMap["min_length"] = helper.IntPtrToAny(attribute.MinLength)
		attributeMap["max_length"] = helper.IntPtrToAny(attribute.MaxLength)
	case "NUMERIC":
		if any(attribute.DefaultValue) != nil {
			attributeMap["default_value"] = helper.ParseFloat(any(attribute.DefaultValue).(float64))
		}
		attributeMap["min"] = helper.Float64PtrToAny(attribute.Min)
		attributeMap["max"] = helper.Float64PtrToAny(attribute.Max)
		attributeMap["scale"] = helper.IntPtrToAny(attribute.Scale)
		attributeMap["precision"] = helper.IntPtrToAny(attribute.Precision)
		attributeMap["unit_id"] = helper.NilIfEmpty(attribute.UnitId)
	case "BOOLEAN":
		if any(attribute.DefaultValue) != nil {
			attributeMap["default_value"] = strconv.FormatBool(any(attribute.DefaultValue).(bool))
		}
	case "TIMESTAMP", "DATE", "TIME":
		if any(attribute.DefaultValue) != nil {
			attributeMap["default_value"] = helper.NilIfEmpty(attribute.DefaultValue)
		}
		attributeMap["before"] = helper.NilIfEmpty(attribute.Before)
		attributeMap["after"] = helper.NilIfEmpty(attribute.After)
	case "ENUM":
		if any(attribute.DefaultValue) != nil {
			attributeMap["default_value"] = helper.ParseFloat(any(attribute.DefaultValue).(float64))
		}
		if attribute.Options != nil {
			attributeMap["options"] = *attribute.Options
		}
	case "GEOPOINT":
		if attribute.Fields != nil {
			fieldMap := make(map[string]any)
			elevationMap := (&attribute.Fields.Elevation).ToMap()
			fieldMap["elevation"] = elevationMap
			latitudeMap := (&attribute.Fields.Latitude).ToMap()
			fieldMap["latitude"] = latitudeMap
			longitudeMap := (&attribute.Fields.Longitude).ToMap()
			fieldMap["longitude"] = longitudeMap
			attributeMap["fields"] = fieldMap
		}
	case "ARRAY":
		attributeMap["min_size"] = helper.IntPtrToAny(attribute.MinSize)
		attributeMap["max_size"] = helper.IntPtrToAny(attribute.MaxSize)
		attributeMap["unique"] = helper.NilIfEmpty(attribute.Unique)
		attributeMap["constraint"] = attribute.Constraint.ToMap()
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

	constraintMap["type"] = helper.NilIfEmpty(constraint.Type)

	constraintMap["required"] = helper.BoolPtrToAny(constraint.Required)

	switch constraint.Type {
	case "TEXT":
		constraintMap["min_length"] = helper.IntPtrToAny(constraint.MinLength)
		constraintMap["max_length"] = helper.IntPtrToAny(constraint.MaxLength)
		constraintMap["pattern"] = helper.NilIfEmpty(constraint.Pattern)
	case "BINARY":
		constraintMap["min_length"] = helper.IntPtrToAny(constraint.MinLength)
		constraintMap["max_length"] = helper.IntPtrToAny(constraint.MaxLength)
	case "NUMERIC":
		constraintMap["min"] = helper.Float64PtrToAny(constraint.Min)
		constraintMap["max"] = helper.Float64PtrToAny(constraint.Max)
		constraintMap["scale"] = helper.IntPtrToAny(constraint.Scale)
		constraintMap["precision"] = helper.IntPtrToAny(constraint.Precision)
		constraintMap["unit_id"] = helper.NilIfEmpty(constraint.UnitId)
	case "BOOLEAN":
		//nothing
	case "TIMESTAMP", "DATE", "TIME":
		constraintMap["before"] = helper.NilIfEmpty(constraint.Before)
		constraintMap["after"] = helper.NilIfEmpty(constraint.After)
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
	fieldMap["min"] = helper.Float64PtrToAny(field.Min)
	fieldMap["max"] = helper.Float64PtrToAny(field.Max)
	fieldMap["scale"] = helper.IntPtrToAny(field.Scale)
	fieldMap["precision"] = helper.IntPtrToAny(field.Precision)
	fieldMap["unit_id"] = helper.NilIfEmpty(field.UnitId)
	return fieldMap
}

func (field *FieldDef[T]) FromMap(fieldMap map[string]any) error {
	field.DefaultValue = fieldMap["default_value"].(T)
	field.Min = helper.CastFloat64Ptr(fieldMap, "min")
	field.Max = helper.CastFloat64Ptr(fieldMap, "max")
	field.Scale = helper.CastIntPtr(fieldMap, "scale")
	field.Precision = helper.CastIntPtr(fieldMap, "precision")
	field.UnitId = helper.CastString(fieldMap, "unit_id")
	return nil
}

func (field *Field[T]) ToMap() map[string]any {
	fieldMap := make(map[string]any)
	// This might be unsafe in the future - for now fields are only used for numbers
	// in the geopoint type so it's alright.
	if any(field.Value) != nil {
		fieldMap["value"] = helper.ParseFloat(any(field.Value).(float64))
	}
	fieldMap["min"] = helper.Float64PtrToAny(field.Min)
	fieldMap["max"] = helper.Float64PtrToAny(field.Max)
	fieldMap["scale"] = helper.IntPtrToAny(field.Scale)
	fieldMap["precision"] = helper.IntPtrToAny(field.Precision)
	fieldMap["unit_id"] = helper.NilIfEmpty(field.UnitId)
	return fieldMap
}

func (field *Field[T]) FromMap(fieldMap map[string]any) error {
	field.Value = fieldMap["value"].(T)
	field.Min = helper.CastFloat64Ptr(fieldMap, "min")
	field.Max = helper.CastFloat64Ptr(fieldMap, "max")
	field.Scale = helper.CastIntPtr(fieldMap, "scale")
	field.Precision = helper.CastIntPtr(fieldMap, "precision")
	field.UnitId = helper.CastString(fieldMap, "unit_id")
	return nil
}

func (a *AuditModel) ToAuditMap() map[string]any {
	m := make(map[string]any)
	m["id"] = helper.NilIfEmpty(a.ID)
	m["created_at"] = helper.NilIfEmpty(a.CreatedAt)
	m["created_by"] = helper.NilIfEmpty(a.CreatedBy)
	m["last_modified_at"] = helper.NilIfEmpty(a.LastModifiedAt)
	m["last_modified_by"] = helper.NilIfEmpty(a.LastModifiedBy)
	return m
}

func (a *AuditModel) FromAuditMap(m map[string]any) {
	a.ID = helper.CastString(m, "id")
	a.CreatedAt = helper.CastString(m, "created_at")
	a.CreatedBy = helper.CastString(m, "created_by")
	a.LastModifiedAt = helper.CastString(m, "last_modified_at")
	a.LastModifiedBy = helper.CastString(m, "last_modified_by")
}
