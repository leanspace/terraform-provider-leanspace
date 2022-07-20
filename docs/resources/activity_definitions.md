---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "leanspace_activity_definitions Resource - terraform-provider-leanspace"
subcategory: ""
description: |-
  
---

# leanspace_activity_definitions (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `activity_definition` (Block List, Min: 1, Max: 1) (see [below for nested schema](#nestedblock--activity_definition))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--activity_definition"></a>
### Nested Schema for `activity_definition`

Required:

- `name` (String)
- `node_id` (String)

Optional:

- `argument_definitions` (Block Set) (see [below for nested schema](#nestedblock--activity_definition--argument_definitions))
- `command_mappings` (Block List) (see [below for nested schema](#nestedblock--activity_definition--command_mappings))
- `description` (String)
- `estimated_duration` (Number)
- `metadata` (Block Set) (see [below for nested schema](#nestedblock--activity_definition--metadata))

Read-Only:

- `created_at` (String)
- `created_by` (String)
- `id` (String) The ID of this resource.
- `last_modified_at` (String)
- `last_modified_by` (String)

<a id="nestedblock--activity_definition--argument_definitions"></a>
### Nested Schema for `activity_definition.argument_definitions`

Required:

- `attributes` (Block List, Min: 1, Max: 1) (see [below for nested schema](#nestedblock--activity_definition--argument_definitions--attributes))
- `name` (String)

Optional:

- `description` (String)

Read-Only:

- `id` (String) The ID of this resource.

<a id="nestedblock--activity_definition--argument_definitions--attributes"></a>
### Nested Schema for `activity_definition.argument_definitions.attributes`

Required:

- `type` (String)

Optional:

- `after` (String)
- `before` (String)
- `default_value` (String)
- `max` (Number)
- `max_length` (Number)
- `min` (Number)
- `min_length` (Number)
- `options` (Map of String)
- `pattern` (String)
- `precision` (Number)
- `required` (Boolean)
- `scale` (Number)
- `unit_id` (String)



<a id="nestedblock--activity_definition--command_mappings"></a>
### Nested Schema for `activity_definition.command_mappings`

Required:

- `command_definition_id` (String)
- `delay_in_milliseconds` (Number)

Optional:

- `argument_mappings` (Block Set) (see [below for nested schema](#nestedblock--activity_definition--command_mappings--argument_mappings))
- `metadata_mappings` (Block Set) (see [below for nested schema](#nestedblock--activity_definition--command_mappings--metadata_mappings))

Read-Only:

- `position` (Number)

<a id="nestedblock--activity_definition--command_mappings--argument_mappings"></a>
### Nested Schema for `activity_definition.command_mappings.argument_mappings`

Required:

- `activity_definition_argument_name` (String)
- `command_definition_argument_name` (String)


<a id="nestedblock--activity_definition--command_mappings--metadata_mappings"></a>
### Nested Schema for `activity_definition.command_mappings.metadata_mappings`

Required:

- `activity_definition_metadata_name` (String)
- `command_definition_argument_name` (String)



<a id="nestedblock--activity_definition--metadata"></a>
### Nested Schema for `activity_definition.metadata`

Required:

- `attributes` (Block List, Min: 1, Max: 1) (see [below for nested schema](#nestedblock--activity_definition--metadata--attributes))
- `name` (String)

Optional:

- `description` (String)

Read-Only:

- `id` (String) The ID of this resource.

<a id="nestedblock--activity_definition--metadata--attributes"></a>
### Nested Schema for `activity_definition.metadata.attributes`

Required:

- `type` (String)

Optional:

- `unit_id` (String)
- `value` (String)

