terraform {
  required_providers {
    leanspace = {
      version = "0.3"
      source  = "leanspace.io/io/leanspace"
    }
  }
}

provider "leanspace" {
  tenant        = "yuri"
  env           = "develop"
  client_id     = "nlbja2p65j8kj7of0tfs29rf4"
  client_secret = "d762kk9862jn0j1qr4c2u3o8bjkv70o45pld3200ek89qtul6kg"
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
  depends_on = [
    module.nodes
  ]
}

module "command_queues" {
  source             = "./commands/command_queues"
  asset_id           = module.nodes.satellite_node.id
  ground_station_ids = [module.nodes.groundstation_node.id]
  depends_on = [
    module.nodes
  ]
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

module "widgets" {
  source            = "./dashboard/widgets"
  text_metric_id    = module.metrics.test_text_metric.id
  numeric_metric_id = module.metrics.test_numeric_metric.id
  depends_on = [
    module.metrics
  ]
}

module "dashboards" {
  source            = "./dashboard/dashboards"
  table_widget_id   = module.widgets.test_table_widget.id
  value_widget_id   = module.widgets.test_value_widget.id
  line_widget_id    = module.widgets.test_line_widget.id
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
  source          = "./teams/members"
  usernames       = ["TerraPaul", "TerraCotta", "TerraKium"]
  access_policies = [module.access_policies.test_access_policy.id]
  depends_on = [
    module.access_policies
  ]
}

module "service_accounts" {
  source          = "./teams/service_accounts"
  usernames       = ["TerraBot 1", "TerraBot 2", "TerraBot 3"]
  access_policies = [module.access_policies.test_access_policy.id]
  depends_on = [
    module.access_policies
  ]
}

module "teams" {
  source          = "./teams/teams"
  members         = [for member in module.members.test_members : member.id]
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

module "plugins" {
  source = "./plugins/plugins"
  path   = abspath("./plugins/plugins/my_plugin.jar")
}

module "action_templates" {
  source = "./monitors/action_templates/"
}
