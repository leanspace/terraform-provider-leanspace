package record_templates

import (
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func (recordTemplate *RecordTemplate) ToMap() map[string]any {
	resourceMap := make(map[string]any)
	resourceMap["id"] = recordTemplate.ID
	resourceMap["name"] = recordTemplate.Name
	resourceMap["description"] = recordTemplate.Description
	resourceMap["stream_id"] = recordTemplate.StreamId
	if recordTemplate.DefaultParsers != nil {
		resourceMap["default_parsers"] = helper.ParseToMaps(recordTemplate.DefaultParsers)
	}
	resourceMap["node_ids"] = recordTemplate.NodeIds
	resourceMap["metric_ids"] = recordTemplate.MetricIds
	resourceMap["command_definition_ids"] = recordTemplate.CommandDefinitionIds
	if recordTemplate.Properties != nil {
		resourceMap["properties"] = helper.ParseToMaps(recordTemplate.Properties)
	}
	resourceMap["tags"] = helper.ParseToMaps(recordTemplate.Tags)
	resourceMap["created_at"] = recordTemplate.CreatedAt
	resourceMap["created_by"] = recordTemplate.CreatedBy
	resourceMap["last_modified_at"] = recordTemplate.LastModifiedAt
	resourceMap["last_modified_by"] = recordTemplate.LastModifiedBy

	return resourceMap
}

func (defaultParser *DefaultParser) ToMap() map[string]any {
	defaultParserMap := make(map[string]any)
	defaultParserMap["id"] = defaultParser.ID
	defaultParserMap["file_type"] = []any{defaultParser.FileType}
	return defaultParserMap
}

func (property *Property[T]) ToMap() map[string]any {
	propertyMap := make(map[string]any)
	propertyMap["name"] = property.Name
	propertyMap["attributes"] = []any{property.Attributes.ToMap()}
	return propertyMap
}

func (recordTemplate *RecordTemplate) FromMap(resourceMap map[string]any) error {
	recordTemplate.ID = resourceMap["id"].(string)
	recordTemplate.Name = resourceMap["name"].(string)
	recordTemplate.Description = resourceMap["description"].(string)
	if resourceMap["default_parsers"] != nil {
		if defaultParsers, err := helper.ParseFromMaps[DefaultParser](
			resourceMap["default_parsers"].(*schema.Set).List(),
		); err != nil {
			return err
		} else {
			recordTemplate.DefaultParsers = defaultParsers
		}
	}
	recordTemplate.StreamId = resourceMap["stream_id"].(string)
	recordTemplate.NodeIds = make([]string, resourceMap["node_ids"].(*schema.Set).Len())
	for index, value := range resourceMap["node_ids"].(*schema.Set).List() {
		recordTemplate.NodeIds[index] = value.(string)
	}
	recordTemplate.MetricIds = make([]string, resourceMap["metric_ids"].(*schema.Set).Len())
	for index, value := range resourceMap["metric_ids"].(*schema.Set).List() {
		recordTemplate.MetricIds[index] = value.(string)
	}
	recordTemplate.CommandDefinitionIds = make([]string, resourceMap["command_definition_ids"].(*schema.Set).Len())
	for index, value := range resourceMap["command_definition_ids"].(*schema.Set).List() {
		recordTemplate.CommandDefinitionIds[index] = value.(string)
	}
	if resourceMap["properties"] != nil {
		if properties, err := helper.ParseFromMaps[Property[any]](
			resourceMap["properties"].(*schema.Set).List(),
		); err != nil {
			return err
		} else {
			recordTemplate.Properties = properties
		}
	}
	if tags, err := helper.ParseFromMaps[general_objects.KeyValue](resourceMap["tags"].(*schema.Set).List()); err != nil {
		return err
	} else {
		recordTemplate.Tags = tags
	}
	recordTemplate.CreatedAt = resourceMap["created_at"].(string)
	recordTemplate.CreatedBy = resourceMap["created_by"].(string)
	recordTemplate.LastModifiedAt = resourceMap["last_modified_at"].(string)
	recordTemplate.LastModifiedBy = resourceMap["last_modified_by"].(string)

	return nil
}

func (defaultParser *DefaultParser) FromMap(defaultParserMap map[string]any) error {
	defaultParser.ID = defaultParserMap["id"].(string)
	defaultParser.FileType = defaultParserMap["file_type"].(string)
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
