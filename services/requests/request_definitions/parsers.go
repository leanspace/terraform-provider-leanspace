package request_definitions

import (
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/services/activities/activity_definitions"
)

func (requestDefinition *RequestDefinition) ToMap() map[string]any {
	requestDefinitionMap := requestDefinition.ToAuditMap()
	requestDefinitionMap["name"] = helper.NilIfEmpty(requestDefinition.Name)
	requestDefinitionMap["description"] = helper.NilIfEmpty(requestDefinition.Description)
	requestDefinitionMap["plan_template_ids"] = helper.NilIfEmpty(requestDefinition.PlanTemplateIds)

	if requestDefinition.FeasibilityConstraintDefinitions != nil {
		requestDefinitionMap["feasibility_constraint_definitions"] = helper.ParseToMaps(requestDefinition.FeasibilityConstraintDefinitions)
	}

	if requestDefinition.ConfigurationArgumentDefinitions != nil {
		requestDefinitionMap["configuration_argument_definitions"] = helper.ParseToMaps(requestDefinition.ConfigurationArgumentDefinitions)
	}

	if requestDefinition.ConfigurationArgumentMappings != nil {
		requestDefinitionMap["configuration_argument_mappings"] = helper.ParseToMaps(requestDefinition.ConfigurationArgumentMappings)
	}

	return requestDefinitionMap
}

func (feasibilityConstraintDefinition *FeasibilityConstraintDefinition) ToMap() map[string]any {
	feasibilityConstraintDefinitionMap := feasibilityConstraintDefinition.ToAuditMap()
	feasibilityConstraintDefinitionMap["name"] = helper.NilIfEmpty(feasibilityConstraintDefinition.Name)
	feasibilityConstraintDefinitionMap["description"] = helper.NilIfEmpty(feasibilityConstraintDefinition.Description)
	feasibilityConstraintDefinitionMap["required"] = helper.NilIfEmpty(feasibilityConstraintDefinition.Required)

	if feasibilityConstraintDefinition.ArgumentDefinitions != nil {
		feasibilityConstraintDefinitionMap["argument_definitions"] = helper.ParseToMaps(feasibilityConstraintDefinition.ArgumentDefinitions)
	}

	return feasibilityConstraintDefinitionMap
}

func (mapping *ArgumentMapping) ToMap() map[string]any {
	mappingMap := make(map[string]any)
	mappingMap["plan_template_id"] = helper.NilIfEmpty(mapping.PlanTemplateId)
	mappingMap["activity_definition_position"] = helper.NilIfEmpty(mapping.ActivityDefinitionPosition)
	mappingMap["configuration_argument_definition_name"] = helper.NilIfEmpty(mapping.ConfigurationArgumentDefinitionName)
	mappingMap["activity_definition_argument_definition_name"] = helper.NilIfEmpty(mapping.ActivityDefinitionArgumentDefinitionName)

	return mappingMap
}

func (requestDefinition *RequestDefinition) FromMap(requestDefinitionMap map[string]any) error {
	requestDefinition.FromAuditMap(requestDefinitionMap)
	requestDefinition.Name = helper.CastString(requestDefinitionMap, "name")
	requestDefinition.Description = helper.CastString(requestDefinitionMap, "description")

	requestDefinition.PlanTemplateIds = make([]string, len(helper.CastSlice(requestDefinitionMap, "plan_template_ids")))
	for i, processorId := range helper.CastSlice(requestDefinitionMap, "plan_template_ids") {
		requestDefinition.PlanTemplateIds[i] = processorId.(string)
	}

	if requestDefinitionMap["feasibility_constraint_definitions"] != nil {
		if feasibilityConstraintDefinitions, err := helper.ParseFromMaps[FeasibilityConstraintDefinition](
			helper.CastSlice(requestDefinitionMap, "feasibility_constraint_definitions"),
		); err != nil {
			return err
		} else {
			requestDefinition.FeasibilityConstraintDefinitions = feasibilityConstraintDefinitions
		}
	}

	if requestDefinitionMap["configuration_argument_definitions"] != nil {
		if argumentDefinitions, err := helper.ParseFromMaps[activity_definitions.ArgumentDefinition[any]](
			helper.CastSlice(requestDefinitionMap, "configuration_argument_definitions"),
		); err != nil {
			return err
		} else {
			requestDefinition.ConfigurationArgumentDefinitions = argumentDefinitions
		}
	}

	if requestDefinitionMap["configuration_argument_mappings"] != nil {
		if configurationArgumentMappings, err := helper.ParseFromMaps[ArgumentMapping](
			helper.CastSlice(requestDefinitionMap, "configuration_argument_mappings"),
		); err != nil {
			return err
		} else {
			requestDefinition.ConfigurationArgumentMappings = configurationArgumentMappings
		}
	}

	return nil
}

func (feasibilityConstraintDefinition *FeasibilityConstraintDefinition) FromMap(feasibilityConstraintDefinitionMap map[string]any) error {
	feasibilityConstraintDefinition.FromAuditMap(feasibilityConstraintDefinitionMap)
	feasibilityConstraintDefinition.Name = helper.CastString(feasibilityConstraintDefinitionMap, "name")
	feasibilityConstraintDefinition.Description = helper.CastString(feasibilityConstraintDefinitionMap, "description")
	feasibilityConstraintDefinition.Required = helper.CastBool(feasibilityConstraintDefinitionMap, "required")

	if feasibilityConstraintDefinitionMap["argument_definitions"] != nil {
		if argumentDefinitions, err := helper.ParseFromMaps[activity_definitions.ArgumentDefinition[any]](
			helper.CastSlice(feasibilityConstraintDefinitionMap, "argument_definitions"),
		); err != nil {
			return err
		} else {
			feasibilityConstraintDefinition.ArgumentDefinitions = argumentDefinitions
		}
	}

	return nil
}

func (mapping *ArgumentMapping) FromMap(mappingMap map[string]any) error {
	mapping.PlanTemplateId = helper.CastString(mappingMap, "plan_template_id")
	mapping.ActivityDefinitionPosition = helper.CastInt(mappingMap, "activity_definition_position")
	mapping.ConfigurationArgumentDefinitionName = helper.CastString(mappingMap, "configuration_argument_definition_name")
	mapping.ActivityDefinitionArgumentDefinitionName = helper.CastString(mappingMap, "activity_definition_argument_definition_name")

	return nil
}
