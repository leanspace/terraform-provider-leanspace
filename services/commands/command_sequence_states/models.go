package command_sequence_states

//go:generate go run github.com/leanspace/terraform-provider-leanspace/tools/gen_tf_models -struct CommandSequenceState

import "github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

type CommandSequenceState struct {
	general_objects.AuditModel
	Name     string `json:"name"`
	ReadOnly bool   `json:"readOnly"`
}
