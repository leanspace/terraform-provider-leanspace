terraform {
  required_providers {
    leanspace = {
      source  = "leanspace/leanspace"
      version = "8.5.8"
    }
  }
}

provider "leanspace" {
  tenant        = "yuri"
  env           = "develop"
  client_id     = "nlbja2p65j8kj7of0tfs29rf4"
  client_secret = "d762kk9862jn0j1qr4c2u3o8bjkv70o45pld3200ek89qtul6kg"
}

resource "leanspace_leaf_space_ground_station_links" "ground_station_link" {
  leafspace_ground_station_id = "d5de2269dc23179929546f41b6239afb"
  leanspace_ground_station_id = "969e157d-8883-43cd-b851-3d7ff3449ec6"

}


data "leanspace_leaf_space_ground_station_links" "all" {
  filters {
    leafspace_ground_station_ids = ["d5de2269dc23179929546f41b6239afb"]
    leanspace_ground_station_ids = ["969e157d-8883-43cd-b851-3d7ff3449ec6"]
    ids                = []
    query              = ""
    page               = 0
    size               = 10
    sort               = ["asc"]
  }
}

output "all" {
  value = data.leanspace_leaf_space_ground_station_links.all
}