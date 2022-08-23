resource "leanspace_metrics" "metric" {
  name        = "Terra Metric"
  description = "A numeric metric, created under terraform."
  node_id     = var.node_id

  attributes {
    type = "NUMERIC"
  }
}
