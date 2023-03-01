package release_queues

import "github.com/leanspace/terraform-provider-leanspace/provider"

var ReleaseQueueDataType = provider.DataSourceType[ReleaseQueue, *ReleaseQueue]{
	ResourceIdentifier: "leanspace_release_queues",
	Path:               "commands-repository/release-queues",
	Schema:             releaseQueueSchema,
	FilterSchema:       dataSourceFilterSchema,
}
