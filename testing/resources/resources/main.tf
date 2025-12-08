terraform {
  required_providers {
    leanspace = {
      source = "leanspace/leanspace"
    }
  }
}

variable "asset_id" {
  type        = string
  description = "The ID of the asset to which the resource will be added."
}

variable "metric_id" {
  type        = string
  description = "The ID of the metric associated to this resource."
}

variable "unit_id" {
  type        = string
  description = "The ID of a unit for a resource defined with no associated metric."
}

data "leanspace_resources" "all" {
  filters {
    asset_ids = [var.asset_id]
    ids       = []
    query     = ""
    tags      = []
    page      = 0
    size      = 10
    sort      = ["name,asc"]
  }
}

resource "leanspace_resources" "a_resource" {
  name        = "Terraform Resource"
  asset_id    = var.asset_id
  metric_id   = var.metric_id
  upper_limit = [50.0]
  lower_limit = [0.0]
  thresholds {
    kind                   = "UPPER"
    value                  = 35.0
    violation_when_reached = true
  }
  thresholds {
    kind                   = "LOWER"
    value                  = 10.0
    violation_when_reached = true
  }
  tags {
    key   = "Mission"
    value = "Terraform"
  }
}

resource "leanspace_resources" "a_second_resource" {
  name        = "Terraform Resource 2"
  asset_id    = var.asset_id
  metric_id   = var.metric_id
  upper_limit = [50.0]
  lower_limit = [0.0]
  tags {
    key   = "Mission"
    value = "Terraform"
  }
}

resource "leanspace_resources" "a_third_resource" {
  name        = "Terraform Resource 3"
  asset_id    = var.asset_id
  metric_id   = var.metric_id
  upper_limit = [50.0]
  lower_limit = [0.0]
  tags {
    key   = "Mission"
    value = "Terraform"
  }
}

resource "leanspace_resources" "a_fourth_resource" {
  name      = "Terraform Resource 4"
  asset_id  = var.asset_id
  metric_id = var.metric_id
  tags {
    key   = "Mission"
    value = "Terraform"
  }
}

resource "leanspace_resources" "a_resource_with_lower_limit_upper_limit_and_thresholds" {
  name          = "Terraform Resource 5"
  asset_id      = var.asset_id
  unit_id = var.unit_id
  default_level = 10.0
  lower_limit   = [5.0]
  upper_limit   = [15.0]
  thresholds {
    name  = "lower threshold not causing violation"
    value = 9.0
    kind  = "LOWER"
  }
  thresholds {
    name                   = "lower threshold causing violation"
    value                  = 6.0
    violation_when_reached = true
    kind                   = "LOWER"
  }
  thresholds {
    name  = "upper threshold not causing violation"
    value = 11.0
    kind  = "UPPER"
  }
  thresholds {
    name                   = "upper threshold causing violation"
    value                  = 14.0
    violation_when_reached = true
    kind                   = "UPPER"
  }
  tags {
    key   = "Mission"
    value = "Terraform"
  }
}

output "a_resource" {
  value = leanspace_resources.a_resource
}

output "a_second_resource" {
  value = leanspace_resources.a_second_resource
}

output "a_third_resource" {
  value = leanspace_resources.a_third_resource
}

output "a_fourth_resource" {
  value = leanspace_resources.a_fourth_resource
}