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
variable "activity_definition_id" {
  type        = string
  description = "The ID of the activity definition to link to the plan template if any"
}
variable "resource_function_id" {
  type        = string
  description = "The ID of the resource function linked to the activity definition if any"
}

data "leanspace_activity_definitions" "all" {
  filters {
    node_ids = [var.asset_id]
    ids      = []
    query    = ""
    page     = 0
    size     = 10
    sort     = ["name,asc"]
  }
}

resource "leanspace_plan_templates" "created" {
  name        = "Terraform PlanTemplateTerraform"
  asset_id    = var.asset_id
  description = "terraform DescriptionTerraform"
  activity_configs {
    activity_definition_id = var.activity_definition_id
    position               = 0
    delay_in_seconds       = 5
    name                   = "terraform ActivityDefinitionTerraformFromPlanTemplate"
    arguments {
      name = "ActivityArgumentNumeric"
      attributes {
        value = "10"
        type  = "NUMERIC"
      }
    }
    tags {
      key = "terraform"
    }
    tags {
      key   = "itworks"
      value = "yes"
    }

    resource_function_formulas {
      formula {
        constant = 5
        rate     = 12
        type     = "LINEAR"
      }
      resource_function_id = var.resource_function_id
    }
  }
}

output "created" {
  value = leanspace_plan_templates.created
}
