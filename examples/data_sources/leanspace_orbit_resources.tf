data "leanspace_orbit_resources" "all" {
  filters {
    satellite_ids                  = [var.satellite_id]
    ids                            = []
    data_sources                   = []
    query                          = ""
    page                           = 0
    size                           = 10
    sort                           = ["name,asc"]
  }
}
