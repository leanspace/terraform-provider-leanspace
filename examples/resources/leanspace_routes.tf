resource "leanspace_routes" "test_create_route" {
  name        = "Terraform Route"
  description = "This is a Route created through terraform!"
  definition {
    configuration = "- route:\r\n    from:\r\n      uri:     'timer:yaml?period=3s'\r\n      steps:\r\n        - set-body:\r\n            simple: 'Timer fired $${header.CamelTimerCounter} times'\r\n        - to:\r\n            uri: 'log:yaml'"
    log_level     = "INFO"
  }
  tags {
    key   = "CreatedBy"
    value = "Terraform"
  }
}
