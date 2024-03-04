data "leanspace_resources" "all" {
  filters {
    asset_ids    = [var.asset_id]
    ids          = []
    data_sources = []
    tags         = []
    query        = ""
    page         = 0
    size         = 10
    sort         = ["name,asc"]
  }
}
