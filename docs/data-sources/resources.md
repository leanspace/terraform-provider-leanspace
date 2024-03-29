---
page_title: "leanspace_resources Data Source - terraform-provider-leanspace"
subcategory: ""
description: |-
  
---

# leanspace_resources (Data Source)



## Example Usage

```terraform
data "leanspace_resources" "all" {
  filters {
    asset_ids    = [var.asset_id]
    ids          = []
    data_sources = []
    tags         = []
    query        = ""
    page         = 0
    size         = 10
    sort         = ["name,asc"]
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

- `asset_ids` (List of String)
- `created_bys` (List of String) Filter on the user who created the Resource. If you have no wish to use this field as a filter, either provide a null value or remove the field.
- `from_created_at` (String) Filter on the Resource creation date. Resources with a creation date greater or equals than the filter value will be selected (if they are not excluded by other filters). If you have no wish to use this field as a filter, either provide a null value or remove the field.
- `from_last_modified_at` (String) Filter on the Resource last modification date. Resources with a last modification date greater or equals than the filter value will be selected (if they are not excluded by other filters). If you have no wish to use this field as a filter, either provide a null value or remove the field.
- `ids` (List of String)
- `last_modified_bys` (List of String) Filter on the user who last modified the Resource. If you have no wish to use this field as a filter, either provide a null value or remove the field.
- `metric_ids` (List of String)
- `page` (Number)
- `query` (String)
- `size` (Number)
- `sort` (List of String)
- `tags` (List of String)
- `to_created_at` (String) Filter on the Resource creation date. Resources with a creation date lower or equals than the filter value will be selected (if they are not excluded by other filters). If you have no wish to use this field as a filter, either provide a null value or remove the field.
- `to_last_modified_at` (String) Filter on the Resource last modification date. Resources with a last modification date lower or equals than the filter value will be selected (if they are not excluded by other filters). If you have no wish to use this field as a filter, either provide a null value or remove the field.
- `unit_ids` (List of String)


<a id="nestedatt--content"></a>
### Nested Schema for `content`

Read-Only:

- `asset_id` (String)
- `constraints` (Set of Object) (see [below for nested schema](#nestedobjatt--content--constraints))
- `created_at` (String)
- `created_by` (String)
- `default_level` (Number)
- `description` (String)
- `id` (String)
- `last_modified_at` (String)
- `last_modified_by` (String)
- `metric_id` (String)
- `name` (String)
- `tags` (Set of Object) (see [below for nested schema](#nestedobjatt--content--tags))
- `unit_id` (String)

<a id="nestedobjatt--content--constraints"></a>
### Nested Schema for `content.constraints`

Read-Only:

- `kind` (String)
- `type` (String)
- `value` (Number)


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
