terraform {
  required_providers {
    leanspace = {
      source = "leanspace/leanspace"
    }
  }
}

resource "leanspace_event_criticalities" "test" {
  name = "TERRAFORM_CRITICALITYT"
}