data "leanspace_metrics" "all" {
  filters {
    node_ids        = var.node_ids
    attribute_types = ["NUMERIC", "TEXT"]
    tags            = []
    ids             = []
    query           = ""
    page            = 0
    size            = 10
    sort            = ["name,asc"]
  }
}
