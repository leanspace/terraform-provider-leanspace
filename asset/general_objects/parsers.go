package general_objects

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Parses a float to a string. Use this method to ensure consistency.
func ParseFloat(num float64) string {
	return strconv.FormatFloat(num, 'g', -1, 64)
}

func TagsInterfaceToStruct(tags any) []Tag {
	tagsStruct := []Tag{}
	if tags != nil {
		for _, tag := range tags.(*schema.Set).List() {
			newTag := Tag{Key: tag.(map[string]any)["key"].(string), Value: tag.(map[string]any)["value"].(string)}
			tagsStruct = append(tagsStruct, newTag)
		}
	}
	return tagsStruct
}

func TagsStructToMap(tags []Tag) []map[string]any {
	if tags != nil {
		tagsList := make([]map[string]any, len(tags))
		for i, tag := range tags {
			tagMap := make(map[string]any)
			tagMap["key"] = tag.Key
			tagMap["value"] = tag.Value
			tagsList[i] = tagMap
		}

		return tagsList
	}
	return make([]map[string]any, 0)
}

func SortStructToMap(sort []Sort) []map[string]any {
	sortList := make([]map[string]any, len(sort))
	for i, sortItem := range sort {
		sortMap := make(map[string]any)
		sortMap["direction"] = sortItem.Direction
		sortMap["property"] = sortItem.Property
		sortMap["ignore_case"] = sortItem.IgnoreCase
		sortMap["null_handling"] = sortItem.NullHandling
		sortMap["ascending"] = sortItem.Ascending
		sortMap["descending"] = sortItem.Descending
		sortList[i] = sortMap
	}

	return sortList
}

func PageableStructToMapList(pageable Pageable, sort any) []map[string]any {
	pageableList := make([]map[string]any, 1)
	pageableMap := make(map[string]any)
	pageableMap["sort"] = sort
	pageableMap["offset"] = pageable.Offset
	pageableMap["page_number"] = pageable.PageNumber
	pageableMap["page_size"] = pageable.PageSize
	pageableMap["paged"] = pageable.Paged
	pageableMap["unpaged"] = pageable.Unpaged
	pageableList[0] = pageableMap

	return pageableList
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
