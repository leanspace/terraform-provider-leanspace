package feasibility_constraint_definitions

import (
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/services/activities/activity_definitions"
)

func (feasibilityConstraintDefinition *FeasibilityConstraintDefinition) ToMap() map[string]any {
	feasibilityConstraintDefinitionMap := feasibilityConstraintDefinition.ToAuditMap()
	feasibilityConstraintDefinitionMap["name"] = helper.NilIfEmpty(feasibilityConstraintDefinition.Name)
	feasibilityConstraintDefinitionMap["description"] = helper.NilIfEmpty(feasibilityConstraintDefinition.Description)

	if feasibilityConstraintDefinition.ArgumentDefinitions != nil {
		feasibilityConstraintDefinitionMap["argument_definitions"] = helper.ParseToMaps(feasibilityConstraintDefinition.ArgumentDefinitions)
	}

	return feasibilityConstraintDefinitionMap
}

func (feasibilityConstraintDefinition *FeasibilityConstraintDefinition) FromMap(feasibilityConstraintDefinitionMap map[string]any) error {
	feasibilityConstraintDefinition.FromAuditMap(feasibilityConstraintDefinitionMap)
	feasibilityConstraintDefinition.Name = helper.CastString(feasibilityConstraintDefinitionMap, "name")
	feasibilityConstraintDefinition.Description = helper.CastString(feasibilityConstraintDefinitionMap, "description")

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
