terraform {
  required_providers {
    leanspace = {
      source = "leanspace/leanspace"
      version = "3.1.0"
    }
  }
}

provider "leanspace" {
  tenant        = "yuri"
  env           = "develop"
  client_id     = "nlbja2p65j8kj7of0tfs29rf4"
  client_secret = "d762kk9862jn0j1qr4c2u3o8bjkv70o45pld3200ek89qtul6kg"
}

variable "satellite_id" {
  type        = string
  description = "The ID of the satellite to which the orbit resource will be added."
}

data "leanspace_orbit_resources" "all" {
  filters {
    satellite_ids                  = [var.satellite_id]
    ids                            = []
    data_sources                   = []
    query                          = ""
    page                           = 0
    size                           = 10
    sort                           = ["name,asc"]
  }
}

resource "leanspace_orbit_resources" "test" {
  name                            = "Terraform Orbit Resource"
  satellite_id                    = var.satellite_id
  data_source                     = "TLE_MANUAL"
}

output "test_orbit_resource" {
  value = leanspace_orbit_resources.test
}
