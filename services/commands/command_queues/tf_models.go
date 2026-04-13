package command_queues

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/leanspace/terraform-provider-leanspace/helper"
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
	return &CommandQueueTF{
		AuditModelTF:                general_objects.AuditModelToTF(&x.AuditModel),
		AssetId:                     helper.TFStringValue(x.AssetId),
		Name:                        helper.TFStringValue(x.Name),
		GroundStationIds:            helper.TFStringsValue(x.GroundStationIds),
		CommandTransformerPluginId:  helper.TFStringValue(x.CommandTransformerPluginId),
		ProtocolTransformerPluginId: helper.TFStringValue(x.ProtocolTransformerPluginId),
		ProtocolTransformerInitData: helper.TFStringValue(x.ProtocolTransformerInitData),
	}
}

func (tf *CommandQueueTF) ToAPI() any {
	return &CommandQueue{
		AuditModel:                  general_objects.AuditModelFromTF(tf.AuditModelTF),
		AssetId:                     helper.FromTFString(tf.AssetId),
		Name:                        helper.FromTFString(tf.Name),
		GroundStationIds:            helper.FromTFStrings(tf.GroundStationIds),
		CommandTransformerPluginId:  helper.FromTFString(tf.CommandTransformerPluginId),
		ProtocolTransformerPluginId: helper.FromTFString(tf.ProtocolTransformerPluginId),
		ProtocolTransformerInitData: helper.FromTFString(tf.ProtocolTransformerInitData),
	}
}
