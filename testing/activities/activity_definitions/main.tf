terraform {
  required_providers {
    leanspace = {
      source = "leanspace/leanspace"
    }
  }
}

variable "node_id" {
  type        = string
  description = "The ID of the node to which the activity definitions will be added."
}

variable "command_definition" {
  type = object({
    id = string
    arguments = set(object({
      name = string
      attributes = list(object({
        type = string
      }))
    }))
  })
  description = "The command definition that will be used for this activity definition"
}

data "leanspace_activity_definitions" "all" {
  filters {
    node_ids = [var.node_id]
    ids = []
    query = ""
    page  = 0
    size  = 10
    sort = ["name,asc"]
  }
}

locals {
  arguments = tolist(var.command_definition.arguments)
}

resource "leanspace_activity_definitions" "test" {
  name               = "Terraform Activity Definition"
  description        = "A complex activity definition, entirely created under terraform."
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
    name        = "ActivityMetadataBinary"
    description = "A binary metadata value"
    attributes {
      value = "62696e617279"
      type  = "BINARY"
    }
  }
  metadata {
    name        = "ActivityMetadataGeoPoint"
    description = "A geopoint metadata value"
    attributes {
      type = "GEOPOINT"
      fields {
        elevation {
          value = 141.0
        }
        latitude {
          value = 48.5
        }
        longitude {
          value = 7.7
        }
      }
    }
  }
  metadata {
    name        = "ActivityMetadataArrayNumeric"
    description = "A Numeric Array metadata value"
    attributes {
      value     = "1,2"
      type      = "ARRAY"
      data_type = "NUMERIC"
    }
  }
  metadata {
    name        = "ActivityMetadataArrayBinary"
    description = "A Binary Array metadata value"
    attributes {
      value     = "62696e617279,62696e617279"
      type      = "ARRAY"
      data_type = "BINARY"
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
      default_value = "test3"
      type          = "TEXT"
    }
  }
  argument_definitions {
    name        = "ActivityArgumentBool"
    description = "A boolean input"
    attributes {
      default_value = true
      type          = "BOOLEAN"
      required      = false
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
    name        = "ActivityArgumentTime"
    description = "A time input"
    attributes {
      default_value = "10:37:19"
      type          = "TIME"
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
    name        = "ActivityArgumentEnum"
    description = "An enum input"
    attributes {
      default_value = 1
      options = {
        1 = "test1"
        2 = "test3"
      }
      type     = "ENUM"
      required = true
    }
  }
  argument_definitions {
    name        = "ActivityArgumentBinary"
    description = "A binary input"
    attributes {
      default_value = "62696e617279"
      type          = "BINARY"
      required      = true
    }
  }
  argument_definitions {
    name        = "ActivityArgumentGeoPoint"
    description = "A geopoint input"
    attributes {
      type = "GEOPOINT"
      fields {
        elevation {
          default_value = 141.0
        }
        latitude {
          default_value = 48.5
        }
        longitude {
          default_value = 7.7
        }
      }
      required = true
    }
  }
  argument_definitions {
    name        = "ActivityArgumentNumericArray"
    description = "A numeric array"
    attributes {
      type          = "ARRAY"
      required      = true
      min_size      = 1
      max_size      = 4
      unique        = true
      default_value = "1,2,3"
      constraint {
        type = "NUMERIC"
        min  = 1
        max  = 10
      }
    }
  }
  argument_definitions {
    name        = "ActivityArgumentBinaryArray"
    description = "A binary array"
    attributes {
      type          = "ARRAY"
      min_size      = 1
      max_size      = 40
      default_value = "62696e617279,62696e617279"
      constraint {
        type       = "BINARY"
        min_length = 1
        max_length = 10
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
  command_mappings {
    command_definition_id = var.command_definition.id
    delay_in_milliseconds = 30
  }
}

output "test_activity_definition" {
  value = leanspace_activity_definitions.test
}
