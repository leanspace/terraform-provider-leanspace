package streams

import "leanspace-terraform-provider/provider"

var StreamDataType = provider.DataSourceType[Stream, *Stream]{
	ResourceIdentifier: "leanspace_streams",
	Name:               "stream",
	Path:               "streams-repository/streams",
	Schema:             streamSchema,
}
