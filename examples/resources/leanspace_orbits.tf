variable "satellite_id" {
  type        = string
  description = "The ID of the satellite to which the orbit will be added."
}

resource "leanspace_orbits" "an_orbit" {
  name         = "Terraform Orbit"
  satellite_id = var.satellite_id
  ideal_orbit {
    type        = "LEO"
    inclination = 97.5
    right_ascension_of_ascending_node = 50.0
    argument_of_perigee = 0.8
    altitude_in_meters = 150.0
    eccentricity = 0.7
  }
  tags {
    key   = "Mission"
    value = "Terraform"
  }
}