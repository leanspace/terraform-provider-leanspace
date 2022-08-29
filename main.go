package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

	"github.com/leanspace/terraform-provider-leanspace/provider"

	"github.com/leanspace/terraform-provider-leanspace/services/activities/activity_definitions"
	"github.com/leanspace/terraform-provider-leanspace/services/agents/remote_agents"
	"github.com/leanspace/terraform-provider-leanspace/services/analyses/analysis_definitions"
	"github.com/leanspace/terraform-provider-leanspace/services/asset/nodes"
	"github.com/leanspace/terraform-provider-leanspace/services/asset/properties"
	"github.com/leanspace/terraform-provider-leanspace/services/asset/units"
	"github.com/leanspace/terraform-provider-leanspace/services/commands/command_definitions"
	"github.com/leanspace/terraform-provider-leanspace/services/commands/command_queues"
	"github.com/leanspace/terraform-provider-leanspace/services/dashboard/dashboards"
	"github.com/leanspace/terraform-provider-leanspace/services/dashboard/widgets"
	"github.com/leanspace/terraform-provider-leanspace/services/metrics/metrics"
	"github.com/leanspace/terraform-provider-leanspace/services/monitors/action_templates"
	"github.com/leanspace/terraform-provider-leanspace/services/monitors/monitors"
	"github.com/leanspace/terraform-provider-leanspace/services/plugins/plugins"
	"github.com/leanspace/terraform-provider-leanspace/services/streams/streams"
	"github.com/leanspace/terraform-provider-leanspace/services/teams/access_policies"
	"github.com/leanspace/terraform-provider-leanspace/services/teams/members"
	"github.com/leanspace/terraform-provider-leanspace/services/teams/service_accounts"
	"github.com/leanspace/terraform-provider-leanspace/services/teams/teams"
)

// Generate the Terraform provider documentation using `tfplugindocs`:
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

func AddDataTypes() {
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
}

func main() {
	AddDataTypes()
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: provider.Provider,
	})
}
