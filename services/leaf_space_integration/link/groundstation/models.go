package leafSpaceGroundstation

type LeafSpaceGroundStationConnection struct {
	ID                         string `json:"id"`
	LeafspaceGroundStationId   string `json:"leafspaceGroundStationId"`
	LeafspaceGroundStationName string `json:"leafspaceGroundStationName"`
	LeanspaceGroundStationId   string `json:"leanspaceGroundStationId"`
	LeanspaceGroundStationName string `json:"leanspaceGroundStationName"`
}

func (leafSpaceGroundStationConnection *LeafSpaceGroundStationConnection) GetID() string {
	return leafSpaceGroundStationConnection.ID
}
