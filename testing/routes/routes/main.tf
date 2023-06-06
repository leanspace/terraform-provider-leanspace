terraform {
  required_providers {
    leanspace = {
      source = "leanspace/leanspace"
    }
  }
}

variable "service_account_id" {
  type        = string
  default     = ""
  description = "The service account id use for creating routes."
}

data "leanspace_routes" "all" {
  filters {
    ids   = []
    tags  = []
    query = ""
    page  = 0
    size  = 10
    sort  = ["name,asc"]
  }
}

resource "leanspace_routes" "test_create_route" {
  name        = "Terraform Route Removed"
  description = "This is a Route created through terraform!"
  definition {
    configuration      = "- route:\r\n    from:\r\n      uri:     'timer:yaml?period=3s'\r\n      steps:\r\n        - set-body:\r\n            simple: 'Timer fired $${header.CamelTimerCounter} times'\r\n        - to:\r\n            uri: 'log:yaml'"
    log_level          = "INFO"
    service_account_id = var.service_account_id == "" ? null : var.service_account_id
  }
  tags {
    key   = "CreatedBy"
    value = "Terraform"
  }
}

output "all_routes" {
  value = data.leanspace_routes.all
}

output "test_create_route" {
  value = leanspace_routes.test_create_route
}
