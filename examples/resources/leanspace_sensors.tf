resource "leanspace_sensors" "circularSensor" {
  name         = "Terraform Circular Sensor"
  satellite_id = var.satellite_id
  aperture_shape {
    type = "CIRCULAR"
    aperture_center {
      x = 1.0
      y = -1.0
      z = 0.0
    }
    half_aperture_angle {
      degrees = 110.0
    }
  }
  tags {
    key   = "Mission"
    value = "Terraform"
  }
}

resource "leanspace_sensors" "rectangularSensor" {
  name         = "Terraform Rectangular Sensor"
  satellite_id = var.satellite_id
  aperture_shape {
    type = "RECTANGULAR"
    first_axis_vector {
      x = -1.0
      y = -2.0
      z = -3.0
    }
    first_axis_half_aperture_angle {
      degrees = 20.0
    }
    second_axis_vector {
      x = 4.0
      y = 5.0
      z = 6.0
    }
    second_axis_half_aperture_angle {
      degrees = 45.0
    }
  }
  tags {
    key   = "Mission"
    value = "Terraform"
  }
}
