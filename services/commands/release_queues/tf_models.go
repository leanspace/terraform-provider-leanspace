package release_queues

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

type ReleaseQueueTF struct {
	general_objects.AuditModelTF
	AssetId                                   types.String                `tfsdk:"asset_id"`
	Name                                      types.String                `tfsdk:"name"`
	Description                               types.String                `tfsdk:"description"`
	CommandTransformerPluginId                types.String                `tfsdk:"command_transformer_plugin_id"`
	CommandTransformationStrategy             types.String                `tfsdk:"command_transformation_strategy"`
	CommandTransformerPluginConfigurationData types.String                `tfsdk:"command_transformer_plugin_configuration_data"`
	GlobalTransmissionMetadata                []general_objects.KeyValueTF `tfsdk:"global_transmission_metadata"`
	LogicalLock                               types.Bool                  `tfsdk:"logical_lock"`
	Tags                                      []general_objects.KeyValueTF `tfsdk:"tags"`
}

func (x *ReleaseQueue) ToTF() any {
	return &ReleaseQueueTF{
		AuditModelTF:                              general_objects.AuditModelToTF(&x.AuditModel),
		AssetId:                                   helper.TFStringValue(x.AssetId),
		Name:                                      helper.TFStringValue(x.Name),
		Description:                               helper.TFStringValue(x.Description),
		CommandTransformerPluginId:                helper.TFStringValue(x.CommandTransformerPluginId),
		CommandTransformationStrategy:             helper.TFStringValue(x.CommandTransformationStrategy),
		CommandTransformerPluginConfigurationData: helper.TFStringValue(x.CommandTransformerPluginConfigurationData),
		GlobalTransmissionMetadata:                general_objects.KeyValuesToTF(x.GlobalTransmissionMetadata),
		LogicalLock:                               helper.TFBoolValue(x.LogicalLock),
		Tags:                                      general_objects.KeyValuesToTF(x.Tags),
	}
}

func (tf *ReleaseQueueTF) ToAPI() any {
	return &ReleaseQueue{
		AuditModel:                                general_objects.AuditModelFromTF(tf.AuditModelTF),
		AssetId:                                   helper.FromTFString(tf.AssetId),
		Name:                                      helper.FromTFString(tf.Name),
		Description:                               helper.FromTFString(tf.Description),
		CommandTransformerPluginId:                helper.FromTFString(tf.CommandTransformerPluginId),
		CommandTransformationStrategy:             helper.FromTFString(tf.CommandTransformationStrategy),
		CommandTransformerPluginConfigurationData: helper.FromTFString(tf.CommandTransformerPluginConfigurationData),
		GlobalTransmissionMetadata:                general_objects.KeyValuesFromTF(tf.GlobalTransmissionMetadata),
		LogicalLock:                               helper.FromTFBool(tf.LogicalLock),
		Tags:                                      general_objects.KeyValuesFromTF(tf.Tags),
	}
}
