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
  name = "MY_TEST"
}

output "created" {
  value = leanspace_contact_states.created
}
