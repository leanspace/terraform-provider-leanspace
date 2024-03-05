variable "resource_id" {
  type        = string
  description = "The ID of the resource to which the resource function is attached."
}

variable "activity_definition_id" {
  type        = string
  description = "The ID of the activity definition to which the resource function is attached."
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