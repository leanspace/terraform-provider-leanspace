package record_templates

import (
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
	"github.com/leanspace/terraform-provider-leanspace/provider"
)

func (recordTemplate *RecordTemplate) ToMap() map[string]any {
	resourceMap := recordTemplate.ToAuditMap()
	resourceMap["name"] = helper.NilIfEmpty(recordTemplate.Name)
	resourceMap["description"] = helper.NilIfEmpty(recordTemplate.Description)
	resourceMap["stream_id"] = helper.NilIfEmpty(recordTemplate.StreamId)
	if recordTemplate.DefaultParsers != nil {
		resourceMap["default_parsers"] = helper.ParseToMaps(recordTemplate.DefaultParsers)
	}
	resourceMap["node_ids"] = helper.NilIfEmpty(recordTemplate.NodeIds)
	resourceMap["metric_ids"] = helper.NilIfEmpty(recordTemplate.MetricIds)
	resourceMap["command_definition_ids"] = helper.NilIfEmpty(recordTemplate.CommandDefinitionIds)
	if recordTemplate.Properties != nil {
		resourceMap["properties"] = helper.ParseToMaps(recordTemplate.Properties)
	}
	resourceMap["tags"] = helper.ParseToMaps(recordTemplate.Tags)

	return resourceMap
}

func (defaultParser *DefaultParser) ToMap() map[string]any {
	defaultParserMap := make(map[string]any)
	defaultParserMap["id"] = helper.NilIfEmpty(defaultParser.ID)
	defaultParserMap["file_type"] = []any{defaultParser.FileType}
	return defaultParserMap
}

func (property *Property[T]) ToMap() map[string]any {
	propertyMap := make(map[string]any)
	propertyMap["name"] = helper.NilIfEmpty(property.Name)
	propertyMap["attributes"] = property.Attributes.ToMap()
	return propertyMap
}

func (recordTemplate *RecordTemplate) FromMap(resourceMap map[string]any) error {
	recordTemplate.FromAuditMap(resourceMap)
	recordTemplate.Name = helper.CastString(resourceMap, "name")
	recordTemplate.Description = helper.CastString(resourceMap, "description")
	if resourceMap["default_parsers"] != nil {
		if defaultParsers, err := helper.ParseFromMaps[DefaultParser](
			helper.CastSlice(resourceMap, "default_parsers"),
		); err != nil {
			return err
		} else {
			recordTemplate.DefaultParsers = defaultParsers
		}
	}
	recordTemplate.StreamId = helper.CastString(resourceMap, "stream_id")
	recordTemplate.NodeIds = make([]string, len(helper.CastSlice(resourceMap, "node_ids")))
	for index, value := range helper.CastSlice(resourceMap, "node_ids") {
		recordTemplate.NodeIds[index] = value.(string)
	}
	recordTemplate.MetricIds = make([]string, len(helper.CastSlice(resourceMap, "metric_ids")))
	for index, value := range helper.CastSlice(resourceMap, "metric_ids") {
		recordTemplate.MetricIds[index] = value.(string)
	}
	recordTemplate.CommandDefinitionIds = make([]string, len(helper.CastSlice(resourceMap, "command_definition_ids")))
	for index, value := range helper.CastSlice(resourceMap, "command_definition_ids") {
		recordTemplate.CommandDefinitionIds[index] = value.(string)
	}
	if resourceMap["properties"] != nil {
		if properties, err := helper.ParseFromMaps[Property[any]](
			helper.CastSlice(resourceMap, "properties"),
		); err != nil {
			return err
		} else {
			recordTemplate.Properties = properties
		}
	}
	if tags, err := helper.ParseFromMaps[general_objects.KeyValue](helper.CastSlice(resourceMap, "tags")); err != nil {
		return err
	} else {
		recordTemplate.Tags = tags
	}
	return nil
}

func (defaultParser *DefaultParser) FromMap(defaultParserMap map[string]any) error {
	defaultParser.ID = helper.CastString(defaultParserMap, "id")
	defaultParser.FileType = helper.CastString(defaultParserMap, "file_type")
	return nil
}

func (property *Property[T]) FromMap(propertyMap map[string]any) error {
	property.Name = helper.CastString(propertyMap, "name")
	if err := property.Attributes.FromMap(helper.CastMapAny(propertyMap, "attributes")); err != nil {
		return err
	}
	return nil
}

// PostReadProcess reorders properties in the API response to match the prior
// state order (matched by property name), preventing perpetual diffs when the API
// returns properties in a different order than they were configured.
// Note: PostUnmarshallProcess has already rebuilt Property from Properties by this point.
func (recordTemplate *RecordTemplate) PostReadProcess(_ *provider.Client, newValue any) error {
	newRecordTemplate, ok := newValue.(*RecordTemplate)
	if !ok || newRecordTemplate == nil {
		return nil
	}
	newRecordTemplate.Properties = helper.ReorderByKey(recordTemplate.Properties, newRecordTemplate.Properties, func(p Property[any]) string { return p.Name })
	return nil
}
