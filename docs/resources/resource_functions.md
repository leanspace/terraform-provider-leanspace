---
page_title: "leanspace_resource_functions Resource - terraform-provider-leanspace"
subcategory: ""
description: |-
  
---

# leanspace_resource_functions (Resource)



## Example Usage

```terraform
variable "resource_id" {
  type        = string
  description = "The ID of the resource to which the resource function is attached."
}

variable "activity_definition_id" {
  type        = string
  description = "The ID of the activity definition to which the resource function is attached."
}

resource "leanspace_resource_functions" "a_resource_function" {
  name                   = "Terraform Resource Function"
  resource_id            = var.resource_id
  activity_definition_id = var.activity_definition_id
  time_unit              = "SECONDS"
  formula {
    constant = 5.0
    rate     = 2.5
    type     = "LINEAR"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `activity_definition_id` (String)
- `formula` (Block List, Min: 1, Max: 1) (see [below for nested schema](#nestedblock--formula))
- `resource_id` (String)
- `time_unit` (String) it must be one of these values: SECONDS, MINUTES, HOURS, DAYS

### Optional

- `name` (String)

### Read-Only

- `created_at` (String) When it was created
- `created_by` (String) Who created it
- `id` (String) The ID of this resource.
- `last_modified_at` (String) When it was last modified
- `last_modified_by` (String) Who modified it the last

<a id="nestedblock--formula"></a>
### Nested Schema for `formula`

Required:

- `constant` (Number)
- `rate` (Number)
- `type` (String) it must be one of these values: LINEAR
