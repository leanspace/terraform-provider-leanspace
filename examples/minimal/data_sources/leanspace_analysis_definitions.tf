data "leanspace_analysis_definitions" "all" {
  filters {
    model_ids  = [var.model_id]
    node_ids   = [var.node_id]
    frameworks = []
    ids        = []
    query      = ""
    page       = 0
    size       = 10
    sort       = ["name,asc"]
  }
}
