package general_objects

import (
	"github.com/leanspace/terraform-provider-leanspace/helper"
	generalobjectstypes "github.com/leanspace/terraform-provider-leanspace/helper/general_objects/types"
)

type KeyValue = generalobjectstypes.KeyValue
type DefinitionAttribute[T any] = generalobjectstypes.DefinitionAttribute[T]
type ArrayConstraint[T any] = generalobjectstypes.ArrayConstraint[T]
type ValueAttribute[T any] = generalobjectstypes.ValueAttribute[T]
type FieldsDef = generalobjectstypes.FieldsDef
type FieldDef[T any] = generalobjectstypes.FieldDef[T]
type Fields = generalobjectstypes.Fields
type Field[T any] = generalobjectstypes.Field[T]

type Sort struct {
	Direction    string `json:"direction"`
	Property     string `json:"property"`
	IgnoreCase   bool   `json:"ignoreCase"`
	NullHandling string `json:"nullHandling"`
	Ascending    bool   `json:"ascending"`
	Descending   bool   `json:"descending"`
}

type Pageable struct {
	Sort       []Sort `json:"sort"`
	Offset     int    `json:"offset"`
	PageNumber int    `json:"pageNumber"`
	PageSize   int    `json:"pageSize"`
	Paged      bool   `json:"paged"`
	Unpaged    bool   `json:"unpaged"`
}

type PaginatedList[T any, PT helper.ParseablePointer[T]] struct {
	Content          []T      `json:"content"`
	TotalElements    int      `json:"totalElements"`
	TotalPages       int      `json:"totalPages"`
	NumberOfElements int      `json:"numberOfElements"`
	Number           int      `json:"number"`
	Size             int      `json:"size"`
	Sort             []Sort   `json:"sort"`
	First            bool     `json:"first"`
	Last             bool     `json:"last"`
	Empty            bool     `json:"empty"`
	Pageable         Pageable `json:"pageable"`
}
