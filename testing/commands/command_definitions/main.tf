terraform {
  required_providers {
    leanspace = {
      source = "leanspace/leanspace"
    }
  }
}

variable "node_id" {
  type        = string
  description = "The ID of the node to which the command definitions will be added."
}

data "leanspace_command_definitions" "all" {
  filters {
    node_ids                    = [var.node_id]
    node_types                  = ["ASSET"]
    node_kinds                  = ["SATELLITE"]
    with_arguments_and_metadata = true
    ids                         = []
    created_bys                 = []
    query                       = ""
    page                        = 0
    size                        = 10
    sort                        = ["name,asc"]
  }
}

resource "leanspace_command_definitions" "test" {
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
      value = "10:37:19.000"
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
      default_value = "10:37:19.000"
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
  arguments {
    name        = "TestArgumentBinary"
    identifier  = "BINARY"
    description = "A binary input"
    attributes {
      default_value = "62696e617279"
      type          = "BINARY"
      required      = true
    }
  }
  arguments {
    name        = "TestArgumentNumericArray"
    identifier  = "Numeric ARRAY"
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
  arguments {
    name        = "TestArgumentTextArray"
    identifier  = "Text ARRAY"
    description = "A text array"
    attributes {
      type          = "ARRAY"
      required      = true
      min_size      = 1
      max_size      = 4
      unique        = true
      default_value = "value1,value2,value3"
      constraint {
        type       = "TEXT"
        min_length = 5
        max_length = 10
      }
    }
  }
  arguments {
    name        = "TestArgumentBooleanArray"
    identifier  = "Boolean ARRAY"
    description = "A boolean array"
    attributes {
      type          = "ARRAY"
      required      = true
      min_size      = 1
      max_size      = 4
      unique        = false
      default_value = "true,false,true"
      constraint {
        type = "BOOLEAN"
      }
    }
  }
  arguments {
    name        = "TestArgumentTimeArray"
    identifier  = "Time ARRAY"
    description = "A time array"
    attributes {
      type          = "ARRAY"
      required      = true
      min_size      = 1
      max_size      = 4
      unique        = true
      default_value = "08:37:19.000,10:37:19.000,15:37:19.000"
      constraint {
        type   = "TIME"
        before = "20:00:00"
        after  = "07:00:00"
      }
    }
  }
  arguments {
    name        = "TestArgumentDateArray"
    identifier  = "Date ARRAY"
    description = "A date array"
    attributes {
      type          = "ARRAY"
      required      = true
      min_size      = 1
      max_size      = 4
      unique        = true
      default_value = "2023-03-30,2023-05-11,2023-07-02"
      constraint {
        type   = "DATE"
        before = "2023-08-01"
        after  = "2023-02-01"
      }
    }
  }
  arguments {
    name        = "TestArgumentTimeStampArray"
    identifier  = "TimeStamp ARRAY"
    description = "A timeStamp array"
    attributes {
      type          = "ARRAY"
      required      = true
      min_size      = 1
      max_size      = 4
      unique        = true
      default_value = "2023-01-30T13:00:00Z,2023-01-29T01:00:00Z,2023-01-31T19:57:23Z"
      constraint {
        type   = "TIMESTAMP"
        before = "2023-01-31T20:00:00Z"
        after  = "2023-01-29T00:00:00Z"
      }
    }
  }
  arguments {
    name        = "TestArgumentEnumArray"
    identifier  = "Enum ARRAY"
    description = "A enum array"
    attributes {
      type          = "ARRAY"
      required      = true
      min_size      = 1
      max_size      = 4
      unique        = false
      default_value = "1,2,3,1"
      constraint {
        type    = "ENUM"
        options = { 1 = "value1", 2 = "value2", 3 = "value3" }
      }
    }
  }
  arguments {
    name        = "TestArgumentBinaryArray"
    identifier  = "Binary ARRAY"
    description = "A binary array"
    attributes {
      type          = "ARRAY"
      required      = true
      min_size      = 1
      max_size      = 4
      unique        = true
      default_value = "62696e617279"
      constraint {
        type       = "BINARY"
        min_length = 1
        max_length = 10
      }
    }
  }
}

output "test_command_definition" {
  value = leanspace_command_definitions.test
}

output "all_command_definitions" {
  value = data.leanspace_command_definitions.all
}
