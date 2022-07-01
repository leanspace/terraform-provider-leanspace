package command_definitions

import (
	"terraform-provider-asset/asset"
)

var CommandDataType = asset.DataSourceType[CommandDefinition, *CommandDefinition]{
	ResourceIdentifier: "leanspace_command_definitions",
	Name:               "command_definition",
	Path:               "commands-repository/command-definitions",
	Schema:             commandDefinitionSchema,
}
