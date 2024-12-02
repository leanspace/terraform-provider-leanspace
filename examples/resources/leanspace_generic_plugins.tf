variable "source_code_path" {
  type        = string
  description = "The path to the generic plugin's source code file (a .jar)."
}

resource "leanspace_generic_plugins" "test" {
  name                                 = "Terraform Generic Plugin"
  description                          = "This is a generic plugin created through terraform!"
  type                                 = "CHECKSUM_FUNCTION"
  language                             = "JAVA"
  source_code_path                     = var.source_code_path
}
