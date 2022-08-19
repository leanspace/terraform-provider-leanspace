resource "leanspace_analysis_definitions" "test" {
  analysis_definition {
    name        = "Terraform Analysis Definition"
    description = "An analysis definition made under terraform!"
    model_id    = var.model_id
    node_id     = var.node_id

    inputs {
      type = "STRUCTURE"
      fields {
        name = "eclipses"
        type = "STRUCTURE"
        fields {
          name   = "earthEclipseEnabled"
          type   = "BOOLEAN"
          source = "STATIC"
          value  = false
        }
        fields {
          name   = "moonEclipseEnabled"
          type   = "BOOLEAN"
          source = "STATIC"
          value  = true
        }
      }
    }
  }
}
