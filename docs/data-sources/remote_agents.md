---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "leanspace_remote_agents Data Source - terraform-provider-leanspace"
subcategory: ""
description: |-
  
---

# leanspace_remote_agents (Data Source)





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

- `connectors` (Set of Object) (see [below for nested schema](#nestedobjatt--content--connectors))
- `created_at` (String)
- `created_by` (String)
- `description` (String)
- `id` (String)
- `last_modified_at` (String)
- `last_modified_by` (String)
- `name` (String)
- `service_account_id` (String)

<a id="nestedobjatt--content--connectors"></a>
### Nested Schema for `content.connectors`

Read-Only:

- `command_queue_id` (String)
- `created_at` (String)
- `created_by` (String)
- `destination` (List of Object) (see [below for nested schema](#nestedobjatt--content--connectors--destination))
- `gateway_id` (String)
- `id` (String)
- `last_modified_at` (String)
- `last_modified_by` (String)
- `socket` (List of Object) (see [below for nested schema](#nestedobjatt--content--connectors--socket))
- `source` (List of Object) (see [below for nested schema](#nestedobjatt--content--connectors--source))
- `stream_id` (String)
- `type` (String)

<a id="nestedobjatt--content--connectors--destination"></a>
### Nested Schema for `content.connectors.destination`

Read-Only:

- `binding` (String)
- `type` (String)


<a id="nestedobjatt--content--connectors--socket"></a>
### Nested Schema for `content.connectors.socket`

Read-Only:

- `host` (String)
- `port` (Number)
- `type` (String)


<a id="nestedobjatt--content--connectors--source"></a>
### Nested Schema for `content.connectors.source`

Read-Only:

- `binding` (String)
- `type` (String)




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

