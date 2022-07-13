package general_objects

type Sort struct {
	Direction    string `json:"direction" terra:"direction"`
	Property     string `json:"property" terra:"property"`
	IgnoreCase   bool   `json:"ignoreCase" terra:"ignore_case"`
	NullHandling string `json:"nullHandling" terra:"null_handling"`
	Ascending    bool   `json:"ascending" terra:"ascending"`
	Descending   bool   `json:"descending" terra:"descending"`
}

type Pageable struct {
	Sort       []Sort `json:"sort" terra:"sort"`
	Offset     int    `json:"offset" terra:"offset"`
	PageNumber int    `json:"pageNumber" terra:"page_number"`
	PageSize   int    `json:"pageSize" terra:"page_size"`
	Paged      bool   `json:"paged" terra:"paged"`
	Unpaged    bool   `json:"unpaged" terra:"unpaged"`
}

type Tag struct {
	Key   string `json:"key" terra:"key"`
	Value string `json:"value,omitempty" terra:"value"`
}

type PaginatedList[T any] struct {
	Content          []T      `json:"content" terra:"content"`
	TotalElements    int      `json:"totalElements" terra:"total_elements"`
	TotalPages       int      `json:"totalPages" terra:"total_pages"`
	NumberOfElements int      `json:"numberOfElements" terra:"number_of_elements"`
	Number           int      `json:"number" terra:"number"`
	Size             int      `json:"size" terra:"size"`
	Sort             []Sort   `json:"sort" terra:"sort"`
	First            bool     `json:"first" terra:"first"`
	Last             bool     `json:"last" terra:"last"`
	Empty            bool     `json:"empty" terra:"empty"`
	Pageable         Pageable `json:"pageable" terra:"pageable"`
}

type DefinitionAttribute[T any] struct {
	// Common
	Type         string `json:"type" terra:"type"`
	Required     *bool  `json:"required,omitempty" terra:"required"`
	DefaultValue T      `json:"defaultValue,omitempty" terra:"default_value"`
	// Text
	MinLength int    `json:"minLength,omitempty" terra:"min_length"`
	MaxLength int    `json:"maxLength,omitempty" terra:"max_length"`
	Pattern   string `json:"pattern,omitempty" terra:"pattern"`
	// Numeric
	Min       float64 `json:"min,omitempty" terra:"min"`
	Max       float64 `json:"max,omitempty" terra:"max"`
	Scale     int     `json:"scale,omitempty" terra:"scale"`
	Precision int     `json:"precision,omitempty" terra:"precision"`
	UnitId    string  `json:"unitId,omitempty" terra:"unit_id"`
	// Date, time, timestamp
	Before string `json:"before,omitempty" terra:"before"`
	After  string `json:"after,omitempty" terra:"after"`
	// Enum
	Options *map[string]any `json:"options,omitempty" terra:"options"`
}
