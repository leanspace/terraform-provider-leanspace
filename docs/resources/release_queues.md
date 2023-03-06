---
page_title: "leanspace_release_queues Resource - terraform-provider-leanspace"
subcategory: ""
description: |-
  
---

# leanspace_release_queues (Resource)



## Example Usage

```terraform
variable "asset_id" {
  type        = string
  description = "The ID of the asset to which the release queue will be added."
}

resource "leanspace_release_queues" "test" {
  name                            = "Terraform Release Queue"
  asset_id                        = var.asset_id
  command_transformation_strategy = "TEST"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `asset_id` (String)
- `name` (String)

### Optional

- `command_transformation_strategy` (String) What transformation strategy shall be applied on created and updated Commands
- `command_transformer_plugin_configuration_data` (String) Configuration data used by the Command Transformer Plugin (coming soon)
- `command_transformer_plugin_id` (String) The Id of the Command Transformer Plugin
- `description` (String)
- `global_transmission_metadata` (Block Set) (see [below for nested schema](#nestedblock--global_transmission_metadata))

### Read-Only

- `created_at` (String) When it was created
- `created_by` (String) Who created it
- `id` (String) The ID of this resource.
- `last_modified_at` (String) When it was last modified
- `last_modified_by` (String) Who modified it the last
- `logical_lock` (Boolean)

<a id="nestedblock--global_transmission_metadata"></a>
### Nested Schema for `global_transmission_metadata`

Required:

- `key` (String)

Optional:

- `value` (String)