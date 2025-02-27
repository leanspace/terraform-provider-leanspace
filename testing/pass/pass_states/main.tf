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
  name = "TERRAFORM_STATE"
}

output "created" {
  value = leanspace_pass_states.created
}
