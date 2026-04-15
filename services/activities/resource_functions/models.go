package resource_functions

//go:generate go run github.com/leanspace/terraform-provider-leanspace/tools/gen_tf_models -struct ResourceFunction

import "github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

type ResourceFunction struct {
	general_objects.AuditModel
	ActivityDefinitionId string                   `json:"activityDefinitionId"`
	ResourceId           string                   `json:"resourceId"`
	Name                 string                   `json:"name"`
	Formula              *ResourceFunctionFormula `json:"formula"`
}

type ResourceFunctionFormula struct {
	Type      string   `json:"type"`
	Amplitude *float64 `json:"amplitude,omitempty"`
	Constant  *float64 `json:"constant,omitempty"`
	Rate      *float64 `json:"rate,omitempty"`
	TimeUnit  *string  `json:"timeUnit,omitempty"`
}
