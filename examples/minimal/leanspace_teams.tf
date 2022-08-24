resource "leanspace_teams" "team" {
  name       = "Terraform Team"
  policy_ids = var.access_policies
  members    = var.members
}
