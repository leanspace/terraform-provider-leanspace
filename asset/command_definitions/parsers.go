package command_definitions

import (
	"strconv"

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
		commandDefinitionMap["metadata"] = make([]any, len(commandDefinition.Metadata))
		for i, metadata := range commandDefinition.Metadata {
			metadataMap := make(map[string]any)
			metadataMap["id"] = metadata.ID
			metadataMap["name"] = metadata.Name
			metadataMap["description"] = metadata.Description
			metadataMap["unit_id"] = metadata.UnitId
			metadataMap["type"] = metadata.Type
			switch metadata.Type {
			case "NUMERIC":
				if metadata.Value != nil {
					metadataMap["value"] = strconv.FormatFloat(metadata.Value.(float64), 'g', -1, 64)
				}
			case "TEXT":
				if metadata.Value != nil {
					metadataMap["value"] = metadata.Value.(string)
				}
			case "TIMESTAMP", "DATE", "TIME":
				if metadata.Value != nil {
					metadataMap["value"] = metadata.Value.(string)
				}
			case "BOOLEAN":
				if metadata.Value != nil {
					metadataMap["value"] = strconv.FormatBool(metadata.Value.(bool))
				}
			}
			commandDefinitionMap["metadata"].([]any)[i] = metadataMap
		}
	}

	if commandDefinition.Arguments != nil {
		commandDefinitionMap["arguments"] = make([]any, len(commandDefinition.Arguments))
		for i, argument := range commandDefinition.Arguments {
			argumentMap := make(map[string]any)
			argumentMap["id"] = argument.ID
			argumentMap["name"] = argument.Name
			argumentMap["identifier"] = argument.Identifier
			argumentMap["description"] = argument.Description
			argumentMap["type"] = argument.Type
			argumentMap["required"] = argument.Required
			switch argument.Type {
			case "NUMERIC":
				if argument.DefaultValue != nil {
					argumentMap["default_value"] = strconv.FormatFloat(argument.DefaultValue.(float64), 'g', -1, 64)
				}
				argumentMap["min"] = argument.Min
				argumentMap["max"] = argument.Max
				argumentMap["scale"] = argument.Scale
				argumentMap["precision"] = argument.Precision
				argumentMap["unit_id"] = argument.UnitId
			case "ENUM":
				if argument.DefaultValue != nil {
					argumentMap["default_value"] = strconv.FormatFloat(argument.DefaultValue.(float64), 'g', -1, 64)
				}
				if argument.Options != nil {
					argumentMap["options"] = *argument.Options
				}
			case "TEXT":
				if argument.DefaultValue != nil {
					argumentMap["default_value"] = argument.DefaultValue.(string)
				}
				argumentMap["min_length"] = argument.MinLength
				argumentMap["max_length"] = argument.MaxLength
				argumentMap["pattern"] = argument.Pattern
			case "TIMESTAMP", "DATE", "TIME":
				if argument.DefaultValue != nil {
					argumentMap["default_value"] = argument.DefaultValue.(string)
				}
				argumentMap["before"] = argument.Before
				argumentMap["after"] = argument.After
			case "BOOLEAN":
				if argument.DefaultValue != nil {
					argumentMap["default_value"] = strconv.FormatBool(argument.DefaultValue.(bool))
				}
			}
			commandDefinitionMap["arguments"].([]any)[i] = argumentMap
		}
	}

	return commandDefinitionMap
}

func getCommandDefinitionData(commandDefinition map[string]any) (CommandDefinition, error) {
	commandDefinitionMap := CommandDefinition{}

	commandDefinitionMap.NodeId = commandDefinition["node_id"].(string)
	commandDefinitionMap.Name = commandDefinition["name"].(string)
	commandDefinitionMap.Description = commandDefinition["description"].(string)
	commandDefinitionMap.Identifier = commandDefinition["identifier"].(string)
	commandDefinitionMap.CreatedAt = commandDefinition["created_at"].(string)
	commandDefinitionMap.CreatedBy = commandDefinition["created_by"].(string)
	commandDefinitionMap.LastModifiedAt = commandDefinition["last_modified_at"].(string)
	commandDefinitionMap.LastModifiedBy = commandDefinition["last_modified_by"].(string)
	if commandDefinition["metadata"] != nil {
		commandDefinitionMap.Metadata = []Metadata[any]{}
		for _, metadata := range commandDefinition["metadata"].(*schema.Set).List() {
			commandDefinitionMap.Metadata = append(commandDefinitionMap.Metadata, metadataInterfaceToStruct(metadata.(map[string]any)))
		}
	}
	if commandDefinition["arguments"] != nil {
		commandDefinitionMap.Arguments = []Argument[any]{}
		for _, argument := range commandDefinition["arguments"].(*schema.Set).List() {
			commandDefinitionMap.Arguments = append(commandDefinitionMap.Arguments, argumentInterfaceToStruct(argument.(map[string]any)))
		}
	}

	return commandDefinitionMap, nil
}

func metadataInterfaceToStruct(metadata map[string]any) Metadata[any] {
	metadataStruct := Metadata[any]{}
	metadataStruct.ID = metadata["id"].(string)
	metadataStruct.Name = metadata["name"].(string)
	metadataStruct.Description = metadata["description"].(string)
	metadataStruct.UnitId = metadata["unit_id"].(string)
	metadataStruct.Value = metadata["value"]
	metadataStruct.Type = metadata["type"].(string)

	return metadataStruct
}

func argumentInterfaceToStruct(argument map[string]any) Argument[any] {
	argumentStruct := Argument[any]{}
	argumentStruct.ID = argument["id"].(string)
	argumentStruct.Name = argument["name"].(string)
	argumentStruct.Identifier = argument["identifier"].(string)
	argumentStruct.Description = argument["description"].(string)
	argumentStruct.Type = argument["type"].(string)
	argumentStruct.Required = argument["required"].(bool)
	switch argumentStruct.Type {
	case "NUMERIC":
		argumentStruct.Min = argument["min"].(float64)
		argumentStruct.Max = argument["max"].(float64)
		argumentStruct.Scale = argument["scale"].(int)
		argumentStruct.Precision = argument["precision"].(int)
		argumentStruct.UnitId = argument["unit_id"].(string)
	case "ENUM":
		if argument["options"] != nil {
			option := argument["options"].(map[string]any)
			argumentStruct.Options = &option
		}
	case "TEXT":
		argumentStruct.MinLength = argument["min_length"].(int)
		argumentStruct.MaxLength = argument["max_length"].(int)
		argumentStruct.Pattern = argument["pattern"].(string)
	case "TIMESTAMP", "DATE", "TIME":
		argumentStruct.Before = argument["before"].(string)
		argumentStruct.After = argument["after"].(string)
	case "BOOLEAN":
		// no extra field
	}
	argumentStruct.DefaultValue = argument["default_value"]

	return argumentStruct
}
