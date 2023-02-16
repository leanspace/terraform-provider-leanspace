package properties

import "github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

type Property[T any] struct {
	ID             string                       `json:"id"`
	Name           string                       `json:"name"`
	Description    string                       `json:"description,omitempty"`
	IsBuiltIn      bool                         `json:"builtIn,omitempty"`
	NodeId         string                       `json:"nodeId"`
	CreatedAt      string                       `json:"createdAt"`
	CreatedBy      string                       `json:"createdBy"`
	LastModifiedAt string                       `json:"lastModifiedAt"`
	LastModifiedBy string                       `json:"lastModifiedBy"`
	Tags           []general_objects.KeyValue   `json:"tags,omitempty"`
	Attributes     PropertyAttribute[T]         `json:"attributes,omitempty"`
}

func (property *Property[T]) GetID() string { return property.ID }

type PropertyAttribute[T any] struct {
	// Common
	Value T      `json:"value,omitempty"`
	Type  string `json:"type"`

	// Geopoint only
	Fields *Fields `json:"fields,omitempty"`

	// Numeric only
	Min       float64 `json:"min,omitempty"`
	Max       float64 `json:"max,omitempty"`
	Scale     int     `json:"scale,omitempty"`
	Precision int     `json:"precision,omitempty"`
	UnitId    string  `json:"unitId,omitempty"`

	// Text only
	MinLength int    `json:"minLength,omitempty"`
	MaxLength int    `json:"maxLength,omitempty"`
	Pattern   string `json:"pattern,omitempty"`

	// Enum only
	Options *map[string]any `json:"options,omitempty"`

	// Date, time, timestamp
	Before string `json:"before,omitempty"`
	After  string `json:"after,omitempty"`
}

type Fields struct {
	Elevation Field[any] `json:"elevation"`
	Latitude  Field[any] `json:"latitude"`
	Longitude Field[any] `json:"longitude"`
}

type Field[T any] struct {
	Value     T       `json:"value,omitempty"`
	Min       float64 `json:"min,omitempty"`
	Max       float64 `json:"max,omitempty"`
	Scale     int     `json:"scale,omitempty"`
	Precision int     `json:"precision,omitempty"`
	UnitId    string  `json:"unitId,omitempty"`
}
