terraform {
  required_providers {
    leanspace = {
      version = "0.3"
      source  = "leanspace.io/io/leanspace"
    }
  }
}

data "leanspace_properties" "all" {}

variable "node_id" {
  type = string
  description = "The ID of the node to which the properties will be added."
}

resource "leanspace_properties" "test_text" {
  property {
    name = "TestTerraformText"
    description = "TestTerraformUpdated2"
    node_id = var.node_id
    type = "TEXT"
    value = "test"
    tags {
      key = "Key1"
      value = "Value1"
    }
    tags {
      key = "Key2"
      value = "Value2"
    }
  }
}

resource "leanspace_properties" "test_numeric" {
  property {
    name = "TestTerraformNumeric"
    description = "TestTerraformUpdated2"
    node_id = var.node_id
    type = "NUMERIC"
    value = 2
    tags {
      key = "Key1"
      value = "Value1"
    }
    tags {
      key = "Key2"
      value = "Value2"
    }
  }
}

resource "leanspace_properties" "test_numeric_large" {
  property {
    name = "TestTerraformNumericLarge"
    description = "TestTerraformUpdated2"
    node_id = var.node_id
    type = "NUMERIC"
    value = "2e+10"
  }
}

resource "leanspace_properties" "test_bool" {
  property {
    name = "TestTerraformBool"
    description = "TestTerraformUpdated2"
    node_id = var.node_id
    type = "BOOLEAN"
    value = true
    tags {
      key = "Key1"
      value = "Value1"
    }
    tags {
      key = "Key2"
      value = "Value2"
    }
  }
}

resource "leanspace_properties" "test_timestamp" {
  property {
    name = "TestTerraformTimestamp"
    description = "TestTerraformUpdated2"
    node_id = var.node_id
    type = "TIMESTAMP"
    value = "2022-06-30T13:57:23Z"
    tags {
      key = "Key1"
      value = "Value1"
    }
    tags {
      key = "Key2"
      value = "Value2"
    }
  }
}

resource "leanspace_properties" "test_date" {
  property {
    name = "TestTerraformDate"
    description = "TestTerraformUpdated2"
    node_id = var.node_id
    type = "DATE"
    value = "2022-06-30"
    tags {
      key = "Key1"
      value = "Value1"
    }
    tags {
      key = "Key2"
      value = "Value2"
    }
  }
}

resource "leanspace_properties" "test_time" {
  property {
    name = "TestTerraformTime"
    description = "TestTerraformUpdated2"
    node_id = var.node_id
    type = "TIME"
    value = "10:37:19"
    tags {
      key = "Key1"
      value = "Value1"
    }
    tags {
      key = "Key2"
      value = "Value2"
    }
  }
}

resource "leanspace_properties" "test_geopoint" {
  property {
    name = "TestTerraformGeoPoint"
    description = "TestTerraformUpdated2"
    node_id = var.node_id
    type = "GEOPOINT"
    fields {
      elevation {
        name = "elevation"
        type = "NUMERIC"
        value = 2.1
      }
      latitude {
        name = "latitude"
        type = "NUMERIC"
        value = 3.2
      }
      longitude {
        name = "longitude"
        type = "NUMERIC"
        value = 4.5
      }
    }
    tags {
      key = "Key1"
      value = "Value1"
    }
    tags {
      key = "Key2"
      value = "Value2"
    }
  }
}

resource "leanspace_properties" "test_enum" {
  property {
    name = "TestTerraformEnum"
    description = "TestTerraformUpdated2"
    node_id = var.node_id
    type = "ENUM"
    value = 1
    options = {1="test"}
    tags {
      key = "Key1"
      value = "Value1"
    }
    tags {
      key = "Key2"
      value = "Value2"
    }
  }
}

output "test_numeric_property" {
  value = leanspace_properties.test_numeric
}

output "test_numeric_property_large" {
  value = leanspace_properties.test_numeric_large
}

output "test_text_property" {
  value = leanspace_properties.test_text
}

output "test_bool_property" {
  value = leanspace_properties.test_bool
}

output "test_timestamp_property" {
  value = leanspace_properties.test_timestamp
}

output "test_date_property" {
  value = leanspace_properties.test_date
}

output "test_time_property" {
  value = leanspace_properties.test_time
}

output "test_geopoint_property" {
  value = leanspace_properties.test_geopoint
}

output "test_enum_property" {
  value = leanspace_properties.test_enum
}