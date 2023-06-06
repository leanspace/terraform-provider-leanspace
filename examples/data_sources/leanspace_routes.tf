data "leanspace_routes" "all" {
  filters {
    ids   = []
    tags  = []
    query = ""
    page  = 0
    size  = 10
    sort  = ["name,asc"]
  }
}
