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
	resourceMap["record_state"] = recordTemplate.RecordState
	resourceMap["start_date_time"] = recordTemplate.StartDateTime
	resourceMap["stop_date_time"] = recordTemplate.StopDateTime
	if recordTemplate.DefaultParsers != nil {
		resourceMap["default_parsers"] = helper.ParseToMaps(recordTemplate.DefaultParsers)
	}
	if recordTemplate.Nodes != nil {
		resourceMap["nodes"] = helper.ParseToMaps(recordTemplate.Nodes)
	}
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
	// TODO
	return defaultParserMap
}

func (node *Node) ToMap() map[string]any {
	nodeMap := make(map[string]any)
	// TODO
	return nodeMap
}

func (property *Property) ToMap() map[string]any {
	propertyMap := make(map[string]any)
	// TODO
	return propertyMap
}

func (recordTemplate *RecordTemplate) FromMap(resourceMap map[string]any) error {
	recordTemplate.ID = resourceMap["id"].(string)
	recordTemplate.Name = resourceMap["name"].(string)
	recordTemplate.Description = resourceMap["description"].(string)
	recordTemplate.RecordState = resourceMap["record_state"].(string)
	recordTemplate.StartDateTime = resourceMap["start_date_time"].(string)
	recordTemplate.StopDateTime = resourceMap["stop_date_time"].(string)
	if resourceMap["default_parsers"] != nil {
		if defaultParsers, err := helper.ParseFromMaps[DefaultParser](
			resourceMap["default_parsers"].(*schema.Set).List(),
		); err != nil {
			return err
		} else {
			recordTemplate.DefaultParsers = defaultParsers
		}
	}
	if resourceMap["nodes"] != nil {
		if nodeSnapshots, err := helper.ParseFromMaps[Node](
			resourceMap["nodes"].(*schema.Set).List(),
		); err != nil {
			return err
		} else {
			recordTemplate.Nodes = nodeSnapshots
		}
	}
	if resourceMap["properties"] != nil {
		if properties, err := helper.ParseFromMaps[Property](
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
	// TODO
	return nil
}

func (node *Node) FromMap(nodeMap map[string]any) error {
	// TODO
	return nil
}

func (property *Property) FromMap(propertyMap map[string]any) error {
	// TODO
	return nil
}
