package resources

import (
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

type Resource struct {
	ID             string                     `json:"id"`
	AssetId        string                     `json:"assetId"`
	UnitId         string                     `json:"unitId"`
	MetricId       string                     `json:"metricId"`
	Name           string                     `json:"name"`
	Description    string                     `json:"description"`
	DefaultLevel   float64                    `json:"defaultLevel"`
	Constraints    []ResourceConstraints      `json:"constraints,omitempty"`
	LowerLimit     float64                    `json:"lowerLimit,omitempty"`
	UpperLimit     float64                    `json:"upperLimit,omitempty"`
	Thresholds     []ResourceThreshold        `json:"thresholds,omitempty"`
	Tags           []general_objects.KeyValue `json:"tags,omitempty"`
	CreatedAt      string                     `json:"createdAt"`
	CreatedBy      string                     `json:"createdBy"`
	LastModifiedAt string                     `json:"lastModifiedAt"`
	LastModifiedBy string                     `json:"lastModifiedBy"`
}

func (resource *Resource) GetID() string { return resource.ID }

type ResourceConstraints struct {
	Type  string  `json:"type"`
	Kind  string  `json:"kind"`
	Value float64 `json:"value"`
	Name  string  `json:"name"`
}

type ResourceThreshold struct {
	Kind                 string  `json:"kind"`
	Name                 string  `json:"name"`
	ViolationWhenReached bool    `json:"violationWhenReached"`
	Value                float64 `json:"value"`
}
