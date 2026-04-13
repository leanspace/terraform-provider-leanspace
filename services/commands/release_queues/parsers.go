package release_queues

import (
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

func (queue *ReleaseQueue) ToMap() map[string]any {
	queueMap := queue.ToAuditMap()
	queueMap["asset_id"] = helper.NilIfEmpty(queue.AssetId)
	queueMap["name"] = helper.NilIfEmpty(queue.Name)
	queueMap["description"] = helper.NilIfEmpty(queue.Description)
	queueMap["command_transformer_plugin_id"] = helper.NilIfEmpty(queue.CommandTransformerPluginId)
	queueMap["command_transformation_strategy"] = helper.NilIfEmpty(queue.CommandTransformationStrategy)
	queueMap["command_transformer_plugin_configuration_data"] = helper.NilIfEmpty(queue.CommandTransformerPluginConfigurationData)
	queueMap["global_transmission_metadata"] = helper.ParseToMaps(queue.GlobalTransmissionMetadata)
	queueMap["logical_lock"] = helper.NilIfEmpty(queue.LogicalLock)
	queueMap["tags"] = helper.ParseToMaps(queue.Tags)
	return queueMap
}

func (queue *ReleaseQueue) FromMap(queueMap map[string]any) error {
	queue.FromAuditMap(queueMap)
	queue.AssetId = helper.CastString(queueMap, "asset_id")
	queue.Name = helper.CastString(queueMap, "name")
	queue.Description = helper.CastString(queueMap, "description")
	queue.CommandTransformerPluginId = helper.CastString(queueMap, "command_transformer_plugin_id")
	queue.CommandTransformationStrategy = helper.CastString(queueMap, "command_transformation_strategy")
	queue.CommandTransformerPluginConfigurationData = helper.CastString(queueMap, "command_transformer_plugin_configuration_data")
	if globalTransmissionMetadata, err := helper.ParseFromMaps[general_objects.KeyValue](helper.CastSlice(queueMap, "global_transmission_metadata")); err != nil {
		return err
	} else {
		queue.GlobalTransmissionMetadata = globalTransmissionMetadata
	}
	queue.LogicalLock = helper.CastBool(queueMap, "logical_lock")
	if tags, err := helper.ParseFromMaps[general_objects.KeyValue](helper.CastSlice(queueMap, "tags")); err != nil {
		return err
	} else {
		queue.Tags = tags
	}
	return nil
}
