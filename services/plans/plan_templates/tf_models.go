package plan_templates

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

var planReasonAttrTypes = map[string]attr.Type{
	"code":    types.StringType,
	"message": types.StringType,
}

type PlanTemplateTF struct {
	general_objects.AuditModelTF
	AssetId                    types.String             `tfsdk:"asset_id"`
	Name                       types.String             `tfsdk:"name"`
	Description                types.String             `tfsdk:"description"`
	IntegrityStatus            types.String             `tfsdk:"integrity_status"`
	ActivityConfigs            []ActivityConfigResultTF `tfsdk:"activity_configs"`
	EstimatedDurationInSeconds types.Int64              `tfsdk:"estimated_duration_in_seconds"`
	InvalidPlanTemplateReasons types.List               `tfsdk:"invalid_plan_template_reasons"`
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
	InvalidDefinitionLinkReasons types.List                          `tfsdk:"invalid_definition_link_reasons"`
}

type ArgumentTF struct {
	Name       types.String                       `tfsdk:"name"`
	Attributes []general_objects.ValueAttributeTF `tfsdk:"attributes"`
}

type ResourceFunctionFormulaOverloadTF struct {
	ResourceFunctionId types.String                `tfsdk:"resource_function_id"`
	Formula            []ResourceFunctionFormulaTF `tfsdk:"formula"`
}

type ResourceFunctionFormulaTF struct {
	Type      types.String  `tfsdk:"type"`
	Amplitude types.Float64 `tfsdk:"amplitude"`
	Constant  types.Float64 `tfsdk:"constant"`
	Rate      types.Float64 `tfsdk:"rate"`
	TimeUnit  types.String  `tfsdk:"time_unit"`
}

func (x *PlanTemplate) ToTF() any {
	activityConfigs := make([]ActivityConfigResultTF, len(x.ActivityConfigs))
	for i, ac := range x.ActivityConfigs {
		arguments := make([]ArgumentTF, len(ac.Arguments))
		for j, arg := range ac.Arguments {
			attrVal := general_objects.ValueAttributeToTF(&arg.Attributes)
			arguments[j] = ArgumentTF{
				Name:       helper.TFStringValue(arg.Name),
				Attributes: []general_objects.ValueAttributeTF{attrVal},
			}
		}

		formulas := make([]ResourceFunctionFormulaOverloadTF, len(ac.ResourceFunctionFormulas))
		for j, rf := range ac.ResourceFunctionFormulas {
			var formulaSlice []ResourceFunctionFormulaTF
			if rf.Formula != nil {
				formulaSlice = []ResourceFunctionFormulaTF{{
					Type:      helper.TFStringValue(rf.Formula.Type),
					Amplitude: helper.TFFloat64PtrValue(rf.Formula.Amplitude),
					Constant:  helper.TFFloat64PtrValue(rf.Formula.Constant),
					Rate:      helper.TFFloat64PtrValue(rf.Formula.Rate),
					TimeUnit:  helper.TFStringPtrValue(rf.Formula.TimeUnit),
				}}
			}
			formulas[j] = ResourceFunctionFormulaOverloadTF{
				ResourceFunctionId: helper.TFStringPtrValue(rf.ResourceFunctionId),
				Formula:            formulaSlice,
			}
		}

		defLinkReasonElems := make([]attr.Value, len(ac.InvalidDefinitionLinkReasons))
		for j, r := range ac.InvalidDefinitionLinkReasons {
			rObj, _ := types.ObjectValue(planReasonAttrTypes, map[string]attr.Value{
				"code":    helper.TFStringValue(r.Code),
				"message": helper.TFStringValue(r.Message),
			})
			defLinkReasonElems[j] = rObj
		}
		defLinkReasons := types.ListValueMust(types.ObjectType{AttrTypes: planReasonAttrTypes}, defLinkReasonElems)

		activityConfigs[i] = ActivityConfigResultTF{
			ActivityDefinitionId:         helper.TFStringValue(ac.ActivityDefinitionId),
			DelayReferenceOnPredecessor:  helper.TFStringPtrValue(ac.DelayReferenceOnPredecessor),
			Position:                     helper.TFInt64Value(ac.Position),
			DelayInSeconds:               helper.TFInt64Value(ac.DelayInSeconds),
			EstimatedDurationInSeconds:   helper.TFIntPtrValue(ac.EstimatedDurationInSeconds),
			Name:                         helper.TFStringPtrValue(ac.Name),
			Arguments:                    arguments,
			ResourceFunctionFormulas:     formulas,
			Tags:                         general_objects.KeyValuesToTF(ac.Tags),
			DefinitionLinkStatus:         helper.TFStringPtrValue(ac.DefinitionLinkStatus),
			InvalidDefinitionLinkReasons: defLinkReasons,
		}
	}

	ptReasonElems := make([]attr.Value, len(x.InvalidPlanTemplateReasons))
	for i, r := range x.InvalidPlanTemplateReasons {
		rObj, _ := types.ObjectValue(planReasonAttrTypes, map[string]attr.Value{
			"code":    helper.TFStringValue(r.Code),
			"message": helper.TFStringValue(r.Message),
		})
		ptReasonElems[i] = rObj
	}
	invalidPlanTemplateReasons := types.ListValueMust(types.ObjectType{AttrTypes: planReasonAttrTypes}, ptReasonElems)

	return &PlanTemplateTF{
		AuditModelTF:               general_objects.AuditModelToTF(&x.AuditModel),
		AssetId:                    helper.TFStringValue(x.AssetId),
		Name:                       helper.TFStringValue(x.Name),
		Description:                helper.TFStringPtrValue(x.Description),
		IntegrityStatus:            helper.TFStringValue(x.IntegrityStatus),
		ActivityConfigs:            activityConfigs,
		EstimatedDurationInSeconds: helper.TFInt64Value(x.EstimatedDurationInSeconds),
		InvalidPlanTemplateReasons: invalidPlanTemplateReasons,
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
				ResourceFunctionId: helper.FromTFStringPtr(rf.ResourceFunctionId),
			}
			if len(rf.Formula) > 0 {
				f := rf.Formula[0]
				formulas[j].Formula = &ResourceFunctionFormula{
					Type:      helper.FromTFString(f.Type),
					Amplitude: helper.FromTFFloat64Ptr(f.Amplitude),
					Constant:  helper.FromTFFloat64Ptr(f.Constant),
					Rate:      helper.FromTFFloat64Ptr(f.Rate),
					TimeUnit:  helper.FromTFStringPtr(f.TimeUnit),
				}
			}
		}

		activityConfigs[i] = ActivityConfigResult{
			ActivityDefinitionId:        helper.FromTFString(ac.ActivityDefinitionId),
			DelayReferenceOnPredecessor: helper.FromTFStringPtr(ac.DelayReferenceOnPredecessor),
			Position:                    helper.FromTFInt64(ac.Position),
			DelayInSeconds:              helper.FromTFInt64(ac.DelayInSeconds),
			EstimatedDurationInSeconds:  helper.FromTFIntPtr(ac.EstimatedDurationInSeconds),
			Name:                        helper.FromTFStringPtr(ac.Name),
			Arguments:                   arguments,
			ResourceFunctionFormulas:    formulas,
			Tags:                        general_objects.KeyValuesFromTF(ac.Tags),
			DefinitionLinkStatus:        helper.FromTFStringPtr(ac.DefinitionLinkStatus),
		}
	}

	return &PlanTemplate{
		AuditModel:                 general_objects.AuditModelFromTF(tf.AuditModelTF),
		AssetId:                    helper.FromTFString(tf.AssetId),
		Name:                       helper.FromTFString(tf.Name),
		Description:                helper.FromTFStringPtr(tf.Description),
		IntegrityStatus:            helper.FromTFString(tf.IntegrityStatus),
		ActivityConfigs:            activityConfigs,
		EstimatedDurationInSeconds: helper.FromTFInt64(tf.EstimatedDurationInSeconds),
	}
}
