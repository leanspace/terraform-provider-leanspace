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

variable "name" {
  type        = string
  description = "The name of the plan template."
}

resource "leanspace_plan_templates" "created" {
  name = var.name
  asset_id = var.asset_id
}

output "created" {
  value = leanspace_plan_templates.created
}
