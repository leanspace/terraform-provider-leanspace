resource "leanspace_nodes" "test_nodes_root" {
  name        = "TerraMission - 0.1"
  description = "This is the root node for an entire mission made in terraform!"
  type        = "GROUP"
  tags {
    key   = "key1"
    value = "Value1"
  }
  tags {
    key   = "key2"
    value = "Value2"
  }
}

resource "leanspace_nodes" "test_nodes_satellite" {
  parent_node_id = leanspace_nodes.test_nodes_root.id
  name           = "TerraSatellite"
  description    = "The satellite responsible for the terraform mission."
  type           = "ASSET"
  kind           = "SATELLITE"
  tags {
    key   = "key1"
    value = "Value1"
  }
  norad_id                 = "33462"
  international_designator = "33462A"
  tle = [
    "1 25544U 98067A   08264.51782528 -.00002182  00000-0 -11606-4 0  2927",
    "2 25544  51.6416 247.4627 0006703 130.5360 325.0288 15.72125391563537"
  ]
}

resource "leanspace_nodes" "test_nodes_groundstation" {
  parent_node_id = leanspace_nodes.test_nodes_root.id
  name           = "TerraGroundStation"
  description    = "The satellite responsible for the terraform mission."
  type           = "ASSET"
  kind           = "GROUND_STATION"
  latitude       = 13.0344
  longitude      = 77.5116
  elevation      = 823
  tags {
    key   = "key1"
    value = "Value1"
  }
}