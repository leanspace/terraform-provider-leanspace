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

output "a_resource" {
  value = leanspace_resources.a_resource
}
