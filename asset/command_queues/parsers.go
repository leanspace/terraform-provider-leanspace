package command_queues

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func (queue *CommandQueue) ToMap() map[string]any {
	queueMap := make(map[string]any)
	queueMap["id"] = queue.ID
	queueMap["asset_id"] = queue.AssetId
	queueMap["name"] = queue.Name
	queueMap["ground_station_ids"] = queue.GroundStationIds
	queueMap["command_transformer_plugin_id"] = queue.CommandTransformerPluginId
	queueMap["protocol_transformer_plugin_id"] = queue.ProtocolTransformerPluginId
	queueMap["protocol_transformer_init_data"] = queue.ProtocolTransformerInitData
	queueMap["created_at"] = queue.CreatedAt
	queueMap["created_by"] = queue.CreatedBy
	queueMap["last_modified_at"] = queue.LastModifiedAt
	queueMap["last_modified_by"] = queue.LastModifiedBy

	return queueMap
}

func (queue *CommandQueue) FromMap(queueMap map[string]any) error {
	queue.ID = queueMap["id"].(string)
	queue.AssetId = queueMap["asset_id"].(string)
	queue.Name = queueMap["name"].(string)
	queue.GroundStationIds = make([]string, queueMap["ground_station_ids"].(*schema.Set).Len())
	for i, value := range queueMap["ground_station_ids"].(*schema.Set).List() {
		queue.GroundStationIds[i] = value.(string)
	}
	queue.CommandTransformerPluginId = queueMap["command_transformer_plugin_id"].(string)
	queue.ProtocolTransformerPluginId = queueMap["protocol_transformer_plugin_id"].(string)
	queue.ProtocolTransformerInitData = queueMap["protocol_transformer_init_data"].(string)
	queue.CreatedAt = queueMap["created_at"].(string)
	queue.CreatedBy = queueMap["created_by"].(string)
	queue.LastModifiedAt = queueMap["last_modified_at"].(string)
	queue.LastModifiedBy = queueMap["last_modified_by"].(string)

	return nil
}
