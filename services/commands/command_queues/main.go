package command_queues

import "leanspace-terraform-provider/provider"

var CommandQueueDataType = provider.DataSourceType[CommandQueue, *CommandQueue]{
	ResourceIdentifier: "leanspace_command_queues",
	Name:               "command_queue",
	Path:               "commands-repository/command-queues",
	Schema:             commandQueueSchema,
}