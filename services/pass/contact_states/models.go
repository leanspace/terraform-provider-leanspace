package contact_states

//go:generate go run github.com/leanspace/terraform-provider-leanspace/tools/gen_tf_models -struct ContactState

import "github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

type ContactState struct {
	general_objects.AuditModel
	Name     string `json:"name"`
	ReadOnly bool   `json:"readOnly"`
}
