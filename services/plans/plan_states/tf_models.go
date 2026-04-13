package plan_states

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

type PlanStateTF struct {
	general_objects.AuditModelTF
	Name     types.String `tfsdk:"name"`
	ReadOnly types.Bool   `tfsdk:"read_only"`
}

func (s *PlanState) ToTF() any {
	return &PlanStateTF{
		AuditModelTF: general_objects.AuditModelToTF(&s.AuditModel),
		Name:         helper.TFStringValue(s.Name),
		ReadOnly:     helper.TFBoolValue(s.ReadOnly),
	}
}

func (tf *PlanStateTF) ToAPI() any {
	return &PlanState{
		AuditModel: general_objects.AuditModelFromTF(tf.AuditModelTF),
		Name:       helper.FromTFString(tf.Name),
		ReadOnly:   helper.FromTFBool(tf.ReadOnly),
	}
}
