package command_queues

import (
	"github.com/leanspace/terraform-provider-leanspace/helper"
)

func (queue *CommandQueue) ToMap() map[string]any {
	queueMap := queue.ToAuditMap()
	queueMap["asset_id"] = helper.NilIfEmpty(queue.AssetId)
	queueMap["name"] = helper.NilIfEmpty(queue.Name)
	queueMap["ground_station_ids"] = helper.NilIfEmpty(queue.GroundStationIds)
	queueMap["command_transformer_plugin_id"] = helper.NilIfEmpty(queue.CommandTransformerPluginId)
	queueMap["protocol_transformer_plugin_id"] = helper.NilIfEmpty(queue.ProtocolTransformerPluginId)
	queueMap["protocol_transformer_init_data"] = helper.NilIfEmpty(queue.ProtocolTransformerInitData)
	return queueMap
}

func (queue *CommandQueue) FromMap(queueMap map[string]any) error {
	queue.FromAuditMap(queueMap)
	queue.AssetId = helper.CastString(queueMap, "asset_id")
	queue.Name = helper.CastString(queueMap, "name")
	queue.GroundStationIds = make([]string, len(helper.CastSlice(queueMap, "ground_station_ids")))
	for i, value := range helper.CastSlice(queueMap, "ground_station_ids") {
		queue.GroundStationIds[i] = value.(string)
	}
	queue.CommandTransformerPluginId = helper.CastString(queueMap, "command_transformer_plugin_id")
	queue.ProtocolTransformerPluginId = helper.CastString(queueMap, "protocol_transformer_plugin_id")
	queue.ProtocolTransformerInitData = helper.CastString(queueMap, "protocol_transformer_init_data")
	return nil
}
