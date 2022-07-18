terraform {
  required_providers {
    leanspace = {
      version = "0.2"
      source  = "leanspace.io/io/asset"
    }
  }
}

variable "usernames" {
  type        = set(string)
  description = "The usernames for the accounts to create."
}

variable "access_policies" {
  type        = list(string)
  description = "The policies to attach to the created users."
}

data "leanspace_members" "all" {}

resource "leanspace_members" "test" {
  for_each = var.usernames
  member {
    name       = each.value
    email      = "${lower(each.value)}@terraform.leanspace.io"
    policy_ids = var.access_policies
  }
}

output "test_members" {
  value = leanspace_members.test
}
