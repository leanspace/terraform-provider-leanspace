package service_accounts

func (serviceAccount *ServiceAccount) ToMap() map[string]any {
	serviceAccountMap := make(map[string]any)
	serviceAccountMap["id"] = serviceAccount.ID
	serviceAccountMap["name"] = serviceAccount.Name
	serviceAccountMap["policy_ids"] = serviceAccount.PolicyIds
	serviceAccountMap["created_at"] = serviceAccount.CreatedAt
	serviceAccountMap["created_by"] = serviceAccount.CreatedBy
	serviceAccountMap["last_modified_at"] = serviceAccount.LastModifiedAt
	serviceAccountMap["last_modified_by"] = serviceAccount.LastModifiedBy

	return serviceAccountMap
}

func (serviceAccount *ServiceAccount) FromMap(serviceAccountMap map[string]any) error {
	serviceAccount.ID = serviceAccountMap["id"].(string)
	serviceAccount.Name = serviceAccountMap["name"].(string)
	serviceAccount.PolicyIds = make([]string, len(serviceAccountMap["policy_ids"].([]any)))
	for i, value := range serviceAccountMap["policy_ids"].([]any) {
		serviceAccount.PolicyIds[i] = value.(string)
	}
	serviceAccount.CreatedAt = serviceAccountMap["created_at"].(string)
	serviceAccount.CreatedBy = serviceAccountMap["created_by"].(string)
	serviceAccount.LastModifiedAt = serviceAccountMap["last_modified_at"].(string)
	serviceAccount.LastModifiedBy = serviceAccountMap["last_modified_by"].(string)

	return nil
}
