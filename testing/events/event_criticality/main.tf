terraform {
  required_providers {
    leanspace = {
      source = "leanspace/leanspace"
    }
  }
}



resource "leanspace_events_criticalities" "test" {
  name        = "TERRFORM_CRITICALITYT"
}