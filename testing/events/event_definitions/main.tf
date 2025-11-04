terraform {
  required_providers {
    leanspace = {
      source = "leanspace/leanspace"
    }
  }
}

resource "leanspace_events_definitions" "test" {
  name        = "Terraform Event Definition"
  source      = "STREAM_DECODED"
  state       = "ACTIVE"
  description = "A complex event definition, entirely created under terraform."
  # criticality = "TEST"  # Optional: uncomment to set criticality
  rules {
    operator = "EQUAL_TO"
    path     = "test"
    comparison_value {
      type  = "NUMERIC"
      value = "3"

    }
  }
  tags {
    key   = "Mission"
    value = "Terraform"
  }

}