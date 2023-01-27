variable "members" {
  type        = list(string)
  description = "The UUIDS of the members to add to the team."
}

variable "access_policies" {
  type        = list(string)
  description = "The policies to attach to the created service accounts."
}

resource "leanspace_teams" "team" {
  name       = "Terraform Team"
  policy_ids = var.access_policies
  members    = var.members
}
