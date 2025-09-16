terraform {
  required_providers {
    leanspace = {
      source = "leanspace/leanspace"
    }
  }
}
resource "leanspace_events_criticalities" "test" {
  name        = "TERRAFORM_CRITICALITYT"
}