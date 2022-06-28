package command_definitions

type CommandDefinition struct {
	ID             string          `json:"id" terra:"id"`
	NodeId         string          `json:"nodeId" terra:"node_id"`
	Name           string          `json:"name" terra:"name"`
	Description    string          `json:"description,omitempty" terra:"description"`
	Identifier     string          `json:"identifier,omitempty" terra:"identifier"`
	Metadata       []Metadata[any] `json:"metadata,omitempty" terra:"metadata"`
	Arguments      []Argument[any] `json:"arguments,omitempty" terra:"arguments"`
	CreatedAt      string          `json:"createdAt" terra:"created_at"`
	CreatedBy      string          `json:"createdBy" terra:"created_by"`
	LastModifiedAt string          `json:"lastModifiedAt" terra:"last_modified_at"`
	LastModifiedBy string          `json:"lastModifiedBy" terra:"last_modified_by"`
}

type Metadata[T any] struct {
	ID          string `json:"id" terra:"id"`
	Name        string `json:"name" terra:"name"`
	Description string `json:"description,omitempty" terra:"description"`
	UnitId      string `json:"unit_id,omitempty" terra:"unit_id"`
	Value       T      `json:"value,omitempty" terra:"value"`
	Required    bool   `json:"required,omitempty" terra:"required"`
	Type        string `json:"type" terra:"type"`
}

type Argument[T any] struct {
	ID           string          `json:"id" terra:"id"`
	Name         string          `json:"name" terra:"name"`
	Identifier   string          `json:"identifier" terra:"identifier"`
	Description  string          `json:"description,omitempty" terra:"description"`
	MinLength    int             `json:"minLength,omitempty" terra:"min_length"`
	MaxLength    int             `json:"maxLength,omitempty" terra:"max_length"`
	Pattern      string          `json:"pattern,omitempty" terra:"pattern"`
	Before       string          `json:"before,omitempty" terra:"before"`
	After        string          `json:"after,omitempty" terra:"after"`
	Options      *map[string]any `json:"options,omitempty" terra:"options"`
	Min          float64         `json:"min,omitempty" terra:"min"`
	Max          float64         `json:"max,omitempty" terra:"max"`
	Scale        int             `json:"scale,omitempty" terra:"scale"`
	Precision    int             `json:"precision,omitempty" terra:"precision"`
	UnitId       string          `json:"unit_id,omitempty" terra:"unit_id"`
	DefaultValue T               `json:"defaultValue,omitempty" terra:"default_value"`
	Required     bool            `json:"required,omitempty" terra:"required"`
	Type         string          `json:"type" terra:"type"`
}
