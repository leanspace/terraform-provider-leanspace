data "leanspace_stream_queues" "all" {
  filters {
    asset_ids = var.asset_ids
    ids       = []
    query     = ""
    page      = 0
    size      = 10
    sort      = ["name,asc"]
  }
}
