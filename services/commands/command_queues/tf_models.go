package command_queues

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

type CommandQueueTF struct {
	general_objects.AuditModelTF
	AssetId                     types.String   `tfsdk:"asset_id"`
	Name                        types.String   `tfsdk:"name"`
	GroundStationIds            []types.String `tfsdk:"ground_station_ids"`
	CommandTransformerPluginId  types.String   `tfsdk:"command_transformer_plugin_id"`
	ProtocolTransformerPluginId types.String   `tfsdk:"protocol_transformer_plugin_id"`
	ProtocolTransformerInitData types.String   `tfsdk:"protocol_transformer_init_data"`
}

func (x *CommandQueue) ToTF() any {
	return general_objects.ReflectToTF[CommandQueueTF](x)
}

func (tf *CommandQueueTF) ToAPI() any {
	return general_objects.ReflectFromTF[CommandQueue](tf)
}
