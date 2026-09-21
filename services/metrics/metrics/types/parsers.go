package types

import (
	generalobjectstypes "github.com/leanspace/terraform-provider-leanspace/helper/general_objects/types"
)

// listable is satisfied by *schema.Set (from terraform-plugin-sdk) without
// importing it: hashicorp's Set.List() has this exact signature. Using
// structural typing here keeps this package free of any SDK dependency
// while still accepting the raw *schema.Set values terraform-plugin-sdk
// puts in ResourceData-derived maps for TypeSet fields (e.g. "tags").
type listable interface {
	List() []any
}

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
	tags := make([]map[string]any, len(metric.Tags))
	for i := range metric.Tags {
		tags[i] = metric.Tags[i].ToMap()
	}
	metricMap["tags"] = tags
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
	if len(metricMap["attributes"].([]any)) > 0 {
		if err := metric.Attributes.FromMap(metricMap["attributes"].([]any)[0].(map[string]any)); err != nil {
			return err
		}
	}

	var tagsList []any
	switch tags := metricMap["tags"].(type) {
	case []any:
		tagsList = tags
	case listable:
		tagsList = tags.List()
	}
	metric.Tags = make([]generalobjectstypes.KeyValue, len(tagsList))
	for i, tag := range tagsList {
		if err := metric.Tags[i].FromMap(tag.(map[string]any)); err != nil {
			return err
		}
	}
	return nil
}
