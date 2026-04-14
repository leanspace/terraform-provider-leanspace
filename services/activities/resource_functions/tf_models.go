package resource_functions

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

type ResourceFunctionTF struct {
	general_objects.AuditModelTF
	ActivityDefinitionId types.String               `tfsdk:"activity_definition_id"`
	ResourceId           types.String               `tfsdk:"resource_id"`
	Name                 types.String               `tfsdk:"name"`
	Formula              *ResourceFunctionFormulaTF `tfsdk:"formula"`
}

type ResourceFunctionFormulaTF struct {
	Type      types.String  `tfsdk:"type"`
	Amplitude types.Float64 `tfsdk:"amplitude"`
	Constant  types.Float64 `tfsdk:"constant"`
	Rate      types.Float64 `tfsdk:"rate"`
	TimeUnit  types.String  `tfsdk:"time_unit"`
}

func (x *ResourceFunction) ToTF() any {
	var formula *ResourceFunctionFormulaTF
	if x.Formula != nil {
		formula = &ResourceFunctionFormulaTF{
			Type:      helper.TFStringValue(x.Formula.Type),
			Amplitude: helper.TFFloat64PtrValue(x.Formula.Amplitude),
			Constant:  helper.TFFloat64PtrValue(x.Formula.Constant),
			Rate:      helper.TFFloat64PtrValue(x.Formula.Rate),
			TimeUnit:  helper.TFStringValue(x.Formula.TimeUnit),
		}
	}
	return &ResourceFunctionTF{
		AuditModelTF:         general_objects.AuditModelToTF(&x.AuditModel),
		ActivityDefinitionId: helper.TFStringValue(x.ActivityDefinitionId),
		ResourceId:           helper.TFStringValue(x.ResourceId),
		Name:                 helper.TFStringValue(x.Name),
		Formula:              formula,
	}
}

func (tf *ResourceFunctionTF) ToAPI() any {
	var formula *ResourceFunctionFormula
	if tf.Formula != nil {
		formula = &ResourceFunctionFormula{
			Type:      helper.FromTFString(tf.Formula.Type),
			Amplitude: helper.FromTFFloat64Ptr(tf.Formula.Amplitude),
			Constant:  helper.FromTFFloat64Ptr(tf.Formula.Constant),
			Rate:      helper.FromTFFloat64Ptr(tf.Formula.Rate),
			TimeUnit:  helper.FromTFString(tf.Formula.TimeUnit),
		}
	}
	return &ResourceFunction{
		AuditModel:           general_objects.AuditModelFromTF(tf.AuditModelTF),
		ActivityDefinitionId: helper.FromTFString(tf.ActivityDefinitionId),
		ResourceId:           helper.FromTFString(tf.ResourceId),
		Name:                 helper.FromTFString(tf.Name),
		Formula:              formula,
	}
}
