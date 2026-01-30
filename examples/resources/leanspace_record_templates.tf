variable "rt_stream_id" {
  type        = string
  description = "The ID of a Stream to which the Record Template will be linked."
}

variable "rt_node_id" {
  type        = string
  description = "The ID of a Node to which the Record Template will be linked."
}

variable "rt_metric_id" {
  type        = string
  description = "The ID of a Metric to which the Record Template will be linked."
}

variable "rt_command_definition_id" {
  type        = string
  description = "The ID of a Command Definition to which the Record Template will be linked."
}

resource "leanspace_record_templates" "record_template" {
  name                   = "My Record Template"
  description            = "Example of Record Template"
  stream_id              = var.rt_stream_id
  node_ids               = [var.rt_node_id]
  metric_ids             = [var.rt_metric_id]
  command_definition_ids = [var.rt_command_definition_id]
  properties {
    name = "Template Numeric"
    attributes {
      type          = "NUMERIC"
      required      = true
      default_value = 1
    }
  }
  tags {
    key   = "My Tag key"
    value = "My Tag value"
  }
}
