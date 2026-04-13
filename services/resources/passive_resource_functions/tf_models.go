package passive_resource_functions

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

type PassiveResourceFunctionTF struct {
	general_objects.AuditModelTF
	ResourceId   types.String                       `tfsdk:"resource_id"`
	Name         types.String                       `tfsdk:"name"`
	Formula      *PassiveResourceFunctionFormulaTF   `tfsdk:"formula"`
	ControlBound types.Float64                      `tfsdk:"control_bound"`
	Tags         []general_objects.KeyValueTF        `tfsdk:"tags"`
}

type PassiveResourceFunctionFormulaTF struct {
	Type     types.String  `tfsdk:"type"`
	Rate     types.Float64 `tfsdk:"rate"`
	TimeUnit types.String  `tfsdk:"time_unit"`
}

func (x *PassiveResourceFunction) ToTF() any {
	var formula *PassiveResourceFunctionFormulaTF
	if x.Formula != nil {
		formula = &PassiveResourceFunctionFormulaTF{
			Type:     helper.TFStringValue(x.Formula.Type),
			Rate:     helper.TFFloat64Value(x.Formula.Rate),
			TimeUnit: helper.TFStringValue(x.Formula.TimeUnit),
		}
	}
	return &PassiveResourceFunctionTF{
		AuditModelTF: general_objects.AuditModelToTF(&x.AuditModel),
		ResourceId:   helper.TFStringValue(x.ResourceId),
		Name:         helper.TFStringValue(x.Name),
		Formula:      formula,
		ControlBound: helper.TFFloat64PtrValue(x.ControlBound),
		Tags:         general_objects.KeyValuesToTF(x.Tags),
	}
}

func (tf *PassiveResourceFunctionTF) ToAPI() any {
	var formula *PassiveResourceFunctionFormula
	if tf.Formula != nil {
		formula = &PassiveResourceFunctionFormula{
			Type:     helper.FromTFString(tf.Formula.Type),
			Rate:     helper.FromTFFloat64(tf.Formula.Rate),
			TimeUnit: helper.FromTFString(tf.Formula.TimeUnit),
		}
	}
	return &PassiveResourceFunction{
		AuditModel:   general_objects.AuditModelFromTF(tf.AuditModelTF),
		ResourceId:   helper.FromTFString(tf.ResourceId),
		Name:         helper.FromTFString(tf.Name),
		Formula:      formula,
		ControlBound: helper.FromTFFloat64Ptr(tf.ControlBound),
		Tags:         general_objects.KeyValuesFromTF(tf.Tags),
	}
}
