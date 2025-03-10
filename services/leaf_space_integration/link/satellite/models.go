package satellite_links

type LeafSpaceSatelliteLink struct {
	ID                     string `json:"id"`
	LeafspaceSatelliteId   string `json:"leafspaceSatelliteId"`
	LeafspaceSatelliteName string `json:"leafspaceSatelliteName"`
	LeanspaceSatelliteId   string `json:"leanspaceSatelliteId"`
	LeanspaceSatelliteName string `json:"leanspaceSatelliteName"`
}

func (leafSpaceSatelliteLink *LeafSpaceSatelliteLink) GetID() string {
	return leafSpaceSatelliteLink.ID
}
