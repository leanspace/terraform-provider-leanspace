
resource "leanspace_events_definitions" "event" {
  name        = "Terraform Event Definition"
  source      = "STREAM_DECODED"
  state       = "ACTIVE"
  description = "A complex event definition, entirely created under terraform."
  # criticality = "HIGH"  # Optional: Define the criticality level
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