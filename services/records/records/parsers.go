package records

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

func (record *Record) ToMap() map[string]any {
	resourceMap := make(map[string]any)
	resourceMap["id"] = record.ID
	resourceMap["record_template_id"] = record.RecordTemplateId
	resourceMap["name"] = record.Name
	resourceMap["record_state"] = record.RecordState
	resourceMap["start_date_time"] = record.StartDateTime
	resourceMap["stop_date_time"] = record.StopDateTime
	resourceMap["stream_id"] = record.StreamId
	resourceMap["node_ids"] = record.NodeIds
	resourceMap["metric_ids"] = record.MetricIds
	if record.Properties != nil {
		resourceMap["properties"] = helper.ParseToMaps(record.Properties)
	}
	resourceMap["command_definition_ids"] = record.CommandDefinitionIds
	resourceMap["tags"] = helper.ParseToMaps(record.Tags)
	resourceMap["comments"] = record.Comments
	resourceMap["created_at"] = record.CreatedAt
	resourceMap["created_by"] = record.CreatedBy
	resourceMap["last_modified_at"] = record.LastModifiedAt
	resourceMap["last_modified_by"] = record.LastModifiedBy

	return resourceMap
}

func (property *Property[T]) ToMap() map[string]any {
	propertyMap := make(map[string]any)
	propertyMap["name"] = property.Name
	propertyMap["attributes"] = []any{property.Attributes.ToMap()}
	return propertyMap
}

func (record *Record) FromMap(resourceMap map[string]any) error {
	record.ID = resourceMap["id"].(string)
	record.RecordTemplateId = resourceMap["record_template_id"].(string)
	record.Name = resourceMap["name"].(string)
	record.RecordState = resourceMap["record_state"].(string)
	record.StartDateTime = resourceMap["start_date_time"].(string)
	record.StopDateTime = resourceMap["stop_date_time"].(string)
	if resourceMap["properties"] != nil {
		if properties, err := helper.ParseFromMaps[Property[any]](
			resourceMap["properties"].(*schema.Set).List(),
		); err != nil {
			return err
		} else {
			record.Properties = properties
		}
	}
	if tags, err := helper.ParseFromMaps[general_objects.KeyValue](resourceMap["tags"].(*schema.Set).List()); err != nil {
		return err
	} else {
		record.Tags = tags
	}
	record.CreatedAt = resourceMap["created_at"].(string)
	record.CreatedBy = resourceMap["created_by"].(string)
	record.LastModifiedAt = resourceMap["last_modified_at"].(string)
	record.LastModifiedBy = resourceMap["last_modified_by"].(string)

	return nil
}

func (property *Property[T]) FromMap(propertyMap map[string]any) error {
	property.Name = propertyMap["name"].(string)
	if len(propertyMap["attributes"].([]any)) > 0 {
		if err := property.Attributes.FromMap(propertyMap["attributes"].([]any)[0].(map[string]any)); err != nil {
			return err
		}
	}
	return nil
}
