variable "asset_id" {
  type        = string
  description = "The ID of the asset to which the resource will be added."
}

variable "activity_definition_id" {
  type        = string
  description = "The ID of the activity definition to link to the plan template if any"
}

variable "resource_function_id" {
  type        = string
  description = "The ID of the resource function linked to the activity definition if any"
}

resource "leanspace_plan_templates" "template" {
  name        = "TerraformPlanTemplate"
  asset_id    = var.asset_id
  description = "Terraform plan optional description"
  activity_configs {
    activity_definition_id = var.activity_definition_id
    position               = 0
    delay_in_seconds       = 5
    name                   = "TerraformActivityTemplateLinkedToPlanTemplate"
    arguments {
      name = "TerraformArgumentNameIfPresentInActivity"
      attributes {
        type  = "NUMERIC"
        value = 10
      }
    }
    tags {
      key   = "terraform"
      value = "optionalValue"
    }

    resource_function_formulas {
      formula {
        type      = "LINEAR"
        constant  = 5
        rate      = 12
        time_unit = "HOURS"
      }
      resource_function_id = var.resource_function_id
    }
  }
}