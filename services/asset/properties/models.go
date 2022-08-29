package properties

import "github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

type Property[T any] struct {
	ID             string                `json:"id"`
	Name           string                `json:"name"`
	Description    string                `json:"description,omitempty"`
	NodeId         string                `json:"nodeId"`
	CreatedAt      string                `json:"createdAt"`
	CreatedBy      string                `json:"createdBy"`
	LastModifiedAt string                `json:"lastModifiedAt"`
	LastModifiedBy string                `json:"lastModifiedBy"`
	Tags           []general_objects.Tag `json:"tags,omitempty"`
	MinLength      int                   `json:"minLength,omitempty"`
	MaxLength      int                   `json:"maxLength,omitempty"`
	Pattern        string                `json:"pattern,omitempty"`
	Before         string                `json:"before,omitempty"`
	After          string                `json:"after,omitempty"`
	Fields         *Fields               `json:"fields,omitempty"`
	Options        *map[string]any       `json:"options,omitempty"`
	Min            float64               `json:"min,omitempty"`
	Max            float64               `json:"max,omitempty"`
	Scale          int                   `json:"scale,omitempty"`
	Precision      int                   `json:"precision,omitempty"`
	UnitId         string                `json:"unitId,omitempty"`
	Value          T                     `json:"value,omitempty"`
	Type           string                `json:"type"`
}

func (property *Property[T]) GetID() string { return property.ID }

type Fields struct {
	Elevation Field[any] `json:"elevation"`
	Latitude  Field[any] `json:"latitude"`
	Longitude Field[any] `json:"longitude"`
}

type Field[T any] struct {
	Type           string                `json:"type"`
	ID             string                `json:"id"`
	Description    string                `json:"description,omitempty"`
	CreatedAt      string                `json:"createdAt"`
	CreatedBy      string                `json:"createdBy"`
	LastModifiedAt string                `json:"lastModifiedAt"`
	LastModifiedBy string                `json:"lastModifiedBy"`
	Value          T                     `json:"value"`
	Tags           []general_objects.Tag `json:"tags,omitempty"`
	Name           string                `json:"name"`
	Min            float64               `json:"min,omitempty"`
	Max            float64               `json:"max,omitempty"`
	Scale          int                   `json:"scale,omitempty"`
	Precision      int                   `json:"precision,omitempty"`
	UnitId         string                `json:"unitId,omitempty"`
}
