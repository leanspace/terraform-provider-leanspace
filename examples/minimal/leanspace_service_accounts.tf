resource "leanspace_service_accounts" "service_account" {
  service_account {
    name       = "My Service Account"
    policy_ids = var.access_policies
  }
}
