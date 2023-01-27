resource "leanspace_members" "test" {
  name       = "John Doe"
  email      = "john@terraform.leanspace.io"
  policy_ids = var.access_policies
}
