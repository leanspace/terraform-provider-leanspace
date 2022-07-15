package general_objects

import (
	"strconv"
)

// Parses a float to a string. Use this method to ensure consistency.
func ParseFloat(num float64) string {
	return strconv.FormatFloat(num, 'g', -1, 64)
}

func (paginatedList *PaginatedList[T, PT]) ToMap() map[string]any {
	paginatedListMap := make(map[string]any)
	content := make([]map[string]any, len(paginatedList.Content))
	for i, value := range paginatedList.Content {
		var pointer PT = &value
		content[i] = pointer.ToMap()
	}
	paginatedListMap["content"] = content
	paginatedListMap["total_elements"] = paginatedList.TotalElements
	paginatedListMap["total_pages"] = paginatedList.TotalPages
	paginatedListMap["number_of_elements"] = paginatedList.NumberOfElements
	paginatedListMap["number"] = paginatedList.Number
	paginatedListMap["size"] = paginatedList.Size
	sort := make([]map[string]any, len(paginatedList.Sort))
	for i, value := range paginatedList.Sort {
		sort[i] = value.ToMap()
	}
	paginatedListMap["sort"] = sort
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
	sorts := make([]map[string]any, len(pageable.Sort))
	for i, sort := range pageable.Sort {
		sorts[i] = sort.ToMap()
	}
	pageableMap["sort"] = sorts
	pageableMap["offset"] = pageable.Offset
	pageableMap["page_number"] = pageable.PageNumber
	pageableMap["page_size"] = pageable.PageSize
	pageableMap["paged"] = pageable.Paged
	pageableMap["unpaged"] = pageable.Unpaged
	return pageableMap
}

func (paginatedList *PaginatedList[T, PT]) FromMap(paginatedListMap map[string]any) error {
	paginatedList.Content = make([]T, len(paginatedListMap["content"].([]any)))
	for i, value := range paginatedListMap["content"].([]any) {
		var pointer PT = &paginatedList.Content[i]
		err := pointer.FromMap(value.(map[string]any))
		if err != nil {
			return err
		}
	}
	paginatedList.TotalElements = paginatedListMap["total_elements"].(int)
	paginatedList.TotalPages = paginatedListMap["total_pages"].(int)
	paginatedList.NumberOfElements = paginatedListMap["number_of_elements"].(int)
	paginatedList.Number = paginatedListMap["number"].(int)
	paginatedList.Size = paginatedListMap["size"].(int)
	paginatedList.Sort = make([]Sort, len(paginatedListMap["sort"].([]any)))
	for i, value := range paginatedListMap["sort"].([]any) {
		paginatedList.Sort[i].FromMap(value.(map[string]any))
	}
	paginatedList.First = paginatedListMap["first"].(bool)
	paginatedList.Last = paginatedListMap["last"].(bool)
	paginatedList.Empty = paginatedListMap["empty"].(bool)
	paginatedList.Pageable.FromMap(paginatedListMap["pageable"].([]any)[0].(map[string]any))
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
	sorts := make([]Sort, len(pageableMap["sorts"].([]any)))
	for i, sort := range pageableMap["sorts"].([]any) {
		err := sorts[i].FromMap(sort.(map[string]any))
		if err != nil {
			return err
		}
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
	if value, ok := attributeMap["default_value"]; ok {
		attribute.DefaultValue = value.(T)
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
			attributeMap["default_value"] = ParseFloat(any(attribute.DefaultValue).(float64))
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
			attributeMap["default_value"] = ParseFloat(any(attribute.DefaultValue).(float64))
		}
		if attribute.Options != nil {
			attributeMap["options"] = *attribute.Options
		}
	}
	return attributeMap
}
