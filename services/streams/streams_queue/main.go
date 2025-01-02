package streams_queue

import (
	"github.com/leanspace/terraform-provider-leanspace/provider"
)

var StreamQueueDataType = provider.DataSourceType[Stream, *Stream]{
	ResourceIdentifier: "leanspace_streams_queue",
	Path:               Path,
	Schema:             streamSchema,
	FilterSchema:       dataSourceFilterSchema,
	CreatePath: func(_ *Stream) string {
		return "streams-repository/stream-queues"
	},
	UpdateFunction: func(client *provider.Client, id string, updatedStream *Stream) (*Stream, error) {
		return updatedStream.UpdateStream(client, id, updatedStream)
	},
}

// need to declare there in order to avoid cyclic dependencies
var Path = "streams-repository/streams"
