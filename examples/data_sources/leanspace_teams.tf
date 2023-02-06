data "leanspace_teams" "all" {
  filters {
    member_ids = var.member_ids
    ids        = []
    query      = ""
    page       = 0
    size       = 10
    sort       = ["name,asc"]
  }
}
