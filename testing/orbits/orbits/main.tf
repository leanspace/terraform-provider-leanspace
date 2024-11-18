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

variable "metric_id_for_latitude" {
  type        = string
  description = "The ID of the metric which hold gps latitude data."
}

variable "metric_id_for_longitude" {
  type        = string
  description = "The ID of the metric which hold gps longitude data."
}

variable "metric_id_for_altitude" {
  type        = string
  description = "The ID of the metric which hold gps altitude data."
}

variable "metric_id_for_ground_speed" {
  type        = string
  description = "The ID of the metric which hold ground speed data."
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
  gps_configuration {
    gps_metrics {
      metric_id_for_latitude     = var.metric_id_for_latitude
      metric_id_for_longitude    = var.metric_id_for_longitude
      metric_id_for_altitude     = var.metric_id_for_altitude
      metric_id_for_ground_speed = var.metric_id_for_ground_speed
    }
    standard_deviations {
      latitude     = 0.2
      longitude    = 0.2
      altitude     = 100
      ground_speed = 10
    }
  }
  satellite_configuration {
    drag_cross_section      = 35.3
    radiation_cross_section = 55.2
  }
  tags {
    key   = "Mission"
    value = "Terraform"
  }
}

output "an_orbit" {
  value = leanspace_orbits.an_orbit
}
