terraform {
  required_providers {
    leanspace = {
      source = "leanspace/leanspace"
    }
  }
}

variable "asset_id" {
  type        = string
  description = "The ID of the asset to which the command queue will be added."
}

variable "ground_station_ids" {
  type        = list(string)
  description = "The list of ground station IDs to which the command queue will be linked."
}

data "leanspace_command_queues" "all" {
  filters {
    asset_ids                       = [var.asset_id]
    ground_station_ids              = var.ground_station_ids
    command_transformer_plugin_ids  = []
    protocol_transformer_plugin_ids = []
    ids                             = []
    query                           = ""
    page                            = 0
    size                            = 10
    sort                            = ["name,asc"]
  }
}

resource "leanspace_command_queues" "test" {
  name               = "Terraform Command Queue"
  asset_id           = var.asset_id
  ground_station_ids = var.ground_station_ids
}

output "test_command_queue" {
  value = leanspace_command_queues.test
}
