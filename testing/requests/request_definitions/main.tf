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

data "leanspace_request_definitions" "all" {
  filters {
    plan_template_ids                     = []
    feasibility_constraint_definition_ids = []
    ids                                   = []
    query                                 = ""
    page                                  = 0
    size                                  = 10
    sort                                  = ["name,asc"]
  }
}

resource "leanspace_request_definitions" "created" {
  name              = "requestDefinitionFromTerraform"
  description       = "requestDefinitionTerraformDescription"
  plan_template_ids = [var.plan_template_id]
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
  configuration_argument_definitions {
    name        = "TextConfigurationArgumentDefinition"
    description = "A text input"
    attributes {
      default_value = "Read boooy"
      type          = "TEXT"
      required      = false
    }
  }
  configuration_argument_definitions {
    name        = "TimeArgumentDefinition"
    description = "A time input"
    attributes {
      default_value = "10:37:19.000"
      type          = "TIME"
      required      = true
    }
  }
  configuration_argument_definitions {
    name        = "GeoPointArgumentDefinition"
    description = "A geopoint input"
    attributes {
      type = "GEOPOINT"
      fields {
        elevation {
          default_value = 141.0
        }
        latitude {
          default_value = 48.5
        }
        longitude {
          default_value = 7.7
        }
      }
      required = true
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
