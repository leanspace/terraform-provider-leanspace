resource "leanspace_properties_v2" "property" {
  name    = "Text Property"
  node_id = var.node_id
  attributes {
    type = "TEXT"
    value   = "Hello World!"
  }
}
