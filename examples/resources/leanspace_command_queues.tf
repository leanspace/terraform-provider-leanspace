variable "asset_id" {
  type        = string
  description = "The ID of the asset to which the command queue will be added."
}

variable "ground_station_ids" {
  type        = list(string)
  description = "The list of ground station IDs to which the command queue will be linked."
}

resource "leanspace_command_queues" "test" {
  name               = "Terraform Command Queue"
  asset_id           = var.asset_id
  ground_station_ids = var.ground_station_ids
}