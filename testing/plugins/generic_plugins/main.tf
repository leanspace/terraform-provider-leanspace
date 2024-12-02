terraform {
  required_providers {
    leanspace = {
      source = "leanspace/leanspace"
    }
  }
}

variable "path" {
  type        = string
  description = "The path to the generic plugin's source code file (a .jar)."
}

data "leanspace_generic_plugins" "all" {
  filters {
    ids = []
    types = [
      "CHECKSUM_FUNCTION",
    ]
    query = ""
    page  = 0
    size  = 10
    sort  = ["name,asc"]
  }
}

resource "leanspace_generic_plugins" "test" {
  source_code_path                     = var.path
  name                                 = "Terraform Generic Plugin"
  description                          = "This is a generic plugin created through terraform!"
  type                                 = "CHECKSUM_FUNCTION"
  language                             = "JAVA"
}

output "all_generic_plugins" {
  value = data.leanspace_generic_plugins.all
}

output "test_generic_plugin" {
  value = leanspace_generic_plugins.test
}
