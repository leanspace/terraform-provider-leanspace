terraform {
  required_providers {
    leanspace = {
      source = "app.terraform.io/leanspace/leanspace"
    }
  }
}

variable "usernames" {
  type        = set(string)
  description = "The usernames for the service accounts to create."
}

variable "access_policies" {
  type        = list(string)
  description = "The policies to attach to the created service accounts."
}

data "leanspace_service_accounts" "all" {}

resource "leanspace_service_accounts" "test" {
  for_each   = var.usernames
  name       = each.value
  policy_ids = var.access_policies
}

output "test_service_accounts" {
  value = leanspace_service_accounts.test
}
