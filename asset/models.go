package asset

type Node struct {
	ID                      string   `json:"id"`
	Name                    string   `json:"name"`
	Description             string   `json:"description,omitempty"`
	CreatedAt               string   `json:"createdAt"`
	CreatedBy               string   `json:"createdBy"`
	LastModifiedAt          string   `json:"lastModifiedAt"`
	LastModifiedBy          string   `json:"lastModifiedBy"`
	ParentNodeId            string   `json:"parentNodeId,omitempty"`
	Tags                    []Tag    `json:"tags,omitempty"`
	Nodes                   []Node   `json:"nodes,omitempty"`
	Type                    string   `json:"type"`
	Kind                    string   `json:"kind,omitempty"`
	NoradId                 string   `json:"noradId,omitempty"`
	InternationalDesignator string   `json:"internationalDesignator,omitempty"`
	Tle                     []string `json:"tle,omitempty"`
}

type PaginatedList[T interface{}] struct {
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

type Tag struct {
	Key   string `json:"key"`
	Value string `json:"value,omitempty"`
}

type Property[T interface{}] struct {
	ID             string                  `json:"id"`
	Name           string                  `json:"name"`
	Description    string                  `json:"description,omitempty"`
	NodeId         string                  `json:"nodeId"`
	CreatedAt      string                  `json:"createdAt"`
	CreatedBy      string                  `json:"createdBy"`
	LastModifiedAt string                  `json:"lastModifiedAt"`
	LastModifiedBy string                  `json:"lastModifiedBy"`
	Tags           []Tag                   `json:"tags,omitempty"`
	MinLength      int                     `json:"minLength,omitempty"`
	MaxLength      int                     `json:"maxLength,omitempty"`
	Pattern        string                  `json:"pattern,omitempty"`
	Before         string                  `json:"before,omitempty"`
	After          string                  `json:"after,omitempty"`
	Fields         *Fields                 `json:"fields,omitempty"`
	Options        *map[string]interface{} `json:"options,omitempty"`
	Min            float64                 `json:"min,omitempty"`
	Max            float64                 `json:"max,omitempty"`
	Scale          int                     `json:"scale,omitempty"`
	Precision      int                     `json:"precision,omitempty"`
	UnitId         string                  `json:"unit_id,omitempty"`
	Value          T                       `json:"value,omitempty"`
	Type           string                  `json:"type"`
}

type Fields struct {
	Elevation Field[interface{}] `json:"elevation"`
	Latitude  Field[interface{}] `json:"latitude"`
	Longitude Field[interface{}] `json:"longitude"`
}

type Field[T interface{}] struct {
	Type           string  `json:"type"`
	ID             string  `json:"id"`
	Description    string  `json:"description,omitempty"`
	CreatedAt      string  `json:"createdAt"`
	CreatedBy      string  `json:"createdBy"`
	LastModifiedAt string  `json:"lastModifiedAt"`
	LastModifiedBy string  `json:"lastModifiedBy"`
	Value          T       `json:"value"`
	Tags           []Tag   `json:"tags,omitempty"`
	Name           string  `json:"name"`
	Min            float64 `json:"min,omitempty"`
	Max            float64 `json:"max,omitempty"`
	Scale          int     `json:"scale,omitempty"`
	Precision      int     `json:"precision,omitempty"`
	UnitId         string  `json:"unit_id,omitempty"`
}

type CommandDefinition struct {
	ID             string                  `json:"id"`
	NodeId         string                  `json:"nodeId"`
	Name           string                  `json:"name"`
	Description    string                  `json:"description,omitempty"`
	Identifier     string                  `json:"identifier,omitempty"`
	Metadata       []Metadata[interface{}] `json:"metadata,omitempty"`
	Arguments      []Argument[interface{}] `json:"arguments,omitempty"`
	CreatedAt      string                  `json:"createdAt"`
	CreatedBy      string                  `json:"createdBy"`
	LastModifiedAt string                  `json:"lastModifiedAt"`
	LastModifiedBy string                  `json:"lastModifiedBy"`
}

type Metadata[T interface{}] struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	UnitId      string `json:"unit_id,omitempty"`
	Value       T      `json:"value,omitempty"`
	Required    bool   `json:"required,omitempty"`
	Type        string `json:"type"`
}

type Argument[T interface{}] struct {
	ID           string                  `json:"id"`
	Name         string                  `json:"name"`
	Identifier   string                  `json:"identifier"`
	Description  string                  `json:"description,omitempty"`
	MinLength    int                     `json:"minLength,omitempty"`
	MaxLength    int                     `json:"maxLength,omitempty"`
	Pattern      string                  `json:"pattern,omitempty"`
	Before       string                  `json:"before,omitempty"`
	After        string                  `json:"after,omitempty"`
	Options      *map[string]interface{} `json:"options,omitempty"`
	Min          float64                 `json:"min,omitempty"`
	Max          float64                 `json:"max,omitempty"`
	Scale        int                     `json:"scale,omitempty"`
	Precision    int                     `json:"precision,omitempty"`
	UnitId       string                  `json:"unit_id,omitempty"`
	DefaultValue T                       `json:"defaultValue,omitempty"`
	Required     bool                    `json:"required,omitempty"`
	Type         string                  `json:"type"`
}
