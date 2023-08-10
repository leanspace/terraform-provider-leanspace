terraform {
  required_providers {
    leanspace = {
      source  = "leanspace/leanspace"
      version = "8.5.10"
    }
  }
}

provider "leanspace" {
  tenant        = "yuri"
  env           = "develop"
  client_id     = "nlbja2p65j8kj7of0tfs29rf4"
  client_secret = "d762kk9862jn0j1qr4c2u3o8bjkv70o45pld3200ek89qtul6kg"
}

resource "leanspace_leaf_space_satellites_link" "ground_station_link" {
  leafspace_satellite_id="c736e7ed36916f5d4c14e454087a3dc2"
  leanspace_satellite_id= "3ff7b5e4-0a5c-4d32-99a6-1f366ee3d069"

}

data "leanspace_leaf_space_satellites_link" "all" {
  filters {
    leafspace_satellite_ids = ["c736e7ed36916f5d4c14e454087a3dc2"]
    leanspace_satellite_ids = ["3ff7b5e4-0a5c-4d32-99a6-1f366ee3d069"]
    ids                = []
    query              = ""
    page               = 0
    size               = 10
    sort               = ["asc"]
  }
}

output "all" {
  value = data.leanspace_leaf_space_satellites_link.all
}