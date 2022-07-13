package streams

import (
	"terraform-provider-asset/asset"
)

var StreamDataType = asset.DataSourceType[Stream, *Stream]{
	ResourceIdentifier: "leanspace_streams",
	Name:               "stream",
	Path:               "streams-repository/streams",
	Schema:             streamSchema,
}
