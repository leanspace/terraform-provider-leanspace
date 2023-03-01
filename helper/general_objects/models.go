package general_objects

import "github.com/leanspace/terraform-provider-leanspace/helper"

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

type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value,omitempty"`
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
type DefinitionAttribute[T any] struct {
	// Common
	Type         string `json:"type"`
	Required     *bool  `json:"required,omitempty"`
	DefaultValue T      `json:"defaultValue,omitempty"`
	// Text & Binary
	MinLength int `json:"minLength,omitempty"`
	MaxLength int `json:"maxLength,omitempty"`
	// Text
	Pattern string `json:"pattern,omitempty"`
	// Numeric
	Min       float64 `json:"min,omitempty"`
	Max       float64 `json:"max,omitempty"`
	Scale     int     `json:"scale,omitempty"`
	Precision int     `json:"precision,omitempty"`
	UnitId    string  `json:"unitId,omitempty"`
	// Date, time, timestamp
	Before string `json:"before,omitempty"`
	After  string `json:"after,omitempty"`
	// Enum
	Options *map[string]any `json:"options,omitempty"`
	// Array
	MinSize    int                  `json:"minSize,omitempty"`
	MaxSize    int                  `json:"maxSize,omitempty"`
	Unique     bool                 `json:"unique,omitempty"`
	Constraint ArrayConstraint[any] `json:"elementConstraint,omitempty"`
}

type ArrayConstraint[T any] struct {
	// Common
	Type         string `json:"type"`
	Required     *bool  `json:"required,omitempty"`
	DefaultValue T      `json:"defaultValue,omitempty"`
	// Text
	MinLength int    `json:"minLength,omitempty"`
	MaxLength int    `json:"maxLength,omitempty"`
	Pattern   string `json:"pattern,omitempty"`
	// Numeric
	Min       float64 `json:"min,omitempty"`
	Max       float64 `json:"max,omitempty"`
	Scale     int     `json:"scale,omitempty"`
	Precision int     `json:"precision,omitempty"`
	UnitId    string  `json:"unitId,omitempty"`
	// Date, time, timestamp
	Before string `json:"before,omitempty"`
	After  string `json:"after,omitempty"`
	// Enum
	Options *map[string]any `json:"options,omitempty"`
}

type ValueAttribute[T any] struct {
	Value T      `json:"value,omitempty"`
	Type  string `json:"type"`
	// Numeric
	UnitId string `json:"unitId,omitempty"`
}
