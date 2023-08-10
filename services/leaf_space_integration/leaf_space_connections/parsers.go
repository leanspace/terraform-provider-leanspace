package leaf_space_connections

func (leafSpaceConnectionIntegration *LeafSpaceConnection) ToMap() map[string]any {
	leafSpaceConnectionIntegrationStateMap := make(map[string]any)
	leafSpaceConnectionIntegrationStateMap["id"] = leafSpaceConnectionIntegration.ID
	leafSpaceConnectionIntegrationStateMap["name"] = leafSpaceConnectionIntegration.Name
	leafSpaceConnectionIntegrationStateMap["domain_url"] = leafSpaceConnectionIntegration.DomainUrl
	leafSpaceConnectionIntegrationStateMap["authentication_token"] = leafSpaceConnectionIntegration.AuthenticationToken
	leafSpaceConnectionIntegrationStateMap["status"] = leafSpaceConnectionIntegration.Status
	leafSpaceConnectionIntegrationStateMap["password"] = leafSpaceConnectionIntegration.Password
	leafSpaceConnectionIntegrationStateMap["username"] = leafSpaceConnectionIntegration.Username
	leafSpaceConnectionIntegrationStateMap["created_at"] = leafSpaceConnectionIntegration.CreatedAt
	leafSpaceConnectionIntegrationStateMap["created_by"] = leafSpaceConnectionIntegration.CreatedBy
	leafSpaceConnectionIntegrationStateMap["last_modified_at"] = leafSpaceConnectionIntegration.LastModifiedAt
	leafSpaceConnectionIntegrationStateMap["last_modified_by"] = leafSpaceConnectionIntegration.LastModifiedBy

	return leafSpaceConnectionIntegrationStateMap
}

func (leafSpaceConnectionIntegration *LeafSpaceConnection) FromMap(leafSpaceIntegrationMap map[string]any) error {
	leafSpaceConnectionIntegration.ID = leafSpaceIntegrationMap["id"].(string)
	leafSpaceConnectionIntegration.Name = leafSpaceIntegrationMap["name"].(string)
	leafSpaceConnectionIntegration.DomainUrl = leafSpaceIntegrationMap["domain_url"].(string)
	leafSpaceConnectionIntegration.AuthenticationToken = leafSpaceIntegrationMap["authentication_token"].(string)
	leafSpaceConnectionIntegration.Password = leafSpaceIntegrationMap["password"].(string)
	leafSpaceConnectionIntegration.Username = leafSpaceIntegrationMap["username"].(string)
	leafSpaceConnectionIntegration.Status = leafSpaceIntegrationMap["status"].(string)
	leafSpaceConnectionIntegration.CreatedAt = leafSpaceIntegrationMap["created_at"].(string)
	leafSpaceConnectionIntegration.CreatedBy = leafSpaceIntegrationMap["created_by"].(string)
	leafSpaceConnectionIntegration.LastModifiedAt = leafSpaceIntegrationMap["last_modified_at"].(string)
	leafSpaceConnectionIntegration.LastModifiedBy = leafSpaceIntegrationMap["last_modified_by"].(string)

	return nil
}
