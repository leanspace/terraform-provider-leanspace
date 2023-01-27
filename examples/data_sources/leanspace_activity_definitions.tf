data "leanspace_activity_definitions" "all" {
  filters {
    node_ids = [var.node_id]
    ids      = []
    query    = ""
    page     = 0
    size     = 10
    sort     = ["name,asc"]
  }
}
