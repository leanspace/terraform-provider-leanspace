package teams

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

type TeamTF struct {
	general_objects.AuditModelTF
	Name      types.String                 `tfsdk:"name"`
	PolicyIds []types.String               `tfsdk:"policy_ids"`
	Members   []types.String               `tfsdk:"members"`
	Tags      []general_objects.KeyValueTF `tfsdk:"tags"`
}

func (x *Team) ToTF() any {
	return &TeamTF{
		AuditModelTF: general_objects.AuditModelToTF(&x.AuditModel),
		Name:         helper.TFStringValue(x.Name),
		PolicyIds:    helper.TFStringsValue(x.PolicyIds),
		Members:      helper.TFStringsValue(x.Members),
		Tags:         general_objects.KeyValuesToTF(x.Tags),
	}
}

func (tf *TeamTF) ToAPI() any {
	return &Team{
		AuditModel: general_objects.AuditModelFromTF(tf.AuditModelTF),
		Name:       helper.FromTFString(tf.Name),
		PolicyIds:  helper.FromTFStrings(tf.PolicyIds),
		Members:    helper.FromTFStrings(tf.Members),
		Tags:       general_objects.KeyValuesFromTF(tf.Tags),
	}
}
