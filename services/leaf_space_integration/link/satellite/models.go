package satellite_links

//go:generate go run github.com/leanspace/terraform-provider-leanspace/tools/gen_tf_models -struct LeafSpaceSatelliteLink

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
