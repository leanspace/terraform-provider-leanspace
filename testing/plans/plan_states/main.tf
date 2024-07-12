terraform {
  required_providers {
    leanspace = {
      source = "leanspace/leanspace"
    }
  }
}

data "leanspace_new_plan_states" "all" {

}

resource "leanspace_new_plan_states" "created" {
  name = "MY_TEST"
}

output "created" {
  value = leanspace_new_plan_states.created
}
