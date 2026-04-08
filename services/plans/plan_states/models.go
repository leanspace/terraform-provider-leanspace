package plan_states

import "github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

type PlanState struct {
	general_objects.AuditModel
	Name     string `json:"name"`
	ReadOnly bool   `json:"readOnly"`
}
