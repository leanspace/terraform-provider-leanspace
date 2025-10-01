package service_accounts

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

func (serviceAccount *ServiceAccount) ToMap() map[string]any {
	serviceAccountMap := make(map[string]any)
	serviceAccountMap["id"] = serviceAccount.ID
	serviceAccountMap["name"] = serviceAccount.Name
	serviceAccountMap["policy_ids"] = serviceAccount.PolicyIds
	serviceAccountMap["credentials"] = []any{serviceAccount.Credentials.ToMap()}
	serviceAccountMap["tags"] = helper.ParseToMaps(serviceAccount.Tags)
	serviceAccountMap["created_at"] = serviceAccount.CreatedAt
	serviceAccountMap["created_by"] = serviceAccount.CreatedBy
	serviceAccountMap["last_modified_at"] = serviceAccount.LastModifiedAt
	serviceAccountMap["last_modified_by"] = serviceAccount.LastModifiedBy
	return serviceAccountMap
}

func (credential *Credentials) FromMap(credentialMap map[string]any) error {
	credential.ClientId = credentialMap["client_id"].(string)
	credential.ClientSecret = credentialMap["client_secret"].(string)
	return nil
}

func (credential *Credentials) ToMap() map[string]any {
	credentialMap := make(map[string]any)
	credentialMap["client_id"] = credential.ClientId
	credentialMap["client_secret"] = credential.ClientSecret
	return credentialMap
}

func (serviceAccount *ServiceAccount) FromMap(serviceAccountMap map[string]any) error {
	serviceAccount.ID = serviceAccountMap["id"].(string)
	serviceAccount.Name = serviceAccountMap["name"].(string)
	serviceAccount.PolicyIds = make([]string, serviceAccountMap["policy_ids"].(*schema.Set).Len())
	for i, value := range serviceAccountMap["policy_ids"].(*schema.Set).List() {
		serviceAccount.PolicyIds[i] = value.(string)
	}
	if tags, err := helper.ParseFromMaps[general_objects.KeyValue](serviceAccountMap["tags"].(*schema.Set).List()); err != nil {
		return err
	} else {
		serviceAccount.Tags = tags
	}
	if len(serviceAccountMap["credentials"].([]any)) > 0 {
		if err := serviceAccount.Credentials.FromMap(serviceAccountMap["credentials"].([]any)[0].(map[string]any)); err != nil {
			return err
		}
	}
	serviceAccount.CreatedAt = serviceAccountMap["created_at"].(string)
	serviceAccount.CreatedBy = serviceAccountMap["created_by"].(string)
	serviceAccount.LastModifiedAt = serviceAccountMap["last_modified_at"].(string)
	serviceAccount.LastModifiedBy = serviceAccountMap["last_modified_by"].(string)

	return nil
}
