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
	Status    types.String   `tfsdk:"status"`
	PolicyIds []types.String `tfsdk:"policy_ids"`
}

func (x *Member) ToTF() any {
	return &MemberTF{
		AuditModelTF: general_objects.AuditModelToTF(&x.AuditModel),
		Name:         helper.TFStringValue(x.Name),
		Email:        helper.TFStringValue(x.Email),
		Status:       helper.TFStringValue(x.Status),
		PolicyIds:    helper.TFStringsValue(x.PolicyIds),
	}
}

func (tf *MemberTF) ToAPI() any {
	return &Member{
		AuditModel: general_objects.AuditModelFromTF(tf.AuditModelTF),
		Name:       helper.FromTFString(tf.Name),
		Email:      helper.FromTFString(tf.Email),
		Status:     helper.FromTFString(tf.Status),
		PolicyIds:  helper.FromTFStrings(tf.PolicyIds),
	}
}
