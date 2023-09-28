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
  metadata {
    name        = "ActivityMetadataText"
    description = "A text metadata value"
    attributes {
      value = "test"
      type  = "TEXT"
    }
  }
  metadata {
    name        = "ActivityMetadataBool"
    description = "A boolean metadata value"
    attributes {
      value = true
      type  = "BOOLEAN"
    }
  }
  metadata {
    name        = "ActivityMetadataTimestamp"
    description = "A timestamp metadata value"
    attributes {
      value = "2022-06-30T13:57:23Z"
      type  = "TIMESTAMP"
    }
  }
  metadata {
    name        = "ActivityMetadataDate"
    description = "A date metadata value"
    attributes {
      value = "2022-06-30"
      type  = "DATE"
    }
  }
  metadata {
    name        = "ActivityMetadataTime"
    description = "A time metadata value"
    attributes {
      value = "10:37:19"
      type  = "TIME"
    }
  }
  metadata {
    name        = "ActivityMetadataTime"
    description = "A time metadata value"
    attributes {
      value = "62696e617279"
      data_type = "BINARY"
      type  = "ARRAY"
    }
  }

  argument_definitions {
    name        = "ActivityArgumentNumeric"
    description = "A numeric input"
    attributes {
      default_value = 2
      type          = "NUMERIC"
      required      = true
    }
  }
  argument_definitions {
    name        = "ActivityArgumentText"
    description = "A text input"
    attributes {
      default_value = "test"
      type          = "TEXT"
    }
  }
  argument_definitions {
    name        = "ActivityArgumentBool"
    description = "A boolean input"
    attributes {
      default_value = true
      type          = "BOOLEAN"
      required      = true
    }
  }
  argument_definitions {
    name        = "ActivityArgumentTimestamp"
    description = "A timestamp input"
    attributes {
      default_value = "2022-06-30T13:57:23Z"
      type          = "TIMESTAMP"
      required      = true
    }
  }
  argument_definitions {
    name        = "ActivityArgumentDate"
    description = "A date input"
    attributes {
      default_value = "2022-06-30"
      type          = "DATE"
      required      = true
    }
  }
  argument_definitions {
    name        = "ActivityArgumentTime"
    description = "A time input"
    attributes {
      default_value = "10:37:19"
      type          = "TIME"
      required      = true
    }
  }
  argument_definitions {
    name        = "ActivityArgumentEnum"
    description = "An enum input"
    attributes {
      default_value = 1
      options       = { 1 = "test" }
      type          = "ENUM"
      required      = true
    }
  }
  argument_definitions {
    name        = "TestArgumentBinaryArray"
    identifier  = "Binary ARRAY"
    description = "A binary array"
    attributes {
      type          = "ARRAY"
      required      = true
      min_size      = 1
      max_size      = 4
      unique        = true
      default_value = "62696e617279,62696e617279"
      constraint {
        type = "BINARY"
        min_length  = 1
        max_length  = 10
      }
    }
  }

  command_mappings {
    command_definition_id = var.command_definition.id
    delay_in_milliseconds = 0
    metadata_mappings {
      activity_definition_metadata_name = "ActivityMetadataText"
      command_definition_argument_name  = local.arguments[index(local.arguments.*.attributes.0.type, "TEXT")].name
    }
    metadata_mappings {
      activity_definition_metadata_name = "ActivityMetadataNumeric"
      command_definition_argument_name  = local.arguments[index(local.arguments.*.attributes.0.type, "NUMERIC")].name
    }
    argument_mappings {
      activity_definition_argument_name = "ActivityArgumentEnum"
      command_definition_argument_name  = local.arguments[index(local.arguments.*.attributes.0.type, "ENUM")].name
    }
  }
}
