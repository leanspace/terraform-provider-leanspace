variable "processor_path" {
  type        = string
  description = "The path to the processor's source code file (a .jar)."
}

resource "leanspace_processors" "test_create_processor" {
  file_path   = var.processor_path
  version     = "1.0"
  name        = "Terraform Processor"
  description = "This is a processor created through terraform!"
}
