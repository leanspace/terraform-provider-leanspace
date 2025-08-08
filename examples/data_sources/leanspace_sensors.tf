data "leanspace_sensors" "circulars" {
  filters {
    ids                  = []
    satellite_ids        = []
    query                = ""
    aperture_shape_types = ["CIRCULAR"]
    tags                 = []
    page                 = 0
    size                 = 10
    sort                 = ["name,asc"]
  }
}
