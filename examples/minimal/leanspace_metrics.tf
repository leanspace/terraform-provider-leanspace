resource "leanspace_metrics" "metric" {
  metric {
    name        = "Terra Metric"
    description = "A numeric metric, created under terraform."
    node_id     = var.node_id

    attributes {
      type = "NUMERIC"
    }
  }
}
