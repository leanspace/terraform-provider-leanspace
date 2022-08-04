terraform {
  required_providers {
    leanspace = {
      source  = "app.terraform.io/leanspace/leanspace"
    }
  }
}

data "leanspace_metrics" "all" {}

variable "node_id" {
  type        = string
  description = "The ID of the node to which the metrics will be added."
}

resource "leanspace_metrics" "test_numeric" {
  metric {
    name        = "Terra Number Metric"
    description = "A numeric metric, entirely created under terraform."
    node_id     = var.node_id

    attributes {
      type          = "NUMERIC"
    }
  }
}

resource "leanspace_metrics" "test_text" {
  metric {
    name        = "Terra Text Metric"
    description = "A text metric, entirely created under terraform."
    node_id     = var.node_id
    attributes {
      type          = "TEXT"
    }
  }
}

resource "leanspace_metrics" "test_bool" {
  metric {
    name        = "Terra Boolean Metric"
    description = "A boolean metric, entirely created under terraform."
    node_id     = var.node_id
    attributes {
      type          = "BOOLEAN"
    }
  }
}


resource "leanspace_metrics" "test_timestamp" {
  metric {
    name        = "Terra Timestamp Metric"
    description = "A timestamp metric, entirely created under terraform."
    node_id     = var.node_id
    attributes {
      type          = "TIMESTAMP"
    }
  }
}

resource "leanspace_metrics" "test_date" {
  metric {
    name        = "Terra Date Metric"
    description = "A date metric, entirely created under terraform."
    node_id     = var.node_id
    attributes {
      type          = "DATE"
    }
  }
}

resource "leanspace_metrics" "test_enum" {
  metric {
    name        = "Terra Enum Metric"
    description = "An enumeration metric, entirely created under terraform."
    node_id     = var.node_id
    attributes {
      options       = { 1 = "test" }
      type          = "ENUM"
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
