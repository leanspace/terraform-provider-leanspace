package command_queues

import "github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

type CommandQueue struct {
	general_objects.AuditModel
	AssetId                     string   `json:"assetId"`
	Name                        string   `json:"name"`
	GroundStationIds            []string `json:"groundStationIds"`
	CommandTransformerPluginId  *string  `json:"commandTransformerPluginId,omitempty"`
	ProtocolTransformerPluginId *string  `json:"protocolTransformerPluginId,omitempty"`
	ProtocolTransformerInitData *string  `json:"protocolTransformerInitData,omitempty"`
}
