package release_queues

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func (queue *ReleaseQueue) ToMap() map[string]any {
	queueMap := make(map[string]any)
	queueMap["id"] = queue.ID
	queueMap["asset_id"] = queue.AssetId
	queueMap["name"] = queue.Name
	queueMap["description"] = queue.Description
	queueMap["command_transformer_plugin_id"] = queue.CommandTransformerPluginId
	queueMap["command_transformation_strategy"] = queue.CommandTransformationStrategy
	queueMap["command_transformer_plugin_configuration_data"] = queue.CommandTransformerPluginConfigurationData
	queueMap["global_transmission_metadata"] = queue.GlobalTransmissionMetadata
	queueMap["logical_lock"] = queue.LogicalLock
	queueMap["created_at"] = queue.CreatedAt
	queueMap["created_by"] = queue.CreatedBy
	queueMap["last_modified_at"] = queue.LastModifiedAt
	queueMap["last_modified_by"] = queue.LastModifiedBy

	return queueMap
}

func (queue *ReleaseQueue) FromMap(queueMap map[string]any) error {
	queue.ID = queueMap["id"].(string)
	queue.AssetId = queueMap["asset_id"].(string)
	queue.Name = queueMap["name"].(string)
	queue.Description = queueMap["description"].(string)
	queue.CommandTransformerPluginId = queueMap["command_transformer_plugin_id"].(string)
	queue.CommandTransformationStrategy = queueMap["command_transformation_strategy"].(string)
	queue.CommandTransformerPluginConfigurationData = queueMap["command_transformer_plugin_configuration_data"].(string)
    if globalTransmissionMetadata, err := helper.ParseFromMaps[general_objects.KeyValue](nodeMap["global_transmission_metadata"].(*schema.Set).List()); err != nil {
        return err
    } else {
        queue.GlobalTransmissionMetadata = globalTransmissionMetadata
    }
	queue.LogicalLock = queueMap["logical_lock"].(string)
	queue.CreatedAt = queueMap["created_at"].(string)
	queue.CreatedBy = queueMap["created_by"].(string)
	queue.LastModifiedAt = queueMap["last_modified_at"].(string)
	queue.LastModifiedBy = queueMap["last_modified_by"].(string)

	return nil
}
