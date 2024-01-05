data "leanspace_orbits" "all" {
  filters {
    satellite_ids = [var.satellite_id]
    ids           = []
    data_sources  = []
    tags          = []
    query         = ""
    page          = 0
    size          = 10
    sort          = ["name,asc"]
  }
}
