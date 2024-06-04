data "leanspace_records" "all" {
  filters {
    ids                 = []
    record_template_ids = []
    names               = []
    query               = ""
    tags                = []
    page                = 0
    size                = 10
    sort                = ["name,asc"]
  }
}
