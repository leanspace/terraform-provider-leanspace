variable "resource_id" {
  type        = string
  description = "The ID of the resource to which the resource function is attached."
}

variable "activity_definition_id" {
  type        = string
  description = "The ID of the activity definition to which the resource function is attached."
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
