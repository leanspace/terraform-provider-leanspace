
resource "leanspace_event_criticalities" "event" {
  name = "EVENT_CRITICALITY"
  tags {
    key   = "Mission"
    value = "Terraform"
  }

}