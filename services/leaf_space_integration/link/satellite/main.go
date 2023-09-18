package leafSpaceSatelliteLink

import (
	"github.com/leanspace/terraform-provider-leanspace/provider"
)

var LeafSpaceSatellitesLinkDataType = provider.DataSourceType[LeafSpaceSatelliteLink, *LeafSpaceSatelliteLink]{
	ResourceIdentifier: "leanspace_leaf_space_satellites_link",
	Path:               "integration-leafspace/satellites/links",
	Schema:             leafSpaceSatelliteLink,
	FilterSchema:       dataSourceFilterSchema,
}
