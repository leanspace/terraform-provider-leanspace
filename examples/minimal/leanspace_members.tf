resource "leanspace_members" "test" {
  member {
    name       = "John Doe"
    email      = "john@terraform.leanspace.io"
    policy_ids = var.access_policies
  }
}
