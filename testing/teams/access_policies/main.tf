terraform {
  required_providers {
    leanspace = {
      source = "leanspace/leanspace"
    }
  }
}

data "leanspace_access_policies" "all" {
  filters {
    action_ids   = []
    action_names = ["updateCommandTransmission"]
    ids          = []
    query        = ""
    page         = 0
    size         = 10
    sort         = ["name,asc"]
  }
}

resource "leanspace_access_policies" "test" {
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

output "test_access_policy" {
  value = leanspace_access_policies.test
}
