terraform {
  required_providers {
    leanspace = {
      source = "leanspace/leanspace"
    }
  }
}


resource "leanspace_leaf_space_integrations" "integration_Connection" {
  name       = "Leafspace Connection Terraform"
  username   = "leanspace01"
  password   = "vVVy5syuUuSk9MVkRfvYm8Xp6Nhy6qAs"
  domain_url = "sandbox.api.leaf.space"
}
