terraform {
  required_providers {
    leanspace = {
      source  = "leanspace/leanspace"
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
    activity_definition_ids = []
    resource_ids            = []
    page                    = 0
    size                    = 10
    sort                    = ["name,asc"]
  }
}

resource "leanspace_resource_functions" "a_resource_function" {
  name                   = "Terraform Resource Function"
  resource_id            = var.resource_id
  activity_definition_id = var.activity_definition_id
  time_unit              = "SECONDS"
  formula {
    constant = 5.0
    rate     = 2.5
    type     = "LINEAR"
  }
}

output "a_resource_function" {
  value = leanspace_resource_functions.a_resource_function
}
