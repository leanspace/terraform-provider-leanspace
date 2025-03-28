package groundstation_links

type LeafSpaceGroundStationLink struct {
	ID                         string `json:"id"`
	LeafspaceGroundStationId   string `json:"leafspaceGroundStationId"`
	LeafspaceGroundStationName string `json:"leafspaceGroundStationName"`
	LeanspaceGroundStationId   string `json:"leanspaceGroundStationId"`
	LeanspaceGroundStationName string `json:"leanspaceGroundStationName"`
}

func (leafSpaceGroundStationLink *LeafSpaceGroundStationLink) GetID() string {
	return leafSpaceGroundStationLink.ID
}
