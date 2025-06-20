resource "leanspace_action_templates" "action_template" {
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

resource "leanspace_action_templates" "leanspace_event" {
  name = "Terraform Action Template for LEANSPACE_EVENT"
  type = "LEANSPACE_EVENT"
}