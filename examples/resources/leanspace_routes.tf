resource "leanspace_routes" "test_create_route" {
  name                                 = "Terraform Route"
  description                          = "This is a Route created through terraform!"
  definition {
    configuration   = "- route"
    log_level       = "INFO"
  }
  tags {
    key   = "CreatedBy"
    value = "Terraform"
  }
}