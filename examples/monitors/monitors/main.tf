terraform {
  required_providers {
    leanspace = {
      version = "0.3"
      source  = "leanspace.io/io/leanspace"
    }
  }
}

variable "metric_id" {
  type        = string
  description = "The UUID of the metric to monitor."
}

variable "action_template_ids" {
  type        = set(string)
  description = "The list of the IDs of the action templates to trigger with these monitors."
}

data "leanspace_monitors" "all" {}

resource "leanspace_monitors" "test_greater_than_monitor" {
  monitor {
    name                         = "Terraform Monitor 1"
    description                  = "A monitor created throug terraform."
    polling_frequency_in_minutes = 60
    metric_id                    = var.metric_id
    expression {
      comparison_operator  = "GREATER_THAN"
      comparison_value     = 200
      aggregation_function = "HIGHEST_VALUE"
      tolerance = 10
    }
    action_template_ids = var.action_template_ids
    tags {
      key   = "Mission"
      value = "Terraform"
    }
  }
}


resource "leanspace_monitors" "test_equals_monitor" {
  monitor {
    name                         = "Terraform Monitor 2"
    description                  = "Another monitor created throug terraform."
    polling_frequency_in_minutes = 60
    metric_id                    = var.metric_id
    expression {
      comparison_operator  = "NOT_EQUAL_TO"
      comparison_value     = 120
      aggregation_function = "COUNT_VALUE"
      tolerance = 10
    }
    action_template_ids = var.action_template_ids
    tags {
      key   = "Mission"
      value = "Terraform"
    }
  }
}

output "test_greater_than_monitor" {
  value = leanspace_monitors.test_greater_than_monitor
}

output "test_equals_monitor" {
    value = leanspace_monitors.test_equals_monitor
}
