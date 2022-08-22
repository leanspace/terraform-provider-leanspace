resource "leanspace_properties" "property" {
  name    = "Text Property"
  node_id = var.node_id
  type    = "TEXT"
  value   = "Hello World!"
}
