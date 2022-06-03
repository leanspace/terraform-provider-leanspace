terraform {
  required_providers {
    leanspace = {
      version = "0.2"
      source  = "leanspace.io/edu/asset"
    }
  }
}

data "leanspace_assets" "all" {}

# Returns all assets
output "all_assets" {
  value = data.leanspace_assets.all.assets
}

output "first_asset" {
  value =  data.leanspace_assets.all.assets[0]
}

resource "leanspace_assets" "test" {
  asset {
    name = "TestTerraform"
    description = "TestTerraformUpdated"
    type = "ASSET"
    kind = "SATELLITE"
    tags {
      key = "Key1"
      value = "Value1"
    }
    tags {
      key = "Key2"
      value = "Value2"
    }
  }
}

output "test_asset" {
  value = leanspace_assets.test
}