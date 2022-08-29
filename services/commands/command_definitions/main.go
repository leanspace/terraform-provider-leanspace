package command_definitions

import "github.com/leanspace/terraform-provider-leanspace/provider"

var CommandDataType = provider.DataSourceType[CommandDefinition, *CommandDefinition]{
	ResourceIdentifier: "leanspace_command_definitions",
	Path:               "commands-repository/command-definitions",
	Schema:             commandDefinitionSchema,
	FilterSchema:       dataSourceFilterSchema,
}
