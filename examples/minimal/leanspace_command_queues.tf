resource "leanspace_command_queues" "command_queue" {
  name               = "Terraform Command Queue"
  asset_id           = var.asset_id
  ground_station_ids = var.ground_station_ids
}
