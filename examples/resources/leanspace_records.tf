variable "record_template_id" {
  type        = string
  description = "The ID of the Record Template to which the Record will be added."
}

resource "leanspace_records" "record" {
  name = "TERRAFORM_RECORD"
  record_template_id = var.record_template_id
}
