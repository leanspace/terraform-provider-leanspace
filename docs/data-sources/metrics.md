---
page_title: "leanspace_metrics Data Source - terraform-provider-leanspace"
subcategory: ""
description: |-
  
---

# leanspace_metrics (Data Source)



## Example Usage

```terraform
data "leanspace_metrics" "all" {
  filters {
    node_ids        = var.node_ids
    attribute_types = ["NUMERIC", "TEXT"]
    tags            = []
    ids             = []
    query           = ""
    page            = 0
    size            = 10
    sort            = ["name,asc"]
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

- `attribute_types` (List of String)
- `ids` (List of String)
- `node_ids` (List of String)
- `page` (Number)
- `query` (String)
- `size` (Number)
- `sort` (List of String)
- `tags` (List of String)


<a id="nestedatt--content"></a>
### Nested Schema for `content`

Read-Only:

- `attributes` (List of Object) (see [below for nested schema](#nestedobjatt--content--attributes))
- `created_at` (String)
- `created_by` (String)
- `description` (String)
- `id` (String)
- `last_modified_at` (String)
- `last_modified_by` (String)
- `name` (String)
- `node_id` (String)
- `tags` (Set of Object) (see [below for nested schema](#nestedobjatt--content--tags))

<a id="nestedobjatt--content--attributes"></a>
### Nested Schema for `content.attributes`

Read-Only:

- `after` (String)
- `before` (String)
- `constraint` (List of Object) (see [below for nested schema](#nestedobjatt--content--attributes--constraint))
- `fields` (List of Object) (see [below for nested schema](#nestedobjatt--content--attributes--fields))
- `max` (Number)
- `max_length` (Number)
- `max_size` (Number)
- `min` (Number)
- `min_length` (Number)
- `min_size` (Number)
- `options` (Map of String)
- `pattern` (String)
- `precision` (Number)
- `scale` (Number)
- `type` (String)
- `unique` (Boolean)
- `unit_id` (String)

<a id="nestedobjatt--content--attributes--constraint"></a>
### Nested Schema for `content.attributes.constraint`

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


<a id="nestedobjatt--content--attributes--fields"></a>
### Nested Schema for `content.attributes.fields`

Read-Only:

- `elevation` (List of Object) (see [below for nested schema](#nestedobjatt--content--attributes--fields--elevation))
- `latitude` (List of Object) (see [below for nested schema](#nestedobjatt--content--attributes--fields--latitude))
- `longitude` (List of Object) (see [below for nested schema](#nestedobjatt--content--attributes--fields--longitude))

<a id="nestedobjatt--content--attributes--fields--elevation"></a>
### Nested Schema for `content.attributes.fields.longitude`

Read-Only:

- `default_value` (String)
- `max` (Number)
- `min` (Number)
- `precision` (Number)
- `scale` (Number)
- `unit_id` (String)


<a id="nestedobjatt--content--attributes--fields--latitude"></a>
### Nested Schema for `content.attributes.fields.longitude`

Read-Only:

- `default_value` (String)
- `max` (Number)
- `min` (Number)
- `precision` (Number)
- `scale` (Number)
- `unit_id` (String)


<a id="nestedobjatt--content--attributes--fields--longitude"></a>
### Nested Schema for `content.attributes.fields.longitude`

Read-Only:

- `default_value` (String)
- `max` (Number)
- `min` (Number)
- `precision` (Number)
- `scale` (Number)
- `unit_id` (String)




<a id="nestedobjatt--content--tags"></a>
### Nested Schema for `content.tags`

Read-Only:

- `key` (String)
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
