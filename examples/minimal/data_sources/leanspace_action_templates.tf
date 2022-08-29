data "leanspace_action_templates" "all" {
  filters {
    types       = ["WEBHOOK"]
    monitor_ids = []
    ids         = []
    query       = ""
    page        = 0
    size        = 10
    sort        = ["name,asc"]
  }
}
