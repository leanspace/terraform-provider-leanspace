data "leanspace_access_policies" "all" {
  filters {
    action_ids   = []
    action_names = ["updateCommandTransmission"]
    ids          = []
    query        = ""
    page         = 0
    size         = 10
    sort         = ["name,asc"]
  }
}
