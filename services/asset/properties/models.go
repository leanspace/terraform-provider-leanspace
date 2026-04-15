package properties

import "github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

type Property[T any] struct {
	general_objects.AuditModel
	Name        string                     `json:"name"`
	Description *string                    `json:"description,omitempty"`
	IsBuiltIn   bool                       `json:"builtIn,omitempty"`
	NodeId      string                     `json:"nodeId"`
	Tags        []general_objects.KeyValue `json:"tags,omitempty"`
	Attributes  PropertyAttribute[T]       `json:"attributes,omitempty"`
}

type PropertyAttribute[T any] struct {
	// Common
	Value T      `json:"value,omitempty"`
	Type  string `json:"type"`

	// Geopoint only
	Fields *general_objects.Fields `json:"fields,omitempty"`

	// Numeric only
	Min       *float64 `json:"min,omitempty"`
	Max       *float64 `json:"max,omitempty"`
	Scale     *int     `json:"scale,omitempty"`
	Precision *int     `json:"precision,omitempty"`
	UnitId    *string `json:"unitId,omitempty"`

	// Text only
	MinLength *int    `json:"minLength,omitempty"`
	MaxLength *int    `json:"maxLength,omitempty"`
	Pattern   *string `json:"pattern,omitempty"`

	// Enum only
	Options *map[string]any `json:"options,omitempty"`

	// Date, time, timestamp
	Before *string `json:"before,omitempty"`
	After  *string `json:"after,omitempty"`
}
