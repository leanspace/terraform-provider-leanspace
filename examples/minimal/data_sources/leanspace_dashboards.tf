data "leanspace_dashboards" "all" {
  filters {
    node_ids   = var.node_ids
    widget_ids = var.widget_ids
    tags       = []
    ids        = []
    query      = ""
    page       = 0
    size       = 10
    sort       = ["name,asc"]
  }
}
