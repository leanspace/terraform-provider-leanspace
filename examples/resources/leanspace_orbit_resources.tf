variable "satellite_id" {
  type        = string
  description = "The ID of the satellite to which the orbit resource will be added."
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