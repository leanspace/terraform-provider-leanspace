package asset

type Node struct {
	ID                      string   `json:"id" terra:"id"`
	Name                    string   `json:"name" terra:"name"`
	Description             string   `json:"description,omitempty" terra:"description"`
	CreatedAt               string   `json:"createdAt" terra:"created_at"`
	CreatedBy               string   `json:"createdBy" terra:"created_by"`
	LastModifiedAt          string   `json:"lastModifiedAt" terra:"last_modified_at"`
	LastModifiedBy          string   `json:"lastModifiedBy" terra:"last_modified_by"`
	ParentNodeId            string   `json:"parentNodeId,omitempty" terra:"parent_node_id"`
	Tags                    []Tag    `json:"tags,omitempty" terra:"tags"`
	Nodes                   []Node   `json:"nodes,omitempty" terra:"nodes"`
	Type                    string   `json:"type" terra:"type"`
	Kind                    string   `json:"kind,omitempty" terra:"kind"`
	NoradId                 string   `json:"noradId,omitempty" terra:"norad_id"`
	InternationalDesignator string   `json:"internationalDesignator,omitempty" terra:"international_designator"`
	Tle                     []string `json:"tle,omitempty" terra:"tle"`
}

type PaginatedList[T interface{}] struct {
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

type Property[T interface{}] struct {
	ID             string                  `json:"id" terra:"id"`
	Name           string                  `json:"name" terra:"name"`
	Description    string                  `json:"description,omitempty" terra:"description"`
	NodeId         string                  `json:"nodeId" terra:"node_id"`
	CreatedAt      string                  `json:"createdAt" terra:"created_at"`
	CreatedBy      string                  `json:"createdBy" terra:"created_by"`
	LastModifiedAt string                  `json:"lastModifiedAt" terra:"last_modified_at"`
	LastModifiedBy string                  `json:"lastModifiedBy" terra:"last_modified_by"`
	Tags           []Tag                   `json:"tags,omitempty" terra:"tags"`
	MinLength      int                     `json:"minLength,omitempty" terra:"min_length"`
	MaxLength      int                     `json:"maxLength,omitempty" terra:"max_length"`
	Pattern        string                  `json:"pattern,omitempty" terra:"pattern"`
	Before         string                  `json:"before,omitempty" terra:"before"`
	After          string                  `json:"after,omitempty" terra:"after"`
	Fields         *Fields                 `json:"fields,omitempty" terra:"fields"`
	Options        *map[string]interface{} `json:"options,omitempty" terra:"options"`
	Min            float64                 `json:"min,omitempty" terra:"min"`
	Max            float64                 `json:"max,omitempty" terra:"max"`
	Scale          int                     `json:"scale,omitempty" terra:"scale"`
	Precision      int                     `json:"precision,omitempty" terra:"precision"`
	UnitId         string                  `json:"unit_id,omitempty" terra:"unit_id"`
	Value          T                       `json:"value,omitempty" terra:"value"`
	Type           string                  `json:"type" terra:"type"`
}

type Fields struct {
	Elevation Field[interface{}] `json:"elevation" terra:"elevation"`
	Latitude  Field[interface{}] `json:"latitude" terra:"latitude"`
	Longitude Field[interface{}] `json:"longitude" terra:"longitude"`
}

type Field[T interface{}] struct {
	Type           string  `json:"type" terra:"type"`
	ID             string  `json:"id" terra:"id"`
	Description    string  `json:"description,omitempty" terra:"description"`
	CreatedAt      string  `json:"createdAt" terra:"created_at"`
	CreatedBy      string  `json:"createdBy" terra:"created_by"`
	LastModifiedAt string  `json:"lastModifiedAt" terra:"last_modified_at"`
	LastModifiedBy string  `json:"lastModifiedBy" terra:"last_modified_by"`
	Value          T       `json:"value" terra:"value"`
	Tags           []Tag   `json:"tags,omitempty" terra:"tags"`
	Name           string  `json:"name" terra:"name"`
	Min            float64 `json:"min,omitempty" terra:"min"`
	Max            float64 `json:"max,omitempty" terra:"max"`
	Scale          int     `json:"scale,omitempty" terra:"scale"`
	Precision      int     `json:"precision,omitempty" terra:"precision"`
	UnitId         string  `json:"unit_id,omitempty" terra:"unit_id"`
}

type CommandDefinition struct {
	ID             string                  `json:"id" terra:"id"`
	NodeId         string                  `json:"nodeId" terra:"node_id"`
	Name           string                  `json:"name" terra:"name"`
	Description    string                  `json:"description,omitempty" terra:"description"`
	Identifier     string                  `json:"identifier,omitempty" terra:"identifier"`
	Metadata       []Metadata[interface{}] `json:"metadata,omitempty" terra:"metadata"`
	Arguments      []Argument[interface{}] `json:"arguments,omitempty" terra:"arguments"`
	CreatedAt      string                  `json:"createdAt" terra:"created_at"`
	CreatedBy      string                  `json:"createdBy" terra:"created_by"`
	LastModifiedAt string                  `json:"lastModifiedAt" terra:"last_modified_at"`
	LastModifiedBy string                  `json:"lastModifiedBy" terra:"last_modified_by"`
}

type Metadata[T interface{}] struct {
	ID          string `json:"id" terra:"id"`
	Name        string `json:"name" terra:"name"`
	Description string `json:"description,omitempty" terra:"description"`
	UnitId      string `json:"unit_id,omitempty" terra:"unit_id"`
	Value       T      `json:"value,omitempty" terra:"value"`
	Required    bool   `json:"required,omitempty" terra:"required"`
	Type        string `json:"type" terra:"type"`
}

type Argument[T interface{}] struct {
	ID           string                  `json:"id" terra:"id"`
	Name         string                  `json:"name" terra:"name"`
	Identifier   string                  `json:"identifier" terra:"identifier"`
	Description  string                  `json:"description,omitempty" terra:"description"`
	MinLength    int                     `json:"minLength,omitempty" terra:"min_length"`
	MaxLength    int                     `json:"maxLength,omitempty" terra:"max_length"`
	Pattern      string                  `json:"pattern,omitempty" terra:"pattern"`
	Before       string                  `json:"before,omitempty" terra:"before"`
	After        string                  `json:"after,omitempty" terra:"after"`
	Options      *map[string]interface{} `json:"options,omitempty" terra:"options"`
	Min          float64                 `json:"min,omitempty" terra:"min"`
	Max          float64                 `json:"max,omitempty" terra:"max"`
	Scale        int                     `json:"scale,omitempty" terra:"scale"`
	Precision    int                     `json:"precision,omitempty" terra:"precision"`
	UnitId       string                  `json:"unit_id,omitempty" terra:"unit_id"`
	DefaultValue T                       `json:"defaultValue,omitempty" terra:"default_value"`
	Required     bool                    `json:"required,omitempty" terra:"required"`
	Type         string                  `json:"type" terra:"type"`
}
