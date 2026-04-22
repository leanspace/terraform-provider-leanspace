package connections

import "github.com/leanspace/terraform-provider-leanspace/provider"

func (leafSpaceConnectionIntegration *LeafSpaceConnection) PostReadProcess(_ *provider.Client, newValue any) error {
	if newConnection, ok := newValue.(*LeafSpaceConnection); ok {
		newConnection.Password = leafSpaceConnectionIntegration.Password
		newConnection.Username = leafSpaceConnectionIntegration.Username
	}
	return nil
}
