data "leanspace_members" "all" {
  filters {
    team_ids = var.team_ids
    states = ["ACTIVE"]
    ids      = []
    query    = ""
    page     = 0
    size     = 10
    sort     = ["name,asc"]
  }
}
