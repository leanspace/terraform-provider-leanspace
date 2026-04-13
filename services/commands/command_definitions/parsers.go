package command_definitions

import (
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/provider"
)

func (commandDefinition *CommandDefinition) ToMap() map[string]any {
	commandDefinitionMap := commandDefinition.ToAuditMap()
	commandDefinitionMap["node_id"] = helper.NilIfEmpty(commandDefinition.NodeId)
	commandDefinitionMap["name"] = helper.NilIfEmpty(commandDefinition.Name)
	commandDefinitionMap["description"] = helper.NilIfEmpty(commandDefinition.Description)
	commandDefinitionMap["identifier"] = helper.NilIfEmpty(commandDefinition.Identifier)
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
	metadataMap["id"] = helper.NilIfEmpty(metadata.ID)
	metadataMap["name"] = helper.NilIfEmpty(metadata.Name)
	metadataMap["description"] = helper.NilIfEmpty(metadata.Description)
	metadataMap["attributes"] = metadata.Attributes.ToMap()
	return metadataMap
}

func (argument Argument[T]) ToMap() map[string]any {
	argumentMap := make(map[string]any)
	argumentMap["id"] = helper.NilIfEmpty(argument.ID)
	argumentMap["name"] = helper.NilIfEmpty(argument.Name)
	argumentMap["identifier"] = helper.NilIfEmpty(argument.Identifier)
	argumentMap["description"] = helper.NilIfEmpty(argument.Description)
	argumentMap["attributes"] = argument.Attributes.ToMap()
	return argumentMap
}

func (commandDefinition *CommandDefinition) FromMap(cmdDefinitionMap map[string]any) error {
	commandDefinition.FromAuditMap(cmdDefinitionMap)
	commandDefinition.NodeId = helper.CastString(cmdDefinitionMap, "node_id")
	commandDefinition.Name = helper.CastString(cmdDefinitionMap, "name")
	commandDefinition.Description = helper.CastString(cmdDefinitionMap, "description")
	commandDefinition.Identifier = helper.CastString(cmdDefinitionMap, "identifier")
	if cmdDefinitionMap["metadata"] != nil {
		if metadata, err := helper.ParseFromMaps[Metadata[any]](
			helper.CastSlice(cmdDefinitionMap, "metadata"),
		); err != nil {
			return err
		} else {
			commandDefinition.Metadata = metadata
		}
	}
	if cmdDefinitionMap["arguments"] != nil {

		if arguments, err := helper.ParseFromMaps[Argument[any]](
			helper.CastSlice(cmdDefinitionMap, "arguments"),
		); err != nil {
			return err
		} else {
			commandDefinition.Arguments = arguments
		}
	}

	return nil
}

func (metadata *Metadata[T]) FromMap(metadataMap map[string]any) error {
	metadata.ID = helper.CastString(metadataMap, "id")
	metadata.Name = helper.CastString(metadataMap, "name")
	metadata.Description = helper.CastString(metadataMap, "description")
	if err := metadata.Attributes.FromMap(helper.CastMapAny(metadataMap, "attributes")); err != nil {
		return err
	}
	return nil
}

func (argument *Argument[T]) FromMap(argumentMap map[string]any) error {
	argument.ID = helper.CastString(argumentMap, "id")
	argument.Name = helper.CastString(argumentMap, "name")
	argument.Identifier = helper.CastString(argumentMap, "identifier")
	argument.Description = helper.CastString(argumentMap, "description")
	if err := argument.Attributes.FromMap(helper.CastMapAny(argumentMap, "attributes")); err != nil {
		return err
	}
	return nil
}

// PostReadProcess reorders the API response's metadata and arguments to match the
// plan/state order (matched by name), preventing "unexpected new value" errors when
// the API returns elements in a different order.
func (commandDefinition *CommandDefinition) PostReadProcess(_ *provider.Client, newValue any) error {
	newCmdDef, ok := newValue.(*CommandDefinition)
	if !ok || newCmdDef == nil {
		return nil
	}
	newCmdDef.Metadata = helper.ReorderByKey(commandDefinition.Metadata, newCmdDef.Metadata, func(m Metadata[any]) string { return m.Name })
	newCmdDef.Arguments = helper.ReorderByKey(commandDefinition.Arguments, newCmdDef.Arguments, func(a Argument[any]) string { return a.Name })
	return nil
}
