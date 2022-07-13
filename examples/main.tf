terraform {
  required_providers {
    leanspace = {
      version = "0.2"
      source  = "leanspace.io/io/asset"
    }
  }
}

provider "leanspace" {
  tenant = "yuri"
  env = "develop"
  client_id = "nlbja2p65j8kj7of0tfs29rf4"
  client_secret = "d762kk9862jn0j1qr4c2u3o8bjkv70o45pld3200ek89qtul6kg"
}

module "assets_node" { 
  source = "./nodes" 
}

module "assets_property" { 
  source = "./property"
  node_id = module.assets_node.satellite_node.id
  depends_on = [
    module.assets_node
  ]
}

module "assets_commands" {
  source = "./command_definition"
  node_id = module.assets_node.satellite_node.id
  depends_on = [
    module.assets_node
  ]
}

module "assets_metrics" {
  source = "./metrics"
  node_id = module.assets_node.satellite_node.id
  depends_on = [
    module.assets_node
  ]
}
