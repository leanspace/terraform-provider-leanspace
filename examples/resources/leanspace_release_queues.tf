variable "asset_id" {
  type        = string
  description = "The ID of the asset to which the release queue will be added."
}

resource "leanspace_release_queues" "test" {
  name                            = "Terraform Release Queue"
  asset_id                        = var.asset_id
  command_transformation_strategy = "TEST"
  global_transmission_metadata {
    key = "mykey"
    value = "myvalue"
  }
}