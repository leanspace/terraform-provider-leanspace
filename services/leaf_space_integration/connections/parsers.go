package connections

import (
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/provider"
)

func (leafSpaceConnectionIntegration *LeafSpaceConnection) ToMap() map[string]any {
	leafSpaceConnectionIntegrationStateMap := leafSpaceConnectionIntegration.ToAuditMap()
	leafSpaceConnectionIntegrationStateMap["name"] = helper.NilIfEmpty(leafSpaceConnectionIntegration.Name)
	leafSpaceConnectionIntegrationStateMap["domain_url"] = helper.NilIfEmpty(leafSpaceConnectionIntegration.DomainUrl)
	leafSpaceConnectionIntegrationStateMap["authentication_token"] = helper.NilIfEmpty(leafSpaceConnectionIntegration.AuthenticationToken)
	leafSpaceConnectionIntegrationStateMap["status"] = helper.NilIfEmpty(leafSpaceConnectionIntegration.Status)
	leafSpaceConnectionIntegrationStateMap["password"] = helper.NilIfEmpty(leafSpaceConnectionIntegration.Password)
	leafSpaceConnectionIntegrationStateMap["username"] = helper.NilIfEmpty(leafSpaceConnectionIntegration.Username)
	return leafSpaceConnectionIntegrationStateMap
}

// PostReadProcess copies write-only fields (password, username) from the prior state
// into the freshly-fetched API response, because the API never returns those values.
func (leafSpaceConnectionIntegration *LeafSpaceConnection) PostReadProcess(_ *provider.Client, newValue any) error {
	if newConnection, ok := newValue.(*LeafSpaceConnection); ok {
		newConnection.Password = leafSpaceConnectionIntegration.Password
		newConnection.Username = leafSpaceConnectionIntegration.Username
	}
	return nil
}

func (leafSpaceConnectionIntegration *LeafSpaceConnection) FromMap(leafSpaceIntegrationMap map[string]any) error {
	leafSpaceConnectionIntegration.FromAuditMap(leafSpaceIntegrationMap)
	leafSpaceConnectionIntegration.Name = helper.CastString(leafSpaceIntegrationMap, "name")
	leafSpaceConnectionIntegration.DomainUrl = helper.CastString(leafSpaceIntegrationMap, "domain_url")
	leafSpaceConnectionIntegration.AuthenticationToken = helper.CastString(leafSpaceIntegrationMap, "authentication_token")
	leafSpaceConnectionIntegration.Password = helper.CastString(leafSpaceIntegrationMap, "password")
	leafSpaceConnectionIntegration.Username = helper.CastString(leafSpaceIntegrationMap, "username")
	leafSpaceConnectionIntegration.Status = helper.CastString(leafSpaceIntegrationMap, "status")
	return nil
}
