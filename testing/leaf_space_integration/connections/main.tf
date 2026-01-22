terraform {
  required_providers {
    leanspace = {
      source = "leanspace/leanspace"
    }
  }
}

variable "password" {
  description = "Password for the Leafspace account"
  type        = string
}


resource "leanspace_leaf_space_integrations" "integration_Connection" {
  name       = "Leafspace Connection Terraform"
  username   = "leanspace01"
  password   = var.password
  domain_url = "apiv2.sandbox.leaf.space"

  lifecycle {
    ignore_changes = [ username, password ]
  }
}
