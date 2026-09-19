package general_objects

import (
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
