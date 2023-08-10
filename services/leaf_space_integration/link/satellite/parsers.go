package leafSpaceSatellite

func (leafSpaceSatellite *LeafSpaceSatellite) ToMap() map[string]any {
	leafSpaceSatelliteStateMap := make(map[string]any)
	leafSpaceSatelliteStateMap["id"] = leafSpaceSatellite.ID
	leafSpaceSatelliteStateMap["leafspace_satellite_id"] = leafSpaceSatellite.LeafspaceSatelliteId
	leafSpaceSatelliteStateMap["leafspace_satellite_name"] = leafSpaceSatellite.LeafspaceSatelliteName
	leafSpaceSatelliteStateMap["leanspace_satellite_id"] = leafSpaceSatellite.LeanspaceSatelliteId
	leafSpaceSatelliteStateMap["leanspace_satellite_name"] = leafSpaceSatellite.LeanspaceSatelliteName

	return leafSpaceSatelliteStateMap
}

func (leafSpaceSatellite *LeafSpaceSatellite) FromMap(leafSpaceIntegrationMap map[string]any) error {
	leafSpaceSatellite.ID = leafSpaceIntegrationMap["id"].(string)
	leafSpaceSatellite.LeafspaceSatelliteId = leafSpaceIntegrationMap["leafspace_satellite_id"].(string)
	leafSpaceSatellite.LeafspaceSatelliteName = leafSpaceIntegrationMap["leafspace_satellite_name"].(string)
	leafSpaceSatellite.LeanspaceSatelliteId = leafSpaceIntegrationMap["leanspace_satellite_id"].(string)
	leafSpaceSatellite.LeanspaceSatelliteName = leafSpaceIntegrationMap["leanspace_satellite_name"].(string)

	return nil
}
