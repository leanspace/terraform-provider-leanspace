package leaf_space_groundstation_link

import (
	"github.com/leanspace/terraform-provider-leanspace/provider"
)

var LeafSpaceGroundStationLinkDataType = provider.DataSourceType[LeafSpaceGroundStationLink, *LeafSpaceGroundStationLink]{
	ResourceIdentifier: "leanspace_leaf_space_ground_station_links",
	Path:               "integration-leafspace/ground-stations/links",
	Schema:             leafSpaceGroundStationLink,
	FilterSchema:       dataSourceFilterSchema,
}
