
resource "leanspace_areas_of_interest" "pointAoi" {
  name = "Terraform Point AoI"
  shape {
    type = "POINT"
    geolocation {
      latitude  = 1
      longitude = 2
      altitude  = 3
    }
  }
  tags {
    key   = "Mission"
    value = "Terraform"
  }
}

resource "leanspace_areas_of_interest" "circleAoI" {
  name = "Terraform Circle AoI"
  shape {
    type = "CIRCLE"
    center_geolocation {
      latitude  = 1.0
      longitude = 2.0
    }
    radius_in_meters = 5.0
  }
  tags {
    key   = "Mission"
    value = "Terraform"
  }
}

resource "leanspace_areas_of_interest" "polygonAoI" {
  name = "Terraform Polygon AoI"
  shape {
    type = "POLYGON"
    vertices_geolocation {
      latitude  = 1.0
      longitude = 2.0
      altitude  = 3.0
    }
    vertices_geolocation {
      latitude  = 4.0
      longitude = 5.0
      altitude  = 6.0
    }
    vertices_geolocation {
      latitude  = 7.0
      longitude = 8.0
      altitude  = 9.0
    }
  }
  tags {
    key   = "Mission"
    value = "Terraform"
  }
}
