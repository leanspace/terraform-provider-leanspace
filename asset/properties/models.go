package properties

import "terraform-provider-asset/asset/general_objects"

type Property[T any] struct {
	ID             string                `json:"id" terra:"id"`
	Name           string                `json:"name" terra:"name"`
	Description    string                `json:"description,omitempty" terra:"description"`
	NodeId         string                `json:"nodeId" terra:"node_id"`
	CreatedAt      string                `json:"createdAt" terra:"created_at"`
	CreatedBy      string                `json:"createdBy" terra:"created_by"`
	LastModifiedAt string                `json:"lastModifiedAt" terra:"last_modified_at"`
	LastModifiedBy string                `json:"lastModifiedBy" terra:"last_modified_by"`
	Tags           []general_objects.Tag `json:"tags,omitempty" terra:"tags"`
	MinLength      int                   `json:"minLength,omitempty" terra:"min_length"`
	MaxLength      int                   `json:"maxLength,omitempty" terra:"max_length"`
	Pattern        string                `json:"pattern,omitempty" terra:"pattern"`
	Before         string                `json:"before,omitempty" terra:"before"`
	After          string                `json:"after,omitempty" terra:"after"`
	Fields         *Fields               `json:"fields,omitempty" terra:"fields"`
	Options        *map[string]any       `json:"options,omitempty" terra:"options"`
	Min            float64               `json:"min,omitempty" terra:"min"`
	Max            float64               `json:"max,omitempty" terra:"max"`
	Scale          int                   `json:"scale,omitempty" terra:"scale"`
	Precision      int                   `json:"precision,omitempty" terra:"precision"`
	UnitId         string                `json:"unit_id,omitempty" terra:"unit_id"`
	Value          T                     `json:"value,omitempty" terra:"value"`
	Type           string                `json:"type" terra:"type"`
}

type Fields struct {
	Elevation Field[any] `json:"elevation" terra:"elevation"`
	Latitude  Field[any] `json:"latitude" terra:"latitude"`
	Longitude Field[any] `json:"longitude" terra:"longitude"`
}

type Field[T any] struct {
	Type           string                `json:"type" terra:"type"`
	ID             string                `json:"id" terra:"id"`
	Description    string                `json:"description,omitempty" terra:"description"`
	CreatedAt      string                `json:"createdAt" terra:"created_at"`
	CreatedBy      string                `json:"createdBy" terra:"created_by"`
	LastModifiedAt string                `json:"lastModifiedAt" terra:"last_modified_at"`
	LastModifiedBy string                `json:"lastModifiedBy" terra:"last_modified_by"`
	Value          T                     `json:"value" terra:"value"`
	Tags           []general_objects.Tag `json:"tags,omitempty" terra:"tags"`
	Name           string                `json:"name" terra:"name"`
	Min            float64               `json:"min,omitempty" terra:"min"`
	Max            float64               `json:"max,omitempty" terra:"max"`
	Scale          int                   `json:"scale,omitempty" terra:"scale"`
	Precision      int                   `json:"precision,omitempty" terra:"precision"`
	UnitId         string                `json:"unit_id,omitempty" terra:"unit_id"`
}
