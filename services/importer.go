package services

import (
	"github.com/leanspace/terraform-provider-leanspace/services/activities/activity_definitions"
	"github.com/leanspace/terraform-provider-leanspace/services/activities/activity_states"
	"github.com/leanspace/terraform-provider-leanspace/services/activities/resource_functions"
	"github.com/leanspace/terraform-provider-leanspace/services/agents/remote_agents"
	"github.com/leanspace/terraform-provider-leanspace/services/asset/nodes"
	"github.com/leanspace/terraform-provider-leanspace/services/asset/properties"
	"github.com/leanspace/terraform-provider-leanspace/services/asset/units"
	"github.com/leanspace/terraform-provider-leanspace/services/commands/command_definitions"
	"github.com/leanspace/terraform-provider-leanspace/services/commands/command_queues"
	"github.com/leanspace/terraform-provider-leanspace/services/commands/command_sequence_states"
	"github.com/leanspace/terraform-provider-leanspace/services/commands/command_states"
	"github.com/leanspace/terraform-provider-leanspace/services/commands/release_queues"
	"github.com/leanspace/terraform-provider-leanspace/services/dashboard/dashboards"
	"github.com/leanspace/terraform-provider-leanspace/services/dashboard/widgets"
	"github.com/leanspace/terraform-provider-leanspace/services/events/events_definitions"
	"github.com/leanspace/terraform-provider-leanspace/services/leaf_space_integration/contact_reservation_status_mappings"
	"github.com/leanspace/terraform-provider-leanspace/services/leaf_space_integration/leaf_space_connections"
	"github.com/leanspace/terraform-provider-leanspace/services/leaf_space_integration/link/groundstation"
	"github.com/leanspace/terraform-provider-leanspace/services/leaf_space_integration/link/satellite"
	"github.com/leanspace/terraform-provider-leanspace/services/metrics/metrics"
	"github.com/leanspace/terraform-provider-leanspace/services/monitors/action_templates"
	"github.com/leanspace/terraform-provider-leanspace/services/monitors/monitors"
	"github.com/leanspace/terraform-provider-leanspace/services/orbits/orbits"
	"github.com/leanspace/terraform-provider-leanspace/services/pass/contact_states"
	"github.com/leanspace/terraform-provider-leanspace/services/pass/pass_states"
	"github.com/leanspace/terraform-provider-leanspace/services/plans/plan_states"
	"github.com/leanspace/terraform-provider-leanspace/services/plans/plan_templates"
	"github.com/leanspace/terraform-provider-leanspace/services/plugins/plugins"
	"github.com/leanspace/terraform-provider-leanspace/services/records/record_templates"
	"github.com/leanspace/terraform-provider-leanspace/services/resources/resources"
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
	resource_functions.ResourceFunctionDataType.Subscribe()
	command_definitions.CommandDataType.Subscribe()
	command_queues.CommandQueueDataType.Subscribe()
	release_queues.ReleaseQueueDataType.Subscribe()
	command_sequence_states.CommandSequenceStateDataType.Subscribe()
	command_states.CommandStateDataType.Subscribe()
	dashboards.DashboardDataType.Subscribe()
	members.MemberDataType.Subscribe()
	metrics.MetricDataType.Subscribe()
	monitors.MonitorDataType.Subscribe()
	nodes.NodeDataType.Subscribe()
	plugins.PluginDataType.Subscribe()
	properties.PropertyDataType.Subscribe()
	remote_agents.RemoteAgentDataType.Subscribe()
	record_templates.RecordTemplateDataType.Subscribe()
	service_accounts.ServiceAccountDataType.Subscribe()
	streams.StreamDataType.Subscribe()
	teams.TeamDataType.Subscribe()
	units.UnitDataType.Subscribe()
	widgets.WidgetDataType.Subscribe()
	plan_states.PlanStateDataType.Subscribe()
	plan_templates.PlanTemplateDataType.Subscribe()
	pass_states.PassStateDataType.Subscribe()
	contact_states.ContactStateDataType.Subscribe()
	activity_states.ActivityStateDataType.Subscribe()
	routes.RouteDataType.Subscribe()
	processors.ProcessorDataType.Subscribe()
	orbits.OrbitDataType.Subscribe()
	leaf_space_connections.LeafSpaceConnectionDataType.Subscribe()
	leaf_space_groundstation_links.LeafSpaceGroundStationLinkDataType.Subscribe()
	leaf_space_satellite_links.LeafSpaceSatellitesLinkDataType.Subscribe()
	leaf_space_contact_reservation_status_mappings.LeafSpaceContactReservationStatusMappingDataType.Subscribe()
	resources.ResourceDataType.Subscribe()
	events_definitions.EventsDefinitionDataType.Subscribe()
}
