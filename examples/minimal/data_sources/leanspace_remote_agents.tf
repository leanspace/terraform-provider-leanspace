data "leanspace_remote_agents" "all" {
  filters {
    gateway_ids         = [var.gateway_id]
    service_account_ids = []
    connector_types     = ["OUTBOUND"]
    ids                 = []
    query               = ""
    page                = 0
    size                = 10
    sort                = ["name,asc"]
  }
}
