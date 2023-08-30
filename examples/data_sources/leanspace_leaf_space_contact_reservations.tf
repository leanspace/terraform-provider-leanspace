data "leanspace_leaf_space_contact_reservations" "all" {
  filters {
    leafspace_statuses = ["Scheduled"]
    ids      = []
    query    = ""
    page     = 0
    size     = 10
    sort     = ["asc"]
  }
}
