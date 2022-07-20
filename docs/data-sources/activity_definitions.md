---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "leanspace_activity_definitions Data Source - terraform-provider-leanspace"
subcategory: ""
description: |-
  
---

# leanspace_activity_definitions (Data Source)





<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `content` (List of Object) (see [below for nested schema](#nestedatt--content))
- `empty` (Boolean)
- `first` (Boolean)
- `id` (String) The ID of this resource.
- `last` (Boolean)
- `number` (Number)
- `number_of_elements` (Number)
- `pageable` (List of Object) (see [below for nested schema](#nestedatt--pageable))
- `size` (Number)
- `sort` (List of Object) (see [below for nested schema](#nestedatt--sort))
- `total_elements` (Number)
- `total_pages` (Number)

<a id="nestedatt--content"></a>
### Nested Schema for `content`

Read-Only:

- `argument_definitions` (Set of Object) (see [below for nested schema](#nestedobjatt--content--argument_definitions))
- `command_mappings` (List of Object) (see [below for nested schema](#nestedobjatt--content--command_mappings))
- `created_at` (String)
- `created_by` (String)
- `description` (String)
- `estimated_duration` (Number)
- `id` (String)
- `last_modified_at` (String)
- `last_modified_by` (String)
- `metadata` (Set of Object) (see [below for nested schema](#nestedobjatt--content--metadata))
- `name` (String)
- `node_id` (String)

<a id="nestedobjatt--content--argument_definitions"></a>
### Nested Schema for `content.argument_definitions`

Read-Only:

- `attributes` (List of Object) (see [below for nested schema](#nestedobjatt--content--argument_definitions--attributes))
- `description` (String)
- `id` (String)
- `name` (String)

<a id="nestedobjatt--content--argument_definitions--attributes"></a>
### Nested Schema for `content.argument_definitions.attributes`

Read-Only:

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
- `type` (String)
- `unit_id` (String)



<a id="nestedobjatt--content--command_mappings"></a>
### Nested Schema for `content.command_mappings`

Read-Only:

- `argument_mappings` (Set of Object) (see [below for nested schema](#nestedobjatt--content--command_mappings--argument_mappings))
- `command_definition_id` (String)
- `delay_in_milliseconds` (Number)
- `metadata_mappings` (Set of Object) (see [below for nested schema](#nestedobjatt--content--command_mappings--metadata_mappings))
- `position` (Number)

<a id="nestedobjatt--content--command_mappings--argument_mappings"></a>
### Nested Schema for `content.command_mappings.argument_mappings`

Read-Only:

- `activity_definition_argument_name` (String)
- `command_definition_argument_name` (String)


<a id="nestedobjatt--content--command_mappings--metadata_mappings"></a>
### Nested Schema for `content.command_mappings.metadata_mappings`

Read-Only:

- `activity_definition_metadata_name` (String)
- `command_definition_argument_name` (String)



<a id="nestedobjatt--content--metadata"></a>
### Nested Schema for `content.metadata`

Read-Only:

- `attributes` (List of Object) (see [below for nested schema](#nestedobjatt--content--metadata--attributes))
- `description` (String)
- `id` (String)
- `name` (String)

<a id="nestedobjatt--content--metadata--attributes"></a>
### Nested Schema for `content.metadata.attributes`

Read-Only:

- `type` (String)
- `unit_id` (String)
- `value` (String)




<a id="nestedatt--pageable"></a>
### Nested Schema for `pageable`

Read-Only:

- `offset` (Number)
- `page_number` (Number)
- `page_size` (Number)
- `paged` (Boolean)
- `sort` (List of Object) (see [below for nested schema](#nestedobjatt--pageable--sort))
- `unpaged` (Boolean)

<a id="nestedobjatt--pageable--sort"></a>
### Nested Schema for `pageable.sort`

Read-Only:

- `ascending` (Boolean)
- `descending` (Boolean)
- `direction` (String)
- `ignore_case` (Boolean)
- `null_handling` (String)
- `property` (String)



<a id="nestedatt--sort"></a>
### Nested Schema for `sort`

Read-Only:

- `ascending` (Boolean)
- `descending` (Boolean)
- `direction` (String)
- `ignore_case` (Boolean)
- `null_handling` (String)
- `property` (String)

