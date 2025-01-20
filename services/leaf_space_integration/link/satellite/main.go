package satellite_links

import (
	"github.com/leanspace/terraform-provider-leanspace/provider"
)

var LeafSpaceSatellitesLinkDataType = provider.DataSourceType[LeafSpaceSatelliteLink, *LeafSpaceSatelliteLink]{
	ResourceIdentifier: "leanspace_leaf_space_satellite_links",
	Path:               "integration-leafspace/satellites/links",
	Schema:             leafSpaceSatelliteLink,
	FilterSchema:       dataSourceFilterSchema,
}
