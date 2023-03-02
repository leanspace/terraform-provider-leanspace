terraform {
  required_providers {
    leanspace = {
      source = "leanspace/leanspace"
    }
  }
}

provider "leanspace" {
  tenant        = "yuri"
  env           = "develop"
  client_id     = "nlbja2p65j8kj7of0tfs29rf4"
  client_secret = "d762kk9862jn0j1qr4c2u3o8bjkv70o45pld3200ek89qtul6kg"
}

variable "asset_id" {
  type        = string
  description = "The ID of the node to which the stream will be added."
  default     = "082642bd-2385-46d6-bf58-26a3df467113"
}

variable "numeric_metric_id" {
  type        = string
  description = "The ID of a numeric metric to which the stream will be mapped."
  default     = "53c0b787-7186-413e-b5d7-627c4382c558"
}

data "leanspace_streams" "all" {
  filters {
    asset_ids = [var.asset_id]
    ids       = []
    query     = ""
    page      = 0
    size      = 10
    sort      = ["name,asc"]
  }
}

resource "leanspace_streams" "test" {
  name        = "Terraform Stream bla bla"
  description = "A complex stream, entirely created under terraform."
  asset_id    = var.asset_id
  configuration {
    endianness = "BE"
    structure {
      elements {
        type      = "FIELD"
        data_type = "UINTEGER"
        name      = "id_field"
        length {
          unit  = "BITS"
          type  = "FIXED"
          value = 8
        }
      }
      elements {
        type = "CONTAINER"
        name = "properties"
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
        elements {
          type      = "FIELD"
          data_type = "INTEGER"
          name      = "version"
          length {
            unit  = "BITS"
            type  = "FIXED"
            value = 8
          }
        }
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
        elements {
          type      = "FIELD"
          data_type = "DECIMAL"
          name      = "solar_w"
          length {
            unit  = "BITS"
            type  = "FIXED"
            value = 32
          }
        }
      }
      elements {
        type = "SWITCH"
        name = "data"
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
          elements {
            type      = "FIELD"
            data_type = "INTEGER"
            name      = "x"
            length {
              unit  = "BITS"
              type  = "FIXED"
              value = 8
            }
          }
          elements {
            type      = "FIELD"
            data_type = "INTEGER"
            name      = "y"
            length {
              unit  = "BITS"
              type  = "FIXED"
              value = 8
            }
          }
        }
        elements {
          type = "CONTAINER"
          name = "rot_data"
          elements {
            type      = "FIELD"
            data_type = "INTEGER"
            name      = "rx"
            length {
              unit  = "BITS"
              type  = "FIXED"
              value = 8
            }
          }
          elements {
            type      = "FIELD"
            data_type = "INTEGER"
            name      = "ry"
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
        name       = "power"
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

output "test_stream" {
  value = leanspace_streams.test
}
