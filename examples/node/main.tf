terraform {
  required_providers {
    leanspace = {
      version = "0.2"
      source  = "leanspace.io/io/asset"
    }
  }
}

resource "leanspace_nodes" "test_nested" {
  node {
    name = "TerraMission - 0.1"
    description = "This is the root node for an entire mission made in terraform!"
    type = "GROUP"
    tags {
      key = "Mission"
      value = "Terraform"
    }
    tags {
      key = "Team"
      value = "Buzz"
    }
    nodes {
      name = "TerraSatellite"
      description = "The satellite responsible for the terraform mission."
      type = "ASSET"
      kind = "SATELLITE"
      tags {
        key = "Mission"
        value = "Terraform"
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

output "test_node" {
  value = leanspace_nodes.test_nested
}