data "leanspace_properties" "all" {
  filters {
    node_ids   = [var.node_id]
    built_in   = []
    names      = []
    tags       = []
    query      = ""
    page       = 0
    size       = 10
    sort       = ["name,asc"]
  }
}
