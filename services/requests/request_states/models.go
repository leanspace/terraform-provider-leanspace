package request_states

//go:generate go run github.com/leanspace/terraform-provider-leanspace/tools/gen_tf_models -struct RequestState

import "github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

type RequestState struct {
	general_objects.AuditModel
	Name string `json:"name"`
}
