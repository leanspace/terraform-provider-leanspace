data "leanspace_access_policies" "all" {
  filters {
    actions = ["updateCommandTransmission"]
    ids          = []
    query        = ""
    page         = 0
    size         = 10
    sort         = ["name,asc"]
  }
}
