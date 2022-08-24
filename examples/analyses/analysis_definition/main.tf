terraform {
  required_providers {
    leanspace = {
      source = "app.terraform.io/leanspace/leanspace"
    }
  }
}

data "leanspace_analysis_definitions" "all" {}

variable "node_id" {
  type        = string
  description = "The ID of the node on which to run the simulation."
}

variable "mass_property_id" {
  type        = string
  description = "The ID of the property containing the mass of the satellite."
}

variable "ground_station_id" {
  type        = string
  description = "The ID of the ground station to check the visibility for."
}

locals {
  node_lrn           = "lrn::leanspace::yuri::topology::node::${var.node_id}"
  ground_station_lrn = "lrn::leanspace::yuri::topology::node::${var.ground_station_id}"
  mass_property_lrn  = "lrn::leanspace::yuri::topology::property::${var.mass_property_id}"
}

resource "leanspace_analysis_definitions" "test" {
  name        = "Terraform Analysis Definition"
  description = "An analysis definition made under terraform!"
  model_id    = "d9f4e470-f4ad-456c-8b8e-47577a8212cd"
  node_id     = var.node_id
  inputs {
    type = "STRUCTURE"
    fields {
      name = "eclipses"
      type = "STRUCTURE"
      fields {
        name   = "earthEclipseEnabled"
        type   = "BOOLEAN"
        source = "STATIC"
        value  = false
      }
      fields {
        name   = "moonEclipseEnabled"
        type   = "BOOLEAN"
        source = "STATIC"
        value  = true
      }
    }

    fields {
      name = "groundStationVisibilities"
      type = "STRUCTURE"
      fields {
        name   = "groundStationVisibilitiesEnabled"
        type   = "BOOLEAN"
        source = "STATIC"
        value  = true
      }
      fields {
        name = "groundStations"
        type = "ARRAY"
        items {
          type = "STRUCTURE"
          fields {
            name   = "id"
            type   = "TEXT"
            source = "STATIC"
            value  = var.ground_station_id
          }
          fields {
            name   = "elevation"
            type   = "NUMERIC"
            source = "REFERENCE"
            ref    = "${local.ground_station_lrn}/elevation"
          }
          fields {
            name   = "latitude"
            type   = "NUMERIC"
            source = "REFERENCE"
            ref    = "${local.ground_station_lrn}/latitude"
          }
          fields {
            name   = "longitude"
            type   = "NUMERIC"
            source = "REFERENCE"
            ref    = "${local.ground_station_lrn}/longitude"
          }
          fields {
            name   = "minimumElevation"
            type   = "NUMERIC"
            source = "STATIC"
            value  = 1
          }
        }
      }
    }

    fields {
      name = "propagation"
      type = "STRUCTURE"
      fields {
        name   = "durationInSeconds"
        type   = "NUMERIC"
        source = "STATIC"
        value  = 20
      }
      fields {
        name   = "tle"
        type   = "TLE"
        source = "REFERENCE"
        ref    = "${local.node_lrn}/tle"
      }
      fields {
        name   = "mass"
        type   = "NUMERIC"
        source = "REFERENCE"
        ref    = "${local.mass_property_lrn}/value"
      }
    }
  }
}

output "test_analysis_definition" {
  value = leanspace_analysis_definitions.test
}
