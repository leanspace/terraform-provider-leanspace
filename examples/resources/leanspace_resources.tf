variable "asset_id" {
  type        = string
  description = "The ID of the asset to which the resource will be added."
}

variable "metric_id" {
  type        = string
  description = "The ID of the metric associated to this resource."
}

resource "leanspace_resources" "a_resource" {
  name      = "Terraform Resource"
  asset_id  = var.asset_id
  metric_id = var.metric_id
  constraints {
    type  = "LIMIT"
    kind  = "UPPER"
    value = 50.0
  }
  constraints {
    type  = "THRESHOLD"
    kind  = "UPPER"
    value = 35.0
  }
  constraints {
    type  = "LIMIT"
    kind  = "LOWER"
    value = 0.0
  }
  constraints {
    type  = "THRESHOLD"
    kind  = "LOWER"
    value = 10.0
  }
  tags {
    key   = "Mission"
    value = "Terraform"
  }
}

resource "leanspace_resources" "a_resource_with_lower_limit_upper_limit_and_thresholds" {
  name      = "Terraform Resource 4"
  asset_id  = var.asset_id
  default_level = 10.0
  lower_limit = 5.0
  upper_limit = 15.0
  thresholds {
    name = "lower threshold not causing violation"
    value = 9.0
    kind = "LOWER"
  }
  thresholds {
    name = "lower threshold causing violation"
    value = 6.0
    violation_when_reached = true
    kind = "LOWER"
  }
  thresholds {
    name = "upper threshold not causing violation"
    value = 11.0
    kind = "UPPER"
  }
  thresholds {
    name = "upper threshold causing violation"
    value = 14.0
    violation_when_reached = true
    kind = "UPPER"
  }
  tags {
    key   = "Mission"
    value = "Terraform"
  }
}
