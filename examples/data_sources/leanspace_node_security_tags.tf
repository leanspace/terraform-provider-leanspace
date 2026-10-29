resource "leanspace_node_security_tags" "mission" {
  node_id       = leanspace_nodes.test_nodes_root.id
  security_tags = ["terraform_mission"]
}