package groundstation_links

func (leafSpaceGroundStationLink *LeafSpaceGroundStationLink) ToMap() map[string]any {
	leafSpaceGroundStationConnectionStateMap := make(map[string]any)
	leafSpaceGroundStationConnectionStateMap["id"] = leafSpaceGroundStationLink.ID
	leafSpaceGroundStationConnectionStateMap["leafspace_ground_station_id"] = leafSpaceGroundStationLink.LeafspaceGroundStationId
	leafSpaceGroundStationConnectionStateMap["leafspace_ground_station_name"] = leafSpaceGroundStationLink.LeafspaceGroundStationName
	leafSpaceGroundStationConnectionStateMap["leanspace_ground_station_id"] = leafSpaceGroundStationLink.LeanspaceGroundStationId
	leafSpaceGroundStationConnectionStateMap["leanspace_ground_station_name"] = leafSpaceGroundStationLink.LeanspaceGroundStationName

	return leafSpaceGroundStationConnectionStateMap
}

func (leafSpaceGroundStationLink *LeafSpaceGroundStationLink) FromMap(leafSpaceIntegrationMap map[string]any) error {
	leafSpaceGroundStationLink.ID = leafSpaceIntegrationMap["id"].(string)
	leafSpaceGroundStationLink.LeafspaceGroundStationId = leafSpaceIntegrationMap["leafspace_ground_station_id"].(string)
	leafSpaceGroundStationLink.LeafspaceGroundStationName = leafSpaceIntegrationMap["leafspace_ground_station_name"].(string)
	leafSpaceGroundStationLink.LeanspaceGroundStationId = leafSpaceIntegrationMap["leanspace_ground_station_id"].(string)
	leafSpaceGroundStationLink.LeanspaceGroundStationName = leafSpaceIntegrationMap["leanspace_ground_station_name"].(string)

	return nil
}
