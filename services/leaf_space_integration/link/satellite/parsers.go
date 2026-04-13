package satellite_links

import (
	"github.com/leanspace/terraform-provider-leanspace/helper"
)

func (leafSpaceSatelliteLink *LeafSpaceSatelliteLink) ToMap() map[string]any {
	leafSpaceSatelliteStateMap := make(map[string]any)
	leafSpaceSatelliteStateMap["id"] = helper.NilIfEmpty(leafSpaceSatelliteLink.ID)
	leafSpaceSatelliteStateMap["leafspace_satellite_id"] = helper.NilIfEmpty(leafSpaceSatelliteLink.LeafspaceSatelliteId)
	leafSpaceSatelliteStateMap["leafspace_satellite_name"] = helper.NilIfEmpty(leafSpaceSatelliteLink.LeafspaceSatelliteName)
	leafSpaceSatelliteStateMap["leanspace_satellite_id"] = helper.NilIfEmpty(leafSpaceSatelliteLink.LeanspaceSatelliteId)
	leafSpaceSatelliteStateMap["leanspace_satellite_name"] = helper.NilIfEmpty(leafSpaceSatelliteLink.LeanspaceSatelliteName)

	return leafSpaceSatelliteStateMap
}

func (leafSpaceSatelliteLink *LeafSpaceSatelliteLink) FromMap(leafSpaceIntegrationMap map[string]any) error {
	leafSpaceSatelliteLink.ID = helper.CastString(leafSpaceIntegrationMap, "id")
	leafSpaceSatelliteLink.LeafspaceSatelliteId = helper.CastString(leafSpaceIntegrationMap, "leafspace_satellite_id")
	leafSpaceSatelliteLink.LeafspaceSatelliteName = helper.CastString(leafSpaceIntegrationMap, "leafspace_satellite_name")
	leafSpaceSatelliteLink.LeanspaceSatelliteId = helper.CastString(leafSpaceIntegrationMap, "leanspace_satellite_id")
	leafSpaceSatelliteLink.LeanspaceSatelliteName = helper.CastString(leafSpaceIntegrationMap, "leanspace_satellite_name")

	return nil
}
