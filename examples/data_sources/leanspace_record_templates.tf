data "leanspace_record_templates" "all" {
  filters {
    page = 0
    size = 10
    sort = ["name,asc"]
  }
}
