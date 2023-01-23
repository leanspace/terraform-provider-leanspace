package properties_v2

import "github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

type Property[T any] struct {
	ID             string                               `json:"id"`
	Name           string                               `json:"name"`
	Description    string                               `json:"description,omitempty"`
	IsBuiltIn      bool                                 `json:"builtIn,omitempty"`
	NodeId         string                               `json:"nodeId"`
	CreatedAt      string                               `json:"createdAt"`
	CreatedBy      string                               `json:"createdBy"`
	LastModifiedAt string                               `json:"lastModifiedAt"`
	LastModifiedBy string                               `json:"lastModifiedBy"`
	Tags           []general_objects.Tag                `json:"tags,omitempty"`
	Attributes     general_objects.PropertyAttribute[T] `json:"attributes,omitempty"`
}

func (property *Property[T]) GetID() string { return property.ID }
