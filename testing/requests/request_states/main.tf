terraform {
  required_providers {
    leanspace = {
      source = "leanspace/leanspace"
    }
  }
}

data "leanspace_request_states" "all" {
  filters {
    ids   = []
    query = ""
    page  = 0
    size  = 10
    sort  = ["name,asc"]
  }
}

resource "leanspace_request_states" "created" {
  name = "TERRAFORM_STATE"
}

output "created" {
  value = leanspace_request_states.created
}
