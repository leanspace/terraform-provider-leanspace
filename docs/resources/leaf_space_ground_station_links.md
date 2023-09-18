---
page_title: "leanspace_leaf_space_ground_station_links Resource - terraform-provider-leanspace"
subcategory: ""
description: |-
  
---

# leanspace_leaf_space_ground_station_links (Resource)



## Example Usage

```terraform
resource "leanspace_leaf_space_ground_station_links" "ground_station_link" {
  leafspace_ground_station_id = "d5de2269dc23179929546f41b6239afb"
  leanspace_ground_station_id = "969e157d-8883-43cd-b851-3d7ff3449ec6"

}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `leafspace_ground_station_id` (String)
- `leanspace_ground_station_id` (String)

### Optional

- `leafspace_ground_station_name` (String)
- `leanspace_ground_station_name` (String)

### Read-Only

- `id` (String) The ID of this resource.