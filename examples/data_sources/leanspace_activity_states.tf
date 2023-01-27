data "leanspace_activity_states" "all" {
  filters {
    ids          = []
    query        = ""
    page         = 0
    size         = 10
    sort         = ["name,asc"]
  }
}
