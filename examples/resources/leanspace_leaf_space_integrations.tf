
resource "leanspace_leaf_space_integrations" "integration_Connection" {
  name       = "Leafspace Connection Terraform"
  username   = "username"
  password   = "password"
  domain_url = "apiv2.sandbox.leaf.space"
}
