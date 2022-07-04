package metrics

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

	return err
}
