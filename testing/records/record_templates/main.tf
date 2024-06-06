terraform {
  required_providers {
    leanspace = {
      source  = "leanspace/leanspace"
    }
  }
}

provider "leanspace" {
  env           = "develop"
  tenant        = "yuri"
  client_id     = "4a4e5cnf4i11rmes6albkqa1st"
  client_secret = "iudp3kn5htosttt11h6753dog6qfvejjs4ge0kmbu93n0d5iju0"
}

variable "node_id" {
  type        = string
  description = "The ID of the node to which the resource will be added."
}

data "leanspace_record_templates" "all" {
  filters {
    related_asset_ids = [var.node_id]
    ids               = []
    names             = []
    query             = ""
    tags              = []
    page              = 0
    size              = 10
    sort              = ["name,asc"]
  }
}

resource "leanspace_record_templates" "a_record_template" {
  name     = "Terraform Record Template"
  node_ids = [var.node_id]
  properties {
    name = "TestPropertyNumeric"
    attributes {
      type          = "NUMERIC"
      required      = true
      default_value = 2
    }
  }
}

output "a_record_template" {
  value = leanspace_record_templates.a_record_template
}
