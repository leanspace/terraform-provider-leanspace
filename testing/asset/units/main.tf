terraform {
  required_providers {
    leanspace = {
      source = "leanspace/leanspace"
    }
  }
}

data "leanspace_units" "all" {
  filters {
    ids   = []
    query = ""
    page  = 0
    size  = 10
    sort  = ["symbol,asc"]
  }
}

resource "leanspace_units" "custom_units" {
  for_each = {
    "k" : "Kilo"
    "h" : "Hecto"
    "da" : "Deca"
    "" : ""
    "d" : "Deci"
    "c" : "Centi"
    "m" : "Milli"
  }
  symbol       = "Terraform${each.key}"
  display_name = "Terraform${each.value}"
}

output "test_units" {
  value = leanspace_units.custom_units
}
