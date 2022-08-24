resource "leanspace_activity_definitions" "activity_definition" {
  name               = "Terraform Activity Definition"
  description        = "An activity definition, created under terraform."
  node_id            = var.node_id
  estimated_duration = 3

  metadata {
    name        = "ActivityMetadataNumeric"
    description = "A numeric metadata value"
    attributes {
      value = 2
      type  = "NUMERIC"
    }
  }

  argument_definitions {
    name        = "ActivityArgumentText"
    description = "A text input"
    attributes {
      type = "TEXT"
    }
  }

  command_mappings {
    command_definition_id = var.command_definition_id
    delay_in_milliseconds = 0
    metadata_mappings {
      activity_definition_metadata_name = "ActivityMetadataNumeric"
      command_definition_argument_name  = "CommandArgumentNumeric"
    }
    argument_mappings {
      activity_definition_argument_name = "ActivityArgumentText"
      command_definition_argument_name  = "CommandArgumentText"
    }
  }
}
