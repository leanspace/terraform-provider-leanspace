package plan_templates

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

func (planTemplate *PlanTemplate) ToMap() map[string]any {
	planTemplateMap := make(map[string]any)
	planTemplateMap["id"] = planTemplate.ID
	planTemplateMap["assetId"] = planTemplate.AssetId
	planTemplateMap["name"] = planTemplate.Name
	planTemplateMap["description"] = planTemplate.Description
	planTemplateMap["integrityStatus"] = planTemplate.IntegrityStatus

	if planTemplate.ActivityConfigs != nil {
		planTemplateMap["activityConfigs"] = helper.ParseToMaps(planTemplate.ActivityConfigs)
	}

	planTemplateMap["estimatedDurationInSeconds"] = planTemplate.EstimatedDurationInSeconds

	if planTemplate.InvalidPlanTemplateReasons != nil {
		planTemplateMap["invalidPlanTemplateReasons"] = helper.ParseToMaps(planTemplate.InvalidPlanTemplateReasons)
	}

	planTemplateMap["createdAt"] = planTemplate.CreatedAt
	planTemplateMap["createdBy"] = planTemplate.CreatedBy
	planTemplateMap["lastModifiedAt"] = planTemplate.LastModifiedAt
	planTemplateMap["lastModifiedBy"] = planTemplate.LastModifiedBy

	return planTemplateMap
}

func (planTemplate *PlanTemplate) FromMap(planTemplateMap map[string]any) error {

	planTemplate.ID = planTemplateMap["id"].(string)
	planTemplate.AssetId = planTemplateMap["assetId"].(string)
	planTemplate.Name = planTemplateMap["name"].(string)
	planTemplate.Description = planTemplateMap["description"].(string)
	planTemplate.IntegrityStatus = planTemplateMap["integrityStatus"].(string)

	if planTemplateMap["activityConfigs"] != nil {
		if activityConfig, err := helper.ParseFromMaps[ActivityConfigResult](
			planTemplateMap["activityConfigs"].(*schema.Set).List(),
		); err != nil {
			return err
		} else {
			planTemplate.ActivityConfigs = activityConfig
		}
	}

	planTemplate.EstimatedDurationInSeconds = planTemplateMap["estimatedDurationInSeconds"].(int)

	if planTemplateMap["invalidPlanTemplateReasons"] != nil {
		if invalidPlanTemplateReason, err := helper.ParseFromMaps[InvalidPlanTemplateReason](
			planTemplateMap["invalidPlanTemplateReasons"].(*schema.Set).List(),
		); err != nil {
			return err
		} else {
			planTemplate.InvalidPlanTemplateReasons = invalidPlanTemplateReason
		}
	}

	planTemplate.CreatedAt = planTemplateMap["createdAt"].(string)
	planTemplate.CreatedBy = planTemplateMap["createdBy"].(string)
	planTemplate.LastModifiedAt = planTemplateMap["lastModifiedAt"].(string)
	planTemplate.LastModifiedBy = planTemplateMap["lastModifiedBy"].(string)

	return nil
}

func (activityConfigResult *ActivityConfigResult) ToMap() map[string]any {
	activityConfigResultMap := make(map[string]any)
	activityConfigResultMap["activityDefinitionId"] = activityConfigResult.ActivityDefinitionId
	activityConfigResultMap["delayReferenceOnPredecessor"] = activityConfigResult.DelayReferenceOnPredecessor
	activityConfigResultMap["position"] = activityConfigResult.Position
	activityConfigResultMap["delayInSeconds"] = activityConfigResult.DelayInSeconds
	activityConfigResultMap["estimatedDurationInSeconds"] = activityConfigResult.EstimatedDurationInSeconds
	activityConfigResultMap["name"] = activityConfigResult.Name

	if activityConfigResult.Arguments != nil {
		activityConfigResultMap["arguments"] = helper.ParseToMaps(activityConfigResult.Arguments)
	}

	activityConfigResultMap["resourceFunctionFormulas"] = []any{activityConfigResult.ResourceFunctionFormulas.ToMap()}
	activityConfigResultMap["tags"] = activityConfigResult.Tags
	activityConfigResultMap["definitionLinkStatus"] = activityConfigResult.DefinitionLinkStatus

	if activityConfigResult.InvalidDefinitionLinkReasons != nil {
		activityConfigResultMap["invalidDefinitionLinkReasons"] = helper.ParseToMaps(activityConfigResult.InvalidDefinitionLinkReasons)
	}

	return activityConfigResultMap

}

func (activityConfigResult *ActivityConfigResult) FromMap(activityConfigResultMap map[string]any) error {

	activityConfigResult.ActivityDefinitionId = activityConfigResultMap["activityDefinitionId"].(string)
	activityConfigResult.DelayReferenceOnPredecessor = activityConfigResultMap["delayReferenceOnPredecessor"].(string)
	activityConfigResult.Position = activityConfigResultMap["position"].(int)
	activityConfigResult.DelayInSeconds = activityConfigResultMap["delayInSeconds"].(int)
	activityConfigResult.EstimatedDurationInSeconds = activityConfigResultMap["estimatedDurationInSeconds"].(int)
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

	if len(activityConfigResultMap["resourceFunctionFormulas"].([]any)) > 0 {
		if err := activityConfigResult.ResourceFunctionFormulas.FromMap(activityConfigResultMap["resourceFunctionFormulas"].([]any)[0].(map[string]any)); err != nil {
			return err
		}
	}

	if activityConfigResultMap["tags"] != nil {
		if tags, err := helper.ParseFromMaps[general_objects.KeyValue](
			activityConfigResultMap["tags"].(*schema.Set).List(),
		); err != nil {
			return err
		} else {
			activityConfigResult.Tags = tags
		}
	}

	activityConfigResult.DefinitionLinkStatus = activityConfigResultMap["definitionLinkStatus"].(string)

	if activityConfigResultMap["invalidDefinitionLinkReasons"] != nil {
		if invalidDefinitionLinkReasons, err := helper.ParseFromMaps[Argument](
			activityConfigResultMap["invalidDefinitionLinkReasons"].(*schema.Set).List(),
		); err != nil {
			return err
		} else {
			activityConfigResult.Arguments = invalidDefinitionLinkReasons
		}
	}

	return nil
}

func (invalidPlanTemplateReason *InvalidPlanTemplateReason) ToMap() map[string]any {
	invalidPlanTemplateReasonMap := make(map[string]any)
	invalidPlanTemplateReasonMap["code"] = invalidPlanTemplateReason.Code
	invalidPlanTemplateReasonMap["message"] = invalidPlanTemplateReason.Message
	return invalidPlanTemplateReasonMap
}

func (invalidPlanTemplateReason *InvalidPlanTemplateReason) FromMap(invalidPlanTemplateReasonMap map[string]any) error {
	invalidPlanTemplateReason.Code = invalidPlanTemplateReasonMap["code"].(string)
	invalidPlanTemplateReason.Message = invalidPlanTemplateReasonMap["message"].(string)
	return nil
}

func (argument *Argument) ToMap() map[string]any {
	argumentMap := make(map[string]any)
	argumentMap["name"] = argument.Name

	if argument.Attributes != nil {
		argumentMap["attributes"] = helper.ParseToMaps(argument.Attributes)
	}

	return argumentMap
}

func (argument *Argument) FromMap(argumentMap map[string]any) error {

	argument.Name = argumentMap["name"].(string)

	if argumentMap["attributes"] != nil {
		if attributes, err := helper.ParseFromMaps[general_objects.ValueAttribute[any]](
			argumentMap["attributes"].(*schema.Set).List(),
		); err != nil {
			return err
		} else {
			argument.Attributes = attributes
		}
	}

	return nil
}

func (resourceFunctionFormulaOverload *ResourceFunctionFormulaOverload) ToMap() map[string]any {
	resourceFunctionFormulaOverloadMap := make(map[string]any)
	resourceFunctionFormulaOverloadMap["resourceFunctionId"] = resourceFunctionFormulaOverload.ResourceFunctionId
	resourceFunctionFormulaOverloadMap["formula"] = []any{resourceFunctionFormulaOverload.Formula.ToMap()}
	return resourceFunctionFormulaOverloadMap
}

func (invalidDefinitionLinkReason *ResourceFunctionFormulaOverload) FromMap(invalidDefinitionLinkMap map[string]any) error {
	invalidDefinitionLinkReason.ResourceFunctionId = invalidDefinitionLinkMap["resourceFunctionId"].(string)

	if len(invalidDefinitionLinkMap["formula"].([]any)) > 0 {
		if err := invalidDefinitionLinkReason.Formula.FromMap(invalidDefinitionLinkMap["formula"].([]any)[0].(map[string]any)); err != nil {
			return err
		}
	}
	return nil
}

func (invalidDefinitionLinkReason *ResourceFunctionFormula) ToMap() map[string]any {
	invalidDefinitionLinkMap := make(map[string]any)
	invalidDefinitionLinkMap["type"] = invalidDefinitionLinkReason.Type
	return invalidDefinitionLinkMap
}

func (invalidDefinitionLinkReason *ResourceFunctionFormula) FromMap(invalidDefinitionLinkMap map[string]any) error {
	invalidDefinitionLinkReason.Type = invalidDefinitionLinkMap["type"].(string)
	return nil
}

func (invalidDefinitionLinkReason *InvalidDefinitionLinkReason) ToMap() map[string]any {
	invalidDefinitionLinkMap := make(map[string]any)
	invalidDefinitionLinkMap["code"] = invalidDefinitionLinkReason.Code
	invalidDefinitionLinkMap["message"] = invalidDefinitionLinkReason.Message
	return invalidDefinitionLinkMap
}

func (invalidDefinitionLinkReason *InvalidDefinitionLinkReason) FromMap(invalidDefinitionLinkMap map[string]any) error {
	invalidDefinitionLinkReason.Code = invalidDefinitionLinkMap["code"].(string)
	invalidDefinitionLinkReason.Message = invalidDefinitionLinkMap["message"].(string)
	return nil
}
