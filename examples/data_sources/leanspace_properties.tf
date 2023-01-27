data "leanspace_properties" "all" {
  filters {
    node_ids   = [var.node_id]
    node_types = ["ASSET"]
    node_kinds = ["SATELLITE"]
    tags       = []
    ids        = []
    query      = ""
    page       = 0
    size       = 10
    sort       = ["name,asc"]
  }
}
