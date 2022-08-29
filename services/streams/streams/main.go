package streams

import "leanspace-terraform-provider/provider"

var StreamDataType = provider.DataSourceType[Stream, *Stream]{
	ResourceIdentifier: "leanspace_streams",
	Path:               "streams-repository/streams",
	Schema:             streamSchema,
	FilterSchema:       dataSourceFilterSchema,
}
