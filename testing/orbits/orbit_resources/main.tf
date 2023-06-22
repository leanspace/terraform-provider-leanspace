terraform {
  required_providers {
    leanspace = {
      source = "leanspace/leanspace"
      version = "3.1.0"
    }
  }
}

provider "leanspace" {
  tenant        = "yuri"
  env           = "develop"
  client_id     = "nlbja2p65j8kj7of0tfs29rf4"
  client_secret = "d762kk9862jn0j1qr4c2u3o8bjkv70o45pld3200ek89qtul6kg"
}

variable "satellite_id" {
  type        = string
  description = "The ID of the satellite to which the orbit resource will be added."
}

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

resource "leanspace_orbit_resources" "test_tle_manual" {
  name         = "Terraform Orbit Resource TLE manual"
  satellite_id = var.satellite_id
  data_source  = "TLE_MANUAL"
}

resource "leanspace_orbit_resources" "test_tle_celestrak" {
  name         = "Terraform Orbit Resource TLE celestrak"
  satellite_id = var.satellite_id
  data_source  = "TLE_CELESTRAK"
  automatic_tle_update = true
  automatic_propagation = true
}

resource "leanspace_orbit_resources" "test_gps_metric" {
  name         = "Terraform Orbit Resource GPS"
  satellite_id = var.satellite_id
  data_source  = "GPS_METRIC"
  gps_metric_ids {
    metric_id_for_position_x = "23e386fe-f24c-460a-a02b-5de14381fcea"
    metric_id_for_position_y = "d0fd9bda-5ba6-40e7-a50a-5af50a39f53d"
    metric_id_for_position_z = "80e7c6e5-7631-46e8-9179-202f0a7a3071"
    metric_id_for_velocity_x = "04da8b40-367c-49b2-ac5f-3f7b1b6ec5f6"
    metric_id_for_velocity_y = "cb456620-52a2-426a-a424-4bd18fbb02bf"
    metric_id_for_velocity_z = "738fbd53-9777-4562-aad9-888048344827"
  }
}

output "tle_manual_orbit_resource" {
  value = leanspace_orbit_resources.test_tle_manual
}

output "tle_celestrak_orbit_resource" {
  value = leanspace_orbit_resources.test_tle_manual
}

output "gps_metric_orbit_resource" {
  value = leanspace_orbit_resources.test_gps_metric
}
