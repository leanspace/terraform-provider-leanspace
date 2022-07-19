package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

	"leanspace-terraform-provider/provider"

	"leanspace-terraform-provider/services/activities/activity_definitions"
	"leanspace-terraform-provider/services/agents/remote_agents"
	"leanspace-terraform-provider/services/asset/nodes"
	"leanspace-terraform-provider/services/asset/properties"
	"leanspace-terraform-provider/services/commands/command_definitions"
	"leanspace-terraform-provider/services/commands/command_queues"
	"leanspace-terraform-provider/services/dashboard/dashboards"
	"leanspace-terraform-provider/services/dashboard/widgets"
	"leanspace-terraform-provider/services/metrics/metrics"
	"leanspace-terraform-provider/services/plugins/plugins"
	"leanspace-terraform-provider/services/streams/streams"
	"leanspace-terraform-provider/services/teams/access_policies"
	"leanspace-terraform-provider/services/teams/members"
	"leanspace-terraform-provider/services/teams/service_accounts"
	"leanspace-terraform-provider/services/teams/teams"
)

func main() {
	access_policies.AccessPolicyDataType.Subscribe()
	activity_definitions.ActivityDefinitionDataType.Subscribe()
	command_definitions.CommandDataType.Subscribe()
	command_queues.CommandQueueDataType.Subscribe()
	dashboards.DashboardDataType.Subscribe()
	members.MemberDataType.Subscribe()
	metrics.MetricDataType.Subscribe()
	nodes.NodeDataType.Subscribe()
	plugins.PluginDataType.Subscribe()
	properties.PropertyDataType.Subscribe()
	remote_agents.RemoteAgentDataType.Subscribe()
	service_accounts.ServiceAccountDataType.Subscribe()
	streams.StreamDataType.Subscribe()
	teams.TeamDataType.Subscribe()
	widgets.WidgetDataType.Subscribe()

	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return provider.Provider()
		},
	})
}
