package command_sequence_states

import "github.com/leanspace/terraform-provider-leanspace/provider"

var CommandSequenceStateDataType = provider.DataSourceType[CommandSequenceState, *CommandSequenceState]{
	ResourceIdentifier: "leanspace_command_sequence_states",
	Path:               "commands-repository/command-sequences/states",
	Schema:             commandSequenceSchema,
	FilterSchema:       nil,
}
