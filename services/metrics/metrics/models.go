package metrics

import "github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

type Metric[T any] struct {
	ID             string                                 `json:"id"`
	Name           string                                 `json:"name"`
	Description    string                                 `json:"description,omitempty"`
	NodeId         string                                 `json:"nodeId"`
	CreatedAt      string                                 `json:"createdAt"`
	CreatedBy      string                                 `json:"createdBy"`
	LastModifiedAt string                                 `json:"lastModifiedAt"`
	LastModifiedBy string                                 `json:"lastModifiedBy"`
	Tags           []general_objects.KeyValue             `json:"tags,omitempty"`
	Attributes     general_objects.DefinitionAttribute[T] `json:"attributes"`
}

func (metric *Metric[T]) GetID() string { return metric.ID }
