terraform {
  required_providers {
    leanspace = {
      source = "leanspace/leanspace"
    }
  }
}

variable "satellite_id" {
  type        = string
  description = "The ID of the satellite to which the orbit will be added."
}

data "leanspace_orbits" "all" {
  filters {
    satellite_ids = [var.satellite_id]
    ids           = []
    query         = ""
    tags          = []
    page          = 0
    size          = 10
    sort          = ["name,asc"]
  }
}

resource "leanspace_orbits" "an_orbit" {
  name         = "Terraform Orbit"
  satellite_id = var.satellite_id
  ideal_orbit {
    type                              = "LEO"
    inclination                       = 97.5
    right_ascension_of_ascending_node = 50.0
    argument_of_perigee               = 0.8
    altitude_in_meters                = 150.0
    eccentricity                      = 0.999
  }
  tags {
    key   = "Mission"
    value = "Terraform"
  }
}

output "an_orbit" {
  value = leanspace_orbits.an_orbit
}
