package service_accounts

import (
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

func (serviceAccount *ServiceAccount) ToMap() map[string]any {
	serviceAccountMap := serviceAccount.ToAuditMap()
	serviceAccountMap["name"] = helper.NilIfEmpty(serviceAccount.Name)
	serviceAccountMap["policy_ids"] = helper.NilIfEmpty(serviceAccount.PolicyIds)
	serviceAccountMap["credentials"] = []any{serviceAccount.Credentials.ToMap()}
	serviceAccountMap["tags"] = helper.ParseToMaps(serviceAccount.Tags)
	return serviceAccountMap
}

func (credential *Credentials) FromMap(credentialMap map[string]any) error {
	credential.ClientId = helper.CastString(credentialMap, "client_id")
	credential.ClientSecret = helper.CastString(credentialMap, "client_secret")
	return nil
}

func (credential *Credentials) ToMap() map[string]any {
	credentialMap := make(map[string]any)
	credentialMap["client_id"] = helper.NilIfEmpty(credential.ClientId)
	credentialMap["client_secret"] = helper.NilIfEmpty(credential.ClientSecret)
	return credentialMap
}

func (serviceAccount *ServiceAccount) FromMap(serviceAccountMap map[string]any) error {
	serviceAccount.ID = helper.CastString(serviceAccountMap, "id")
	serviceAccount.Name = helper.CastString(serviceAccountMap, "name")
	serviceAccount.PolicyIds = make([]string, len(helper.CastSlice(serviceAccountMap, "policy_ids")))
	for i, value := range helper.CastSlice(serviceAccountMap, "policy_ids") {
		serviceAccount.PolicyIds[i] = value.(string)
	}
	if tags, err := helper.ParseFromMaps[general_objects.KeyValue](helper.CastSlice(serviceAccountMap, "tags")); err != nil {
		return err
	} else {
		serviceAccount.Tags = tags
	}
	if len(helper.CastSlice(serviceAccountMap, "credentials")) > 0 {
		if err := serviceAccount.Credentials.FromMap(helper.CastSlice(serviceAccountMap, "credentials")[0].(map[string]any)); err != nil {
			return err
		}
	}
	serviceAccount.CreatedAt = helper.CastString(serviceAccountMap, "created_at")
	serviceAccount.CreatedBy = helper.CastString(serviceAccountMap, "created_by")
	serviceAccount.LastModifiedAt = helper.CastString(serviceAccountMap, "last_modified_at")
	serviceAccount.LastModifiedBy = helper.CastString(serviceAccountMap, "last_modified_by")

	return nil
}
