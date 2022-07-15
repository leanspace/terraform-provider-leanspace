terraform {
  required_providers {
    leanspace = {
      version = "0.2"
      source  = "leanspace.io/io/asset"
    }
  }
}

variable "members" {
  type        = list(string)
  description = "The UUIDS of the members to add to the team."
}

variable "access_policies" {
  type        = list(string)
  description = "The policies to attach to the created service accounts."
}

data "leanspace_teams" "all" {}

resource "leanspace_teams" "test" {
  team {
    name       = "Terraform Team"
    policy_ids = var.access_policies
    members    = var.members
  }
}

output "test_team" {
  value = leanspace_teams.test
}
