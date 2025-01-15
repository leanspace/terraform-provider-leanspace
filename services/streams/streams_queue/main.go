package streams_queue

import (
	"github.com/leanspace/terraform-provider-leanspace/provider"
	"github.com/leanspace/terraform-provider-leanspace/services/streams/streams"
)

var StreamQueueDataType = provider.DataSourceType[streams.Stream, *streams.Stream]{
	ResourceIdentifier: "leanspace_stream_queues",
	Path:               Path,
	Schema:             streams.StreamSchema,
	FilterSchema:       streams.DataSourceFilterSchema,
	CreatePath: func(_ *streams.Stream) string {
		return "streams-repository/stream-queues"
	},
	UpdateFunction: func(client *provider.Client, id string, updatedStream *streams.Stream) (*streams.Stream, error) {
		return UpdateStream(updatedStream, client, id)
	},
	CreateFunction: func(client *provider.Client, stream *streams.Stream) (*streams.Stream, error) {
		return CreateStream(stream, client)
	},
}

// need to declare there in order to avoid cyclic dependencies
var Path = "streams-repository/streams"
