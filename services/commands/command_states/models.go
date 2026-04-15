package command_states

//go:generate go run github.com/leanspace/terraform-provider-leanspace/tools/gen_tf_models -struct CommandState

import "github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

type CommandState struct {
	general_objects.AuditModel
	Name     string                     `json:"name"`
	ReadOnly bool                       `json:"readOnly"`
	Tags     []general_objects.KeyValue `json:"tags,omitempty"`
}
