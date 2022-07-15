terraform {
  required_providers {
    leanspace = {
      version = "0.2"
      source  = "leanspace.io/io/asset"
    }
  }
}

variable "usernames" {
  type        = list(string)
  description = "The usernames for the service accounts to create."
}

variable "access_policies" {
  type        = list(string)
  description = "The policies to attach to the created service accounts."
}

data "leanspace_service_accounts" "all" {}

resource "leanspace_service_accounts" "test" {
  for_each = { for u in var.usernames : u => u }
  service_account {
    name       = each.value
    policy_ids = var.access_policies
  }
}

output "test_service_accounts" {
  value = leanspace_service_accounts.test
}
