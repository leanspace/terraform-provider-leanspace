terraform {
  required_providers {
    leanspace = {
      source = "leanspace/leanspace"
    }
  }
}

data "leanspace_command_states" "all" {

}

resource "leanspace_command_states" "created" {
  name = "MY_TEST"
}

output "created" {
  value = leanspace_command_states.created
}