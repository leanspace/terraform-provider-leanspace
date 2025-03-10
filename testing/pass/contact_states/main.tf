terraform {
  required_providers {
    leanspace = {
      source = "leanspace/leanspace"
    }
  }
}

data "leanspace_contact_states" "all" {

}

resource "leanspace_contact_states" "created" {
  name = "TERRAFORM_STATE"
}

output "created" {
  value = leanspace_contact_states.created
}
