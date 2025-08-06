
data "leanspace_areas_of_interest" "points" {
  filters {
    ids   = []
    query = ""
    types = ["POINT"]
    tags  = []
    page  = 0
    size  = 10
    sort  = ["name,asc"]
  }
}
