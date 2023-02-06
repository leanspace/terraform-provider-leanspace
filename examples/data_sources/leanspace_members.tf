data "leanspace_members" "all" {
  filters {
    team_ids = var.team_ids
    statuses = ["ACTIVE"]
    ids      = []
    query    = ""
    page     = 0
    size     = 10
    sort     = ["name,asc"]
  }
}
