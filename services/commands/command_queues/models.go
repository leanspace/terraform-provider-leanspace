package command_queues

type CommandQueue struct {
	ID                          string   `json:"id"`
	AssetId                     string   `json:"assetId"`
	Name                        string   `json:"name"`
	GroundStationIds            []string `json:"groundStationIds,"`
	CommandTransformerPluginId  string   `json:"commandTransformerPluginId"`
	ProtocolTransformerPluginId string   `json:"protocolTransformerPluginId"`
	ProtocolTransformerInitData string   `json:"protocolTransformerInitData"`
	CreatedAt                   string   `json:"createdAt"`
	CreatedBy                   string   `json:"createdBy"`
	LastModifiedAt              string   `json:"lastModifiedAt"`
	LastModifiedBy              string   `json:"lastModifiedBy"`
}

func (queue *CommandQueue) GetID() string { return queue.ID }
