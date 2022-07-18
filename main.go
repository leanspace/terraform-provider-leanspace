package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

	"terraform-provider-asset/asset"

	"terraform-provider-asset/asset/access_policies"
	"terraform-provider-asset/asset/activity_definitions"
	"terraform-provider-asset/asset/command_definitions"
	"terraform-provider-asset/asset/command_queues"
	"terraform-provider-asset/asset/dashboards"
	"terraform-provider-asset/asset/members"
	"terraform-provider-asset/asset/metrics"
	"terraform-provider-asset/asset/nodes"
	"terraform-provider-asset/asset/plugins"
	"terraform-provider-asset/asset/properties"
	"terraform-provider-asset/asset/remote_agents"
	"terraform-provider-asset/asset/service_accounts"
	"terraform-provider-asset/asset/streams"
	"terraform-provider-asset/asset/teams"
	"terraform-provider-asset/asset/widgets"
)

func main() {
	nodes.NodeDataType.Subscribe()
	command_definitions.CommandDataType.Subscribe()
	command_queues.CommandQueueDataType.Subscribe()
	properties.PropertyDataType.Subscribe()
	metrics.MetricDataType.Subscribe()
	streams.StreamDataType.Subscribe()
	widgets.WidgetDataType.Subscribe()
	dashboards.DashboardDataType.Subscribe()
	remote_agents.RemoteAgentDataType.Subscribe()
	access_policies.AccessPolicyDataType.Subscribe()
	members.MemberDataType.Subscribe()
	service_accounts.ServiceAccountDataType.Subscribe()
	teams.TeamDataType.Subscribe()
	activity_definitions.ActivityDefinitionDataType.Subscribe()
	plugins.PluginDataType.Subscribe()

	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return asset.Provider()
		},
	})
}
