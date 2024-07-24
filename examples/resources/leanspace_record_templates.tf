variable "rt_node_id" {
  type        = string
  description = "The ID of a Node to which the Record Template will be linked."
}

variable "rt_metric_id" {
  type        = string
  description = "The ID of a Metric to which the Record Template will be linked."
}

resource "leanspace_record_templates" "record_template" {
  name       = "TERRAFORM_RECORD_TEMPLATE"
  node_ids   = [var.rt_node_id]
  metric_ids = [var.rt_metric_id]
}
