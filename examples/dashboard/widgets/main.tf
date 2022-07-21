terraform {
  required_providers {
    leanspace = {
      version = "0.3.0"
      source  = "app.terraform.io/leanspace/leanspace"
    }
  }
}

data "leanspace_widgets" "all" {}

variable "text_metric_id" {
  type        = string
  description = "The ID of the text metric to create widgets for."
}

variable "numeric_metric_id" {
  type        = string
  description = "The ID of the numeric metric to create widgets for."
}

resource "leanspace_widgets" "test_table" {
  widget {
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
}

resource "leanspace_widgets" "test_line" {
  widget {
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
      y_axis_label = "This is a label"
    }
    tags {
      key   = "Mission"
      value = "Terraform"
    }
  }
}

resource "leanspace_widgets" "test_bar" {
  widget {
    name        = "Terraform Bar Widget"
    description = "A bar widget created with Terraform"
    type        = "BAR"
    granularity = "hour"
    series {
      id          = "error_code"
      datasource  = "raw_stream"
      aggregation = "avg"
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
}

resource "leanspace_widgets" "test_area" {
  widget {
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
}

resource "leanspace_widgets" "test_value" {
  widget {
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
}

output "test_table_widget" {
  value = leanspace_widgets.test_table
}

output "test_line_widget" {
  value = leanspace_widgets.test_line
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
