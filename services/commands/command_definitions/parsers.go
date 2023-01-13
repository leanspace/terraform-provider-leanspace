package command_definitions

import (
	"github.com/leanspace/terraform-provider-leanspace/helper"

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
		commandDefinitionMap["metadata"] = helper.ParseToMaps(commandDefinition.Metadata)
	}
	if commandDefinition.Arguments != nil {
		commandDefinitionMap["arguments"] = helper.ParseToMaps(commandDefinition.Arguments)
	}

	return commandDefinitionMap
}

func (metadata Metadata[T]) ToMap() map[string]any {
	metadataMap := make(map[string]any)
	metadataMap["id"] = metadata.ID
	metadataMap["name"] = metadata.Name
	metadataMap["description"] = metadata.Description
	metadataMap["attributes"] = []any{metadata.Attributes.ToMap()}
	return metadataMap
}

func (argument Argument[T]) ToMap() map[string]any {
	argumentMap := make(map[string]any)
	argumentMap["id"] = argument.ID
	argumentMap["name"] = argument.Name
	argumentMap["identifier"] = argument.Identifier
	argumentMap["description"] = argument.Description
	argumentMap["attributes"] = []any{argument.Attributes.ToMap()}
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
		if metadata, err := helper.ParseFromMaps[Metadata[any]](
			cmdDefinitionMap["metadata"].(*schema.Set).List(),
		); err != nil {
			return err
		} else {
			commandDefinition.Metadata = metadata
		}
	}
	if cmdDefinitionMap["arguments"] != nil {

		if arguments, err := helper.ParseFromMaps[Argument[any]](
			cmdDefinitionMap["arguments"].(*schema.Set).List(),
		); err != nil {
			return err
		} else {
			commandDefinition.Arguments = arguments
		}
	}

	return nil
}

func (metadata *Metadata[T]) FromMap(metadataMap map[string]any) error {
	metadata.ID = metadataMap["id"].(string)
	metadata.Name = metadataMap["name"].(string)
	metadata.Description = metadataMap["description"].(string)
	if len(metadataMap["attributes"].([]any)) > 0 {
		if err := metadata.Attributes.FromMap(metadataMap["attributes"].([]any)[0].(map[string]any)); err != nil {
			return err
		}
	}
	return nil
}

func (argument *Argument[T]) FromMap(argumentMap map[string]any) error {
	argument.ID = argumentMap["id"].(string)
	argument.Name = argumentMap["name"].(string)
	argument.Identifier = argumentMap["identifier"].(string)
	argument.Description = argumentMap["description"].(string)
	if len(argumentMap["attributes"].([]any)) > 0 {
		if err := argument.Attributes.FromMap(argumentMap["attributes"].([]any)[0].(map[string]any)); err != nil {
			return err
		}
	}
	return nil
}
