terraform {
  required_providers {
    leanspace = {
      source = "leanspace/leanspace"
    }
  }
}

resource "leanspace_nodes" "imported_node" {}

output "sample_node" {
  value = leanspace_nodes.imported_node
}

resource "leanspace_properties" "imported_property" {}

output "sample_property" {
  value = leanspace_properties.imported_property
}

resource "leanspace_command_definitions" "imported_command_definition" {}

output "sample_command_definition" {
  value = leanspace_command_definitions.imported_command_definition
}
