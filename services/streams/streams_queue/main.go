package streams_queue

import "github.com/leanspace/terraform-provider-leanspace/provider"

var StreamQueueDataType = provider.DataSourceType[Stream, *Stream]{
	ResourceIdentifier: "leanspace_streams_queue",
	Path:               "streams-repository/streams",
	Schema:             streamSchema,
	FilterSchema:       dataSourceFilterSchema,
	CreatePath: func(_ *Stream) string {
		return "streams-repository/stream-queues"
	},
	UpdatePath: func(streamQueue *Stream) string {
		return "streams-repository/stream-queues/" + streamQueue.StreamQueueId
	},
}
