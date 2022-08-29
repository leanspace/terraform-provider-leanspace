data "leanspace_nodes" "all" {
  filters {
    parent_node_ids = []
    property_ids    = []
    metric_ids      = []
    types           = ["ASSET"]
    kinds           = ["SATELLITE"]
    tags            = []
    ids             = []
    query           = ""
    page            = 0
    size            = 10
    sort            = ["name,asc"]
  }
}
