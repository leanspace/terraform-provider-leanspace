package plan_templates

import (
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

func (planTemplate *PlanTemplate) ToMap() map[string]any {
	planTemplateMap := planTemplate.ToAuditMap()
	planTemplateMap["asset_id"] = helper.NilIfEmpty(planTemplate.AssetId)
	planTemplateMap["name"] = helper.NilIfEmpty(planTemplate.Name)
	planTemplateMap["description"] = helper.NilIfEmpty(planTemplate.Description)
	planTemplateMap["integrity_status"] = helper.NilIfEmpty(planTemplate.IntegrityStatus)

	if planTemplate.ActivityConfigs != nil {
		planTemplateMap["activity_configs"] = helper.ParseToMaps(planTemplate.ActivityConfigs)
	}

	planTemplateMap["estimated_duration_in_seconds"] = helper.NilIfEmpty(planTemplate.EstimatedDurationInSeconds)

	if planTemplate.InvalidPlanTemplateReasons != nil {
		planTemplateMap["invalid_plan_template_reasons"] = helper.ParseToMaps(planTemplate.InvalidPlanTemplateReasons)
	}

	return planTemplateMap
}

func (activityConfigResult *ActivityConfigResult) ToMap() map[string]any {
	activityConfigResultMap := make(map[string]any)
	activityConfigResultMap["activity_definition_id"] = helper.NilIfEmpty(activityConfigResult.ActivityDefinitionId)
	activityConfigResultMap["delay_reference_on_predecessor"] = helper.NilIfEmpty(activityConfigResult.DelayReferenceOnPredecessor)
	activityConfigResultMap["position"] = helper.NilIfEmpty(activityConfigResult.Position)
	activityConfigResultMap["delay_in_seconds"] = helper.NilIfEmpty(activityConfigResult.DelayInSeconds)
	activityConfigResultMap["estimated_duration_in_seconds"] = helper.IntPtrToAny(activityConfigResult.EstimatedDurationInSeconds)
	activityConfigResultMap["name"] = helper.NilIfEmpty(activityConfigResult.Name)

	if activityConfigResult.Arguments != nil {
		activityConfigResultMap["arguments"] = helper.ParseToMaps(activityConfigResult.Arguments)
	}

	if activityConfigResult.ResourceFunctionFormulas != nil {
		activityConfigResultMap["resource_function_formulas"] = helper.ParseToMaps(activityConfigResult.ResourceFunctionFormulas)
	}

	activityConfigResultMap["tags"] = helper.ParseToMaps(activityConfigResult.Tags)
	activityConfigResultMap["definition_link_status"] = helper.NilIfEmpty(activityConfigResult.DefinitionLinkStatus)

	activityConfigResultMap["invalid_definition_link_reasons"] = helper.ParseToMaps(activityConfigResult.InvalidDefinitionLinkReasons)

	return activityConfigResultMap

}

func (invalidPlanTemplateReason *InvalidPlanTemplateReason) ToMap() map[string]any {
	invalidPlanTemplateReasonMap := make(map[string]any)
	invalidPlanTemplateReasonMap["code"] = helper.NilIfEmpty(invalidPlanTemplateReason.Code)
	invalidPlanTemplateReasonMap["message"] = helper.NilIfEmpty(invalidPlanTemplateReason.Message)
	return invalidPlanTemplateReasonMap
}

func (argument *Argument) ToMap() map[string]any {
	argumentMap := make(map[string]any)
	argumentMap["name"] = helper.NilIfEmpty(argument.Name)
	argumentMap["attributes"] = []any{argument.Attributes.ToMap()}

	return argumentMap
}

func (resourceFunctionFormulaOverload *ResourceFunctionFormulaOverload) ToMap() map[string]any {
	resourceFunctionFormulaOverloadMap := make(map[string]any)
	resourceFunctionFormulaOverloadMap["resource_function_id"] = helper.NilIfEmpty(resourceFunctionFormulaOverload.ResourceFunctionId)
	resourceFunctionFormulaOverloadMap["formula"] = []map[string]any{resourceFunctionFormulaOverload.Formula.ToMap()}
	return resourceFunctionFormulaOverloadMap
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

func (invalidDefinitionLinkReason *InvalidDefinitionLinkReason) ToMap() map[string]any {
	invalidDefinitionLinkMap := make(map[string]any)
	invalidDefinitionLinkMap["code"] = helper.NilIfEmpty(invalidDefinitionLinkReason.Code)
	invalidDefinitionLinkMap["message"] = helper.NilIfEmpty(invalidDefinitionLinkReason.Message)
	return invalidDefinitionLinkMap
}

func (planTemplate *PlanTemplate) FromMap(planTemplateMap map[string]any) error {

	planTemplate.FromAuditMap(planTemplateMap)
	planTemplate.AssetId = helper.CastString(planTemplateMap, "asset_id")
	planTemplate.Name = helper.CastString(planTemplateMap, "name")
	planTemplate.Description = helper.CastString(planTemplateMap, "description")
	planTemplate.IntegrityStatus = helper.CastString(planTemplateMap, "integrity_status")

	if planTemplateMap["activity_configs"] != nil {
		if activityConfigs, err := helper.ParseFromMaps[ActivityConfigResult](
			helper.CastSlice(planTemplateMap, "activity_configs"),
		); err != nil {
			return err
		} else {
			planTemplate.ActivityConfigs = activityConfigs
		}
	}

	planTemplate.EstimatedDurationInSeconds = helper.CastInt(planTemplateMap, "estimated_duration_in_seconds")

	if planTemplateMap["invalid_plan_template_reasons"] != nil {
		if invalidPlanTemplateReason, err := helper.ParseFromMaps[InvalidPlanTemplateReason](
			helper.CastSlice(planTemplateMap, "invalid_plan_template_reasons"),
		); err != nil {
			return err
		} else {
			planTemplate.InvalidPlanTemplateReasons = invalidPlanTemplateReason
		}
	}

	return nil
}

func (activityConfigResult *ActivityConfigResult) FromMap(activityConfigResultMap map[string]any) error {

	activityConfigResult.ActivityDefinitionId = helper.CastString(activityConfigResultMap, "activity_definition_id")
	activityConfigResult.DelayReferenceOnPredecessor = helper.CastString(activityConfigResultMap, "delay_reference_on_predecessor")
	activityConfigResult.Position = helper.CastInt(activityConfigResultMap, "position")
	activityConfigResult.DelayInSeconds = helper.CastInt(activityConfigResultMap, "delay_in_seconds")
	activityConfigResult.EstimatedDurationInSeconds = helper.CastIntPtr(activityConfigResultMap, "estimated_duration_in_seconds")
	activityConfigResult.Name = helper.CastString(activityConfigResultMap, "name")

	if activityConfigResultMap["arguments"] != nil {
		if arguments, err := helper.ParseFromMaps[Argument](
			helper.CastSlice(activityConfigResultMap, "arguments"),
		); err != nil {
			return err
		} else {
			activityConfigResult.Arguments = arguments
		}
	}

	if activityConfigResultMap["resource_function_formulas"] != nil {
		if resourceFunctionFormulaOverload, err := helper.ParseFromMaps[ResourceFunctionFormulaOverload](
			helper.CastSlice(activityConfigResultMap, "resource_function_formulas"),
		); err != nil {
			return err
		} else {
			activityConfigResult.ResourceFunctionFormulas = resourceFunctionFormulaOverload
		}
	}

	if tags, err := helper.ParseFromMaps[general_objects.KeyValue](helper.CastSlice(activityConfigResultMap, "tags")); err != nil {
		return err
	} else {
		activityConfigResult.Tags = tags
	}

	activityConfigResult.DefinitionLinkStatus = helper.CastString(activityConfigResultMap, "definition_link_status")

	if invalidDefinitionLinkReasons, err := helper.ParseFromMaps[InvalidDefinitionLinkReason](helper.CastSlice(activityConfigResultMap, "invalid_definition_link_reasons")); err != nil {
		return err
	} else {
		activityConfigResult.InvalidDefinitionLinkReasons = invalidDefinitionLinkReasons
	}

	return nil
}

func (invalidPlanTemplateReason *InvalidPlanTemplateReason) FromMap(invalidPlanTemplateReasonMap map[string]any) error {
	invalidPlanTemplateReason.Code = helper.CastString(invalidPlanTemplateReasonMap, "code")
	invalidPlanTemplateReason.Message = helper.CastString(invalidPlanTemplateReasonMap, "message")
	return nil
}

func (argument *Argument) FromMap(argumentMap map[string]any) error {

	argument.Name = helper.CastString(argumentMap, "name")

	if len(helper.CastSlice(argumentMap, "attributes")) > 0 {
		if err := argument.Attributes.FromMap(helper.CastSlice(argumentMap, "attributes")[0].(map[string]any)); err != nil {
			return err
		}
	}

	return nil
}

func (resourceFunctionFormulaOverload *ResourceFunctionFormulaOverload) FromMap(resourceFunctionFormulaOverloadMap map[string]any) error {
	resourceFunctionFormulaOverload.ResourceFunctionId = helper.CastString(resourceFunctionFormulaOverloadMap, "resource_function_id")

	if len(helper.CastSlice(resourceFunctionFormulaOverloadMap, "formula")) > 0 && helper.CastSlice(resourceFunctionFormulaOverloadMap, "formula")[0] != nil {
		resourceFunctionFormulaOverload.Formula = new(ResourceFunctionFormula)
		if err := resourceFunctionFormulaOverload.Formula.FromMap(helper.CastSlice(resourceFunctionFormulaOverloadMap, "formula")[0].(map[string]any)); err != nil {
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

func (invalidDefinitionLinkReason *InvalidDefinitionLinkReason) FromMap(invalidDefinitionLinReasonMap map[string]any) error {
	invalidDefinitionLinkReason.Code = helper.CastString(invalidDefinitionLinReasonMap, "code")
	invalidDefinitionLinkReason.Message = helper.CastString(invalidDefinitionLinReasonMap, "message")
	return nil
}
