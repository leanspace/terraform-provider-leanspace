---
page_title: "leanspace_pass_states Resource - terraform-provider-leanspace"
subcategory: ""
description: |-
  
---

# leanspace_pass_states (Resource)



## Example Usage

```terraform
resource "leanspace_pass_states" "state" {
  name = "TERRAFORM_STATE"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String)

### Read-Only

- `created_at` (String) When it was created
- `created_by` (String) Who created it
- `id` (String) The ID of this resource.
- `last_modified_at` (String) When it was last modified
- `last_modified_by` (String) Who modified it the last
- `read_only` (Boolean)