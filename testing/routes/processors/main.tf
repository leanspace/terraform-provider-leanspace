terraform {
  required_providers {
    leanspace = {
      source = "leanspace/leanspace"
    }
  }
}

variable "path" {
  type        = string
  description = "The path to the processor's source code file (a .jar)."
}

data "leanspace_processors" "all" {
  filters {
    ids = []
    query = ""
    page  = 0
    size  = 10
    sort  = ["name,asc"]
  }
}

resource "leanspace_processors" "test_create_processor" {
  file_path                            = var.path
  version                              = "1.0"
  name                                 = "Terraform Processor"
  description                          = "This is a processor created through terraform!"
}

output "all_processors" {
  value = data.leanspace_processors.all
}

output "test_create_processor" {
  value = leanspace_processors.test_create_processor
}