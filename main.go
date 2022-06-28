package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

	"terraform-provider-asset/asset"

	"terraform-provider-asset/asset/command_definitions"
	"terraform-provider-asset/asset/nodes"
	"terraform-provider-asset/asset/properties"
)

func main() {
	nodes.NodeDataType.Subscribe()
	command_definitions.CommandDataType.Subscribe()
	properties.PropertyDataType.Subscribe()

	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return asset.Provider()
		},
	})
}
