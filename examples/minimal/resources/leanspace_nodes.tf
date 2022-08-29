resource "leanspace_nodes" "group_node" {
  name = "Group Node"
  type = "GROUP"
}

resource "leanspace_nodes" "satellite_node" {
  parent_node_id = leanspace_nodes.group_node.id
  name           = "My Satellite"
  description    = "A satellite made with Terraform."
  type           = "ASSET"
  kind           = "SATELLITE"

  norad_id                 = "33462"
  international_designator = "33462A"
  tle = [
    "1 25544U 98067A   08264.51782528 -.00002182  00000-0 -11606-4 0  2927",
    "2 25544  51.6416 247.4627 0006703 130.5360 325.0288 15.72125391563537"
  ]
}

resource "leanspace_nodes" "groundstation_node" {
  parent_node_id = leanspace_nodes.group_node.id
  name           = "My Ground Station"
  type           = "ASSET"
  kind           = "GROUND_STATION"
  latitude       = 13.0344
  longitude      = 77.5116
  elevation      = 823
}
