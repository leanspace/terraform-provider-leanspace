terraform {
  required_providers {
    leanspace = {
      source  = "app.terraform.io/leanspace/leanspace"
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
    node_ids    = var.attached_node_ids
    widget_info {
      id    = var.value_widget_id
      type  = "VALUE"
      x     = 0
      y     = 0
      w     = 1
      h     = 5
      min_w = 1
      min_h = 5
    }
    widget_info {
      id    = var.line_widget_id
      type  = "LINE"
      x     = 1
      y     = 0
      w     = 2
      h     = 5
      min_w = 1
      min_h = 5
    }
    widget_info {
      id    = var.table_widget_id
      type  = "TABLE"
      x     = 0
      y     = 1
      w     = 3
      h     = 13
      min_w = 2
      min_h = 13
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
