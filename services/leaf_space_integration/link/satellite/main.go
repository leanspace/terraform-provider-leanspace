package leafSpaceSatellite

import (
	"github.com/leanspace/terraform-provider-leanspace/provider"
)

var LeafSpaceSatellitesLink = provider.DataSourceType[LeafSpaceSatellite, *LeafSpaceSatellite]{
	ResourceIdentifier: "leanspace_leaf_space_satellites_link",
	Path:               "integration-leafspace/satellites/links",
	Schema:             leafSpaceSatelliteLink,
	FilterSchema:       dataSourceFilterSchema,
}
