package command_queues

import (
	"terraform-provider-asset/asset"
)

var CommandQueueDataType = asset.DataSourceType[CommandQueue, *CommandQueue]{
	ResourceIdentifier: "leanspace_command_queues",
	Name:               "command_queue",
	Path:               "commands-repository/command-queues",
	Schema:             commandQueueSchema,
}
