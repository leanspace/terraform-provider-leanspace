package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

	"terraform-provider-asset/asset"

	"terraform-provider-asset/asset/command_definitions"
	"terraform-provider-asset/asset/command_queues"
	"terraform-provider-asset/asset/metrics"
	"terraform-provider-asset/asset/nodes"
	"terraform-provider-asset/asset/properties"
)

func main() {
	nodes.NodeDataType.Subscribe()
	command_definitions.CommandDataType.Subscribe()
	command_queues.CommandQueueDataType.Subscribe()
	properties.PropertyDataType.Subscribe()
	metrics.MetricDataType.Subscribe()

	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return asset.Provider()
		},
	})
}
