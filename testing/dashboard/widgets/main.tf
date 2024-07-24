terraform {
  required_providers {
    leanspace = {
      source = "leanspace/leanspace"
    }
  }
}

variable "text_metric_id" {
  type        = string
  description = "The ID of the text metric to create widgets for."
}

variable "numeric_metric_id" {
  type        = string
  description = "The ID of the numeric metric to create widgets for."
}

variable "enum_metric_id" {
  type        = string
  description = "The ID of the enum metric to create widgets for."
}

variable "resource_id" {
  type        = string
  description = "The ID of the text resources to create widgets for."
}

variable "topology_id" {
  type        = string
  description = "The ID of the asset to create widgets for."
}

data "leanspace_widgets" "all" {
  filters {
    types          = ["LINE"]
    tags           = []
    dashboard_ids  = []
    datasource_ids = [var.text_metric_id]
    datasources    = ["metric"]
    ids            = []
    query          = ""
    page           = 0
    size           = 10
    sort           = ["name,asc"]
  }
}

resource "leanspace_widgets" "test_table" {
  name        = "Terraform Table Widget"
  description = "A table widget created with Terraform"
  type        = "TABLE"
  granularity = "raw"
  series {
    id          = var.text_metric_id
    datasource  = "metric"
    aggregation = "none"
  }
  tags {
    key   = "Mission"
    value = "Terraform"
  }
}

resource "leanspace_widgets" "test_line" {
  name        = "Terraform Line Widget"
  description = "A line widget created with Terraform"
  type        = "LINE"
  granularity = "second"
  series {
    id          = var.numeric_metric_id
    datasource  = "metric"
    aggregation = "avg"
    filters {
      filter_by = var.numeric_metric_id
      operator  = "gt"
      value     = 3
    }
  }
  metadata {
    y_axis_range_max = [100]
    y_axis_label     = "This is a label"
  }
  tags {
    key   = "Mission"
    value = "Terraform"
  }
}

resource "leanspace_widgets" "test_enum" {
  name        = "Terraform Enum Widget"
  description = "An enum widget created with Terraform"
  type        = "ENUM"
  granularity = "second"
  series {
    id          = var.enum_metric_id
    datasource  = "metric"
    aggregation = "avg"
    filters {
      filter_by = var.enum_metric_id
      operator  = "gt"
      value     = 3
    }
  }
  metadata {
    y_axis_range_max = [100]
    y_axis_label     = "This is a label"
  }
  tags {
    key   = "Mission"
    value = "Terraform"
  }
}

resource "leanspace_widgets" "test_earth" {
  name        = "Terraform Earth Widget"
  description = "An earth widget created with Terraform"
  type        = "EARTH"
  granularity = "second"
  series {
    id          = var.topology_id
    datasource  = "topology"
    aggregation = "count"
  }
  metadata {
    y_axis_range_max = [100]
    y_axis_label     = "This is a label"
  }
  tags {
    key   = "Mission"
    value = "Terraform"
  }
}

resource "leanspace_widgets" "test_gauge" {
  name        = "Terraform Gauge Widget"
  description = "A gauge widget created with Terraform"
  type        = "GAUGE"
  granularity = "second"
  series {
    id          = var.numeric_metric_id
    datasource  = "metric"
    aggregation = "avg"
  }
  metadata {
    y_axis_range_max = [100]
    y_axis_label     = "This is a label"
    thresholds {
      to    = 49
      color = "#52C31A"
    }
    thresholds {
      from  = 49
      to    = 500
      color = "#FAAD14"
    }
    thresholds {
      from  = 500
      color = "#FF4D4F"
    }
  }
  tags {
    key   = "Mission"
    value = "Terraform"
  }
}

resource "leanspace_widgets" "test_bar" {
  name        = "Terraform Bar Widget"
  description = "A bar widget created with Terraform"
  type        = "BAR"
  granularity = "hour"
  series {
    id          = "error_code"
    datasource  = "raw_stream"
    aggregation = "count"
    filters {
      filter_by = "error_code"
      operator  = "notEquals"
      value     = 500
    }
  }
  metadata {
    y_axis_range_min = [200]
    y_axis_range_max = [600]
  }
  tags {
    key   = "Mission"
    value = "Terraform"
  }
}

resource "leanspace_widgets" "test_area" {
  name        = "Terraform Area Widget"
  description = "An area widget created with Terraform"
  type        = "AREA"
  granularity = "day"
  series {
    id          = var.numeric_metric_id
    datasource  = "metric"
    aggregation = "max"
    filters {
      filter_by = var.numeric_metric_id
      operator  = "lt"
      value     = 1000
    }
    filters {
      filter_by = var.numeric_metric_id
      operator  = "gt"
      value     = 0
    }
  }
  tags {
    key   = "Mission"
    value = "Terraform"
  }
}

resource "leanspace_widgets" "test_value" {
  name        = "Terraform Value Widget"
  description = "A value widget created with Terraform"
  type        = "VALUE"
  granularity = "minute"
  series {
    id          = var.text_metric_id
    datasource  = "metric"
    aggregation = "max"
  }
  tags {
    key   = "Mission"
    value = "Terraform"
  }
}

resource "leanspace_widgets" "test_resources" {
  name        = "Terraform Resources Widget"
  description = "A resources widget created with Terraform"
  type        = "RESOURCES"
  granularity = "raw"
  series {
    id          = var.resource_id
    datasource  = "resources"
    aggregation = "none"
  }
  tags {
    key   = "Mission"
    value = "Terraform"
  }
}

output "test_table_widget" {
  value = leanspace_widgets.test_table
}

output "test_line_widget" {
  value = leanspace_widgets.test_line
}

output "test_enum_widget" {
  value = leanspace_widgets.test_enum
}

output "test_earth_widget" {
  value = leanspace_widgets.test_earth
}

output "test_gauge_widget" {
  value = leanspace_widgets.test_gauge
}

output "test_bar_widget" {
  value = leanspace_widgets.test_bar
}

output "test_area_widget" {
  value = leanspace_widgets.test_area
}

output "test_value_widget" {
  value = leanspace_widgets.test_value
}

output "test_resources_widget" {
  value = leanspace_widgets.test_resources
}

