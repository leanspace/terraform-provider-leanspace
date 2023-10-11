package command_states

import "github.com/leanspace/terraform-provider-leanspace/provider"

var CommandStateDataType = provider.DataSourceType[CommandState, *CommandState]{
	ResourceIdentifier: "leanspace_command_states",
	Path:               "commands-repository/command-sequences/commands/states",
	Schema:             commandStateSchema,
	FilterSchema:       nil,
}
