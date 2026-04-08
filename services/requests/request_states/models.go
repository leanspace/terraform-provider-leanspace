package request_states

import "github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

type RequestState struct {
	general_objects.AuditModel
	Name string `json:"name"`
}
