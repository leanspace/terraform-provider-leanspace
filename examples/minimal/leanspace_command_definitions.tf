resource "leanspace_command_definitions" "command_definition" {
  command_definition {
    name        = "Terraform Command"
    description = "A command definition, created under terraform."
    node_id     = var.node_id
    identifier  = "TERRA_CMD"

    metadata {
      name        = "TestMetadataText"
      description = "A text metadata value"
      attributes {
        value = "test"
        type  = "TEXT"
      }
    }

    arguments {
      name        = "TestArgumentNumeric"
      identifier  = "NUMERIC"
      description = "A numeric input"
      attributes {
        default_value = 2
        type          = "NUMERIC"
        required      = true
      }
    }
  }
}
