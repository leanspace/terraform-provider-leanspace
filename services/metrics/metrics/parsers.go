package metrics

import (
	"leanspace-terraform-provider/helper"
	"leanspace-terraform-provider/helper/general_objects"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func (metric *Metric[T]) ToMap() map[string]any {
	metricMap := make(map[string]any)
	metricMap["id"] = metric.ID
	metricMap["node_id"] = metric.NodeId
	metricMap["name"] = metric.Name
	metricMap["description"] = metric.Description
	metricMap["created_at"] = metric.CreatedAt
	metricMap["created_by"] = metric.CreatedBy
	metricMap["last_modified_at"] = metric.LastModifiedAt
	metricMap["last_modified_by"] = metric.LastModifiedBy
	metricMap["attributes"] = []map[string]any{metric.Attributes.ToMap()}
	metricMap["tags"] = helper.ParseToMaps(metric.Tags)
	return metricMap
}

func (metric *Metric[T]) FromMap(metricMap map[string]any) error {
	metric.ID = metricMap["id"].(string)
	metric.NodeId = metricMap["node_id"].(string)
	metric.Name = metricMap["name"].(string)
	metric.Description = metricMap["description"].(string)
	metric.CreatedAt = metricMap["created_at"].(string)
	metric.CreatedBy = metricMap["created_by"].(string)
	metric.LastModifiedAt = metricMap["last_modified_at"].(string)
	metric.LastModifiedBy = metricMap["last_modified_by"].(string)
	err := metric.Attributes.FromMap(metricMap["attributes"].([]any)[0].(map[string]any))
	if err != nil {
		return err
	}
	if tags, err := helper.ParseFromMaps[general_objects.Tag](metricMap["tags"].(*schema.Set).List()); err != nil {
		return err
	} else {
		metric.Tags = tags
	}
	return nil
}
