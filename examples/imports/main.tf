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

resource "leanspace_assets" "imported_asset" {}

output "sample_asset" {
  value = leanspace_assets.imported_asset
}

resource "leanspace_properties" "imported_property" {}

output "sample_property" {
  value = leanspace_properties.imported_property
}

resource "leanspace_command_definitions" "imported_command_definition" {}

output "sample_command_definition" {
  value = leanspace_command_definitions.imported_command_definition
}