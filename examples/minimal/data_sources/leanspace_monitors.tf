data "leanspace_monitors" "all" {
  filters {
    metric_ids          = [var.metric_id]
    name                = ""
    node_ids            = []
    statuses            = ["UNKNOWN", "TRIGGERED"]
    tags                = []
    action_template_ids = []
    ids                 = []
    query               = ""
    page                = 0
    size                = 10
    sort                = ["name,asc"]
  }
}
