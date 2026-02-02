terraform {
  required_providers {
    leanspace = {
      source = "leanspace/leanspace"
    }
  }
}

variable "node_id" {
  type        = string
  description = "The ID of the node to which the resource will be added."
}

variable "metric_id" {
  type        = string
  description = "The ID of the node to which the resource will be added."
}

variable "stream_id" {
  type        = string
  description = "The ID of the node to which the resource will be added."
}

variable "command_definition_id" {
  type        = string
  description = "The ID of the node to which the resource will be added."
}

data "leanspace_record_templates" "all" {
  filters {
    ids        = []
    names      = []
    node_ids   = [var.node_id]
    metric_ids = [var.metric_id]
    query      = ""
    tags       = []
    page       = 0
    size       = 10
    sort       = ["name,asc"]
  }
}

resource "leanspace_record_templates" "a_record_template" {
  name                   = "Terraform Record Template"
  description            = "Test Record Template using Terraform"
  stream_id              = var.stream_id
  node_ids               = [var.node_id]
  metric_ids             = [var.metric_id]
  command_definition_ids = [var.command_definition_id]
  properties {
    name = "Numeric"
    attributes {
      type          = "NUMERIC"
      required      = true
      default_value = 1
    }
  }
  properties {
    name = "Text"
    attributes {
      type          = "TEXT"
      required      = true
      default_value = "text"
    }
  }
  properties {
    name = "Boolean"
    attributes {
      type          = "BOOLEAN"
      required      = true
      default_value = true
    }
  }
  properties {
    name = "Enum"
    attributes {
      type          = "ENUM"
      required      = true
      default_value = 1
      options       = { 1 = "key1", 2 = "key2" }
    }
  }
  properties {
    name = "Timestamp"
    attributes {
      type          = "TIMESTAMP"
      required      = true
      default_value = "2025-01-01T01:01:01.001Z"
    }
  }
  properties {
    name = "Date"
    attributes {
      type          = "DATE"
      required      = true
      default_value = "2025-01-01"
    }
  }
  properties {
    name = "Time"
    attributes {
      type          = "TIME"
      required      = true
      default_value = "01:01:01.001"
    }
  }
  properties {
    name = "Array"
    attributes {
      type          = "ARRAY"
      required      = true
      default_value = "1,2,3"
      min_size      = 1
      max_size      = 4
      unique        = true
      constraint {
        type = "NUMERIC"
        min  = 1
        max  = 10
      }
    }
  }
  tags {
    key   = "Test key"
    value = "Test value"
  }
}

output "a_record_template" {
  value = leanspace_record_templates.a_record_template
}
