terraform {
  required_providers {
    leanspace = {
      source = "leanspace/leanspace"
    }
  }
}

variable "resource1_id" {
  type        = string
  description = "The ID of the resource to which the resource function is attached."
}

variable "resource2_id" {
  type        = string
  description = "The ID of the resource to which the resource function is attached."
}

data "leanspace_passive_resource_functions" "all" {
  filters {
    ids                     = []
    resource_ids            = [var.resource1_id, var.resource2_id]
    query                   = ""
    tags                    = []
    page                    = 0
    size                    = 10
    sort                    = ["name,asc"]
  }
}

resource "leanspace_passive_resource_functions" "a_linear_resource_function" {
  name                   = "Terraform Linear Passive Resource Function"
  resource_id            = var.resource1_id
  control_bound          = [25.0]
  formula {
    constant  = 5.0
    rate      = 2.5
    type      = "LINEAR"
    time_unit = "SECONDS"
  }
  tags {
    key   = "Mission"
    value = "Terraform"
  }
}

resource "leanspace_passive_resource_functions" "a_linear_resource_function_with_0_constant" {
  name                   = "Terraform Linear Passive Resource Function With Constant At Zero"
  resource_id            = var.resource2_id
  formula {
    constant  = 0.0
    rate      = 2.5
    type      = "LINEAR"
    time_unit = "SECONDS"
  }
}

resource "leanspace_passive_resource_functions" "a_linear_resource_function_with_0_rate" {
  name                   = "Terraform Linear Passive Resource Function With Rate At Zero"
  resource_id            = var.resource2_id
  formula {
    constant  = 5.0
    rate      = 0.0
    type      = "LINEAR"
    time_unit = "SECONDS"
  }
}

resource "leanspace_passive_resource_functions" "a_linear_resource_function_with_0_controlBound" {
  name                   = "Terraform Linear Passive Resource Function With Control Bound At Zero"
  resource_id            = var.resource2_id
  control_bound          = [0.0]
  formula {
    constant  = 5.0
    rate      = 1.0
    type      = "LINEAR"
    time_unit = "SECONDS"
  }
}

output "a_linear_passive_resource_function" {
  value = leanspace_passive_resource_functions.a_linear_resource_function
}
