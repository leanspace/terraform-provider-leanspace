package passive_resource_functions

//go:generate go run github.com/leanspace/terraform-provider-leanspace/tools/gen_tf_models -struct PassiveResourceFunction

import "github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

type PassiveResourceFunction struct {
	general_objects.AuditModel
	ResourceId   string                         `json:"resourceId"`
	Name         string                         `json:"name"`
	Formula      PassiveResourceFunctionFormula `json:"formula"`
	ControlBound *float64                       `json:"controlBound,omitempty"`
	Tags         []general_objects.KeyValue     `json:"tags,omitempty"`
}

type PassiveResourceFunctionFormula struct {
	Type     string  `json:"type"`
	Rate     float64 `json:"rate"`
	TimeUnit string  `json:"timeUnit"`
}
