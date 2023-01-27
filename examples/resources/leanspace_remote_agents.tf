variable "command_queue_id" {
  type        = string
  description = "The ID of the command queue to be attached to the remote agent."
}

variable "stream_id" {
  type        = string
  description = "The ID of the stream to be attached to the remote agent."
}

variable "ground_station_id" {
  type        = string
  description = "The ground station ID that will be used as a gateway by the remote agent."
}

resource "leanspace_remote_agents" "remote_agent" {
  name        = "Terraform Remote Agent"
  description = "A basic remote agent made with terraform."
  connectors {
    type             = "OUTBOUND"
    gateway_id       = var.ground_station_id
    command_queue_id = var.command_queue_id
    socket {
      type = "UDP"
      host = "myhost"
      port = 456
    }
  }
  connectors {
    type       = "INBOUND"
    gateway_id = var.ground_station_id
    stream_id  = var.stream_id
    socket {
      type = "TCP"
      port = 123
    }
  }

}
