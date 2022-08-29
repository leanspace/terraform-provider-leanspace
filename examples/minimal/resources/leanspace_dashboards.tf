resource "leanspace_dashboards" "dashboard" {
  name        = "Terraform Dashboard"
  description = "A dashboard created through terraform!"
  node_ids    = var.attached_node_ids

  widget_info {
    id    = var.line_widget_id
    type  = "LINE"
    x     = 1
    y     = 0
    w     = 2
    h     = 5
    min_w = 1
    min_h = 5
  }
}
