terraform {
  required_providers {
    leanspace = {
      source = "leanspace/leanspace"
    }
  }
}

data "leanspace_plan_states" "all" {

}

resource "leanspace_plan_states" "created" {
  name = "MY_TEST"
}

output "created" {
  value = leanspace_plan_states.created
}
