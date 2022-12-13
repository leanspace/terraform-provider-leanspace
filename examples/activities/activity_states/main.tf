terraform {
  required_providers {
    leanspace = {
      source = "app.terraform.io/leanspace/leanspace"
    }
  }
}

data "leanspace_activity_states" "all" {

}

resource "leanspace_activity_states" "created" {
  name = "MY_TEST"
}

output "created" {
  value = leanspace_activity_states.created
}