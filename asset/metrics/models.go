package metrics

import "terraform-provider-asset/asset/general_objects"

type Metric[T any] struct {
	ID             string                                 `json:"id" terra:"id"`
	Name           string                                 `json:"name" terra:"name"`
	Description    string                                 `json:"description,omitempty" terra:"description"`
	NodeId         string                                 `json:"nodeId" terra:"node_id"`
	CreatedAt      string                                 `json:"createdAt" terra:"created_at"`
	CreatedBy      string                                 `json:"createdBy" terra:"created_by"`
	LastModifiedAt string                                 `json:"lastModifiedAt" terra:"last_modified_at"`
	LastModifiedBy string                                 `json:"lastModifiedBy" terra:"last_modified_by"`
	Tags           []general_objects.Tag                  `json:"tags,omitempty" terra:"tags"`
	Attributes     general_objects.DefinitionAttribute[T] `json:"attributes" terra:"attributes"`
}

func (prop *Metric[T]) GetID() string { return prop.ID }
