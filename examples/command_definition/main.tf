terraform {
  required_providers {
    leanspace = {
      version = "0.2"
      source  = "leanspace.io/edu/asset"
    }
  }
}

data "leanspace_command_definitions" "all" {}

# Returns all properties
output "all_command_definition" {
  value = data.leanspace_command_definitions.all.command_definitions
}

output "first_command_definition" {
  value =  data.leanspace_command_definitions.all.command_definitions[0]
}

resource "leanspace_command_definitions" "test" {
  command_definition {
    name = "TestTerraform"
    description = "TestTerraformUpdated"
    node_id = "62e5449d-57a3-46bc-906b-b51e970bfdc9"
    identifier = "TEST"
    metadata {
        name = "TestMetadataNumeric"
        description = "TestMetadataUpdated"
        value = 2
        required = true
        type = "NUMERIC"
    }
    metadata {
        name = "TestMetadataText"
        description = "TestMetadataUpdated"
        value = "test"
        type = "TEXT"
    }
    metadata {
        name = "TestMetadataBool"
        description = "TestMetadataUpdated"
        value = true
        required = true
        type = "BOOLEAN"
    }
    metadata {
        name = "TestMetadataTimestamp"
        description = "TestMetadataUpdated"
        value = "2022-06-30T13:57:23Z"
        required = true
        type = "TIMESTAMP"
    }
    metadata {
        name = "TestMetadataDate"
        description = "TestMetadataUpdated"
        value = "2022-06-30"
        required = true
        type = "DATE"
    }
    metadata {
        name = "TestMetadataTime"
        description = "TestMetadataUpdated"
        value = "10:37:19"
        required = true
        type = "TIME"
    }
    arguments {
        name = "TestArgumentNumeric"
        identifier = "NUMERIC"
        description = "TestArgumentUpdated"
        default_value = 2
        type = "NUMERIC"
        required = true
    }
    arguments {
        name = "TestArgumentText"
        identifier = "TEXT"
        description = "TestArgumentUpdated"
        default_value = "test"
        type = "TEXT"
    }
    arguments {
        name = "TestArgumentBool"
        identifier = "BOOL"
        description = "TestArgumentUpdated"
        default_value = true
        type = "BOOLEAN"
        required = true
    }
    arguments {
        name = "TestArgumentTimestamp"
        identifier = "TIMESTAMP"
        description = "TestArgumentUpdated"
        default_value = "2022-06-30T13:57:23Z"
        type = "TIMESTAMP"
        required = true
    }
    arguments {
        name = "TestArgumentDate"
        identifier = "DATE"
        description = "TestArgumentUpdated"
        default_value = "2022-06-30"
        type = "DATE"
        required = true
    }
    arguments {
        name = "TestArgumentTime"
        identifier = "TIME"
        description = "TestArgumentUpdated"
        default_value = "10:37:19"
        type = "TIME"
        required = true
    }
    arguments {
        name = "TestArgumentEnum"
        identifier = "ENUM"
        description = "TestArgumentUpdated"
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