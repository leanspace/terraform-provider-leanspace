resource "leanspace_service_accounts" "service_account" {
  security_tags = ["terraform_mission"]
  name          = "My Service Account"
  policy_ids    = var.access_policies
}
