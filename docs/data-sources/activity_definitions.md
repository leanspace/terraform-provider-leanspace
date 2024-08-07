---
page_title: "leanspace_activity_definitions Data Source - terraform-provider-leanspace"
subcategory: ""
description: |-
  
---

# leanspace_activity_definitions (Data Source)



## Example Usage

```terraform
data "leanspace_activity_definitions" "all" {
  filters {
    node_ids = [var.node_id]
    ids      = []
    query    = ""
    page     = 0
    size     = 10
    sort     = ["name,asc"]
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `filters` (Block List, Max: 1) (see [below for nested schema](#nestedblock--filters))

### Read-Only

- `content` (List of Object) (see [below for nested schema](#nestedatt--content))
- `empty` (Boolean) True if the content is empty
- `first` (Boolean) True if this is the first page
- `id` (String) The ID of this resource.
- `last` (Boolean) True if this is the last page
- `number` (Number) Page number
- `number_of_elements` (Number) Number of elements fetched in this page
- `pageable` (List of Object) (see [below for nested schema](#nestedatt--pageable))
- `size` (Number) Size of this page
- `sort` (List of Object) (see [below for nested schema](#nestedatt--sort))
- `total_elements` (Number) Number of elements in total
- `total_pages` (Number) Number of pages in total

<a id="nestedblock--filters"></a>
### Nested Schema for `filters`

Optional:

- `ids` (List of String)
- `node_ids` (List of String)
- `page` (Number)
- `query` (String)
- `size` (Number)
- `sort` (List of String)


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
- `mapping_status` (String)
- `metadata` (Set of Object) (see [below for nested schema](#nestedobjatt--content--metadata))
- `name` (String)
- `node_id` (String)

<a id="nestedobjatt--content--argument_definitions"></a>
### Nested Schema for `content.argument_definitions`

Read-Only:

- `attributes` (List of Object) (see [below for nested schema](#nestedobjatt--content--argument_definitions--attributes))
- `description` (String)
- `name` (String)

<a id="nestedobjatt--content--argument_definitions--attributes"></a>
### Nested Schema for `content.argument_definitions.attributes`

Read-Only:

- `after` (String)
- `before` (String)
- `constraint` (List of Object) (see [below for nested schema](#nestedobjatt--content--argument_definitions--attributes--constraint))
- `default_value` (String)
- `fields` (List of Object) (see [below for nested schema](#nestedobjatt--content--argument_definitions--attributes--fields))
- `max` (Number)
- `max_length` (Number)
- `max_size` (Number)
- `min` (Number)
- `min_length` (Number)
- `min_size` (Number)
- `options` (Map of String)
- `pattern` (String)
- `precision` (Number)
- `required` (Boolean)
- `scale` (Number)
- `type` (String)
- `unique` (Boolean)
- `unit_id` (String)

<a id="nestedobjatt--content--argument_definitions--attributes--constraint"></a>
### Nested Schema for `content.argument_definitions.attributes.unit_id`

Read-Only:

- `after` (String)
- `before` (String)
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


<a id="nestedobjatt--content--argument_definitions--attributes--fields"></a>
### Nested Schema for `content.argument_definitions.attributes.unit_id`

Read-Only:

- `elevation` (List of Object) (see [below for nested schema](#nestedobjatt--content--argument_definitions--attributes--unit_id--elevation))
- `latitude` (List of Object) (see [below for nested schema](#nestedobjatt--content--argument_definitions--attributes--unit_id--latitude))
- `longitude` (List of Object) (see [below for nested schema](#nestedobjatt--content--argument_definitions--attributes--unit_id--longitude))

<a id="nestedobjatt--content--argument_definitions--attributes--unit_id--elevation"></a>
### Nested Schema for `content.argument_definitions.attributes.unit_id.elevation`

Read-Only:

- `default_value` (String)
- `max` (Number)
- `min` (Number)
- `precision` (Number)
- `scale` (Number)
- `unit_id` (String)


<a id="nestedobjatt--content--argument_definitions--attributes--unit_id--latitude"></a>
### Nested Schema for `content.argument_definitions.attributes.unit_id.latitude`

Read-Only:

- `default_value` (String)
- `max` (Number)
- `min` (Number)
- `precision` (Number)
- `scale` (Number)
- `unit_id` (String)


<a id="nestedobjatt--content--argument_definitions--attributes--unit_id--longitude"></a>
### Nested Schema for `content.argument_definitions.attributes.unit_id.longitude`

Read-Only:

- `default_value` (String)
- `max` (Number)
- `min` (Number)
- `precision` (Number)
- `scale` (Number)
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
- `mapping_status` (String)


<a id="nestedobjatt--content--command_mappings--metadata_mappings"></a>
### Nested Schema for `content.command_mappings.metadata_mappings`

Read-Only:

- `activity_definition_metadata_name` (String)
- `command_definition_argument_name` (String)
- `mapping_status` (String)



<a id="nestedobjatt--content--metadata"></a>
### Nested Schema for `content.metadata`

Read-Only:

- `attributes` (List of Object) (see [below for nested schema](#nestedobjatt--content--metadata--attributes))
- `description` (String)
- `name` (String)

<a id="nestedobjatt--content--metadata--attributes"></a>
### Nested Schema for `content.metadata.attributes`

Read-Only:

- `data_type` (String)
- `fields` (List of Object) (see [below for nested schema](#nestedobjatt--content--metadata--attributes--fields))
- `type` (String)
- `unit_id` (String)
- `value` (String)

<a id="nestedobjatt--content--metadata--attributes--fields"></a>
### Nested Schema for `content.metadata.attributes.value`

Read-Only:

- `elevation` (List of Object) (see [below for nested schema](#nestedobjatt--content--metadata--attributes--value--elevation))
- `latitude` (List of Object) (see [below for nested schema](#nestedobjatt--content--metadata--attributes--value--latitude))
- `longitude` (List of Object) (see [below for nested schema](#nestedobjatt--content--metadata--attributes--value--longitude))

<a id="nestedobjatt--content--metadata--attributes--value--elevation"></a>
### Nested Schema for `content.metadata.attributes.value.elevation`

Read-Only:

- `max` (Number)
- `min` (Number)
- `precision` (Number)
- `scale` (Number)
- `unit_id` (String)
- `value` (String)


<a id="nestedobjatt--content--metadata--attributes--value--latitude"></a>
### Nested Schema for `content.metadata.attributes.value.latitude`

Read-Only:

- `max` (Number)
- `min` (Number)
- `precision` (Number)
- `scale` (Number)
- `unit_id` (String)
- `value` (String)


<a id="nestedobjatt--content--metadata--attributes--value--longitude"></a>
### Nested Schema for `content.metadata.attributes.value.longitude`

Read-Only:

- `max` (Number)
- `min` (Number)
- `precision` (Number)
- `scale` (Number)
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
