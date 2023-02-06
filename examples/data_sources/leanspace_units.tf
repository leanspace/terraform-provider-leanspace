data "leanspace_units" "all" {
  filters {
    ids   = []
    query = ""
    page  = 0
    size  = 10
    sort  = ["symbol,asc"]
  }
}
