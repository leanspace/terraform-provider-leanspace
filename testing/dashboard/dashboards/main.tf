terraform {
  required_providers {
    leanspace = {
      source = "leanspace/leanspace"
    }
  }
}

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

variable "enum_widget_id" {
  type        = string
  description = "The ID of the enum widget to display."
}

variable "earth_widget_id" {
  type        = string
  description = "The ID of the earth widget to display."
}

variable "gauge_widget_id" {
  type        = string
  description = "The ID of the gauge widget to display."
}

variable "bar_widget_id" {
  type        = string
  description = "The ID of the bar widget to display."
}

variable "area_widget_id" {
  type        = string
  description = "The ID of the area widget to display."
}

data "leanspace_dashboards" "all" {
  filters {
    node_ids   = var.attached_node_ids
    widget_ids = [var.table_widget_id, var.value_widget_id]
    tags       = []
    ids        = []
    query      = ""
    page       = 0
    size       = 10
    sort       = ["name,asc"]
  }
}

resource "leanspace_dashboards" "test" {
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
  widget_info {
    id    = var.enum_widget_id
    type  = "TABLE"
    x     = 0
    y     = 2
    w     = 1
    h     = 5
    min_w = 1
    min_h = 5
  }
  widget_info {
    id    = var.earth_widget_id
    type  = "TABLE"
    x     = 1
    y     = 2
    w     = 1
    h     = 5
    min_w = 1
    min_h = 5
  }
  widget_info {
    id    = var.gauge_widget_id
    type  = "TABLE"
    x     = 2
    y     = 2
    w     = 1
    h     = 10
    min_w = 1
    min_h = 10
  }
  widget_info {
    id    = var.bar_widget_id
    type  = "TABLE"
    x     = 0
    y     = 3
    w     = 1
    h     = 5
    min_w = 1
    min_h = 5
  }
  widget_info {
    id    = var.area_widget_id
    type  = "TABLE"
    x     = 1
    y     = 3
    w     = 1
    h     = 5
    min_w = 1
    min_h = 5
  }
  tags {
    key   = "Mission"
    value = "Terraform"
  }
}

output "test_dashboard" {
  value = leanspace_dashboards.test
}
