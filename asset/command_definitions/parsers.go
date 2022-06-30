package command_definitions

import (
	"strconv"
	"terraform-provider-asset/asset"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func commandDefinitionStructToInterface(commandDefinition CommandDefinition) map[string]any {
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
		commandDefinitionMap["metadata"] = asset.Map(
			commandDefinition.Metadata,
			metadataStructToInterface,
		)
	}
	if commandDefinition.Arguments != nil {
		commandDefinitionMap["arguments"] = asset.Map(
			commandDefinition.Arguments,
			argumentStructToInterface,
		)
	}

	return commandDefinitionMap
}

func metadataStructToInterface(metadata Metadata[any]) map[string]any {
	metadataMap := make(map[string]any)
	metadataMap["id"] = metadata.ID
	metadataMap["name"] = metadata.Name
	metadataMap["description"] = metadata.Description

	attributes := make(map[string]any)

	attributes["type"] = metadata.Attributes.Type
	if metadata.Attributes.Type == "NUMERIC" {
		attributes["unit_id"] = metadata.Attributes.UnitId
	}
	if metadata.Attributes.Value != nil {
		switch metadata.Attributes.Type {
		case "NUMERIC":
			attributes["value"] = strconv.FormatFloat(metadata.Attributes.Value.(float64), 'g', -1, 64)
		case "TEXT":
			attributes["value"] = metadata.Attributes.Value.(string)
		case "BOOLEAN":
			attributes["value"] = strconv.FormatBool(metadata.Attributes.Value.(bool))
		case "TIMESTAMP", "DATE", "TIME":
			attributes["value"] = metadata.Attributes.Value.(string)
		}
	}
	metadataMap["attributes"] = []map[string]any{attributes}

	return metadataMap
}

func argumentStructToInterface(argument Argument[any]) map[string]any {
	argumentMap := make(map[string]any)
	argumentMap["id"] = argument.ID
	argumentMap["name"] = argument.Name
	argumentMap["identifier"] = argument.Identifier
	argumentMap["description"] = argument.Description

	attributes := make(map[string]any)

	attributes["type"] = argument.Attributes.Type
	attributes["required"] = argument.Attributes.Required

	switch argument.Attributes.Type {
	case "TEXT":
		if argument.Attributes.DefaultValue != nil {
			attributes["default_value"] = argument.Attributes.DefaultValue.(string)
		}
		attributes["min_length"] = argument.Attributes.MinLength
		attributes["max_length"] = argument.Attributes.MaxLength
		attributes["pattern"] = argument.Attributes.Pattern
	case "NUMERIC":
		if argument.Attributes.DefaultValue != nil {
			attributes["default_value"] = asset.ParseFloat(argument.Attributes.DefaultValue.(float64))
		}
		attributes["min"] = argument.Attributes.Min
		attributes["max"] = argument.Attributes.Max
		attributes["scale"] = argument.Attributes.Scale
		attributes["precision"] = argument.Attributes.Precision
		attributes["unit_id"] = argument.Attributes.UnitId
	case "BOOLEAN":
		if argument.Attributes.DefaultValue != nil {
			attributes["default_value"] = strconv.FormatBool(argument.Attributes.DefaultValue.(bool))
		}
	case "TIMESTAMP", "DATE", "TIME":
		if argument.Attributes.DefaultValue != nil {
			attributes["default_value"] = argument.Attributes.DefaultValue.(string)
		}
		attributes["before"] = argument.Attributes.Before
		attributes["after"] = argument.Attributes.After
	case "ENUM":
		if argument.Attributes.DefaultValue != nil {
			attributes["default_value"] = asset.ParseFloat(argument.Attributes.DefaultValue.(float64))
		}
		if argument.Attributes.Options != nil {
			attributes["options"] = *argument.Attributes.Options
		}
	}

	argumentMap["attributes"] = []map[string]any{attributes}

	return argumentMap
}

func getCommandDefinitionData(commandDefinition map[string]any) (CommandDefinition, error) {
	commandDefinitionMap := CommandDefinition{}

	commandDefinitionMap.ID = commandDefinition["id"].(string)
	commandDefinitionMap.NodeId = commandDefinition["node_id"].(string)
	commandDefinitionMap.Name = commandDefinition["name"].(string)
	commandDefinitionMap.Description = commandDefinition["description"].(string)
	commandDefinitionMap.Identifier = commandDefinition["identifier"].(string)
	commandDefinitionMap.CreatedAt = commandDefinition["created_at"].(string)
	commandDefinitionMap.CreatedBy = commandDefinition["created_by"].(string)
	commandDefinitionMap.LastModifiedAt = commandDefinition["last_modified_at"].(string)
	commandDefinitionMap.LastModifiedBy = commandDefinition["last_modified_by"].(string)
	if commandDefinition["metadata"] != nil {
		commandDefinitionMap.Metadata = asset.CastMap(
			commandDefinition["metadata"].(*schema.Set).List(),
			metadataInterfaceToStruct,
		)
	}
	if commandDefinition["arguments"] != nil {
		commandDefinitionMap.Arguments = asset.CastMap(
			commandDefinition["arguments"].(*schema.Set).List(),
			argumentInterfaceToStruct,
		)
	}

	return commandDefinitionMap, nil
}

func metadataInterfaceToStruct(metadata map[string]any) Metadata[any] {
	metadataStruct := Metadata[any]{}
	metadataStruct.ID = metadata["id"].(string)
	metadataStruct.Name = metadata["name"].(string)
	metadataStruct.Description = metadata["description"].(string)

	attributes := metadata["attributes"].([]any)[0].(map[string]any)
	metadataStruct.Attributes.Value = attributes["value"]
	metadataStruct.Attributes.Type = attributes["type"].(string)
	if metadata["type"] == "NUMERIC" {
		metadataStruct.Attributes.UnitId = attributes["unit_id"].(string)
	}

	return metadataStruct
}

func argumentInterfaceToStruct(argument map[string]any) Argument[any] {
	argumentStruct := Argument[any]{}
	argumentStruct.ID = argument["id"].(string)
	argumentStruct.Name = argument["name"].(string)
	argumentStruct.Identifier = argument["identifier"].(string)
	argumentStruct.Description = argument["description"].(string)

	attributes := argument["attributes"].([]any)[0].(map[string]any)
	argumentStruct.Attributes.Type = attributes["type"].(string)
	argumentStruct.Attributes.Required = attributes["required"].(bool)
	argumentStruct.Attributes.DefaultValue = attributes["default_value"]
	switch argumentStruct.Attributes.Type {
	case "NUMERIC":
		argumentStruct.Attributes.Min = attributes["min"].(float64)
		argumentStruct.Attributes.Max = attributes["max"].(float64)
		argumentStruct.Attributes.Scale = attributes["scale"].(int)
		argumentStruct.Attributes.Precision = attributes["precision"].(int)
		argumentStruct.Attributes.UnitId = attributes["unit_id"].(string)
	case "ENUM":
		if attributes["options"] != nil {
			option := attributes["options"].(map[string]any)
			argumentStruct.Attributes.Options = &option
		}
	case "TEXT":
		argumentStruct.Attributes.MinLength = attributes["min_length"].(int)
		argumentStruct.Attributes.MaxLength = attributes["max_length"].(int)
		argumentStruct.Attributes.Pattern = attributes["pattern"].(string)
	case "TIMESTAMP", "DATE", "TIME":
		argumentStruct.Attributes.Before = attributes["before"].(string)
		argumentStruct.Attributes.After = attributes["after"].(string)
	case "BOOLEAN":
		// no extra field
	}

	return argumentStruct
}
