variable "asset_id" {
  type        = string
  description = "The ID of the node to which the stream will be added."
}

variable "numeric_metric_id" {
  type        = string
  description = "The ID of a numeric metric to which the stream will be mapped."
}

resource "leanspace_streams" "stream" {
  name     = "Terraform Stream"
  asset_id = var.asset_id

  configuration {
    endianness = "BE"
    structure {
      elements {
        type      = "FIELD"
        data_type = "UINTEGER"
        name      = "version"
        length {
          unit  = "BITS"
          type  = "FIXED"
          value = 8
        }
      }
      elements {
        type = "SWITCH"
        name = "data"
        expression {
          switch_on = "structure.version"
          options {
            component = "structure.data.data_0"
            value {
              data_type = "INTEGER"
              data      = 0
            }
          }
          options {
            component = "structure.data.data_1"
            value {
              data_type = "INTEGER"
              data      = 1
            }
          }
        }
        elements {
          type = "CONTAINER"
          name = "data_0"
          elements {
            type      = "FIELD"
            data_type = "TEXT"
            name      = "name"
            length {
              unit  = "BITS"
              type  = "FIXED"
              value = 32
            }
            processor = "zlib"
          }
        }
        elements {
          type = "CONTAINER"
          name = "data_1"
          elements {
            type      = "FIELD"
            data_type = "BOOLEAN"
            name      = "is_active"
            length {
              unit  = "BITS"
              type  = "FIXED"
              value = 8
            }
          }
        }
      }
    }
    metadata {
      timestamp {
        expression = "(ctx, raw) => ctx['metadata.received_at'];"
      }
    }
    computations {
      elements {
        data_type  = "UINTEGER"
        name       = "is_version_0"
        expression = <<-EOT
            (ctx) => ctx['structure.version'] === 0
          EOT
      }
    }
  }
  mappings {
    metric_id  = var.numeric_metric_id
    expression = "$..x"
  }
}
