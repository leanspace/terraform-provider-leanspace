terraform {
  required_providers {
    leanspace = {
      source = "leanspace/leanspace"
    }
  }
}

provider "leanspace" {
  tenant        = "my-org"
  client_id     = "client-id"
  client_secret = "client-secret"
}

resource "leanspace_nodes" "my_node" {
  name        = "MySatellite"
  description = "Using terraform is so easy!"
  type        = "ASSET"
  kind        = "SATELLITE"
}

resource "leanspace_properties" "mass_property" {
  name        = "Mass"
  description = "The mass of this satellite"
  node_id     = leanspace_nodes.my_node.id
  type        = "NUMERIC"
  value       = 800
}
