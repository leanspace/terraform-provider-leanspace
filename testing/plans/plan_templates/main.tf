terraform {
  required_providers {
    leanspace = {
      source  = "leanspace/leanspace"
      version = "5.9.0"
    }
  }
}

provider "leanspace" {
  tenant        = "yuri"
  env           = "develop"
  client_id     = "nlbja2p65j8kj7of0tfs29rf4"
  client_secret = "d762kk9862jn0j1qr4c2u3o8bjkv70o45pld3200ek89qtul6kg"
}

//variable "asset_id" {
//  type        = string
//  description = "The ID of the asset to which the resource will be added."
//}
//
//variable "name" {
//  type        = string
//  description = "The name of the plan template."
//}

resource "leanspace_plan_templates" "created" {
  //name = var.name
  name = "marcPlanTemplateTerraform"
  //asset_id = var.asset_id
  asset_id    = "3ff7b5e4-0a5c-4d32-99a6-1f366ee3d069"
  description = "marcDescriptionTerraform"
  activity_configs {
    activity_definition_id = "3e2ec036-4bb2-4286-95f8-8db1dd7c08d8"
    position               = 0
    delay_in_seconds       = 5
    name = "marcActivityDefinitionTerraformFromPlanTemplate"
    arguments {
      name = "binaryArgument"
      attributes {
        type = "BINARY"
        value = "f6d9c3"
      }
    }
    tags {
      key = "terraform"
    }
    tags {
      key = "itworks"
      value = "no"
    }
  }
}

output "created" {
  value = leanspace_plan_templates.created
}
