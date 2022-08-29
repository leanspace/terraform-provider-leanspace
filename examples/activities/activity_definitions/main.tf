terraform {
  required_providers {
    leanspace = {
      source = "app.terraform.io/leanspace/leanspace"
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
    ids      = []
    query    = ""
    page     = 0
    size     = 10
    sort     = ["name,asc"]
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
