package resource_functions

import (
	"github.com/leanspace/terraform-provider-leanspace/helper"
)

func (resourceFunction *ResourceFunction) ToMap() map[string]any {
	resourceFunctionMap := resourceFunction.ToAuditMap()
	resourceFunctionMap["activity_definition_id"] = helper.NilIfEmpty(resourceFunction.ActivityDefinitionId)
	resourceFunctionMap["resource_id"] = helper.NilIfEmpty(resourceFunction.ResourceId)
	resourceFunctionMap["name"] = helper.NilIfEmpty(resourceFunction.Name)
	resourceFunctionMap["formula"] = resourceFunction.Formula.ToMap()
	return resourceFunctionMap
}

func (formula *ResourceFunctionFormula) ToMap() map[string]any {
	formulaMap := make(map[string]any)

	formulaMap["type"] = helper.NilIfEmpty(formula.Type)

	if formula.Type == "LINEAR" {
		formulaMap["constant"] = helper.NilIfEmpty(formula.Constant)
		formulaMap["rate"] = helper.NilIfEmpty(formula.Rate)
		formulaMap["time_unit"] = helper.NilIfEmpty(formula.TimeUnit)
	}

	if formula.Type == "RECTANGULAR" {
		formulaMap["amplitude"] = helper.NilIfEmpty(formula.Amplitude)
	}

	return formulaMap
}

func (resourceFunction *ResourceFunction) FromMap(resourceFunctionMap map[string]any) error {
	resourceFunction.FromAuditMap(resourceFunctionMap)
	resourceFunction.ActivityDefinitionId = helper.CastString(resourceFunctionMap, "activity_definition_id")
	resourceFunction.ResourceId = helper.CastString(resourceFunctionMap, "resource_id")
	resourceFunction.Name = helper.CastString(resourceFunctionMap, "name")
	if resourceFunctionMap["formula"] != nil {
		resourceFunction.Formula = new(ResourceFunctionFormula)
		if err := resourceFunction.Formula.FromMap(helper.CastMapAny(resourceFunctionMap, "formula")); err != nil {
			return err
		}
	}
	return nil
}

func (formula *ResourceFunctionFormula) FromMap(formulaMap map[string]any) error {
	formula.Type = helper.CastString(formulaMap, "type")

	if formula.Type == "LINEAR" {
		formula.Constant = helper.CastFloat64(formulaMap, "constant")
		formula.Rate = helper.CastFloat64(formulaMap, "rate")
		formula.TimeUnit = helper.CastString(formulaMap, "time_unit")
	}

	if formula.Type == "RECTANGULAR" {
		formula.Amplitude = helper.CastFloat64(formulaMap, "amplitude")
	}

	return nil
}
