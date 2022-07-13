package activity_definitions

import (
	"terraform-provider-asset/asset"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func (activityDefinition *ActivityDefinition) ToMap() map[string]any {
	actDefinitionMap := make(map[string]any)
	actDefinitionMap["id"] = activityDefinition.ID
	actDefinitionMap["node_id"] = activityDefinition.NodeId
	actDefinitionMap["name"] = activityDefinition.Name
	actDefinitionMap["description"] = activityDefinition.Description
	actDefinitionMap["estimated_duration"] = activityDefinition.EstimatedDuration
	actDefinitionMap["created_at"] = activityDefinition.CreatedAt
	actDefinitionMap["created_by"] = activityDefinition.CreatedBy
	actDefinitionMap["last_modified_at"] = activityDefinition.LastModifiedAt
	actDefinitionMap["last_modified_by"] = activityDefinition.LastModifiedBy
	if activityDefinition.Metadata != nil {
		actDefinitionMap["metadata"] = asset.ParseToMaps(activityDefinition.Metadata)
	}
	if activityDefinition.ArgumentDefinitions != nil {
		actDefinitionMap["argument_definitions"] = asset.ParseToMaps(activityDefinition.ArgumentDefinitions)
	}
	if activityDefinition.CommandMappings != nil {
		actDefinitionMap["command_mappings"] = asset.ParseToMaps(activityDefinition.CommandMappings)
	}
	return actDefinitionMap
}

func (metadata *Metadata[T]) ToMap() map[string]any {
	metadataMap := make(map[string]any)
	metadataMap["id"] = metadata.ID
	metadataMap["name"] = metadata.Name
	metadataMap["description"] = metadata.Description
	metadataMap["attributes"] = []map[string]any{metadata.Attributes.ToMap()}
	return metadataMap
}

func (argument *ArgumentDefinition[T]) ToMap() map[string]any {
	argumentMap := make(map[string]any)
	argumentMap["id"] = argument.ID
	argumentMap["name"] = argument.Name
	argumentMap["description"] = argument.Description
	argumentMap["attributes"] = []any{argument.Attributes.ToMap()}
	return argumentMap
}

func (commandMapping *CommandMapping) ToMap() map[string]any {
	commandMappingMap := make(map[string]any)
	commandMappingMap["command_definition_id"] = commandMapping.CommandDefinitionId
	commandMappingMap["position"] = commandMapping.Position
	commandMappingMap["delay_in_milliseconds"] = commandMapping.DelayInMilliseconds
	commandMappingMap["argument_mappings"] = asset.ParseToMaps(commandMapping.ArgumentMappings)
	commandMappingMap["metadata_mappings"] = asset.ParseToMaps(commandMapping.MetadataMappings)
	return commandMappingMap
}

func (argumentMapping *ArgumentMapping) ToMap() map[string]any {
	argumentMappingMap := make(map[string]any)
	argumentMappingMap["activity_definition_argument_name"] = argumentMapping.ActivityDefinitionArgumentName
	argumentMappingMap["command_definition_argument_name"] = argumentMapping.CommandDefinitionArgumentName
	return argumentMappingMap
}

func (metadataMapping *MetadataMapping) ToMap() map[string]any {
	argumentMappingMap := make(map[string]any)
	argumentMappingMap["activity_definition_metadata_name"] = metadataMapping.ActivityDefinitionMetadataName
	argumentMappingMap["command_definition_argument_name"] = metadataMapping.CommandDefinitionArgumentName
	return argumentMappingMap
}

func (activityDefinition *ActivityDefinition) FromMap(actDefinitionMap map[string]any) error {
	activityDefinition.ID = actDefinitionMap["id"].(string)
	activityDefinition.NodeId = actDefinitionMap["node_id"].(string)
	activityDefinition.Name = actDefinitionMap["name"].(string)
	activityDefinition.EstimatedDuration = actDefinitionMap["estimated_duration"].(int)
	activityDefinition.Description = actDefinitionMap["description"].(string)
	activityDefinition.CreatedAt = actDefinitionMap["created_at"].(string)
	activityDefinition.CreatedBy = actDefinitionMap["created_by"].(string)
	activityDefinition.LastModifiedAt = actDefinitionMap["last_modified_at"].(string)
	activityDefinition.LastModifiedBy = actDefinitionMap["last_modified_by"].(string)
	if actDefinitionMap["metadata"] != nil {
		if metadata, err := asset.ParseFromMaps[Metadata[any]](
			actDefinitionMap["metadata"].(*schema.Set).List(),
		); err != nil {
			return err
		} else {
			activityDefinition.Metadata = metadata
		}
	}
	if actDefinitionMap["argument_definitions"] != nil {
		if argumentDefinitions, err := asset.ParseFromMaps[ArgumentDefinition[any]](
			actDefinitionMap["argument_definitions"].(*schema.Set).List(),
		); err != nil {
			return err
		} else {
			activityDefinition.ArgumentDefinitions = argumentDefinitions
		}
	}
	if actDefinitionMap["command_mappings"] != nil {
		if commandMappings, err := asset.ParseFromMaps[CommandMapping](
			actDefinitionMap["command_mappings"].([]any),
		); err != nil {
			return err
		} else {
			activityDefinition.CommandMappings = commandMappings
		}
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

func (argument *ArgumentDefinition[T]) FromMap(argumentMap map[string]any) error {
	argument.ID = argumentMap["id"].(string)
	argument.Name = argumentMap["name"].(string)
	argument.Description = argumentMap["description"].(string)

	attributeMap := argumentMap["attributes"].([]any)[0].(map[string]any)
	err := argument.Attributes.FromMap(attributeMap)
	return err
}

func (commandMapping *CommandMapping) FromMap(commandMappingMap map[string]any) error {
	commandMapping.CommandDefinitionId = commandMappingMap["command_definition_id"].(string)
	commandMapping.Position = commandMappingMap["position"].(int)
	commandMapping.DelayInMilliseconds = commandMappingMap["delay_in_milliseconds"].(int)
	if argumentMappings, err := asset.ParseFromMaps[ArgumentMapping](commandMappingMap["argument_mappings"].(*schema.Set).List()); err != nil {
		return err
	} else {
		commandMapping.ArgumentMappings = argumentMappings
	}
	if metadataMappings, err := asset.ParseFromMaps[MetadataMapping](commandMappingMap["metadata_mappings"].(*schema.Set).List()); err != nil {
		return err
	} else {
		commandMapping.MetadataMappings = metadataMappings
	}
	return nil
}

func (argumentMapping *ArgumentMapping) FromMap(argumentMappingMap map[string]any) error {
	argumentMapping.ActivityDefinitionArgumentName = argumentMappingMap["activity_definition_argument_name"].(string)
	argumentMapping.CommandDefinitionArgumentName = argumentMappingMap["command_definition_argument_name"].(string)
	return nil
}

func (metadataMapping *MetadataMapping) FromMap(metadataMappingMap map[string]any) error {
	metadataMapping.ActivityDefinitionMetadataName = metadataMappingMap["activity_definition_metadata_name"].(string)
	metadataMapping.CommandDefinitionArgumentName = metadataMappingMap["command_definition_argument_name"].(string)
	return nil
}

func (activityDefinition *ActivityDefinition) PreMarshallProcess() error {
	for i := range activityDefinition.CommandMappings {
		activityDefinition.CommandMappings[i].Position = i
	}
	return nil
}
