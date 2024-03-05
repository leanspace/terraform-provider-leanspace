package resource_functions

func (resourceFunction *ResourceFunction) ToMap() map[string]any {
	resourceFunctionMap := make(map[string]any)
	resourceFunctionMap["id"] = resourceFunction.ID
	resourceFunctionMap["activity_definition_id"] = resourceFunction.ActivityDefinitionId
	resourceFunctionMap["resource_id"] = resourceFunction.ResourceId
	resourceFunctionMap["name"] = resourceFunction.Name
	resourceFunctionMap["time_unit"] = resourceFunction.TimeUnit
	resourceFunctionMap["formula"] = []map[string]any{resourceFunction.Formula.ToMap()}
	resourceFunctionMap["created_at"] = resourceFunction.CreatedAt
	resourceFunctionMap["created_by"] = resourceFunction.CreatedBy
	resourceFunctionMap["last_modified_at"] = resourceFunction.LastModifiedAt
	resourceFunctionMap["last_modified_by"] = resourceFunction.LastModifiedBy
	return resourceFunctionMap
}

func (formula *ResourceFunctionFormula) ToMap() map[string]any {
	formulaMap := make(map[string]any)
	formulaMap["constant"] = formula.Constant
	formulaMap["rate"] = formula.Rate
	formulaMap["type"] = formula.Type
	return formulaMap
}

func (resourceFunction *ResourceFunction) FromMap(resourceFunctionMap map[string]any) error {
	resourceFunction.ID = resourceFunctionMap["id"].(string)
	resourceFunction.ActivityDefinitionId = resourceFunctionMap["activity_definition_id"].(string)
	resourceFunction.ResourceId = resourceFunctionMap["resource_id"].(string)
	resourceFunction.Name = resourceFunctionMap["name"].(string)
	resourceFunction.TimeUnit = resourceFunctionMap["time_unit"].(string)
	if len(resourceFunctionMap["formula"].([]any)) > 0 && resourceFunctionMap["formula"].([]any)[0] != nil {
		resourceFunction.Formula = new(ResourceFunctionFormula)
		if err := resourceFunction.Formula.FromMap(resourceFunctionMap["formula"].([]any)[0].(map[string]any)); err != nil {
			return err
		}
	}
	resourceFunction.CreatedAt = resourceFunctionMap["created_at"].(string)
	resourceFunction.CreatedBy = resourceFunctionMap["created_by"].(string)
	resourceFunction.LastModifiedAt = resourceFunctionMap["last_modified_at"].(string)
	resourceFunction.LastModifiedBy = resourceFunctionMap["last_modified_by"].(string)
	return nil
}

func (formula *ResourceFunctionFormula) FromMap(formulaMap map[string]any) error {
	formula.Constant = formulaMap["constant"].(float64)
	formula.Rate = formulaMap["rate"].(float64)
	formula.Type = formulaMap["type"].(string)
	return nil
}
