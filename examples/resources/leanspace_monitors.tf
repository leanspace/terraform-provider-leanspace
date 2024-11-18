variable "metric_id" {
  type        = string
  description = "The UUID of the metric to monitor."
}

variable "action_template_ids" {
  type        = set(string)
  description = "The list of the IDs of the action templates to trigger with these monitors."
}

resource "leanspace_monitors" "test_greater_than_monitor" {
  name        = "Terraform Monitor 1"
  description = "A monitor created throug terraform."
  metric_id   = var.metric_id
  rule {
    comparison_operator = "GREATER_THAN"
    comparison_value    = 200
  }
  dynamic "action_template_links" {
    for_each = var.action_template_ids
    content {
      id = action_template_links.value
    }
  }
  tags {
    key   = "Mission"
    value = "Terraform"
  }
}


resource "leanspace_monitors" "test_equals_monitor" {
  name        = "Terraform Monitor 2"
  description = "Another monitor created throug terraform."
  metric_id   = var.metric_id
  rule {
    comparison_operator = "NOT_EQUAL_TO"
    comparison_value    = 120
    tolerance           = 10
  }
  dynamic "action_template_links" {
    for_each = var.action_template_ids
    content {
      id = action_template_links.value
    }
  }
  tags {
    key   = "Mission"
    value = "Terraform"
  }
}
