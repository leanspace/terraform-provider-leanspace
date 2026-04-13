package service_accounts

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

// credentialsAttrTypes matches credentialSchema attribute names and types.
var credentialsAttrTypes = map[string]attr.Type{
	"client_id":     types.StringType,
	"client_secret": types.StringType,
}

type ServiceAccountTF struct {
	general_objects.AuditModelTF
	Name        types.String                 `tfsdk:"name"`
	PolicyIds   []types.String               `tfsdk:"policy_ids"`
	Credentials types.List                   `tfsdk:"credentials"`
	Tags        []general_objects.KeyValueTF `tfsdk:"tags"`
}

func (x *ServiceAccount) ToTF() any {
	elemType := types.ObjectType{AttrTypes: credentialsAttrTypes}
	credObj, _ := types.ObjectValue(credentialsAttrTypes, map[string]attr.Value{
		"client_id":     helper.TFStringValue(x.Credentials.ClientId),
		"client_secret": helper.TFStringValue(x.Credentials.ClientSecret),
	})
	credentials, _ := types.ListValue(elemType, []attr.Value{credObj})

	return &ServiceAccountTF{
		AuditModelTF: general_objects.AuditModelToTF(&x.AuditModel),
		Name:         helper.TFStringValue(x.Name),
		PolicyIds:    helper.TFStringsValue(x.PolicyIds),
		Credentials:  credentials,
		Tags:         general_objects.KeyValuesToTF(x.Tags),
	}
}

func (tf *ServiceAccountTF) ToAPI() any {
	var creds Credentials
	elems := tf.Credentials.Elements()
	if len(elems) > 0 {
		if obj, ok := elems[0].(types.Object); ok {
			attrs := obj.Attributes()
			creds = Credentials{
				ClientId:     helper.FromTFString(attrs["client_id"].(types.String)),
				ClientSecret: helper.FromTFString(attrs["client_secret"].(types.String)),
			}
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
