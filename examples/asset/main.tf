terraform {
  required_providers {
    leanspace = {
      version = "0.2"
      source  = "leanspace.io/io/asset"
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
    type = "GROUP"
    parent_node_id = "ec8dedb2-67bf-40c8-a00a-97e604b5c1cd"
    tags {
      key = "Key1"
      value = "Value1"
    }
    tags {
      key = "Key2"
      value = "Value2"
    }
    norad_id = "33462"
    international_designator = "33462A"
    tle = [
      "1 25544U 98067A   08264.51782528 -.00002182  00000-0 -11606-4 0  2927",
      "2 25544  51.6416 247.4627 0006703 130.5360 325.0288 15.72125391563537"
    ]
  }
}

resource "leanspace_assets" "test_nested" {
  asset {
    name = "TestTerraformNested"
    description = "TestTerraformUpdated"
    type = "GROUP"
    parent_node_id = "ec8dedb2-67bf-40c8-a00a-97e604b5c1cd"
    tags {
      key = "Key1"
      value = "Value1"
    }
    tags {
      key = "Key2"
      value = "Value2"
    }
    norad_id = "33462"
    international_designator = "33462A"
    tle = [
      "1 25544U 98067A   08264.51782528 -.00002182  00000-0 -11606-4 0  2927",
      "2 25544  51.6416 247.4627 0006703 130.5360 325.0288 15.72125391563537"
    ]
    nodes {
      name = "TestTerraformIner"
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
      norad_id = "33462"
      international_designator = "33462A"
      tle = [
        "1 25544U 98067A   08264.51782528 -.00002182  00000-0 -11606-4 0  2927",
        "2 25544  51.6416 247.4627 0006703 130.5360 325.0288 15.72125391563537"
      ]
    }
  }
}

output "test_asset" {
  value = leanspace_assets.test
}