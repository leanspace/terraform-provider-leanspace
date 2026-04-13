package access_policies

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

type AccessPolicyTF struct {
	general_objects.AuditModelTF
	Name        types.String                 `tfsdk:"name"`
	Description types.String                 `tfsdk:"description"`
	ReadOnly    types.Bool                   `tfsdk:"read_only"`
	Statements  []StatementTF                `tfsdk:"statements"`
	Tags        []general_objects.KeyValueTF `tfsdk:"tags"`
}

type StatementTF struct {
	Name    types.String   `tfsdk:"name"`
	Actions []types.String `tfsdk:"actions"`
}

func (x *AccessPolicy) ToTF() any {
	statements := make([]StatementTF, len(x.Statements))
	for i, s := range x.Statements {
		statements[i] = StatementTF{
			Name:    helper.TFStringValue(s.Name),
			Actions: helper.TFStringsValue(s.Actions),
		}
	}
	return &AccessPolicyTF{
		AuditModelTF: general_objects.AuditModelToTF(&x.AuditModel),
		Name:         helper.TFStringValue(x.Name),
		Description:  helper.TFStringValue(x.Description),
		ReadOnly:     helper.TFBoolValue(x.ReadOnly),
		Statements:   statements,
		Tags:         general_objects.KeyValuesToTF(x.Tags),
	}
}

func (tf *AccessPolicyTF) ToAPI() any {
	statements := make([]Statement, len(tf.Statements))
	for i, s := range tf.Statements {
		statements[i] = Statement{
			Name:    helper.FromTFString(s.Name),
			Actions: helper.FromTFStrings(s.Actions),
		}
	}
	return &AccessPolicy{
		AuditModel:  general_objects.AuditModelFromTF(tf.AuditModelTF),
		Name:        helper.FromTFString(tf.Name),
		Description: helper.FromTFString(tf.Description),
		ReadOnly:    helper.FromTFBool(tf.ReadOnly),
		Statements:  statements,
		Tags:        general_objects.KeyValuesFromTF(tf.Tags),
	}
}
