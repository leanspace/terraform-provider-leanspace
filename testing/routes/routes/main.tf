terraform {
  required_providers {
    leanspace = {
      source = "leanspace/leanspace"
    }
  }
}

data "leanspace_routes" "all" {
  filters {
    ids             = []
    tags            = []
    query           = ""
    page            = 0
    size            = 10
    sort            = ["name,asc"]
  }
}

resource "leanspace_routes" "test_create_route" {
  name                                 = "Terraform Route"
  description                          = "This is a Route created through terraform!"
  definition {
    configuration   = "- route"
    log_level       = "INFO"
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
