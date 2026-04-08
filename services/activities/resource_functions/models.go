package resource_functions

import "github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

type ResourceFunction struct {
	general_objects.AuditModel
	ActivityDefinitionId string                   `json:"activityDefinitionId"`
	ResourceId           string                   `json:"resourceId"`
	Name                 string                   `json:"name"`
	Formula              *ResourceFunctionFormula `json:"formula"`
}

type ResourceFunctionFormula struct {
	Type      string  `json:"type"`
	Amplitude float64 `json:"amplitude"`
	Constant  float64 `json:"constant"`
	Rate      float64 `json:"rate"`
	TimeUnit  string  `json:"timeUnit,omitempty"`
}
