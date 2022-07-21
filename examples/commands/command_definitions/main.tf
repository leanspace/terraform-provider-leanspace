terraform {
  required_providers {
    leanspace = {
      version = "0.3.0"
      source  = "app.terraform.io/leanspace/leanspace"
    }
  }
}

data "leanspace_command_definitions" "all" {}

variable "node_id" {
  type        = string
  description = "The ID of the node to which the command definitions will be added."
}

resource "leanspace_command_definitions" "test" {
  command_definition {
    name        = "Terraform Command"
    description = "A complex command definition, entirely created under terraform."
    node_id     = var.node_id
    identifier  = "TERRA_CMD"

    metadata {
      name        = "TestMetadataNumeric"
      description = "A numeric metadata value"
      attributes {
        value = 2
        type  = "NUMERIC"
      }
    }
    metadata {
      name        = "TestMetadataText"
      description = "A text metadata value"
      attributes {
        value = "test"
        type  = "TEXT"
      }
    }
    metadata {
      name        = "TestMetadataBool"
      description = "A boolean metadata value"
      attributes {
        value = true
        type  = "BOOLEAN"
      }
    }
    metadata {
      name        = "TestMetadataTimestamp"
      description = "A timestamp metadata value"
      attributes {
        value = "2022-06-30T13:57:23Z"
        type  = "TIMESTAMP"
      }
    }
    metadata {
      name        = "TestMetadataDate"
      description = "A date metadata value"
      attributes {
        value = "2022-06-30"
        type  = "DATE"
      }
    }
    metadata {
      name        = "TestMetadataTime"
      description = "A time metadata value"
      attributes {
        value = "10:37:19"
        type  = "TIME"
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
    arguments {
      name        = "TestArgumentText"
      identifier  = "TEXT"
      description = "A text input"
      attributes {
        default_value = "test"
        type          = "TEXT"
      }
    }
    arguments {
      name        = "TestArgumentBool"
      identifier  = "BOOL"
      description = "A boolean input"
      attributes {
        default_value = true
        type          = "BOOLEAN"
        required      = true
      }
    }
    arguments {
      name        = "TestArgumentTimestamp"
      identifier  = "TIMESTAMP"
      description = "A timestamp input"
      attributes {
        default_value = "2022-06-30T13:57:23Z"
        type          = "TIMESTAMP"
        required      = true
      }
    }
    arguments {
      name        = "TestArgumentDate"
      identifier  = "DATE"
      description = "A date input"
      attributes {
        default_value = "2022-06-30"
        type          = "DATE"
        required      = true
      }
    }
    arguments {
      name        = "TestArgumentTime"
      identifier  = "TIME"
      description = "A time input"
      attributes {
        default_value = "10:37:19"
        type          = "TIME"
        required      = true
      }
    }
    arguments {
      name        = "TestArgumentEnum"
      identifier  = "ENUM"
      description = "An enum input"
      attributes {
        default_value = 1
        options       = { 1 = "test" }
        type          = "ENUM"
        required      = true
      }
    }
  }
}

output "test_command_definition" {
  value = leanspace_command_definitions.test
}
