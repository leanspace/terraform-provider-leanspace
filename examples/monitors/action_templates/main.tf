terraform {
  required_providers {
    leanspace = {
      source = "leanspace/leanspace"
    }
  }
}

data "leanspace_action_templates" "all" {
  filters {
    types       = ["WEBHOOK"]
    monitor_ids = []
    ids         = []
    query       = ""
    page        = 0
    size        = 10
    sort        = ["name,asc"]
  }
}

resource "leanspace_action_templates" "test" {
  name = "Terraform Action Template"
  type = "WEBHOOK"
  url  = "https://my-custom-webhook.com"
  payload = jsonencode({
    pi   = 3.1415
    data = "This is my action template."
  })
  headers = {
    "Authorization" : "Bearer token",
    "Content-Type" : "application/json",
  }
}

output "test_action_template" {
  value = leanspace_action_templates.test
}
