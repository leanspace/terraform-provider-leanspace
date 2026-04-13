package groundstation_links

import (
	"github.com/leanspace/terraform-provider-leanspace/helper"
)

func (leafSpaceGroundStationLink *LeafSpaceGroundStationLink) ToMap() map[string]any {
	leafSpaceGroundStationConnectionStateMap := make(map[string]any)
	leafSpaceGroundStationConnectionStateMap["id"] = helper.NilIfEmpty(leafSpaceGroundStationLink.ID)
	leafSpaceGroundStationConnectionStateMap["leafspace_ground_station_id"] = helper.NilIfEmpty(leafSpaceGroundStationLink.LeafspaceGroundStationId)
	leafSpaceGroundStationConnectionStateMap["leafspace_ground_station_name"] = helper.NilIfEmpty(leafSpaceGroundStationLink.LeafspaceGroundStationName)
	leafSpaceGroundStationConnectionStateMap["leanspace_ground_station_id"] = helper.NilIfEmpty(leafSpaceGroundStationLink.LeanspaceGroundStationId)
	leafSpaceGroundStationConnectionStateMap["leanspace_ground_station_name"] = helper.NilIfEmpty(leafSpaceGroundStationLink.LeanspaceGroundStationName)

	return leafSpaceGroundStationConnectionStateMap
}

func (leafSpaceGroundStationLink *LeafSpaceGroundStationLink) FromMap(leafSpaceIntegrationMap map[string]any) error {
	leafSpaceGroundStationLink.ID = helper.CastString(leafSpaceIntegrationMap, "id")
	leafSpaceGroundStationLink.LeafspaceGroundStationId = helper.CastString(leafSpaceIntegrationMap, "leafspace_ground_station_id")
	leafSpaceGroundStationLink.LeafspaceGroundStationName = helper.CastString(leafSpaceIntegrationMap, "leafspace_ground_station_name")
	leafSpaceGroundStationLink.LeanspaceGroundStationId = helper.CastString(leafSpaceIntegrationMap, "leanspace_ground_station_id")
	leafSpaceGroundStationLink.LeanspaceGroundStationName = helper.CastString(leafSpaceIntegrationMap, "leanspace_ground_station_name")

	return nil
}
