package plan_templates

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

func (planTemplate *PlanTemplate) ToMap() map[string]any {
	planTemplateMap := make(map[string]any)
	planTemplateMap["id"] = planTemplate.ID
	planTemplateMap["asset_id"] = planTemplate.AssetId
	planTemplateMap["name"] = planTemplate.Name
	planTemplateMap["description"] = planTemplate.Description
	planTemplateMap["integrity_status"] = planTemplate.IntegrityStatus

	if planTemplate.ActivityConfigs != nil {
		planTemplateMap["activity_configs"] = helper.ParseToMaps(planTemplate.ActivityConfigs)
	}

	planTemplateMap["estimated_duration_in_seconds"] = planTemplate.EstimatedDurationInSeconds

	if planTemplate.InvalidPlanTemplateReasons != nil {
		planTemplateMap["invalid_plan_template_reasons"] = helper.ParseToMaps(planTemplate.InvalidPlanTemplateReasons)
	}

	planTemplateMap["created_at"] = planTemplate.CreatedAt
	planTemplateMap["created_by"] = planTemplate.CreatedBy
	planTemplateMap["last_modified_at"] = planTemplate.LastModifiedAt
	planTemplateMap["last_modified_by"] = planTemplate.LastModifiedBy

	helper.Logger.Printf("%s", planTemplateMap)
	return planTemplateMap
}

func (activityConfigResult *ActivityConfigResult) ToMap() map[string]any {
	activityConfigResultMap := make(map[string]any)
	activityConfigResultMap["activity_definition_id"] = activityConfigResult.ActivityDefinitionId
	activityConfigResultMap["delay_reference_on_predecessor"] = activityConfigResult.DelayReferenceOnPredecessor
	activityConfigResultMap["position"] = activityConfigResult.Position
	activityConfigResultMap["delay_in_seconds"] = activityConfigResult.DelayInSeconds
	activityConfigResultMap["estimated_duration_in_seconds"] = activityConfigResult.EstimatedDurationInSeconds
	activityConfigResultMap["name"] = activityConfigResult.Name

	if activityConfigResult.Arguments != nil {
		activityConfigResultMap["arguments"] = helper.ParseToMaps(activityConfigResult.Arguments)
	}

	if activityConfigResult.ResourceFunctionFormulas != nil {
		activityConfigResultMap["resource_function_formulas"] = helper.ParseToMaps(activityConfigResult.ResourceFunctionFormulas)
	}

	activityConfigResultMap["tags"] = helper.ParseToMaps(activityConfigResult.Tags)
	activityConfigResultMap["definition_link_status"] = activityConfigResult.DefinitionLinkStatus

	activityConfigResultMap["invalid_definition_link_reasons"] = helper.ParseToMaps(activityConfigResult.InvalidDefinitionLinkReasons)

	return activityConfigResultMap

}

func (invalidPlanTemplateReason *InvalidPlanTemplateReason) ToMap() map[string]any {
	invalidPlanTemplateReasonMap := make(map[string]any)
	invalidPlanTemplateReasonMap["code"] = invalidPlanTemplateReason.Code
	invalidPlanTemplateReasonMap["message"] = invalidPlanTemplateReason.Message
	return invalidPlanTemplateReasonMap
}

func (argument *Argument) ToMap() map[string]any {
	argumentMap := make(map[string]any)
	argumentMap["name"] = argument.Name
	argumentMap["attributes"] = []any{argument.Attributes.ToMap()}

	return argumentMap
}

func (resourceFunctionFormulaOverload *ResourceFunctionFormulaOverload) ToMap() map[string]any {
	resourceFunctionFormulaOverloadMap := make(map[string]any)
	resourceFunctionFormulaOverloadMap["resource_function_id"] = resourceFunctionFormulaOverload.ResourceFunctionId
	resourceFunctionFormulaOverloadMap["formula"] = []map[string]any{resourceFunctionFormulaOverload.Formula.ToMap()}
	return resourceFunctionFormulaOverloadMap
}

func (formula *ResourceFunctionFormula) ToMap() map[string]any {
	formulaMap := make(map[string]any)
	formulaMap["type"] = formula.Type
	formulaMap["constant"] = formula.Constant
	formulaMap["rate"] = formula.Rate
	return formulaMap
}

func (invalidDefinitionLinkReason *InvalidDefinitionLinkReason) ToMap() map[string]any {
	invalidDefinitionLinkMap := make(map[string]any)
	invalidDefinitionLinkMap["code"] = invalidDefinitionLinkReason.Code
	invalidDefinitionLinkMap["message"] = invalidDefinitionLinkReason.Message
	return invalidDefinitionLinkMap
}

func (planTemplate *PlanTemplate) FromMap(planTemplateMap map[string]any) error {

	planTemplate.ID = planTemplateMap["id"].(string)
	planTemplate.AssetId = planTemplateMap["asset_id"].(string)
	planTemplate.Name = planTemplateMap["name"].(string)
	planTemplate.Description = planTemplateMap["description"].(string)
	planTemplate.IntegrityStatus = planTemplateMap["integrity_status"].(string)

	if planTemplateMap["activity_configs"] != nil {
		if activityConfigs, err := helper.ParseFromMaps[ActivityConfigResult](
			planTemplateMap["activity_configs"].([]any),
		); err != nil {
			return err
		} else {
			planTemplate.ActivityConfigs = activityConfigs
		}
	}

	planTemplate.EstimatedDurationInSeconds = planTemplateMap["estimated_duration_in_seconds"].(int)

	if planTemplateMap["invalid_plan_template_reasons"] != nil {
		if invalidPlanTemplateReason, err := helper.ParseFromMaps[InvalidPlanTemplateReason](
			planTemplateMap["invalid_plan_template_reasons"].([]any),
		); err != nil {
			return err
		} else {
			planTemplate.InvalidPlanTemplateReasons = invalidPlanTemplateReason
		}
	}

	planTemplate.CreatedAt = planTemplateMap["created_at"].(string)
	planTemplate.CreatedBy = planTemplateMap["created_by"].(string)
	planTemplate.LastModifiedAt = planTemplateMap["last_modified_at"].(string)
	planTemplate.LastModifiedBy = planTemplateMap["last_modified_by"].(string)

	return nil
}

func (activityConfigResult *ActivityConfigResult) FromMap(activityConfigResultMap map[string]any) error {

	activityConfigResult.ActivityDefinitionId = activityConfigResultMap["activity_definition_id"].(string)
	activityConfigResult.DelayReferenceOnPredecessor = activityConfigResultMap["delay_reference_on_predecessor"].(string)
	activityConfigResult.Position = activityConfigResultMap["position"].(int)
	activityConfigResult.DelayInSeconds = activityConfigResultMap["delay_in_seconds"].(int)
	activityConfigResult.EstimatedDurationInSeconds = activityConfigResultMap["estimated_duration_in_seconds"].(int)
	activityConfigResult.Name = activityConfigResultMap["name"].(string)

	if activityConfigResultMap["arguments"] != nil {
		if arguments, err := helper.ParseFromMaps[Argument](
			activityConfigResultMap["arguments"].(*schema.Set).List(),
		); err != nil {
			return err
		} else {
			activityConfigResult.Arguments = arguments
		}
	}

	if activityConfigResultMap["resource_function_formulas"] != nil {
		if resourceFunctionFormulaOverload, err := helper.ParseFromMaps[ResourceFunctionFormulaOverload](
			activityConfigResultMap["resource_function_formulas"].(*schema.Set).List(),
		); err != nil {
			return err
		} else {
			activityConfigResult.ResourceFunctionFormulas = resourceFunctionFormulaOverload
		}
	}

	if tags, err := helper.ParseFromMaps[general_objects.KeyValue](activityConfigResultMap["tags"].(*schema.Set).List()); err != nil {
		return err
	} else {
		activityConfigResult.Tags = tags
	}

	activityConfigResult.DefinitionLinkStatus = activityConfigResultMap["definition_link_status"].(string)

	if invalidDefinitionLinkReasons, err := helper.ParseFromMaps[InvalidDefinitionLinkReason](activityConfigResultMap["invalid_definition_link_reasons"].(*schema.Set).List()); err != nil {
		return err
	} else {
		activityConfigResult.InvalidDefinitionLinkReasons = invalidDefinitionLinkReasons
	}

	return nil
}

func (invalidPlanTemplateReason *InvalidPlanTemplateReason) FromMap(invalidPlanTemplateReasonMap map[string]any) error {
	invalidPlanTemplateReason.Code = invalidPlanTemplateReasonMap["code"].(string)
	invalidPlanTemplateReason.Message = invalidPlanTemplateReasonMap["message"].(string)
	return nil
}

func (argument *Argument) FromMap(argumentMap map[string]any) error {

	argument.Name = argumentMap["name"].(string)

	if len(argumentMap["attributes"].(*schema.Set).List()) > 0 {
		if err := argument.Attributes.FromMap(argumentMap["attributes"].(*schema.Set).List()[0].(map[string]any)); err != nil {
			return err
		}
	}

	return nil
}

func (resourceFunctionFormulaOverload *ResourceFunctionFormulaOverload) FromMap(resourceFunctionFormulaOverloadMap map[string]any) error {
	resourceFunctionFormulaOverload.ResourceFunctionId = resourceFunctionFormulaOverloadMap["resource_function_id"].(string)

	if len(resourceFunctionFormulaOverloadMap["formula"].(*schema.Set).List()) > 0 && resourceFunctionFormulaOverloadMap["formula"].(*schema.Set).List()[0] != nil {
		resourceFunctionFormulaOverload.Formula = new(ResourceFunctionFormula)
		if err := resourceFunctionFormulaOverload.Formula.FromMap(resourceFunctionFormulaOverloadMap["formula"].(*schema.Set).List()[0].(map[string]any)); err != nil {
			return err
		}
	}

	return nil
}

func (formula *ResourceFunctionFormula) FromMap(formulaMap map[string]any) error {
	formula.Type = formulaMap["type"].(string)
	formula.Constant = formulaMap["constant"].(float64)
	formula.Rate = formulaMap["rate"].(float64)
	return nil
}

func (invalidDefinitionLinkReason *InvalidDefinitionLinkReason) FromMap(invalidDefinitionLinReasonMap map[string]any) error {
	invalidDefinitionLinkReason.Code = invalidDefinitionLinReasonMap["code"].(string)
	invalidDefinitionLinkReason.Message = invalidDefinitionLinReasonMap["message"].(string)
	return nil
}
