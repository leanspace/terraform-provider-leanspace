package activity_definitions

import (
	"github.com/leanspace/terraform-provider-leanspace/helper"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func (activityDefinition *ActivityDefinition) ToMap() map[string]any {
	actDefinitionMap := make(map[string]any)
	actDefinitionMap["id"] = activityDefinition.ID
	actDefinitionMap["node_id"] = activityDefinition.NodeId
	actDefinitionMap["name"] = activityDefinition.Name
	actDefinitionMap["description"] = activityDefinition.Description
	actDefinitionMap["estimated_duration"] = activityDefinition.EstimatedDuration
	actDefinitionMap["mapping_status"] = activityDefinition.MappingStatus
	actDefinitionMap["created_at"] = activityDefinition.CreatedAt
	actDefinitionMap["created_by"] = activityDefinition.CreatedBy
	actDefinitionMap["last_modified_at"] = activityDefinition.LastModifiedAt
	actDefinitionMap["last_modified_by"] = activityDefinition.LastModifiedBy
	if activityDefinition.Metadata != nil {
		actDefinitionMap["metadata"] = helper.ParseToMaps(activityDefinition.Metadata)
	}
	if activityDefinition.ArgumentDefinitions != nil {
		actDefinitionMap["argument_definitions"] = helper.ParseToMaps(activityDefinition.ArgumentDefinitions)
	}
	if activityDefinition.CommandMappings != nil {
		actDefinitionMap["command_mappings"] = helper.ParseToMaps(activityDefinition.CommandMappings)
	}
	return actDefinitionMap
}

func (metadata *Metadata[T]) ToMap() map[string]any {
	metadataMap := make(map[string]any)
	metadataMap["id"] = metadata.ID
	metadataMap["name"] = metadata.Name
	metadataMap["description"] = metadata.Description
	metadataMap["attributes"] = []any{metadata.Attributes.ToMap()}
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
	commandMappingMap["argument_mappings"] = helper.ParseToMaps(commandMapping.ArgumentMappings)
	commandMappingMap["metadata_mappings"] = helper.ParseToMaps(commandMapping.MetadataMappings)
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
	activityDefinition.MappingStatus = actDefinitionMap["mapping_status"].(string)
	activityDefinition.CreatedAt = actDefinitionMap["created_at"].(string)
	activityDefinition.CreatedBy = actDefinitionMap["created_by"].(string)
	activityDefinition.LastModifiedAt = actDefinitionMap["last_modified_at"].(string)
	activityDefinition.LastModifiedBy = actDefinitionMap["last_modified_by"].(string)
	if actDefinitionMap["metadata"] != nil {
		if metadata, err := helper.ParseFromMaps[Metadata[any]](
			actDefinitionMap["metadata"].(*schema.Set).List(),
		); err != nil {
			return err
		} else {
			activityDefinition.Metadata = metadata
		}
	}
	if actDefinitionMap["argument_definitions"] != nil {
		if argumentDefinitions, err := helper.ParseFromMaps[ArgumentDefinition[any]](
			actDefinitionMap["argument_definitions"].(*schema.Set).List(),
		); err != nil {
			return err
		} else {
			activityDefinition.ArgumentDefinitions = argumentDefinitions
		}
	}
	if actDefinitionMap["command_mappings"] != nil {
		if commandMappings, err := helper.ParseFromMaps[CommandMapping](
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
	if len(metadataMap["attributes"].([]any)) > 0 {
		if err := metadata.Attributes.FromMap(metadataMap["attributes"].([]any)[0].(map[string]any)); err != nil {
			return err
		}
	}
	return nil
}

func (argument *ArgumentDefinition[T]) FromMap(argumentMap map[string]any) error {
	argument.ID = argumentMap["id"].(string)
	argument.Name = argumentMap["name"].(string)
	argument.Description = argumentMap["description"].(string)

	if len(argumentMap["attributes"].([]any)) > 0 {
		if err := argument.Attributes.FromMap(argumentMap["attributes"].([]any)[0].(map[string]any)); err != nil {
			return err
		}
	}
	return nil
}

func (commandMapping *CommandMapping) FromMap(commandMappingMap map[string]any) error {
	commandMapping.CommandDefinitionId = commandMappingMap["command_definition_id"].(string)
	commandMapping.Position = commandMappingMap["position"].(int)
	commandMapping.DelayInMilliseconds = commandMappingMap["delay_in_milliseconds"].(int)
	if argumentMappings, err := helper.ParseFromMaps[ArgumentMapping](commandMappingMap["argument_mappings"].(*schema.Set).List()); err != nil {
		return err
	} else {
		commandMapping.ArgumentMappings = argumentMappings
	}
	if metadataMappings, err := helper.ParseFromMaps[MetadataMapping](commandMappingMap["metadata_mappings"].(*schema.Set).List()); err != nil {
		return err
	} else {
		commandMapping.MetadataMappings = metadataMappings
	}
	return nil
}

func (argumentMapping *ArgumentMapping) FromMap(argumentMappingMap map[string]any) error {
	argumentMapping.ActivityDefinitionArgumentName = argumentMappingMap["activity_definition_argument_name"].(string)
	argumentMapping.CommandDefinitionArgumentName = argumentMappingMap["command_definition_argument_name"].(string)
	argumentMapping.MappingStatus = argumentMappingMap["mapping_status"].(string)
	return nil
}

func (metadataMapping *MetadataMapping) FromMap(metadataMappingMap map[string]any) error {
	metadataMapping.ActivityDefinitionMetadataName = metadataMappingMap["activity_definition_metadata_name"].(string)
	metadataMapping.CommandDefinitionArgumentName = metadataMappingMap["command_definition_argument_name"].(string)
	metadataMapping.MappingStatus = metadataMappingMap["mapping_status"].(string)
	return nil
}

func (activityDefinition *ActivityDefinition) PreMarshallProcess() error {
	for i := range activityDefinition.CommandMappings {
		activityDefinition.CommandMappings[i].Position = i
	}
	return nil
}
