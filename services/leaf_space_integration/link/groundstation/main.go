package leafSpaceGroundstation

import (
	"github.com/leanspace/terraform-provider-leanspace/provider"
)

var LeafSpaceGroundStationLink = provider.DataSourceType[LeafSpaceGroundStationConnection, *LeafSpaceGroundStationConnection]{
	ResourceIdentifier: "leanspace_leaf_space_ground_station_links",
	Path:               "integration-leafspace/ground-stations/links",
	Schema:             leafSpaceGroundStationLink,
	FilterSchema:       dataSourceFilterSchema,
}
