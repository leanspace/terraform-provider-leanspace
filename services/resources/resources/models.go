package resources

import (
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

type Resource struct {
	general_objects.AuditModel
	AssetId      string                     `json:"assetId"`
	UnitId       string                     `json:"unitId"`
	MetricId     *string                    `json:"metricId,omitempty"`
	Name         string                     `json:"name"`
	Description  *string                    `json:"description,omitempty"`
	DefaultLevel float64                    `json:"defaultLevel"`
	LowerLimit   *float64                   `json:"lowerLimit"`
	UpperLimit   *float64                   `json:"upperLimit"`
	Thresholds   []ResourceThreshold        `json:"thresholds,omitempty"`
	Tags         []general_objects.KeyValue `json:"tags,omitempty"`
}

type ResourceThreshold struct {
	Kind                 string  `json:"kind"`
	Name                 *string `json:"name,omitempty"`
	ViolationWhenReached bool    `json:"violationWhenReached"`
	Value                float64 `json:"value"`
}
