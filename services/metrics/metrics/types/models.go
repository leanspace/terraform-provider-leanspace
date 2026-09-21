// Package types holds the pure data shape of a Leanspace metric (the JSON
// shape of the leanspace_metrics resource/data source), with no dependency
// on terraform-plugin-sdk, so it can be imported by non-provider tools (e.g.
// xtceimporter) that only need the shape for compile-time drift checking,
// without pulling in the SDK.
package types

import (
	generalobjectstypes "github.com/leanspace/terraform-provider-leanspace/helper/general_objects/types"
)

type Metric[T any] struct {
	ID             string                                     `json:"id"`
	Name           string                                     `json:"name"`
	Description    string                                     `json:"description,omitempty"`
	NodeId         string                                     `json:"nodeId"`
	CreatedAt      string                                     `json:"createdAt"`
	CreatedBy      string                                     `json:"createdBy"`
	LastModifiedAt string                                     `json:"lastModifiedAt"`
	LastModifiedBy string                                     `json:"lastModifiedBy"`
	Tags           []generalobjectstypes.KeyValue             `json:"tags,omitempty"`
	Attributes     generalobjectstypes.DefinitionAttribute[T] `json:"attributes"`
}

func (metric *Metric[T]) GetID() string { return metric.ID }
