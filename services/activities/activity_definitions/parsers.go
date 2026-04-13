package activity_definitions

import (
	"fmt"

	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
	"github.com/leanspace/terraform-provider-leanspace/provider"
)

func (activityDefinition *ActivityDefinition) ToMap() map[string]any {
	actDefinitionMap := activityDefinition.ToAuditMap()
	actDefinitionMap["node_id"] = helper.NilIfEmpty(activityDefinition.NodeId)
	actDefinitionMap["name"] = helper.NilIfEmpty(activityDefinition.Name)
	actDefinitionMap["description"] = helper.NilIfEmpty(activityDefinition.Description)
	actDefinitionMap["estimated_duration"] = helper.NilIfEmpty(activityDefinition.EstimatedDuration)
	actDefinitionMap["mapping_status"] = helper.NilIfEmpty(activityDefinition.MappingStatus)
	actDefinitionMap["tags"] = helper.ParseToMaps(activityDefinition.Tags)
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
	metadataMap["name"] = helper.NilIfEmpty(metadata.Name)
	metadataMap["description"] = helper.NilIfEmpty(metadata.Description)
	metadataMap["attributes"] = metadata.Attributes.ToMap()
	return metadataMap
}

func (argument *ArgumentDefinition[T]) ToMap() map[string]any {
	argumentMap := make(map[string]any)
	argumentMap["name"] = helper.NilIfEmpty(argument.Name)
	argumentMap["description"] = helper.NilIfEmpty(argument.Description)
	argumentMap["attributes"] = argument.Attributes.ToMap()
	return argumentMap
}

func (commandMapping *CommandMapping) ToMap() map[string]any {
	commandMappingMap := make(map[string]any)
	commandMappingMap["command_definition_id"] = helper.NilIfEmpty(commandMapping.CommandDefinitionId)
	commandMappingMap["position"] = helper.NilIfEmpty(commandMapping.Position)
	commandMappingMap["delay_in_milliseconds"] = helper.NilIfEmpty(commandMapping.DelayInMilliseconds)
	commandMappingMap["argument_mappings"] = helper.ParseToMaps(commandMapping.ArgumentMappings)
	commandMappingMap["metadata_mappings"] = helper.ParseToMaps(commandMapping.MetadataMappings)
	return commandMappingMap
}

func (argumentMapping *ArgumentMapping) ToMap() map[string]any {
	argumentMappingMap := make(map[string]any)
	argumentMappingMap["activity_definition_argument_name"] = helper.NilIfEmpty(argumentMapping.ActivityDefinitionArgumentName)
	argumentMappingMap["command_definition_argument_name"] = helper.NilIfEmpty(argumentMapping.CommandDefinitionArgumentName)
	return argumentMappingMap
}

func (metadataMapping *MetadataMapping) ToMap() map[string]any {
	metadataMappingMap := make(map[string]any)
	metadataMappingMap["activity_definition_metadata_name"] = helper.NilIfEmpty(metadataMapping.ActivityDefinitionMetadataName)
	metadataMappingMap["command_definition_argument_name"] = helper.NilIfEmpty(metadataMapping.CommandDefinitionArgumentName)
	return metadataMappingMap
}

func (activityDefinition *ActivityDefinition) FromMap(actDefinitionMap map[string]any) error {
	activityDefinition.FromAuditMap(actDefinitionMap)
	activityDefinition.NodeId = helper.CastString(actDefinitionMap, "node_id")
	activityDefinition.Name = helper.CastString(actDefinitionMap, "name")
	activityDefinition.EstimatedDuration = helper.CastInt(actDefinitionMap, "estimated_duration")
	activityDefinition.Description = helper.CastString(actDefinitionMap, "description")
	activityDefinition.MappingStatus = helper.CastString(actDefinitionMap, "mapping_status")
	if tags, err := helper.ParseFromMaps[general_objects.KeyValue](helper.CastSlice(actDefinitionMap, "tags")); err != nil {
		return err
	} else {
		activityDefinition.Tags = tags
	}
	if actDefinitionMap["metadata"] != nil {
		if metadata, err := helper.ParseFromMaps[Metadata[any]](
			helper.CastSlice(actDefinitionMap, "metadata"),
		); err != nil {
			return err
		} else {
			activityDefinition.Metadata = metadata
		}
	}
	if actDefinitionMap["argument_definitions"] != nil {
		if argumentDefinitions, err := helper.ParseFromMaps[ArgumentDefinition[any]](
			helper.CastSlice(actDefinitionMap, "argument_definitions"),
		); err != nil {
			return err
		} else {
			activityDefinition.ArgumentDefinitions = argumentDefinitions
		}
	}
	if actDefinitionMap["command_mappings"] != nil {
		if commandMappings, err := helper.ParseFromMaps[CommandMapping](
			helper.CastSlice(actDefinitionMap, "command_mappings"),
		); err != nil {
			return err
		} else {
			activityDefinition.CommandMappings = commandMappings
		}
	}
	return nil
}

func (metadata *Metadata[T]) FromMap(metadataMap map[string]any) error {
	metadata.Name = helper.CastString(metadataMap, "name")
	metadata.Description = helper.CastString(metadataMap, "description")
	if err := metadata.Attributes.FromMap(helper.CastMapAny(metadataMap, "attributes")); err != nil {
		return err
	}
	return nil
}

func (argument *ArgumentDefinition[T]) FromMap(argumentMap map[string]any) error {
	argument.Name = helper.CastString(argumentMap, "name")
	argument.Description = helper.CastString(argumentMap, "description")

	if err := argument.Attributes.FromMap(helper.CastMapAny(argumentMap, "attributes")); err != nil {
		return err
	}
	return nil
}

func (commandMapping *CommandMapping) FromMap(commandMappingMap map[string]any) error {
	commandMapping.CommandDefinitionId = helper.CastString(commandMappingMap, "command_definition_id")
	commandMapping.Position = helper.CastInt(commandMappingMap, "position")
	commandMapping.DelayInMilliseconds = helper.CastInt(commandMappingMap, "delay_in_milliseconds")
	if argumentMappings, err := helper.ParseFromMaps[ArgumentMapping](helper.CastSlice(commandMappingMap, "argument_mappings")); err != nil {
		return err
	} else {
		commandMapping.ArgumentMappings = argumentMappings
	}
	if metadataMappings, err := helper.ParseFromMaps[MetadataMapping](helper.CastSlice(commandMappingMap, "metadata_mappings")); err != nil {
		return err
	} else {
		commandMapping.MetadataMappings = metadataMappings
	}
	return nil
}

func (argumentMapping *ArgumentMapping) FromMap(argumentMappingMap map[string]any) error {
	argumentMapping.ActivityDefinitionArgumentName = helper.CastString(argumentMappingMap, "activity_definition_argument_name")
	argumentMapping.CommandDefinitionArgumentName = helper.CastString(argumentMappingMap, "command_definition_argument_name")
	argumentMapping.MappingStatus = helper.CastString(argumentMappingMap, "mapping_status")
	return nil
}

func (metadataMapping *MetadataMapping) FromMap(metadataMappingMap map[string]any) error {
	metadataMapping.ActivityDefinitionMetadataName = helper.CastString(metadataMappingMap, "activity_definition_metadata_name")
	metadataMapping.CommandDefinitionArgumentName = helper.CastString(metadataMappingMap, "command_definition_argument_name")
	metadataMapping.MappingStatus = helper.CastString(metadataMappingMap, "mapping_status")
	return nil
}

func (activityDefinition *ActivityDefinition) PreMarshallProcess() error {
	for i := range activityDefinition.CommandMappings {
		activityDefinition.CommandMappings[i].Position = i
	}
	return nil
}

func (activityDefinition *ActivityDefinition) PostReadProcess(_ *provider.Client, newValue any) error {
	newDef, ok := newValue.(*ActivityDefinition)
	if !ok || newDef == nil {
		return nil
	}
	newDef.Metadata = helper.ReorderByKey(activityDefinition.Metadata, newDef.Metadata,
		func(m Metadata[any]) string { return m.Name })
	newDef.ArgumentDefinitions = helper.ReorderByKey(activityDefinition.ArgumentDefinitions, newDef.ArgumentDefinitions,
		func(a ArgumentDefinition[any]) string { return a.Name })
	newDef.CommandMappings = helper.ReorderByKey(activityDefinition.CommandMappings, newDef.CommandMappings,
		func(c CommandMapping) string { return fmt.Sprintf("%s:%d", c.CommandDefinitionId, c.Position) })
	// Reorder mappings within each command mapping to match state order.
	for i, stateCmd := range activityDefinition.CommandMappings {
		if i >= len(newDef.CommandMappings) {
			break
		}
		newDef.CommandMappings[i].ArgumentMappings = helper.ReorderByKey(
			stateCmd.ArgumentMappings, newDef.CommandMappings[i].ArgumentMappings,
			func(m ArgumentMapping) string { return m.ActivityDefinitionArgumentName },
		)
		newDef.CommandMappings[i].MetadataMappings = helper.ReorderByKey(
			stateCmd.MetadataMappings, newDef.CommandMappings[i].MetadataMappings,
			func(m MetadataMapping) string { return m.ActivityDefinitionMetadataName },
		)
	}
	return nil
}
