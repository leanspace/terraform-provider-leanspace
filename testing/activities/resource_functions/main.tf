terraform {
  required_providers {
    leanspace = {
      source = "leanspace/leanspace"
    }
  }
}

variable "resource_id" {
  type        = string
  description = "The ID of the resource to which the resource function is attached."
}

variable "activity_definition_id" {
  type        = string
  description = "The ID of the activity definition to which the resource function is attached."
}

data "leanspace_resource_functions" "all" {
  filters {
    ids                     = []
    activity_definition_ids = [var.activity_definition_id]
    resource_ids            = [var.resource_id]
    page                    = 0
    size                    = 10
    sort                    = ["name,asc"]
  }
}

resource "leanspace_resource_functions" "a_linear_resource_function" {
  name                   = "Terraform Linear Resource Function"
  resource_id            = var.resource_id
  activity_definition_id = var.activity_definition_id
  formula {
    constant  = 5.0
    rate      = 2.5
    type      = "LINEAR"
    time_unit = "SECONDS"
  }
}

resource "leanspace_resource_functions" "a_rectangular_resource_function" {
  name                   = "Terraform Rectangular Resource Function"
  resource_id            = var.resource_id
  activity_definition_id = var.activity_definition_id
  formula {
    type      = "RECTANGULAR"
    amplitude = 5.0
  }
}

output "a_linear_resource_function" {
  value = leanspace_resource_functions.a_linear_resource_function
}

output "a_rectangular_resource_function" {
  value = leanspace_resource_functions.a_rectangular_resource_function
}
