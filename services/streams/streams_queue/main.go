package streams_queue

import (
	"github.com/leanspace/terraform-provider-leanspace/provider"
)

var StreamQueueDataType = provider.DataSourceType[streamQueue, *streamQueue]{
	ResourceIdentifier: "leanspace_stream_queues",
	Path:               path,
	Schema:             streamSchema,
	FilterSchema:       dataSourceFilterSchema,
	CreatePath: func(_ *streamQueue) string {
		return "streams-repository/stream-queues"
	},
	UpdateFunction: func(client *provider.Client, id string, updatedStreamQueue *streamQueue) (*streamQueue, error) {
		return UpdateStream(updatedStreamQueue, client, id)
	},
	CreateFunction: func(client *provider.Client, streamQueue *streamQueue) (*streamQueue, error) {
		return CreateStream(streamQueue, client)
	},
}

// need to declare there in order to avoid cyclic dependencies
var path = "streams-repository/streams"
