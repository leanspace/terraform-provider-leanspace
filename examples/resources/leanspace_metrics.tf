variable "node_id" {
  type        = string
  description = "The ID of the node to which the metrics will be added."
}

variable "unit_id" {
  type        = string
  description = "The ID of the unit to create numeric metrics for."
}

resource "leanspace_metrics" "test_numeric" {
  name        = "Terra Number Metric"
  description = "A numeric metric, entirely created under terraform."
  node_id     = var.node_id

  attributes {
    type    = "NUMERIC"
    unit_id = var.unit_id
  }
}

resource "leanspace_metrics" "test_text" {
  name        = "Terra Text Metric"
  description = "A text metric, entirely created under terraform."
  node_id     = var.node_id
  attributes {
    type = "TEXT"
  }
}

resource "leanspace_metrics" "test_bool" {
  name        = "Terra Boolean Metric"
  description = "A boolean metric, entirely created under terraform."
  node_id     = var.node_id
  attributes {
    type = "BOOLEAN"
  }
}


resource "leanspace_metrics" "test_timestamp" {
  name        = "Terra Timestamp Metric"
  description = "A timestamp metric, entirely created under terraform."
  node_id     = var.node_id
  attributes {
    type = "TIMESTAMP"
  }
}

resource "leanspace_metrics" "test_date" {
  name        = "Terra Date Metric"
  description = "A date metric, entirely created under terraform."
  node_id     = var.node_id
  attributes {
    type = "DATE"
  }
}

resource "leanspace_metrics" "test_enum" {
  name        = "Terra Enum Metric"
  description = "An enumeration metric, entirely created under terraform."
  node_id     = var.node_id
  attributes {
    options = { 1 = "test" }
    type    = "ENUM"
  }
}

resource "leanspace_metrics" "test_binary" {
  name        = "Terra Binary Metric"
  description = "A binary metric, entirely created under terraform."
  node_id     = var.node_id
  attributes {
    type = "BINARY"
  }
}

resource "leanspace_metrics" "test_polynomial_calibration" {
  name        = "Terra Polynomial Calibrated Metric"
  description = "A numeric metric with a polynomial calibration, entirely created under terraform."
  node_id     = var.node_id

  attributes {
    type    = "NUMERIC"
    unit_id = var.unit_id
  }

  calibration {
    type         = "POLYNOMIAL"
    description  = "Converts raw counts to volts."
    coefficients = [0.5, 0.01]
    capping {
      min = [0]
      max = [5]
    }
  }
}

resource "leanspace_metrics" "test_log_calibration" {
  name        = "Terra Log Calibrated Metric"
  description = "A numeric metric with a logarithmic calibration, entirely created under terraform."
  node_id     = var.node_id

  attributes {
    type    = "NUMERIC"
    unit_id = var.unit_id
  }

  calibration {
    type         = "LOG"
    base         = 10
    coefficients = [1, 2]
  }
}

resource "leanspace_metrics" "test_inverse_log_calibration" {
  name        = "Terra Inverse Log Calibrated Metric"
  description = "A numeric metric with an inverse logarithmic calibration, entirely created under terraform."
  node_id     = var.node_id

  attributes {
    type    = "NUMERIC"
    unit_id = var.unit_id
  }

  calibration {
    type         = "INVERSE_LOG"
    base         = 2
    coefficients = [0, 1]
    capping {
      max = [1000]
    }
  }
}

resource "leanspace_metrics" "test_capping_calibration" {
  name        = "Terra Capping Calibrated Metric"
  description = "A numeric metric with a capping calibration, entirely created under terraform."
  node_id     = var.node_id

  attributes {
    type    = "NUMERIC"
    unit_id = var.unit_id
  }

  calibration {
    type = "CAPPING"
    min  = [-10]
    max  = [10]
  }
}

resource "leanspace_metrics" "test_range_calibration" {
  name        = "Terra Range Calibrated Metric"
  description = "A numeric metric with a range calibration, entirely created under terraform."
  node_id     = var.node_id

  attributes {
    type    = "NUMERIC"
    unit_id = var.unit_id
  }

  calibration {
    type = "RANGE"
    ranges {
      max   = [0]
      value = -1
    }
    ranges {
      min   = [0]
      max   = [100]
      value = 0
    }
    ranges {
      min   = [100]
      value = 1
    }
  }
}

resource "leanspace_metrics" "test_text_mapping_calibration" {
  name        = "Terra Text Mapping Calibrated Metric"
  description = "A text metric with a text mapping calibration, entirely created under terraform."
  node_id     = var.node_id

  attributes {
    type = "TEXT"
  }

  calibration {
    type = "TEXT_MAPPING"
    mappings = {
      "0" = "OFF"
      "1" = "ON"
    }
  }
}
