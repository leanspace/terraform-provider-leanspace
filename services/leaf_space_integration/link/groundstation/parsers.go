package leafSpaceGroundstation

func (leafSpaceGroundStationConnection *LeafSpaceGroundStationConnection) ToMap() map[string]any {
	leafSpaceGroundStationConnectionStateMap := make(map[string]any)
	leafSpaceGroundStationConnectionStateMap["id"] = leafSpaceGroundStationConnection.ID
	leafSpaceGroundStationConnectionStateMap["leafspace_ground_station_id"] = leafSpaceGroundStationConnection.LeafspaceGroundStationId
	leafSpaceGroundStationConnectionStateMap["leafspace_ground_station_name"] = leafSpaceGroundStationConnection.LeafspaceGroundStationName
	leafSpaceGroundStationConnectionStateMap["leanspace_ground_station_id"] = leafSpaceGroundStationConnection.LeanspaceGroundStationId
	leafSpaceGroundStationConnectionStateMap["leanspace_ground_station_name"] = leafSpaceGroundStationConnection.LeanspaceGroundStationName

	return leafSpaceGroundStationConnectionStateMap
}

func (leafSpaceGroundStationConnection *LeafSpaceGroundStationConnection) FromMap(leafSpaceIntegrationMap map[string]any) error {
	leafSpaceGroundStationConnection.ID = leafSpaceIntegrationMap["id"].(string)
	leafSpaceGroundStationConnection.LeafspaceGroundStationId = leafSpaceIntegrationMap["leafspace_ground_station_id"].(string)
	leafSpaceGroundStationConnection.LeafspaceGroundStationName = leafSpaceIntegrationMap["leafspace_ground_station_name"].(string)
	leafSpaceGroundStationConnection.LeanspaceGroundStationId = leafSpaceIntegrationMap["leanspace_ground_station_id"].(string)
	leafSpaceGroundStationConnection.LeanspaceGroundStationName = leafSpaceIntegrationMap["leanspace_ground_station_name"].(string)

	return nil
}
