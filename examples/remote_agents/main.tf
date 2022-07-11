terraform {
  required_providers {
    leanspace = {
      version = "0.2"
      source  = "leanspace.io/io/asset"
    }
  }
}

data "leanspace_remote_agents" "all" {}

variable "command_queue_id" {
  type        = string
  description = "The ID of the asset to which the command queue will be added."
}

variable "stream_id" {
  type        = string
  description = ""
}

variable "ground_station_id" {
  type        = string
  description = "The list of ground station IDs to which the command queue will be linked."
}

resource "leanspace_remote_agents" "test" {
  remote_agent {
    name        = "Terraform Remote Agent"
    description = "A basic remote agent made with terraform."
    connectors {
      gateway_id       = var.ground_station_id
      type             = "OUTBOUND"
      command_queue_id = var.command_queue_id
      socket {
        type = "UDP"
        host = "myhost"
        port = 456
      }
    }
    connectors {
      gateway_id = var.ground_station_id
      type       = "INBOUND"
      stream_id  = var.stream_id
      socket {
        type = "TCP"
        port = 123
      }
    }

  }
}

output "test_remote_agent" {
  value = leanspace_remote_agents.test
}
