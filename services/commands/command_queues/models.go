package command_queues

type CommandQueue struct {
	ID                          string   `json:"id" terra:"id"`
	AssetId                     string   `json:"assetId" terra:"asset_id"`
	Name                        string   `json:"name" terra:"name"`
	GroundStationIds            []string `json:"groundStationIds," terra:"ground_station_ids"`
	CommandTransformerPluginId  string   `json:"commandTransformerPluginId" terra:"command_transformer_plugin_id"`
	ProtocolTransformerPluginId string   `json:"protocolTransformerPluginId" terra:"protocol_transformer_plugin_id"`
	ProtocolTransformerInitData string   `json:"protocolTransformerInitData" terra:"protocol_transformer_init_data"`
	CreatedAt                   string   `json:"createdAt" terra:"created_at"`
	CreatedBy                   string   `json:"createdBy" terra:"created_by"`
	LastModifiedAt              string   `json:"lastModifiedAt" terra:"last_modified_at"`
	LastModifiedBy              string   `json:"lastModifiedBy" terra:"last_modified_by"`
}

func (queue *CommandQueue) GetID() string { return queue.ID }
