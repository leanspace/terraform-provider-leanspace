data "leanspace_widgets" "all" {
  filters {
    types          = ["LINE"]
    tags           = []
    dashboard_ids  = []
    datasource_ids = [var.metric_id]
    datasources    = ["metric"]
    ids            = []
    query          = ""
    page           = 0
    size           = 10
    sort           = ["name,asc"]
  }
}
