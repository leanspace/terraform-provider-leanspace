package services

import (
	"github.com/leanspace/terraform-provider-leanspace/services/activities/activity_definitions"
	"github.com/leanspace/terraform-provider-leanspace/services/activities/activity_states"
	"github.com/leanspace/terraform-provider-leanspace/services/agents/remote_agents"
	"github.com/leanspace/terraform-provider-leanspace/services/analyses/analysis_definitions"
	"github.com/leanspace/terraform-provider-leanspace/services/asset/nodes"
	"github.com/leanspace/terraform-provider-leanspace/services/asset/properties"
	"github.com/leanspace/terraform-provider-leanspace/services/asset/units"
	"github.com/leanspace/terraform-provider-leanspace/services/commands/command_definitions"
	"github.com/leanspace/terraform-provider-leanspace/services/commands/command_queues"
	"github.com/leanspace/terraform-provider-leanspace/services/commands/command_sequence_states"
	"github.com/leanspace/terraform-provider-leanspace/services/commands/release_queues"
	"github.com/leanspace/terraform-provider-leanspace/services/dashboard/dashboards"
	"github.com/leanspace/terraform-provider-leanspace/services/dashboard/widgets"
	"github.com/leanspace/terraform-provider-leanspace/services/metrics/metrics"
	"github.com/leanspace/terraform-provider-leanspace/services/monitors/action_templates"
	"github.com/leanspace/terraform-provider-leanspace/services/monitors/monitors"
	"github.com/leanspace/terraform-provider-leanspace/services/plans/plan_states"
	"github.com/leanspace/terraform-provider-leanspace/services/plugins/plugins"
	"github.com/leanspace/terraform-provider-leanspace/services/routes/processors"
	"github.com/leanspace/terraform-provider-leanspace/services/routes/routes"
	"github.com/leanspace/terraform-provider-leanspace/services/streams/streams"
	"github.com/leanspace/terraform-provider-leanspace/services/teams/access_policies"
	"github.com/leanspace/terraform-provider-leanspace/services/teams/members"
	"github.com/leanspace/terraform-provider-leanspace/services/teams/service_accounts"
	"github.com/leanspace/terraform-provider-leanspace/services/teams/teams"
)

func AddDataTypes() {
	access_policies.AccessPolicyDataType.Subscribe()
	action_templates.ActionTemplateDataType.Subscribe()
	activity_definitions.ActivityDefinitionDataType.Subscribe()
	analysis_definitions.AnalysisDefinitionDataType.Subscribe()
	command_definitions.CommandDataType.Subscribe()
	command_queues.CommandQueueDataType.Subscribe()
	release_queues.ReleaseQueueDataType.Subscribe()
	command_sequence_states.CommandSequenceStateDataType.Subscribe()
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
	plan_states.PlanStateDataType.Subscribe()
	activity_states.ActivityStateDataType.Subscribe()
	routes.RouteDataType.Subscribe()
	processors.ProcessorDataType.Subscribe()
}