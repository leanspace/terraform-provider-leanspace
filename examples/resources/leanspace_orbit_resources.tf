variable "satellite_id" {
  type        = string
  description = "The ID of the satellite to which the orbit resource will be added."
}

resource "leanspace_orbit_resources" "test" {
  name                            = "Terraform Orbit Resource"
  satellite_id                    = var.satellite_id
  data_source                     = "TLE_MANUAL"
}