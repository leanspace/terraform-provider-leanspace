resource "leanspace_monitors" "monitor" {
  monitor {
    name                         = "Terraform Monitor"
    description                  = "A monitor created throug terraform."
    polling_frequency_in_minutes = 60
    metric_id                    = var.metric_id
    expression {
      comparison_operator  = "GREATER_THAN"
      comparison_value     = 200
      aggregation_function = "HIGHEST_VALUE"
    }
  }
}
