package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

	"leanspace-terraform-provider/provider"

	"leanspace-terraform-provider/services/activities/activity_definitions"
	"leanspace-terraform-provider/services/agents/remote_agents"
	"leanspace-terraform-provider/services/analyses/analysis_definitions"
	"leanspace-terraform-provider/services/asset/nodes"
	"leanspace-terraform-provider/services/asset/properties"
	"leanspace-terraform-provider/services/asset/units"
	"leanspace-terraform-provider/services/commands/command_definitions"
	"leanspace-terraform-provider/services/commands/command_queues"
	"leanspace-terraform-provider/services/dashboard/dashboards"
	"leanspace-terraform-provider/services/dashboard/widgets"
	"leanspace-terraform-provider/services/metrics/metrics"
	"leanspace-terraform-provider/services/monitors/action_templates"
	"leanspace-terraform-provider/services/monitors/monitors"
	"leanspace-terraform-provider/services/plugins/plugins"
	"leanspace-terraform-provider/services/streams/streams"
	"leanspace-terraform-provider/services/teams/access_policies"
	"leanspace-terraform-provider/services/teams/members"
	"leanspace-terraform-provider/services/teams/service_accounts"
	"leanspace-terraform-provider/services/teams/teams"
)

// Generate the Terraform provider documentation using `tfplugindocs`:
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

func main() {
	access_policies.AccessPolicyDataType.Subscribe()
	action_templates.ActionTemplateDataType.Subscribe()
	activity_definitions.ActivityDefinitionDataType.Subscribe()
	analysis_definitions.AnalysisDefinitionDataType.Subscribe()
	command_definitions.CommandDataType.Subscribe()
	command_queues.CommandQueueDataType.Subscribe()
	dashboards.DashboardDataType.Subscribe()
	members.MemberDataType.Subscribe()
	metrics.MetricDataType.Subscribe()
	monitors.MonitorDataType.Subscribe()
	nodes.NodeDataType.Subscribe()
	plugins.PluginDataType.Subscribe()
	properties.PropertyDataType.Subscribe()
	remote_agents.RemoteAgentDataType.Subscribe()
	service_accounts.ServiceAccountDataType.Subscribe()
	streams.StreamDataType.Subscribe()
	teams.TeamDataType.Subscribe()
	units.UnitDataType.Subscribe()
	widgets.WidgetDataType.Subscribe()

	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return provider.Provider()
		},
	})
}
