package metrics

import "github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

type Metric[T any] struct {
	general_objects.AuditModel
	Name            string                                 `json:"name"`
	Description     *string                                `json:"description,omitempty"`
	NodeId          string                                 `json:"nodeId"`
	AncestorAssetId *string                                `json:"ancestorAssetId,omitempty"`
	Tags            []general_objects.KeyValue             `json:"tags,omitempty"`
	Attributes      general_objects.DefinitionAttribute[T] `json:"attributes"`
}
