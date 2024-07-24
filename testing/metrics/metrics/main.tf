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
