terraform {
  required_providers {
    leanspace = {
      source = "app.terraform.io/leanspace/leanspace"
    }
  }
}

data "leanspace_action_templates" "all" {}

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
