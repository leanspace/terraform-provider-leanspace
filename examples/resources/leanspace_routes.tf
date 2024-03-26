variable "processor_ids" {
  type        = list(string)
  default     = []
  description = "The list of processors to attach to the route."
}

resource "leanspace_routes" "test_create_route" {
  name        = "Terraform Route"
  description = "This is a Route created through terraform!"
  definition {
    configuration = trimspace(<<EOT
- route:
    from:
      uri: 'timer:yaml?period=3s'
      steps:
        - set-body:
            simple: 'Timer fired $${header.CamelTimerCounter} times'
        - to:
            uri: 'log:yaml'
EOT
)
    log_level     = "INFO"
  }
  processor_ids = var.processor_ids
  tags {
    key   = "CreatedBy"
    value = "Terraform"
  }
}
