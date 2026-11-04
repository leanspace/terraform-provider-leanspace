resource "leanspace_members" "test" {
  security_tags = ["terraform_mission"]
  name          = "John Doe"
  email         = "john@terraform.leanspace.io"
  policy_ids    = var.access_policies
}
