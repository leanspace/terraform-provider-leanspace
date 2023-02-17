terraform {
  required_providers {
    leanspace = {
      source = "leanspace/leanspace"
    }
  }
}

variable "asset_id" {
  type        = string
  description = "The ID of the asset to which the release queue will be added."
}

data "leanspace_release_queues" "all" {
  filters {
    asset_ids                       = [var.asset_id]
    command_transformer_plugin_ids  = []
    ids                             = []
    logical_lock                    = true
    query                           = ""
    page                            = 0
    size                            = 10
    sort                            = ["name,asc"]
  }
}

resource "leanspace_release_queues" "test" {
  name                            = "Terraform Release Queue"
  asset_id                        = var.asset_id
  command_transformation_strategy = "TEST"
}

output "test_release_queue" {
  value = leanspace_release_queues.test
}
