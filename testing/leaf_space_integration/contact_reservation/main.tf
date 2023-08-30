terraform {
  required_providers {
    leanspace = {
      source = "leanspace/leanspace"
    }
  }
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