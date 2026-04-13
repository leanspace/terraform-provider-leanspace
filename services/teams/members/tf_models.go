package members

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

type MemberTF struct {
	general_objects.AuditModelTF
	Name      types.String   `tfsdk:"name"`
	Email     types.String   `tfsdk:"email"`
	State     types.String   `tfsdk:"state"`
	PolicyIds []types.String `tfsdk:"policy_ids"`
}

func (x *Member) ToTF() any {
	return &MemberTF{
		AuditModelTF: general_objects.AuditModelToTF(&x.AuditModel),
		Name:         helper.TFStringValue(x.Name),
		Email:        helper.TFStringValue(x.Email),
		State:        helper.TFStringValue(x.State),
		PolicyIds:    helper.TFStringsValue(x.PolicyIds),
	}
}

func (tf *MemberTF) ToAPI() any {
	return &Member{
		AuditModel: general_objects.AuditModelFromTF(tf.AuditModelTF),
		Name:       helper.FromTFString(tf.Name),
		Email:      helper.FromTFString(tf.Email),
		State:      helper.FromTFString(tf.State),
		PolicyIds:  helper.FromTFStrings(tf.PolicyIds),
	}
}
