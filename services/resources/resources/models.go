package resources

//go:generate go run github.com/leanspace/terraform-provider-leanspace/tools/gen_tf_models -struct Resource

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
	DefaultLevel *float64                   `json:"defaultLevel,omitempty"`
	LowerLimit   *float64                   `json:"lowerLimit,omitempty"`
	UpperLimit   *float64                   `json:"upperLimit,omitempty"`
	Thresholds   []ResourceThreshold        `json:"thresholds,omitempty"`
	Tags         []general_objects.KeyValue `json:"tags,omitempty"`
}

type ResourceThreshold struct {
	Kind                 string  `json:"kind"`
	Name                 *string `json:"name,omitempty"`
	ViolationWhenReached bool    `json:"violationWhenReached"`
	Value                float64 `json:"value"`
}
