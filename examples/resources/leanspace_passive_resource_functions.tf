variable "resource_id" {
  type        = string
  description = "The ID of the resource to which the passive resource function is attached."
}

resource "leanspace_passive_resource_functions" "a_linear_passive_resource_function" {
  name          = "Terraform Linear Passive Resource Function"
  resource_id   = var.resource_id
  control_bound = 100.0
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