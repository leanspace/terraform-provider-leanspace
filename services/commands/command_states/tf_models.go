package command_states

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

type CommandStateTF struct {
	general_objects.AuditModelTF
	Name     types.String                `tfsdk:"name"`
	ReadOnly types.Bool                  `tfsdk:"read_only"`
	Tags     []general_objects.KeyValueTF `tfsdk:"tags"`
}

func (s *CommandState) ToTF() any {
	return &CommandStateTF{
		AuditModelTF: general_objects.AuditModelToTF(&s.AuditModel),
		Name:         helper.TFStringValue(s.Name),
		ReadOnly:     helper.TFBoolValue(s.ReadOnly),
		Tags:         general_objects.KeyValuesToTF(s.Tags),
	}
}

func (tf *CommandStateTF) ToAPI() any {
	return &CommandState{
		AuditModel: general_objects.AuditModelFromTF(tf.AuditModelTF),
		Name:       helper.FromTFString(tf.Name),
		ReadOnly:   helper.FromTFBool(tf.ReadOnly),
		Tags:       general_objects.KeyValuesFromTF(tf.Tags),
	}
}
