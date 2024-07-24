terraform {
  required_providers {
    leanspace = {
      source = "leanspace/leanspace"
    }
  }
}

resource "leanspace_leaf_space_satellite_links" "satellite_link" {
  leafspace_satellite_id = "c736e7ed36916f5d4c14e454087a3dc2"
  leanspace_satellite_id = "3ff7b5e4-0a5c-4d32-99a6-1f366ee3d069"

}

data "leanspace_leaf_space_satellite_links" "all" {
  filters {
    leafspace_satellite_ids = ["c736e7ed36916f5d4c14e454087a3dc2"]
    leanspace_satellite_ids = ["3ff7b5e4-0a5c-4d32-99a6-1f366ee3d069"]
    ids                     = []
    query                   = ""
    page                    = 0
    size                    = 10
    sort                    = ["asc"]
  }
}

output "all" {
  value = data.leanspace_leaf_space_satellite_links.all
}