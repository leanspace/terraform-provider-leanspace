terraform {
  required_providers {
    leanspace = {
      source = "leanspace/leanspace"
    }
  }
}

data "leanspace_feasibility_constraint_definitions" "all" {
  filters {
    ids      = []
    query = ""
    page     = 0
    size     = 10
    sort     = ["name,asc"]
  }
}

resource "leanspace_feasibility_constraint_definitions" "created" {
  name        = "feasibilityConstraintDefinitionFromTerraform"
  description = "feasibilityConstraintDefinitionTerraformDescription"
  cloned = false
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
    name        = "ActivityArgumentTime"
    description = "A time input"
    attributes {
      default_value = "10:37:19"
      type          = "TIME"
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
}

output "created" {
  value = leanspace_feasibility_constraint_definitions.created
}
