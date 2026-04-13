package plan_templates

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

type PlanTemplateTF struct {
	general_objects.AuditModelTF
	AssetId                    types.String                     `tfsdk:"asset_id"`
	Name                       types.String                     `tfsdk:"name"`
	Description                types.String                     `tfsdk:"description"`
	IntegrityStatus            types.String                     `tfsdk:"integrity_status"`
	ActivityConfigs            []ActivityConfigResultTF         `tfsdk:"activity_configs"`
	EstimatedDurationInSeconds types.Int64                      `tfsdk:"estimated_duration_in_seconds"`
	InvalidPlanTemplateReasons []InvalidPlanTemplateReasonTF    `tfsdk:"invalid_plan_template_reasons"`
}

type ActivityConfigResultTF struct {
	ActivityDefinitionId         types.String                        `tfsdk:"activity_definition_id"`
	DelayReferenceOnPredecessor  types.String                        `tfsdk:"delay_reference_on_predecessor"`
	Position                     types.Int64                         `tfsdk:"position"`
	DelayInSeconds               types.Int64                         `tfsdk:"delay_in_seconds"`
	EstimatedDurationInSeconds   types.Int64                         `tfsdk:"estimated_duration_in_seconds"`
	Name                         types.String                        `tfsdk:"name"`
	Arguments                    []ArgumentTF                        `tfsdk:"arguments"`
	ResourceFunctionFormulas     []ResourceFunctionFormulaOverloadTF `tfsdk:"resource_function_formulas"`
	Tags                         []general_objects.KeyValueTF        `tfsdk:"tags"`
	DefinitionLinkStatus         types.String                        `tfsdk:"definition_link_status"`
	InvalidDefinitionLinkReasons []InvalidDefinitionLinkReasonTF     `tfsdk:"invalid_definition_link_reasons"`
}

type InvalidPlanTemplateReasonTF struct {
	Code    types.String `tfsdk:"code"`
	Message types.String `tfsdk:"message"`
}

type ArgumentTF struct {
	Name       types.String                       `tfsdk:"name"`
	Attributes []general_objects.ValueAttributeTF `tfsdk:"attributes"`
}

type ResourceFunctionFormulaOverloadTF struct {
	ResourceFunctionId types.String                  `tfsdk:"resource_function_id"`
	Formula            []ResourceFunctionFormulaTF   `tfsdk:"formula"`
}

type ResourceFunctionFormulaTF struct {
	Type      types.String  `tfsdk:"type"`
	Amplitude types.Float64 `tfsdk:"amplitude"`
	Constant  types.Float64 `tfsdk:"constant"`
	Rate      types.Float64 `tfsdk:"rate"`
	TimeUnit  types.String  `tfsdk:"time_unit"`
}

type InvalidDefinitionLinkReasonTF struct {
	Code    types.String `tfsdk:"code"`
	Message types.String `tfsdk:"message"`
}

func (x *PlanTemplate) ToTF() any {
	activityConfigs := make([]ActivityConfigResultTF, len(x.ActivityConfigs))
	for i, ac := range x.ActivityConfigs {
		arguments := make([]ArgumentTF, len(ac.Arguments))
		for j, arg := range ac.Arguments {
			attr := general_objects.ValueAttributeToTF(&arg.Attributes)
			arguments[j] = ArgumentTF{
				Name:       helper.TFStringValue(arg.Name),
				Attributes: []general_objects.ValueAttributeTF{attr},
			}
		}

		formulas := make([]ResourceFunctionFormulaOverloadTF, len(ac.ResourceFunctionFormulas))
		for j, rf := range ac.ResourceFunctionFormulas {
			var formulaSlice []ResourceFunctionFormulaTF
			if rf.Formula != nil {
				formulaSlice = []ResourceFunctionFormulaTF{{
					Type:      helper.TFStringValue(rf.Formula.Type),
					Amplitude: helper.TFFloat64Value(rf.Formula.Amplitude),
					Constant:  helper.TFFloat64Value(rf.Formula.Constant),
					Rate:      helper.TFFloat64Value(rf.Formula.Rate),
					TimeUnit:  helper.TFStringValue(rf.Formula.TimeUnit),
				}}
			}
			formulas[j] = ResourceFunctionFormulaOverloadTF{
				ResourceFunctionId: helper.TFStringValue(rf.ResourceFunctionId),
				Formula:            formulaSlice,
			}
		}

		invalidReasons := make([]InvalidDefinitionLinkReasonTF, len(ac.InvalidDefinitionLinkReasons))
		for j, r := range ac.InvalidDefinitionLinkReasons {
			invalidReasons[j] = InvalidDefinitionLinkReasonTF{
				Code:    helper.TFStringValue(r.Code),
				Message: helper.TFStringValue(r.Message),
			}
		}

		activityConfigs[i] = ActivityConfigResultTF{
			ActivityDefinitionId:         helper.TFStringValue(ac.ActivityDefinitionId),
			DelayReferenceOnPredecessor:  helper.TFStringValue(ac.DelayReferenceOnPredecessor),
			Position:                     helper.TFInt64Value(ac.Position),
			DelayInSeconds:               helper.TFInt64Value(ac.DelayInSeconds),
			EstimatedDurationInSeconds:   helper.TFIntPtrValue(ac.EstimatedDurationInSeconds),
			Name:                         helper.TFStringValue(ac.Name),
			Arguments:                    arguments,
			ResourceFunctionFormulas:     formulas,
			Tags:                         general_objects.KeyValuesToTF(ac.Tags),
			DefinitionLinkStatus:         helper.TFStringValue(ac.DefinitionLinkStatus),
			InvalidDefinitionLinkReasons: invalidReasons,
		}
	}

	invalidReasons := make([]InvalidPlanTemplateReasonTF, len(x.InvalidPlanTemplateReasons))
	for i, r := range x.InvalidPlanTemplateReasons {
		invalidReasons[i] = InvalidPlanTemplateReasonTF{
			Code:    helper.TFStringValue(r.Code),
			Message: helper.TFStringValue(r.Message),
		}
	}

	return &PlanTemplateTF{
		AuditModelTF:               general_objects.AuditModelToTF(&x.AuditModel),
		AssetId:                    helper.TFStringValue(x.AssetId),
		Name:                       helper.TFStringValue(x.Name),
		Description:                helper.TFStringValue(x.Description),
		IntegrityStatus:            helper.TFStringValue(x.IntegrityStatus),
		ActivityConfigs:            activityConfigs,
		EstimatedDurationInSeconds: helper.TFInt64Value(x.EstimatedDurationInSeconds),
		InvalidPlanTemplateReasons: invalidReasons,
	}
}

func (tf *PlanTemplateTF) ToAPI() any {
	activityConfigs := make([]ActivityConfigResult, len(tf.ActivityConfigs))
	for i, ac := range tf.ActivityConfigs {
		arguments := make([]Argument, len(ac.Arguments))
		for j, arg := range ac.Arguments {
			arguments[j] = Argument{Name: helper.FromTFString(arg.Name)}
			if len(arg.Attributes) > 0 {
				arguments[j].Attributes = general_objects.ValueAttributeFromTF(arg.Attributes[0])
			}
		}

		formulas := make([]ResourceFunctionFormulaOverload, len(ac.ResourceFunctionFormulas))
		for j, rf := range ac.ResourceFunctionFormulas {
			formulas[j] = ResourceFunctionFormulaOverload{
				ResourceFunctionId: helper.FromTFString(rf.ResourceFunctionId),
			}
			if len(rf.Formula) > 0 {
				f := rf.Formula[0]
				formulas[j].Formula = &ResourceFunctionFormula{
					Type:      helper.FromTFString(f.Type),
					Amplitude: helper.FromTFFloat64(f.Amplitude),
					Constant:  helper.FromTFFloat64(f.Constant),
					Rate:      helper.FromTFFloat64(f.Rate),
					TimeUnit:  helper.FromTFString(f.TimeUnit),
				}
			}
		}

		invalidReasons := make([]InvalidDefinitionLinkReason, len(ac.InvalidDefinitionLinkReasons))
		for j, r := range ac.InvalidDefinitionLinkReasons {
			invalidReasons[j] = InvalidDefinitionLinkReason{
				Code:    helper.FromTFString(r.Code),
				Message: helper.FromTFString(r.Message),
			}
		}

		activityConfigs[i] = ActivityConfigResult{
			ActivityDefinitionId:         helper.FromTFString(ac.ActivityDefinitionId),
			DelayReferenceOnPredecessor:  helper.FromTFString(ac.DelayReferenceOnPredecessor),
			Position:                     helper.FromTFInt64(ac.Position),
			DelayInSeconds:               helper.FromTFInt64(ac.DelayInSeconds),
			EstimatedDurationInSeconds:   helper.FromTFIntPtr(ac.EstimatedDurationInSeconds),
			Name:                         helper.FromTFString(ac.Name),
			Arguments:                    arguments,
			ResourceFunctionFormulas:     formulas,
			Tags:                         general_objects.KeyValuesFromTF(ac.Tags),
			DefinitionLinkStatus:         helper.FromTFString(ac.DefinitionLinkStatus),
			InvalidDefinitionLinkReasons: invalidReasons,
		}
	}

	invalidReasons := make([]InvalidPlanTemplateReason, len(tf.InvalidPlanTemplateReasons))
	for i, r := range tf.InvalidPlanTemplateReasons {
		invalidReasons[i] = InvalidPlanTemplateReason{
			Code:    helper.FromTFString(r.Code),
			Message: helper.FromTFString(r.Message),
		}
	}

	return &PlanTemplate{
		AuditModel:                 general_objects.AuditModelFromTF(tf.AuditModelTF),
		AssetId:                    helper.FromTFString(tf.AssetId),
		Name:                       helper.FromTFString(tf.Name),
		Description:                helper.FromTFString(tf.Description),
		IntegrityStatus:            helper.FromTFString(tf.IntegrityStatus),
		ActivityConfigs:            activityConfigs,
		EstimatedDurationInSeconds: helper.FromTFInt64(tf.EstimatedDurationInSeconds),
		InvalidPlanTemplateReasons: invalidReasons,
	}
}
