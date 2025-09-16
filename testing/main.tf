terraform {
  required_providers {
    leanspace = {
      source  = "leanspace/leanspace"
    }
  }
}

provider "leanspace" {
  tenant        = var.tenant
  env           = var.env
  client_id     = var.client_id
  client_secret = var.client_secret
}

module "nodes" {
  source = "./asset/nodes"
}

module "properties" {
  source  = "./asset/properties"
  node_id = module.nodes.satellite_node.id
  depends_on = [
    module.nodes
  ]
}

module "event_definitions" {
  source = "./events/event_definitions"
}

module "event_criticality" {
  source = "./events/event_criticality"
}
module "command_states" {
  source = "./commands/command_states"
}

module "pass_contact_states" {
  source = "./pass/contact_states"
}

module "pass_states" {
  source = "./pass/pass_states"
}

module "leaf_space_connection" {
  source   = "./leaf_space_integration/connections"
  password = var.leaf_space_password
}

module "leaf_space_contact_reservation_status_mappings" {
  source           = "./leaf_space_integration/contact_reservation_status_mappings"
  contact_state_id = module.pass_contact_states.created.id
}

module "command_definitions" {
  source  = "./commands/command_definitions"
  node_id = module.nodes.satellite_node.id
  depends_on = [
    module.nodes
  ]
}

module "metrics" {
  source  = "./metrics/metrics"
  node_id = module.nodes.satellite_node.id
  unit_id = module.units.test_units["k"].id
  depends_on = [
    module.nodes,
    module.units
  ]
}

module "command_queues" {
  source   = "./commands/command_queues"
  asset_id = module.nodes.satellite_node.id
  ground_station_ids = [module.nodes.groundstation_node.id]
  depends_on = [
    module.nodes
  ]
}

module "release_queues" {
  source   = "./commands/release_queues"
  asset_id = module.nodes.satellite_node.id
  depends_on = [
    module.nodes
  ]
}

module "command_sequence_states" {
  source = "./commands/command_sequence_states"
}

module "streams" {
  source            = "./streams/streams"
  asset_id          = module.nodes.satellite_node.id
  numeric_metric_id = module.metrics.test_numeric_metric.id
  depends_on = [
    module.nodes,
    module.metrics
  ]
}

/* module "streams_queue" {
  source            = "./streams/stream_queues"
  asset_id          = module.nodes.satellite_node.id
  numeric_metric_id = module.metrics.test_numeric_metric.id
  depends_on = [
    module.nodes,
    module.metrics
  ]
} */ // Disabled until stream_queues is fixed

module "widgets" {
  source            = "./dashboard/widgets"
  text_metric_id    = module.metrics.test_text_metric.id
  numeric_metric_id = module.metrics.test_numeric_metric.id
  enum_metric_id    = module.metrics.test_enum_metric.id
  resource_id       = module.resources.a_resource.id
  topology_id       = module.nodes.groundstation_node.id
  depends_on = [
    module.metrics
  ]
}

module "dashboards" {
  source          = "./dashboard/dashboards"
  table_widget_id = module.widgets.test_table_widget.id
  value_widget_id = module.widgets.test_value_widget.id
  line_widget_id  = module.widgets.test_line_widget.id
  enum_widget_id  = module.widgets.test_enum_widget.id
  earth_widget_id = module.widgets.test_earth_widget.id
  gauge_widget_id = module.widgets.test_gauge_widget.id
  bar_widget_id   = module.widgets.test_bar_widget.id
  area_widget_id  = module.widgets.test_area_widget.id
  attached_node_ids = [module.nodes.satellite_node.id]
  depends_on = [
    module.widgets,
    module.nodes
  ]
}

module "remote_agents" {
  source            = "./agents/remote_agents"
  ground_station_id = module.nodes.groundstation_node.id
  command_queue_id  = module.command_queues.test_command_queue.id
  stream_id         = module.streams.test_stream.id
  depends_on = [
    module.nodes,
    module.command_queues,
    module.streams
  ]
}

module "access_policies" {
  source = "./teams/access_policies"
}

module "members" {
  source = "./teams/members"
  usernames = ["TerraformPaul", "TerraformCotta", "TerraformKium"]
  access_policies = [module.access_policies.test_access_policy.id]
  depends_on = [
    module.access_policies
  ]
}

module "service_accounts" {
  source = "./teams/service_accounts"
  usernames = ["TerraformBot 1", "TerraformBot 2", "TerraformBot 3"]
  access_policies = [module.access_policies.test_access_policy.id]
  depends_on = [
    module.access_policies
  ]
}

module "teams" {
  source  = "./teams/teams"
  members = [for member in module.members.test_members : member.id]
  access_policies = [module.access_policies.test_access_policy.id]
  depends_on = [
    module.access_policies,
    module.members
  ]
}

module "activity_definitions" {
  source             = "./activities/activity_definitions"
  node_id            = module.nodes.satellite_node.id
  command_definition = module.command_definitions.test_command_definition
  depends_on = [
    module.nodes,
    module.command_definitions
  ]
}

module "activity_states" {
  source = "./activities/activity_states"
}

module "generic_plugins" {
  source = "./plugins/generic_plugins"
  path = abspath("./plugins/generic_plugins/checksum_function.jar")
}

module "plugins" {
  source = "./plugins/plugins"
  path = abspath("./plugins/plugins/my_plugin.jar")
}

module "action_templates" {
  source = "./monitors/action_templates/"
}

module "monitors" {
  source    = "./monitors/monitors/"
  metric_id = module.metrics.test_numeric_metric.id
  action_template_ids = [module.action_templates.test_action_template.id]
}

module "units" {
  source = "./asset/units"
}

module "plan_states" {
  source = "./plans/plan_states"
}

module "plan_templates" {
  source                 = "./plans/plan_templates"
  asset_id               = module.nodes.satellite_node.id
  activity_definition_id = module.activity_definitions.test_activity_definition.id
  resource_function_id   = module.resource_functions.a_resource_function.id
  depends_on = [
    module.nodes,
    module.activity_definitions,
    module.resource_functions
  ]
}

module "routes" {
  source             = "./routes/routes"
  processor_ids = [module.processors.test_create_processor.id]
  service_account_id = values(module.service_accounts.test_service_accounts)[0].id
}

module "processors" {
  source = "./routes/processors"
  path = abspath("./routes/processors/processor.jar")
}

module "orbits" {
  source                     = "./orbits/orbits"
  satellite_id               = module.nodes.satellite_node.id
  metric_id_for_latitude     = module.metrics.test_numeric_1.id
  metric_id_for_longitude    = module.metrics.test_numeric_2.id
  metric_id_for_altitude     = module.metrics.test_numeric_3.id
  metric_id_for_ground_speed = module.metrics.test_numeric_4.id
  depends_on = [
    module.nodes,
    module.metrics
  ]
}

module "resources" {
  source    = "./resources/resources"
  asset_id  = module.nodes.satellite_node.id
  metric_id = module.metrics.test_numeric_metric.id
  depends_on = [
    module.nodes,
    module.metrics
  ]
}

module "resource_functions" {
  source                 = "./activities/resource_functions"
  resource_id            = module.resources.a_resource.id
  activity_definition_id = module.activity_definitions.test_activity_definition.id
  depends_on = [
    module.activity_definitions,
    module.resources
  ]
}

module "record_templates" {
  source  = "./records/record_templates"
  node_id = module.nodes.satellite_node.id
  depends_on = [
    module.nodes,
    module.metrics
  ]
}

module "feasibility_constraint_definitions" {
  source = "./requests/feasibility_constraint_definitions"
}

module "request_definitions" {
  source                    = "./requests/request_definitions"
  plan_template_id          = module.plan_templates.created.id
  feasibility_constraint_id = module.feasibility_constraint_definitions.created.id
  depends_on = [
    module.feasibility_constraint_definitions,
    module.plan_templates
  ]
}

module "request_states" {
  source = "./requests/request_states/"
}

module "leanspace_pass_delay_configuration" {
  source = "./pass/pass_delay_configuration"
}

module "leanspace_areas_of_interest" {
  source = "./orbits/areas_of_interest"
}

module "leanspace_sensors" {
  source       = "./orbits/sensors"
  satellite_id = module.nodes.satellite_node.id
}