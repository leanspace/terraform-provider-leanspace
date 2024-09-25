variable "satellite_id" {
  type        = string
  description = "The ID of the satellite to which the orbit will be added."
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
  gps_configuration {
    gps_metrics {
      metric_id_for_position_x = var.metric_id_for_position_x
      metric_id_for_position_y = var.metric_id_for_position_y
      metric_id_for_position_z = var.metric_id_for_position_z
      metric_id_for_velocity_x = var.metric_id_for_velocity_x
      metric_id_for_velocity_y = var.metric_id_for_velocity_y
      metric_id_for_velocity_z = var.metric_id_for_velocity_z
    }
    standard_deviations {
      latitude  = 0.2
      longitude = 0.2
      altitude  = 100
      ground_speed = 10
    }
  }
  satellite_configuration {
    drag_cross_section = 35.3
    radiation_cross_section = 55.2
  }
  tags {
    key   = "Mission"
    value = "Terraform"
  }
}