terraform {
  required_providers {
    leanspace = {
      version = "0.2"
      source  = "leanspace.io/io/asset"
    }
  }
}

data "leanspace_dashboards" "all" {}

variable "attached_node_ids" {
  type        = list(string)
  description = "A list of nodes to attach to the dashboard."
}

variable "table_widget_id" {
  type        = string
  description = "The ID of the table widget to display."
}

variable "line_widget_id" {
  type        = string
  description = "The ID of the line widget to display."
}

variable "value_widget_id" {
  type        = string
  description = "The ID of the value widget to display."
}

resource "leanspace_dashboards" "test" {
  dashboard {
    name        = "Terraform Dashboard"
    description = "A whole dashboard created through terraform!"
    node_ids       = var.attached_node_ids
    widget_info {
      id   = var.value_widget_id
      type = "value"
      x    = 0
      y    = 0
      w    = 1
      h    = 1
    }
    widget_info {
      id   = var.line_widget_id
      type = "line"
      x    = 1
      y    = 0
      w    = 2
      h    = 1
    }
    widget_info {
      id    = var.table_widget_id
      type  = "table"
      x     = 0
      y     = 1
      w     = 3
      h     = 5
      min_w = 2
      min_h = 3
    }
    tags {
      key   = "Mission"
      value = "Terraform"
    }
  }
}

output "test_dashboard" {
  value = leanspace_dashboards.test
}
