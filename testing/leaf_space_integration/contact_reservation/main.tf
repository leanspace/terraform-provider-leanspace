terraform {
  required_providers {
    leanspace = {
      source = "leanspace/leanspace"
      version = "8.5.5"
    }
  }
}

provider "leanspace" {
  tenant        = "yuri"
  env           = "develop"
  client_id     = "nlbja2p65j8kj7of0tfs29rf4"
  client_secret = "d762kk9862jn0j1qr4c2u3o8bjkv70o45pld3200ek89qtul6kg"
}

resource "leanspace_leaf_space_contact_reservations" "contact_reservation" {
  contact_state_id ="5f7f8218-56b4-4c40-9703-e92c5e68772c"
  leafspace_status= "terraform test"

}

data "leanspace_leaf_space_contact_reservations" "all" {
  filters {
    leafspace_statuses = ["Scheduled"]
    ids      = []
    query    = ""
    page     = 0
    size     = 10
    sort     = ["asc"]
  }
}

output "all" {
  value = data.leanspace_leaf_space_contact_reservations.all
}