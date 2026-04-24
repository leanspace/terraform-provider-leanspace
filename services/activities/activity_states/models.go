package activity_states

//go:generate go run github.com/leanspace/terraform-provider-leanspace/tools/gen_tf_models -struct ActivityState

import "github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

type ActivityState struct {
	general_objects.AuditModel
	Name     string `json:"name"`
	ReadOnly bool   `json:"readOnly"`
}
