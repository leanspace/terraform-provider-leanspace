// Package types holds the pure data-shape types used to describe generic
// attributes (metrics, structure fields, etc). It has no dependency on
// terraform-plugin-sdk so it can be imported by non-provider tools (e.g.
// xtceimporter) that only need the JSON shape, without pulling in the SDK.
package types

type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value,omitempty"`
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
	// Geopoint
	Fields *FieldsDef `json:"fields,omitempty"`
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
	// Text & binary
	MinLength int `json:"minLength,omitempty"`
	MaxLength int `json:"maxLength,omitempty"`
	// Text only
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
}

type ValueAttribute[T any] struct {
	Value T      `json:"value,omitempty"`
	Type  string `json:"type"`
	// Numeric
	UnitId string `json:"unitId,omitempty"`
	// Geopoint
	Fields *Fields `json:"fields,omitempty"`
	// Array
	DataType string `json:"dataType,omitempty"`
}

type FieldsDef struct {
	Elevation FieldDef[any] `json:"elevation"`
	Latitude  FieldDef[any] `json:"latitude"`
	Longitude FieldDef[any] `json:"longitude"`
}

type FieldDef[T any] struct {
	DefaultValue T       `json:"defaultValue,omitempty"`
	Min          float64 `json:"min,omitempty"`
	Max          float64 `json:"max,omitempty"`
	Scale        int     `json:"scale,omitempty"`
	Precision    int     `json:"precision,omitempty"`
	UnitId       string  `json:"unitId,omitempty"`
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
