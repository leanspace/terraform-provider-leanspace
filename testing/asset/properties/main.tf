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

data "leanspace_properties" "all" {
  filters {
    node_ids              = [var.node_id]
    category              = ""
    created_by            = null
    from_created_at       = null
    last_modified_by      = null
    to_created_at         = null
    from_last_modified_at = null
    to_last_modified_at   = null
    ids                   = []
    kinds                 = []
    node_types            = []
    query                 = ""
    tags                  = []
    page                  = 0
    size                  = 10
    sort                  = ["name,asc"]
  }
}

resource "leanspace_properties" "numeric_node_property" {
  name        = "TestTerraformNumeric"
  description = "TestTerraformNumericDescription"
  node_id     = var.node_id
  type = "NUMERIC"
  value = 100
  min = 50
  max = 200
  scale = 0
  precision = 0
  unit_id = null
    tags {
    key   = "Key1"
    value = "Value1"
  }
  tags {
    key   = "Key2"
    value = "Value2"
  }
}

resource "leanspace_properties" "text_node_property" {
  name        = "TestTerraformText"
  description = "TestTerraformTextDescription"
  node_id     = var.node_id
  type = "TEXT"
  value = "leanspace"
  min_length = 2
  max_length = 15
  tags {
    key   = "Key1"
    value = "Value1"
  }
  tags {
    key   = "Key2"
    value = "Value2"
  }
}

resource "leanspace_properties" "enum_node_property" {
  name        = "TestTerraformEnum"
  description = "TestTerraformEnumDescription"
  node_id     = var.node_id
  type = "ENUM"
  value = 2
  options  = { 
      1 = "value1"
      2 = "value2"
      3 = "value3"
  }
  tags {
    key   = "Key1"
    value = "Value1"
  }
  tags {
    key   = "Key2"
    value = "Value2"
  }
}

resource "leanspace_properties" "timestamp_node_property" {
  name        = "TestTerraformTimestamp"
  description = "TestTerraformTimestampDescription"
  node_id     = var.node_id
  type = "TIMESTAMP"
  value = "2023-01-30T00:00:00Z"
  before  = "2023-01-31T20:00:00Z"
  after = "2023-01-29T00:00:00Z"
  tags {
    key   = "Key1"
    value = "Value1"
  }
  tags {
    key   = "Key2"
    value = "Value2"
  }
}

resource "leanspace_properties" "date_node_property" {
  name        = "TestTerraformDate"
  description = "TestTerraformDateDescription"
  node_id     = var.node_id
  type = "DATE"
  value = "2023-05-01"
  before  = "2023-08-01"
  after = "2023-01-01"
  tags {
    key   = "Key1"
    value = "Value1"
  }
  tags {
    key   = "Key2"
    value = "Value2"
  }
}

resource "leanspace_properties" "time_node_property" {
  name        = "TestTerraformTime"
  description = "TestTerraformTimeDescription"
  node_id     = var.node_id
  type = "TIME"
  value = "10:00:00"
  before  = "20:00:00"
  after = "08:00:00"
  tags {
    key   = "Key1"
    value = "Value1"
  }
  tags {
    key   = "Key2"
    value = "Value2"
  }
}

resource "leanspace_properties" "boolean_node_property" {
  name        = "TestTerraformBoolean"
  description = "TestTerraformBooleanDescription"
  node_id     = var.node_id
  type = "BOOLEAN"
  value = true
  tags {
    key   = "Key1"
    value = "Value1"
  }
  tags {
    key   = "Key2"
    value = "Value2"
  }
}

resource "leanspace_properties" "geopoint_node_property" {
  name        = "TestTerraformGeopoint"
  description = "TestTerraformGeopointDescription"
  node_id     = var.node_id
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
  tags {
    key   = "Key1"
    value = "Value1"
  }
  tags {
    key   = "Key2"
    value = "Value2"
  }
}

# Only works when you import the built-in property before "terraform apply"
# Import command : `terraform import leanspace_properties.tle_node_property <property_id>`
#resource "leanspace_properties" "tle_node_property" {
#  name        = "TLE"
#  description = "Built-in property for satellite TLE"
#  node_id     = var.node_id
#  type = "TLE"
#  value = "1 99944U 98067A   20097.18686503  .00000920  00000-0  25115-4 0  9994,2 99944  51.6465 344.7546 0003971  92.6495  47.0504 15.48684294220999"
#  tags {
#    key   = "Key1"
#    value = "Value1"
#  }
#  tags {
#    key   = "Key2"
#    value = "Value2"
#  }
#}

output "all_properties" {
  value = data.leanspace_properties.all
}