terraform {
  required_providers {
    leanspace = {
      version = "0.2"
      source  = "leanspace.io/io/asset"
    }
  }
}

data "leanspace_command_definitions" "all" {}

variable "node_id" {
  type = string
  description = "The ID of the node to which the properties will be added."
}

resource "leanspace_command_definitions" "test" {
  command_definition {
    name = "Terraform Command"
    description = "A complex command definition, entirely created under terraform."
    node_id = var.node_id
    identifier = "TERRA_CMD"
    metadata {
        name = "TestMetadataNumeric"
        description = "A numeric metadata value"
        value = 2
        required = true
        type = "NUMERIC"
    }
    metadata {
        name = "TestMetadataText"
        description = "A text metadata value"
        value = "test"
        type = "TEXT"
    }
    metadata {
        name = "TestMetadataBool"
        description = "A boolean metadata value"
        value = true
        required = true
        type = "BOOLEAN"
    }
    metadata {
        name = "TestMetadataTimestamp"
        description = "A timestamp metadata value"
        value = "2022-06-30T13:57:23Z"
        required = true
        type = "TIMESTAMP"
    }
    metadata {
        name = "TestMetadataDate"
        description = "A date metadata value"
        value = "2022-06-30"
        required = true
        type = "DATE"
    }
    metadata {
        name = "TestMetadataTime"
        description = "A time metadata value"
        value = "10:37:19"
        required = true
        type = "TIME"
    }
    arguments {
        name = "TestArgumentNumeric"
        identifier = "NUMERIC"
        description = "A numeric input"
        default_value = 2
        type = "NUMERIC"
        required = true
    }
    arguments {
        name = "TestArgumentText"
        identifier = "TEXT"
        description = "A text input"
        default_value = "test"
        type = "TEXT"
    }
    arguments {
        name = "TestArgumentBool"
        identifier = "BOOL"
        description = "A boolean input"
        default_value = true
        type = "BOOLEAN"
        required = true
    }
    arguments {
        name = "TestArgumentTimestamp"
        identifier = "TIMESTAMP"
        description = "A timestamp input"
        default_value = "2022-06-30T13:57:23Z"
        type = "TIMESTAMP"
        required = true
    }
    arguments {
        name = "TestArgumentDate"
        identifier = "DATE"
        description = "A date input"
        default_value = "2022-06-30"
        type = "DATE"
        required = true
    }
    arguments {
        name = "TestArgumentTime"
        identifier = "TIME"
        description = "A time input"
        default_value = "10:37:19"
        type = "TIME"
        required = true
    }
    arguments {
        name = "TestArgumentEnum"
        identifier = "ENUM"
        description = "An enum input"
        default_value = 1
        options = {1="test"}
        type = "ENUM"
        required = true
    }
  }
}

output "test_command_definition" {
  value = leanspace_command_definitions.test
}