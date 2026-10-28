terraform {
  required_providers {
    leanspace = {
      source = "leanspace/leanspace"
    }
  }
}

variable "node_id" {
  type        = string
  description = "The ID of the node to which the metrics will be added."
}

variable "unit_id" {
  type        = string
  description = "The ID of the unit to create numeric metrics for."
}

data "leanspace_metrics" "all" {
  filters {
    node_ids        = []
    attribute_types = ["NUMERIC", "TEXT"]
    tags            = []
    ids             = []
    query           = ""
    page            = 0
    size            = 10
    sort            = ["name,asc"]
  }
}

resource "leanspace_metrics" "test_numeric" {
  name        = "Terraform Number Metric"
  description = "A numeric metric, entirely created under terraform."
  node_id     = var.node_id

  attributes {
    type    = "NUMERIC"
    unit_id = var.unit_id
  }
}

resource "leanspace_metrics" "test_text" {
  name        = "Terraform Text Metric"
  description = "A text metric, entirely created under terraform."
  node_id     = var.node_id
  attributes {
    type = "TEXT"
  }
}

resource "leanspace_metrics" "test_bool" {
  name        = "Terraform Boolean Metric"
  description = "A boolean metric, entirely created under terraform."
  node_id     = var.node_id
  attributes {
    type = "BOOLEAN"
  }
}


resource "leanspace_metrics" "test_timestamp" {
  name        = "Terraform Timestamp Metric"
  description = "A timestamp metric, entirely created under terraform."
  node_id     = var.node_id
  attributes {
    type = "TIMESTAMP"
  }
}

resource "leanspace_metrics" "test_date" {
  name        = "Terraform Date Metric"
  description = "A date metric, entirely created under terraform."
  node_id     = var.node_id
  attributes {
    type = "DATE"
  }
}

resource "leanspace_metrics" "test_enum" {
  name        = "Terraform Enum Metric"
  description = "An enumeration metric, entirely created under terraform."
  node_id     = var.node_id
  attributes {
    options = { 1 = "test" }
    type    = "ENUM"
  }
}

resource "leanspace_metrics" "test_binary" {
  name        = "Terraform Binary Metric"
  description = "A binary metric, entirely created under terraform."
  node_id     = var.node_id
  attributes {
    type = "BINARY"
  }
}

resource "leanspace_metrics" "test_numeric_1" {
  name        = "Terraform Number Metric 1"
  description = "A numeric metric, entirely created under terraform."
  node_id     = var.node_id

  attributes {
    type    = "NUMERIC"
    unit_id = var.unit_id
  }
}

resource "leanspace_metrics" "test_numeric_2" {
  name        = "Terraform Number Metric 2"
  description = "A numeric metric, entirely created under terraform."
  node_id     = var.node_id

  attributes {
    type    = "NUMERIC"
    unit_id = var.unit_id
  }
}

resource "leanspace_metrics" "test_numeric_3" {
  name        = "Terraform Number Metric 3"
  description = "A numeric metric, entirely created under terraform."
  node_id     = var.node_id

  attributes {
    type    = "NUMERIC"
    unit_id = var.unit_id
  }
}

resource "leanspace_metrics" "test_numeric_4" {
  name        = "Terraform Number Metric 4"
  description = "A numeric metric, entirely created under terraform."
  node_id     = var.node_id

  attributes {
    type    = "NUMERIC"
    unit_id = var.unit_id
  }
}

resource "leanspace_metrics" "test_numeric_5" {
  name        = "Terraform Number Metric 5"
  description = "A numeric metric, entirely created under terraform."
  node_id     = var.node_id

  attributes {
    type    = "NUMERIC"
    unit_id = var.unit_id
  }
}

resource "leanspace_metrics" "test_numeric_6" {
  name        = "Terraform Number Metric 6"
  description = "A numeric metric, entirely created under terraform."
  node_id     = var.node_id

  attributes {
    type    = "NUMERIC"
    unit_id = var.unit_id
  }
}


resource "leanspace_metrics" "test_polynomial_calibration" {
  name        = "Terraform Polynomial Calibrated Metric"
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
  name        = "Terraform Log Calibrated Metric"
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
  name        = "Terraform Inverse Log Calibrated Metric"
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
  name        = "Terraform Capping Calibrated Metric"
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
  name        = "Terraform Range Calibrated Metric"
  description = "A numeric metric with a range calibration, entirely created under terraform."
  node_id     = var.node_id

  attributes {
    type    = "NUMERIC"
    unit_id = var.unit_id
  }

  calibration {
    type = "RANGE"
    ranges {
      min   = [-100]
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
  name        = "Terraform Text Mapping Calibrated Metric"
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

output "test_numeric_metric" {
  value = leanspace_metrics.test_numeric
}

output "test_text_metric" {
  value = leanspace_metrics.test_text
}

output "test_bool_metric" {
  value = leanspace_metrics.test_bool
}

output "test_timestamp_metric" {
  value = leanspace_metrics.test_timestamp
}

output "test_date_metric" {
  value = leanspace_metrics.test_date
}

output "test_enum_metric" {
  value = leanspace_metrics.test_enum
}

output "test_binary_metric" {
  value = leanspace_metrics.test_binary
}

output "test_numeric_1" {
  value = leanspace_metrics.test_numeric_1
}

output "test_numeric_2" {
  value = leanspace_metrics.test_numeric_2
}

output "test_numeric_3" {
  value = leanspace_metrics.test_numeric_3
}

output "test_numeric_4" {
  value = leanspace_metrics.test_numeric_4
}

output "test_numeric_5" {
  value = leanspace_metrics.test_numeric_5
}

output "test_numeric_6" {
  value = leanspace_metrics.test_numeric_6
}

output "test_polynomial_calibration" {
  value = leanspace_metrics.test_polynomial_calibration
}

output "test_log_calibration" {
  value = leanspace_metrics.test_log_calibration
}

output "test_inverse_log_calibration" {
  value = leanspace_metrics.test_inverse_log_calibration
}

output "test_capping_calibration" {
  value = leanspace_metrics.test_capping_calibration
}

output "test_range_calibration" {
  value = leanspace_metrics.test_range_calibration
}

output "test_text_mapping_calibration" {
  value = leanspace_metrics.test_text_mapping_calibration
}
