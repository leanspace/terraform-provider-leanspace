package satellite_links

func (leafSpaceSatelliteLink *LeafSpaceSatelliteLink) ToMap() map[string]any {
	leafSpaceSatelliteStateMap := make(map[string]any)
	leafSpaceSatelliteStateMap["id"] = leafSpaceSatelliteLink.ID
	leafSpaceSatelliteStateMap["leafspace_satellite_id"] = leafSpaceSatelliteLink.LeafspaceSatelliteId
	leafSpaceSatelliteStateMap["leafspace_satellite_name"] = leafSpaceSatelliteLink.LeafspaceSatelliteName
	leafSpaceSatelliteStateMap["leanspace_satellite_id"] = leafSpaceSatelliteLink.LeanspaceSatelliteId
	leafSpaceSatelliteStateMap["leanspace_satellite_name"] = leafSpaceSatelliteLink.LeanspaceSatelliteName

	return leafSpaceSatelliteStateMap
}

func (leafSpaceSatelliteLink *LeafSpaceSatelliteLink) FromMap(leafSpaceIntegrationMap map[string]any) error {
	leafSpaceSatelliteLink.ID = leafSpaceIntegrationMap["id"].(string)
	leafSpaceSatelliteLink.LeafspaceSatelliteId = leafSpaceIntegrationMap["leafspace_satellite_id"].(string)
	leafSpaceSatelliteLink.LeafspaceSatelliteName = leafSpaceIntegrationMap["leafspace_satellite_name"].(string)
	leafSpaceSatelliteLink.LeanspaceSatelliteId = leafSpaceIntegrationMap["leanspace_satellite_id"].(string)
	leafSpaceSatelliteLink.LeanspaceSatelliteName = leafSpaceIntegrationMap["leanspace_satellite_name"].(string)

	return nil
}
