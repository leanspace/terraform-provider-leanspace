terraform {
  required_providers {
    leanspace = {
      source = "leanspace/leanspace"
      version = "0.4.0"
    }
  }
}

data "leanspace_plan_templates" "all" {

}

resource "leanspace_plan_templates" "created" {
  name = "MY_TEST"
}

output "created" {
  value = leanspace_plan_templates.created
}
