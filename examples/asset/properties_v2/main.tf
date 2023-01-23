terraform {
  required_providers {
    leanspace = {
      source = "leanspace/leanspace"
    }
  }
}

variable "node_id" {
  type        = string
  description = "The ID of the node to which the properties will be added."
}

data "leanspace_properties_v2" "all" {
  filters {
    node_ids   = [var.node_id]
    node_types = ["ASSET"]
    node_kinds = ["SATELLITE"]
    tags       = []
    ids        = []
    query      = ""
    page       = 0
    size       = 10
    sort       = ["name,asc"]
  }
}

resource "leanspace_properties_v2" "test_numeric" {
  name        = "TestTerraformNumeric"
  description = "TestTerraformNumericDescription"
  node_id     = var.node_id
  tags {
    key   = "Key1"
    value = "Value1"
  }
  tags {
    key   = "Key2"
    value = "Value2"
  }
  attributes {
    type = "NUMERIC"
    value = 100
    min = 50
    max = 200
    scale = 0
    precision = 0
    unit_id = null
  }
}

resource "leanspace_properties_v2" "test_text" {
  name        = "TestTerraformText"
  description = "TestTerraformTextDescription"
  node_id     = var.node_id
  tags {
    key   = "Key1"
    value = "Value1"
  }
  tags {
    key   = "Key2"
    value = "Value2"
  }
  attributes {
    type = "TEXT"
    value = "leanspace"
    min_length = 2
    max_length = 15
  }
}

resource "leanspace_properties_v2" "test_enum" {
  name        = "TestTerraformEnum"
  description = "TestTerraformEnumDescription"
  node_id     = var.node_id
  tags {
    key   = "Key1"
    value = "Value1"
  }
  tags {
    key   = "Key2"
    value = "Value2"
  }
  attributes {
    type = "ENUM"
    value = 2
    options       = { 
      1 = "value1"
      2 = "value2"
      3 = "value3"
    }
  }
}

resource "leanspace_properties_v2" "test_timestamp" {
  name        = "TestTerraformTimestamp"
  description = "TestTerraformTimestampDescription"
  node_id     = var.node_id
  tags {
    key   = "Key1"
    value = "Value1"
  }
  tags {
    key   = "Key2"
    value = "Value2"
  }
  attributes {
    type = "TIMESTAMP"
    value = "2023-01-30T00:00:00Z"
    before  = "2023-01-31T20:00:00Z"
    after = "2023-01-29T00:00:00Z"
  }
}

resource "leanspace_properties_v2" "test_date" {
  name        = "TestTerraformDate"
  description = "TestTerraformDateDescription"
  node_id     = var.node_id
  tags {
    key   = "Key1"
    value = "Value1"
  }
  tags {
    key   = "Key2"
    value = "Value2"
  }
  attributes {
    type = "DATE"
    value = "2023-05-01"
    before  = "2023-08-01"
    after = "2023-01-01"
  }
}

resource "leanspace_properties_v2" "test_time" {
  name        = "TestTerraformTime"
  description = "TestTerraformTimeDescription"
  node_id     = var.node_id
  tags {
    key   = "Key1"
    value = "Value1"
  }
  tags {
    key   = "Key2"
    value = "Value2"
  }
  attributes {
    type = "TIME"
    value = "10:00:00"
    before  = "20:00:00"
    after = "08:00:00"
  }
}

resource "leanspace_properties_v2" "test_boolean" {
  name        = "TestTerraformBoolean"
  description = "TestTerraformBooleanDescription"
  node_id     = var.node_id
  tags {
    key   = "Key1"
    value = "Value1"
  }
  tags {
    key   = "Key2"
    value = "Value2"
  }
  attributes {
    type = "BOOLEAN"
    value = true
  }
}

resource "leanspace_properties_v2" "test_geopoint" {
  name        = "TestTerraformGeopoint"
  description = "TestTerraformGeopointDescription"
  node_id     = var.node_id
  tags {
    key   = "Key1"
    value = "Value1"
  }
  tags {
    key   = "Key2"
    value = "Value2"
  }
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