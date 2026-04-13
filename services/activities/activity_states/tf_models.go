package activity_states

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

type ActivityStateTF struct {
	general_objects.AuditModelTF
	Name     types.String `tfsdk:"name"`
	ReadOnly types.Bool   `tfsdk:"read_only"`
}

func (s *ActivityState) ToTF() any {
	return &ActivityStateTF{
		AuditModelTF: general_objects.AuditModelToTF(&s.AuditModel),
		Name:         helper.TFStringValue(s.Name),
		ReadOnly:     helper.TFBoolValue(s.ReadOnly),
	}
}

func (tf *ActivityStateTF) ToAPI() any {
	return &ActivityState{
		AuditModel: general_objects.AuditModelFromTF(tf.AuditModelTF),
		Name:       helper.FromTFString(tf.Name),
		ReadOnly:   helper.FromTFBool(tf.ReadOnly),
	}
}
