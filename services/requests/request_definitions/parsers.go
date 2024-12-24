package request_definitions

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/leanspace/terraform-provider-leanspace/helper"
)

func (feasibilityConstraintDefinition *RequestDefinition) ToMap() map[string]any {
	feasibilityConstraintDefinitionMap := make(map[string]any)
	feasibilityConstraintDefinitionMap["id"] = feasibilityConstraintDefinition.ID
	feasibilityConstraintDefinitionMap["name"] = feasibilityConstraintDefinition.Name
	feasibilityConstraintDefinitionMap["description"] = feasibilityConstraintDefinition.Description
	feasibilityConstraintDefinitionMap["cloned"] = feasibilityConstraintDefinition.FeasibilityConstraintDefinitions

	if feasibilityConstraintDefinition.PlanTemplateIds != nil {
		feasibilityConstraintDefinitionMap["planTemplateIds"] = []any{feasibilityConstraintDefinition.PlanTemplateIds}
	}

	if feasibilityConstraintDefinition.FeasibilityConstraintDefinitions != nil {
		feasibilityConstraintDefinitionMap["feasibilityConstraintDefinitions"] = helper.ParseToMaps(feasibilityConstraintDefinition.FeasibilityConstraintDefinitions)
	}

	if feasibilityConstraintDefinition.ConfigurationArgumentDefinitions != nil {
		feasibilityConstraintDefinitionMap["configurationArgumentDefinitions"] = helper.ParseToMaps(feasibilityConstraintDefinition.ConfigurationArgumentDefinitions)
	}

	if feasibilityConstraintDefinition.ConfigurationArgumentMappings != nil {
		feasibilityConstraintDefinitionMap["configurationArgumentMappings"] = helper.ParseToMaps(feasibilityConstraintDefinition.ConfigurationArgumentMappings)
	}

	feasibilityConstraintDefinitionMap["created_at"] = feasibilityConstraintDefinition.CreatedAt
	feasibilityConstraintDefinitionMap["created_by"] = feasibilityConstraintDefinition.CreatedBy
	feasibilityConstraintDefinitionMap["last_modified_at"] = feasibilityConstraintDefinition.LastModifiedAt
	feasibilityConstraintDefinitionMap["last_modified_by"] = feasibilityConstraintDefinition.LastModifiedBy

	helper.Logger.Printf("%s", feasibilityConstraintDefinitionMap)
	return feasibilityConstraintDefinitionMap
}

func (feasibilityConstraintDefinition *FeasibilityConstraintDefinition) ToMap() map[string]any {
	feasibilityConstraintDefinitionMap := make(map[string]any)
	feasibilityConstraintDefinitionMap["id"] = feasibilityConstraintDefinition.ID
	feasibilityConstraintDefinitionMap["name"] = feasibilityConstraintDefinition.Name
	feasibilityConstraintDefinitionMap["description"] = feasibilityConstraintDefinition.Description
	feasibilityConstraintDefinitionMap["cloned"] = feasibilityConstraintDefinition.Cloned

	if feasibilityConstraintDefinition.ArgumentDefinitions != nil {
		feasibilityConstraintDefinitionMap["argument_definitions"] = helper.ParseToMaps(feasibilityConstraintDefinition.ArgumentDefinitions)
	}

	feasibilityConstraintDefinitionMap["created_at"] = feasibilityConstraintDefinition.CreatedAt
	feasibilityConstraintDefinitionMap["created_by"] = feasibilityConstraintDefinition.CreatedBy
	feasibilityConstraintDefinitionMap["last_modified_at"] = feasibilityConstraintDefinition.LastModifiedAt
	feasibilityConstraintDefinitionMap["last_modified_by"] = feasibilityConstraintDefinition.LastModifiedBy

	helper.Logger.Printf("%s", feasibilityConstraintDefinitionMap)
	return feasibilityConstraintDefinitionMap
}

func (argument *ArgumentDefinition[T]) ToMap() map[string]any {
	argumentMap := make(map[string]any)
	argumentMap["name"] = argument.Name
	argumentMap["description"] = argument.Description
	argumentMap["attributes"] = []any{argument.Attributes.ToMap()}
	return argumentMap
}

func (mapping *ArgumentMapping) ToMap() map[string]any {
	mappingMap := make(map[string]any)
	mappingMap["planTemplateId"] = mapping.PlanTemplateId
	mappingMap["activityDefinitionPosition"] = mapping.ActivityDefinitionPosition
	mappingMap["configurationArgumentDefinitionName"] = mapping.ConfigurationArgumentDefinitionName
	mappingMap["activityDefinitionArgumentDefinitionName"] = mapping.ActivityDefinitionArgumentDefinitionName

	return mappingMap
}

func (requestDefinition *RequestDefinition) FromMap(requestDefinitionMap map[string]any) error {
	requestDefinition.ID = requestDefinitionMap["id"].(string)
	requestDefinition.Name = requestDefinitionMap["name"].(string)
	requestDefinition.Description = requestDefinitionMap["description"].(string)

	requestDefinition.PlanTemplateIds = make([]string, requestDefinitionMap["planTemplateIds"].(*schema.Set).Len())
	for i, value := range requestDefinitionMap["planTemplateIds"].(*schema.Set).List() {
		requestDefinition.PlanTemplateIds[i] = value.(string)
	}

	if requestDefinitionMap["feasibilityConstraintDefinitions"] != nil {
		if feasibilityConstraintDefinitions, err := helper.ParseFromMaps[FeasibilityConstraintDefinition](
			requestDefinitionMap["feasibilityConstraintDefinitions"].(*schema.Set).List(),
		); err != nil {
			return err
		} else {
			requestDefinition.FeasibilityConstraintDefinitions = feasibilityConstraintDefinitions
		}
	}

	if requestDefinitionMap["configurationArgumentDefinitions"] != nil {
		if argumentDefinitions, err := helper.ParseFromMaps[ArgumentDefinition[any]](
			requestDefinitionMap["configurationArgumentDefinitions"].(*schema.Set).List(),
		); err != nil {
			return err
		} else {
			requestDefinition.ConfigurationArgumentDefinitions = argumentDefinitions
		}
	}

	if requestDefinitionMap["configurationArgumentMappings"] != nil {
		if configurationArgumentMappings, err := helper.ParseFromMaps[ArgumentMapping](
			requestDefinitionMap["configurationArgumentMappings"].(*schema.Set).List(),
		); err != nil {
			return err
		} else {
			requestDefinition.ConfigurationArgumentMappings = configurationArgumentMappings
		}
	}

	requestDefinition.CreatedAt = requestDefinitionMap["created_at"].(string)
	requestDefinition.CreatedBy = requestDefinitionMap["created_by"].(string)
	requestDefinition.LastModifiedAt = requestDefinitionMap["last_modified_at"].(string)
	requestDefinition.LastModifiedBy = requestDefinitionMap["last_modified_by"].(string)
	return nil
}

func (feasibilityConstraintDefinition *FeasibilityConstraintDefinition) FromMap(feasibilityConstraintDefinitionMap map[string]any) error {
	feasibilityConstraintDefinition.ID = feasibilityConstraintDefinitionMap["id"].(string)
	feasibilityConstraintDefinition.Name = feasibilityConstraintDefinitionMap["name"].(string)
	feasibilityConstraintDefinition.Description = feasibilityConstraintDefinitionMap["description"].(string)
	feasibilityConstraintDefinition.Cloned = feasibilityConstraintDefinitionMap["cloned"].(bool)

	if feasibilityConstraintDefinitionMap["argument_definitions"] != nil {
		if argumentDefinitions, err := helper.ParseFromMaps[ArgumentDefinition[any]](
			feasibilityConstraintDefinitionMap["argument_definitions"].(*schema.Set).List(),
		); err != nil {
			return err
		} else {
			feasibilityConstraintDefinition.ArgumentDefinitions = argumentDefinitions
		}
	}

	feasibilityConstraintDefinition.CreatedAt = feasibilityConstraintDefinitionMap["created_at"].(string)
	feasibilityConstraintDefinition.CreatedBy = feasibilityConstraintDefinitionMap["created_by"].(string)
	feasibilityConstraintDefinition.LastModifiedAt = feasibilityConstraintDefinitionMap["last_modified_at"].(string)
	feasibilityConstraintDefinition.LastModifiedBy = feasibilityConstraintDefinitionMap["last_modified_by"].(string)
	return nil
}

func (argument *ArgumentDefinition[T]) FromMap(argumentMap map[string]any) error {
	argument.Name = argumentMap["name"].(string)
	argument.Description = argumentMap["description"].(string)

	if len(argumentMap["attributes"].([]any)) > 0 {
		if err := argument.Attributes.FromMap(argumentMap["attributes"].([]any)[0].(map[string]any)); err != nil {
			return err
		}
	}
	return nil
}

func (mapping *ArgumentMapping) FromMap(mappingMap map[string]any) error {
	mapping.PlanTemplateId = mappingMap["planTemplateId"].(string)
	mapping.ActivityDefinitionPosition = mappingMap["activityDefinitionPosition"].(int)
	mapping.ConfigurationArgumentDefinitionName = mappingMap["configurationArgumentDefinitionName"].(string)
	mapping.ActivityDefinitionArgumentDefinitionName = mappingMap["activityDefinitionArgumentDefinitionName"].(string)

	return nil
}
