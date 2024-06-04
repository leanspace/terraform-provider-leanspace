data "leanspace_record_templates" "all" {
  filters {
    related_asset_ids = []
    ids               = []
    names             = []
    query             = ""
    tags              = []
    page              = 0
    size              = 10
    sort              = ["name,asc"]
  }
}
