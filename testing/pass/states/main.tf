terraform {
  required_providers {
    leanspace = {
      source = "leanspace/leanspace"
    }
  }
}

data "leanspace_pass_states" "all" {

}

resource "leanspace_pass_states" "created" {
  name = "MY_TEST"
}

output "created" {
  value = leanspace_pass_states.created
}
