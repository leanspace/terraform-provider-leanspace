data "leanspace_plan_states" "all" {
  filters {
    ids          = []
    query        = ""
    page         = 0
    size         = 10
    sort         = ["name,asc"]
  }
}
