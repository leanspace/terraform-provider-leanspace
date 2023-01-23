data "leanspace_properties_v2" "all" {
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
