package feasibility_constraint_definitions

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/services/activities/activity_definitions"
)

func (feasibilityConstraintDefinition *FeasibilityConstraintDefinition) ToMap() map[string]any {
	feasibilityConstraintDefinitionMap := make(map[string]any)
	feasibilityConstraintDefinitionMap["id"] = feasibilityConstraintDefinition.ID
	feasibilityConstraintDefinitionMap["name"] = feasibilityConstraintDefinition.Name
	feasibilityConstraintDefinitionMap["description"] = feasibilityConstraintDefinition.Description

	if feasibilityConstraintDefinition.ArgumentDefinitions != nil {
		feasibilityConstraintDefinitionMap["argument_definitions"] = helper.ParseToMaps(feasibilityConstraintDefinition.ArgumentDefinitions)
	}

	feasibilityConstraintDefinitionMap["created_at"] = feasibilityConstraintDefinition.CreatedAt
	feasibilityConstraintDefinitionMap["created_by"] = feasibilityConstraintDefinition.CreatedBy
	feasibilityConstraintDefinitionMap["last_modified_at"] = feasibilityConstraintDefinition.LastModifiedAt
	feasibilityConstraintDefinitionMap["last_modified_by"] = feasibilityConstraintDefinition.LastModifiedBy

	return feasibilityConstraintDefinitionMap
}

func (feasibilityConstraintDefinition *FeasibilityConstraintDefinition) FromMap(feasibilityConstraintDefinitionMap map[string]any) error {
	feasibilityConstraintDefinition.ID = feasibilityConstraintDefinitionMap["id"].(string)
	feasibilityConstraintDefinition.Name = feasibilityConstraintDefinitionMap["name"].(string)
	feasibilityConstraintDefinition.Description = feasibilityConstraintDefinitionMap["description"].(string)

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
