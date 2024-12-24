package feasibility_constraint_definitions

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/leanspace/terraform-provider-leanspace/helper"
)

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
