terraform {
  required_providers {
    leanspace = {
      source  = "leanspace/leanspace"
    }
  }
}

variable "record_template_id" {
  type        = string
  description = "The ID of the Record Template to which the resource will be linked."
}

variable "start_date_time" {
  type        = string
  description = "The start date of the Record, at ISO-8601 format."
}

data "leanspace_records" "all" {
  filters {
    ids                 = []
    record_template_ids = [var.record_template_id]
    names               = []
    query               = ""
    tags                = []
    page                = 0
    size                = 10
    sort                = ["name,asc"]
  }
}

resource "leanspace_records" "a_record" {
  name               = "Terraform Record"
  record_template_id = var.record_template_id
  start_date_time    = var.start_date_time
  properties {
    name = "TestPropertyNumeric"
    attributes {
      type  = "NUMERIC"
      value = 3
    }
  }
}

output "a_records" {
  value = leanspace_records.a_record
}
