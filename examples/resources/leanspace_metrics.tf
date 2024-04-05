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
    type = "NUMERIC"
    unit_id     = var.unit_id
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
