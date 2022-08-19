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
      y_axis_label     = "This is a label"
    }
  }
}
