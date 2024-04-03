terraform {
  required_providers {
    leanspace = {
      source = "leanspace/leanspace"
      version = "0.0.1" // TODO to remove
    }
  }
}

provider "leanspace" { // TODO to remove
  env           = "develop"
  tenant        = "yuri"
  client_id     = "4a4e5cnf4i11rmes6albkqa1st"
  client_secret = "iudp3kn5htosttt11h6753dog6qfvejjs4ge0kmbu93n0d5iju0"
}

variable "asset_id" {
  type        = string
  description = "The ID of the asset to which the resource will be added."
}

variable "metric_id" {
  type        = string
  description = "The ID of the metric associated to this resource."
}

data "leanspace_record_templates" "all" {
  filters {
    related_asset_ids = [var.asset_id]
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
  name      = "Terraform Record Template"
  // TODO add assetId and metricId
}

output "a_record_template" {
  value = leanspace_record_templates.a_record_template
}
