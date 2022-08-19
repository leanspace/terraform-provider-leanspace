resource "leanspace_access_policies" "access_policy" {
  access_policy {
    name        = "Terraform Access Policy"
    description = "An access policy made through Terraform, for easy team management."
    statements {
      name    = "Dashboard Full Access"
      actions = ["dashboards:*"]
    }
    statements {
      name = "MonitorsReadAccess"
      actions = [
        "monitors:getMonitor",
        "monitors:searchActionTemplates",
        "monitors:searchMonitors",
        "monitors:searchMonitorsStatesHistory",
        "monitors:searchMonitorTags"
      ]
    }
  }
}
