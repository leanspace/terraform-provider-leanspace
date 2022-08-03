package command_definitions

import "leanspace-terraform-provider/provider"

var CommandDataType = provider.DataSourceType[CommandDefinition, *CommandDefinition]{
	ResourceIdentifier: "leanspace_command_definitions",
	Name:               "command_definition",
	Path:               "commands-repository/command-definitions",
	Schema:             commandDefinitionSchema,
}
