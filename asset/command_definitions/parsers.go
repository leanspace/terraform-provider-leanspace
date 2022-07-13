package command_definitions

import (
	"strconv"
	"terraform-provider-asset/asset"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func (commandDefinition *CommandDefinition) ToMap() map[string]any {
	commandDefinitionMap := make(map[string]any)
	commandDefinitionMap["id"] = commandDefinition.ID
	commandDefinitionMap["node_id"] = commandDefinition.NodeId
	commandDefinitionMap["name"] = commandDefinition.Name
	commandDefinitionMap["description"] = commandDefinition.Description
	commandDefinitionMap["identifier"] = commandDefinition.Identifier
	commandDefinitionMap["created_at"] = commandDefinition.CreatedAt
	commandDefinitionMap["created_by"] = commandDefinition.CreatedBy
	commandDefinitionMap["last_modified_at"] = commandDefinition.LastModifiedAt
	commandDefinitionMap["last_modified_by"] = commandDefinition.LastModifiedBy
	if commandDefinition.Metadata != nil {
		commandDefinitionMap["metadata"] = asset.ParseToMaps(commandDefinition.Metadata)
	}
	if commandDefinition.Arguments != nil {
		commandDefinitionMap["arguments"] = asset.ParseToMaps(commandDefinition.Arguments)
	}

	return commandDefinitionMap
}

func (metadata Metadata[T]) ToMap() map[string]any {
	metadataMap := make(map[string]any)
	metadataMap["id"] = metadata.ID
	metadataMap["name"] = metadata.Name
	metadataMap["description"] = metadata.Description

	attributes := make(map[string]any)

	attributes["type"] = metadata.Attributes.Type
	if metadata.Attributes.Type == "NUMERIC" {
		attributes["unit_id"] = metadata.Attributes.UnitId
	}
	switch metadata.Attributes.Type {
	case "NUMERIC":
		attributes["value"] = strconv.FormatFloat(any(metadata.Attributes.Value).(float64), 'g', -1, 64)
	case "TEXT":
		attributes["value"] = metadata.Attributes.Value
	case "BOOLEAN":
		attributes["value"] = strconv.FormatBool(any(metadata.Attributes.Value).(bool))
	case "TIMESTAMP", "DATE", "TIME":
		attributes["value"] = metadata.Attributes.Value
	}
	metadataMap["attributes"] = []map[string]any{attributes}

	return metadataMap
}

func (argument Argument[T]) ToMap() map[string]any {
	argumentMap := make(map[string]any)
	argumentMap["id"] = argument.ID
	argumentMap["name"] = argument.Name
	argumentMap["identifier"] = argument.Identifier
	argumentMap["description"] = argument.Description
	argumentMap["attributes"] = argument.Attributes.ToMap()

	return argumentMap
}

func (commandDefinition *CommandDefinition) FromMap(cmdDefinitionMap map[string]any) error {
	commandDefinition.ID = cmdDefinitionMap["id"].(string)
	commandDefinition.NodeId = cmdDefinitionMap["node_id"].(string)
	commandDefinition.Name = cmdDefinitionMap["name"].(string)
	commandDefinition.Description = cmdDefinitionMap["description"].(string)
	commandDefinition.Identifier = cmdDefinitionMap["identifier"].(string)
	commandDefinition.CreatedAt = cmdDefinitionMap["created_at"].(string)
	commandDefinition.CreatedBy = cmdDefinitionMap["created_by"].(string)
	commandDefinition.LastModifiedAt = cmdDefinitionMap["last_modified_at"].(string)
	commandDefinition.LastModifiedBy = cmdDefinitionMap["last_modified_by"].(string)
	if cmdDefinitionMap["metadata"] != nil {
		commandDefinition.Metadata = asset.ParseFromMaps[Metadata[any]](
			cmdDefinitionMap["metadata"].(*schema.Set).List(),
		)
	}
	if cmdDefinitionMap["arguments"] != nil {
		commandDefinition.Arguments = asset.ParseFromMaps[Argument[any]](
			cmdDefinitionMap["arguments"].(*schema.Set).List(),
		)
	}

	return nil
}

func (metadata *Metadata[T]) FromMap(metadataMap map[string]any) error {
	metadata.ID = metadataMap["id"].(string)
	metadata.Name = metadataMap["name"].(string)
	metadata.Description = metadataMap["description"].(string)

	attributes := metadataMap["attributes"].([]any)[0].(map[string]any)
	metadata.Attributes.Value = attributes["value"].(T)
	metadata.Attributes.Type = attributes["type"].(string)
	if metadataMap["type"] == "NUMERIC" {
		metadata.Attributes.UnitId = attributes["unit_id"].(string)
	}

	return nil
}

func (argument *Argument[T]) FromMap(argumentMap map[string]any) error {
	argument.ID = argumentMap["id"].(string)
	argument.Name = argumentMap["name"].(string)
	argument.Identifier = argumentMap["identifier"].(string)
	argument.Description = argumentMap["description"].(string)

	attributeMap := argumentMap["attributes"].([]any)[0].(map[string]any)
	err := argument.Attributes.FromMap(attributeMap)
	return err
}
