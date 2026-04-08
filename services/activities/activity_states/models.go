package activity_states

import "github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

type ActivityState struct {
	general_objects.AuditModel
	Name     string `json:"name"`
	ReadOnly bool   `json:"readOnly"`
}
