package request_definitions

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/services/activities/activity_definitions"
)

func (requestDefinition *RequestDefinition) ToMap() map[string]any {
	requestDefinitionMap := make(map[string]any)
	requestDefinitionMap["id"] = requestDefinition.ID
	requestDefinitionMap["name"] = requestDefinition.Name
	requestDefinitionMap["description"] = requestDefinition.Description

	if requestDefinition.PlanTemplateIds != nil {
		requestDefinitionMap["planTemplateIds"] = []any{requestDefinition.PlanTemplateIds}
	}

	if requestDefinition.FeasibilityConstraintDefinitions != nil {
		requestDefinitionMap["feasibilityConstraintDefinitions"] = helper.ParseToMaps(requestDefinition.FeasibilityConstraintDefinitions)
	}

	if requestDefinition.ConfigurationArgumentDefinitions != nil {
		requestDefinitionMap["configurationArgumentDefinitions"] = helper.ParseToMaps(requestDefinition.ConfigurationArgumentDefinitions)
	}

	if requestDefinition.ConfigurationArgumentMappings != nil {
		requestDefinitionMap["configurationArgumentMappings"] = helper.ParseToMaps(requestDefinition.ConfigurationArgumentMappings)
	}

	requestDefinitionMap["created_at"] = requestDefinition.CreatedAt
	requestDefinitionMap["created_by"] = requestDefinition.CreatedBy
	requestDefinitionMap["last_modified_at"] = requestDefinition.LastModifiedAt
	requestDefinitionMap["last_modified_by"] = requestDefinition.LastModifiedBy

	helper.Logger.Printf("%s", requestDefinitionMap)
	return requestDefinitionMap
}

func (feasibilityConstraintDefinition *FeasibilityConstraintDefinition) ToMap() map[string]any {
	feasibilityConstraintDefinitionMap := make(map[string]any)
	feasibilityConstraintDefinitionMap["id"] = feasibilityConstraintDefinition.ID
	feasibilityConstraintDefinitionMap["name"] = feasibilityConstraintDefinition.Name
	feasibilityConstraintDefinitionMap["description"] = feasibilityConstraintDefinition.Description
	feasibilityConstraintDefinitionMap["required"] = feasibilityConstraintDefinition.Required

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
		if argumentDefinitions, err := helper.ParseFromMaps[activity_definitions.ArgumentDefinition[any]](
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
	feasibilityConstraintDefinition.Required = feasibilityConstraintDefinitionMap["required"].(bool)

	if feasibilityConstraintDefinitionMap["argument_definitions"] != nil {
		if argumentDefinitions, err := helper.ParseFromMaps[activity_definitions.ArgumentDefinition[any]](
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

func (mapping *ArgumentMapping) FromMap(mappingMap map[string]any) error {
	mapping.PlanTemplateId = mappingMap["planTemplateId"].(string)
	mapping.ActivityDefinitionPosition = mappingMap["activityDefinitionPosition"].(int)
	mapping.ConfigurationArgumentDefinitionName = mappingMap["configurationArgumentDefinitionName"].(string)
	mapping.ActivityDefinitionArgumentDefinitionName = mappingMap["activityDefinitionArgumentDefinitionName"].(string)

	return nil
}
