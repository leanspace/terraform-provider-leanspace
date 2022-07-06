terraform {
  required_providers {
    leanspace = {
      version = "0.2"
      source  = "leanspace.io/io/asset"
    }
  }
}

data "leanspace_streams" "all" {}

variable "asset_id" {
  type        = string
  description = "The ID of the node to which the stream will be added."
}

variable "numeric_metric_id" {
  type        = string
  description = "The ID of a numeric metric to which the stream will be mapped."
}

resource "leanspace_streams" "test" {
  stream {
    name        = "Terraform Stream"
    description = "A complex stream, entirely created under terraform."
    asset_id    = var.asset_id
    configuration {
      endianness = "BE"
      structure {
        elements {
          type           = "FIELD"
          data_type      = "UINTEGER"
          name           = "id_field"
          order          = 0
          path           = "structure.id_field"
          length_in_bits = 8
        }
        elements {
          type  = "CONTAINER"
          name  = "properties"
          order = 1
          path  = "structure.properties"
          elements {
            type           = "FIELD"
            data_type      = "TEXT"
            name           = "name"
            order          = 0
            path           = "structure.properties.name"
            length_in_bits = 32
            processor      = "zlib"
          }
          elements {
            type           = "FIELD"
            data_type      = "INTEGER"
            name           = "version"
            order          = 1
            path           = "structure.properties.version"
            length_in_bits = 8
          }
          elements {
            type           = "FIELD"
            data_type      = "BOOLEAN"
            name           = "is_active"
            order          = 2
            path           = "structure.properties.is_active"
            length_in_bits = 8
          }
          elements {
            type           = "FIELD"
            data_type      = "DECIMAL"
            name           = "solar_w"
            order          = 3
            path           = "structure.properties.solar_w"
            length_in_bits = 32
          }
        }
        elements {
          type  = "SWITCH"
          name  = "data"
          order = 2
          path  = "structure.data"
          expression {
            switch_on = "structure.properties.version"
            options {
              component = "structure.data.pos_data"
              value {
                data_type = "INTEGER"
                data      = 0
              }
            }
            options {
              component = "structure.data.rot_data"
              value {
                data_type = "INTEGER"
                data      = 1
              }
            }
            options {
              component = "structure.data.rot_data"
              value {
                data_type = "INTEGER"
                data      = 2
              }
            }
          }
          elements {
            type = "CONTAINER"
            name = "pos_data"
            path = "structure.data.pos_data"
            elements {
              type           = "FIELD"
              data_type      = "INTEGER"
              name           = "x"
              order          = 0
              path           = "structure.data.pos_data.x"
              length_in_bits = 8
            }
            elements {
              type           = "FIELD"
              data_type      = "INTEGER"
              name           = "y"
              order          = 1
              path           = "structure.data.pos_data.y"
              length_in_bits = 8
            }
          }
          elements {
            type = "CONTAINER"
            name = "rot_data"
            path = "structure.data.rot_data"
            elements {
              type           = "FIELD"
              data_type      = "INTEGER"
              name           = "rx"
              order          = 0
              path           = "structure.data.rot_data.x"
              length_in_bits = 8
            }
            elements {
              type           = "FIELD"
              data_type      = "INTEGER"
              name           = "ry"
              order          = 1
              path           = "structure.data.rot_data.y"
              length_in_bits = 8
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
          type       = "COMPUTATION"
          data_type  = "UINTEGER"
          name       = "power"
          order      = 0
          expression = <<-EOT
            (ctx) => {
              const voltage = ctx['structure.properties.solar_w'];
              var power = voltage * 15;
              return (power);
            }
          EOT
        }
      }
    }
    mappings {
      metric_id = var.numeric_metric_id
      component = "power"
    }
  }
}

output "test_stream" {
  value = leanspace_streams.test
}
