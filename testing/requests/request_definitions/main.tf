terraform {
  required_providers {
    leanspace = {
      source = "leanspace/leanspace"
    }
  }
}

variable "plan_template_id" {
  type        = string
  description = "The ID of the plan template that can execute the request definition."
}

variable "feasibility_constraint_id" {
  type        = string
  description = "The ID of the pre-defined feasibility constraint to link to the request definition"
}

resource "leanspace_request_definitions" "created" {
  name              = "requestDefinitionFromTerraform"
  description       = "requestDefinitionTerraformDescription"
  plan_template_ids = [ var.plan_template_id ]
  feasibility_constraint_definitions {
    id       = var.feasibility_constraint_id
    required = false
  }
  configuration_argument_definitions {
    name        = "NumericConfigurationArgumentDefinition"
    description = "A numeric input"
    attributes {
      default_value = 2
      type          = "NUMERIC"
      required      = false
    }
  }
  configuration_argument_mappings {
    plan_template_id                             = var.plan_template_id
    activity_definition_position                 = 0
    configuration_argument_definition_name       = "NumericConfigurationArgumentDefinition"
    activity_definition_argument_definition_name = "ActivityArgumentNumeric"
  }
}

output "created" {
  value = leanspace_request_definitions.created
}
