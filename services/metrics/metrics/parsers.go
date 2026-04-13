package metrics

import (
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

func (metric *Metric[T]) ToMap() map[string]any {
	metricMap := metric.ToAuditMap()
	metricMap["node_id"] = helper.NilIfEmpty(metric.NodeId)
	metricMap["name"] = helper.NilIfEmpty(metric.Name)
	metricMap["description"] = helper.NilIfEmpty(metric.Description)
	metricMap["attributes"] = metric.Attributes.ToMap()
	metricMap["tags"] = helper.ParseToMaps(metric.Tags)
	return metricMap
}

func (metric *Metric[T]) FromMap(metricMap map[string]any) error {
	metric.FromAuditMap(metricMap)
	metric.NodeId = helper.CastString(metricMap, "node_id")
	metric.Name = helper.CastString(metricMap, "name")
	metric.Description = helper.CastString(metricMap, "description")
	if err := metric.Attributes.FromMap(helper.CastMapAny(metricMap, "attributes")); err != nil {
		return err
	}
	if tags, err := helper.ParseFromMaps[general_objects.KeyValue](helper.CastSlice(metricMap, "tags")); err != nil {
		return err
	} else {
		metric.Tags = tags
	}
	return nil
}
