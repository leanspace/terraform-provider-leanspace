
data "leanspace_event_criticalities" "all" {
  filters {
    ids   = []
    query = ""
    tags  = ["Mission"]
    page  = 0
    size  = 10
    sort  = ["name,asc"]
  }
}