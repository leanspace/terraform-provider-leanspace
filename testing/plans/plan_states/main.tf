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
  name = "TERRAFORM_STATE"
}

output "created" {
  value = leanspace_plan_states.created
}
