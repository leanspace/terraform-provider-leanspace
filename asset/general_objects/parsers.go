package general_objects

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

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
