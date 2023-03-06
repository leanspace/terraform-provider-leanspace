data "leanspace_release_queues" "all" {
  filters {
    asset_ids                      = [var.asset_id]
    command_transformer_plugin_ids = []
    ids                            = []
    logical_lock                   = true
    query                          = ""
    page                           = 0
    size                           = 10
    sort                           = ["name,asc"]
  }
}
