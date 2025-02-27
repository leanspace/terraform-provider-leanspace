terraform {
  required_providers {
    leanspace = {
      source = "leanspace/leanspace"
    }
  }
}

variable "contact_state_id" {
  type        = string
  description = "The ID of the node to which the command definitions will be added."
}

resource "leanspace_leaf_space_contact_reservation_status_mappings" "contact_reservation_status_mapping" {
  contact_state_id = var.contact_state_id
  leafspace_status = "Terraform test"

}

data "leanspace_leaf_space_contact_reservation_status_mappings" "all" {
  filters {
    leafspace_statuses = ["Scheduled"]
    ids                = []
    query              = ""
    page               = 0
    size               = 10
    sort               = ["asc"]
  }
}

output "all" {
  value = data.leanspace_leaf_space_contact_reservation_status_mappings.all
}