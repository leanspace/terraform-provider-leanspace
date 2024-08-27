package streams

import "github.com/leanspace/terraform-provider-leanspace/provider"

var StreamDataType = provider.DataSourceType[Stream, *Stream]{
	ResourceIdentifier: "leanspace_streams",
	Path:               "streams-repository/streams",
	Schema:             streamSchema,
	FilterSchema:       dataSourceFilterSchema,
	CreatePath: func(_ *Stream) string {
		return "streams-repository/stream-queues"
	},
}
