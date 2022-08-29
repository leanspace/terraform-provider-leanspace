package command_queues

import "github.com/leanspace/terraform-provider-leanspace/provider"

var CommandQueueDataType = provider.DataSourceType[CommandQueue, *CommandQueue]{
	ResourceIdentifier: "leanspace_command_queues",
	Path:               "commands-repository/command-queues",
	Schema:             commandQueueSchema,
	FilterSchema:       dataSourceFilterSchema,
}
