package command_definitions

import (
	"fmt"
	"terraform-provider-asset/asset"
)

var CommandDataType = asset.DataSourceType[CommandDefinition]{
	ResourceIdentifier: "leanspace_command_definitions",
	Name:               "command_definition",
	Path:               "asset-repository/command-definitions",
	CreatePath: func(c CommandDefinition) string {
		return fmt.Sprintf("asset-repository/nodes/%s/command-definitions", c.NodeId)
	},

	Schema:     commandDefinitionSchema,
	RootSchema: commandDefinitionSchema,

	GetID:       func(c *CommandDefinition) string { return c.ID },
	MapToStruct: getCommandDefinitionData,
	StructToMap: commandDefinitionStructToInterface,
}
