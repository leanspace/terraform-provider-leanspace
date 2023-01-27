data "leanspace_command_queues" "all" {
  filters {
    asset_ids                       = [var.asset_id]
    ground_station_ids              = var.ground_station_ids
    command_transformer_plugin_ids  = []
    protocol_transformer_plugin_ids = []
    ids                             = []
    query                           = ""
    page                            = 0
    size                            = 10
    sort                            = ["name,asc"]
  }
}
