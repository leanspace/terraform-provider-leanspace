terraform {
  required_providers {
    leanspace = {
      source = "leanspace/leanspace"
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

data "leanspace_teams" "all" {
  filters {
    member_ids = var.members
    ids        = []
    query      = ""
    page       = 0
    size       = 10
    sort       = ["name,asc"]
  }
}

resource "leanspace_teams" "test" {
  name       = "Terraform Team"
  policy_ids = var.access_policies
  members    = var.members
}


output "test_team" {
  value = leanspace_teams.test
}
