terraform {
  required_providers {
    leanspace = {
      source = "leanspace/leanspace"
    }
  }
}

variable "service_account_id" {
  type        = string
  description = "The service account id use for creating routes."
}

variable "processor_ids" {
  type        = list(string)
  description = "The list of processors to attach to the route."
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
  name        = "Terraform Route"
  description = "This is a Route created through terraform!"
  definition {
    configuration = trimspace(<<EOT
- route:
    from:
      uri: 'timer:yaml?period=3s'
      steps:
        - set-body:
            simple: 'Timer fired $${header.CamelTimerCounter} times'
        - to:
            uri: 'log:yaml'
EOT
    )
    log_level          = "INFO"
    service_account_id = var.service_account_id
  }
  processor_ids = var.processor_ids
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
