package leafSpaceSatellite

type LeafSpaceSatellite struct {
	ID                     string `json:"id"`
	LeafspaceSatelliteId   string `json:"leafspaceSatelliteId"`
	LeafspaceSatelliteName string `json:"leafspaceSatelliteName"`
	LeanspaceSatelliteId   string `json:"leanspaceSatelliteId"`
	LeanspaceSatelliteName string `json:"leanspaceSatelliteName"`
}

func (leafSpaceSatellite *LeafSpaceSatellite) GetID() string {
	return leafSpaceSatellite.ID
}
