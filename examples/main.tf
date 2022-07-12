terraform {
  required_providers {
    leanspace = {
      version = "0.2"
      source  = "leanspace.io/io/asset"
    }
  }
}

provider "leanspace" {
  tenant        = "yuri"
  env           = "develop"
  client_id     = "nlbja2p65j8kj7of0tfs29rf4"
  client_secret = "d762kk9862jn0j1qr4c2u3o8bjkv70o45pld3200ek89qtul6kg"
}

module "assets_node" {
  source = "./nodes"
}

module "assets_property" {
  source  = "./property"
  node_id = module.assets_node.satellite_node.id
  depends_on = [
    module.assets_node
  ]
}

module "assets_command_definitions" {
  source  = "./command_definition"
  node_id = module.assets_node.satellite_node.id
  depends_on = [
    module.assets_node
  ]
}

module "assets_metrics" {
  source  = "./metrics"
  node_id = module.assets_node.satellite_node.id
  depends_on = [
    module.assets_node
  ]
}

module "assets_command_queues" {
  source             = "./command_queues"
  asset_id           = module.assets_node.satellite_node.id
  ground_station_ids = [module.assets_node.groundstation_node.id]
  depends_on = [
    module.assets_node
  ]
}

module "assets_streams" {
  source            = "./streams"
  asset_id          = module.assets_node.satellite_node.id
  numeric_metric_id = module.assets_metrics.test_numeric_metric.id
  depends_on = [
    module.assets_node,
    module.assets_metrics
  ]
}

module "assets_widgets" {
  source            = "./widgets"
  text_metric_id    = module.assets_metrics.test_text_metric.id
  numeric_metric_id = module.assets_metrics.test_numeric_metric.id
  depends_on = [
    module.assets_metrics
  ]
}

module "assets_dashboards" {
  source            = "./dashboards"
  table_widget_id   = module.assets_widgets.test_table_widget.id
  value_widget_id   = module.assets_widgets.test_value_widget.id
  line_widget_id    = module.assets_widgets.test_line_widget.id
  attached_node_ids = [module.assets_node.satellite_node.id]
  depends_on = [
    module.assets_widgets,
    module.assets_node
  ]
}

module "assets_remote_agents" {
  source            = "./remote_agents"
  ground_station_id = module.assets_node.groundstation_node.id
  command_queue_id  = module.assets_command_queues.test_command_queue.id
  stream_id         = module.assets_streams.test_stream.id
  depends_on = [
    module.assets_node,
    module.assets_command_queues,
    module.assets_streams
  ]
}

module "assets_access_policies" {
  source            = "./access_policies"
}
