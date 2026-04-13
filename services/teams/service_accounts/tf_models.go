package service_accounts

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

type ServiceAccountTF struct {
	general_objects.AuditModelTF
	Name        types.String                `tfsdk:"name"`
	PolicyIds   []types.String              `tfsdk:"policy_ids"`
	Credentials *CredentialsTF              `tfsdk:"credentials"`
	Tags        []general_objects.KeyValueTF `tfsdk:"tags"`
}

type CredentialsTF struct {
	ClientId     types.String `tfsdk:"client_id"`
	ClientSecret types.String `tfsdk:"client_secret"`
}

func (x *ServiceAccount) ToTF() any {
	var creds *CredentialsTF
	creds = &CredentialsTF{
		ClientId:     helper.TFStringValue(x.Credentials.ClientId),
		ClientSecret: helper.TFStringValue(x.Credentials.ClientSecret),
	}
	return &ServiceAccountTF{
		AuditModelTF: general_objects.AuditModelToTF(&x.AuditModel),
		Name:         helper.TFStringValue(x.Name),
		PolicyIds:    helper.TFStringsValue(x.PolicyIds),
		Credentials:  creds,
		Tags:         general_objects.KeyValuesToTF(x.Tags),
	}
}

func (tf *ServiceAccountTF) ToAPI() any {
	var creds Credentials
	if tf.Credentials != nil {
		creds = Credentials{
			ClientId:     helper.FromTFString(tf.Credentials.ClientId),
			ClientSecret: helper.FromTFString(tf.Credentials.ClientSecret),
		}
	}
	return &ServiceAccount{
		AuditModel:  general_objects.AuditModelFromTF(tf.AuditModelTF),
		Name:        helper.FromTFString(tf.Name),
		PolicyIds:   helper.FromTFStrings(tf.PolicyIds),
		Credentials: creds,
		Tags:        general_objects.KeyValuesFromTF(tf.Tags),
	}
}
