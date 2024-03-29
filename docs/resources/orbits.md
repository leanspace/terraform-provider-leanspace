---
page_title: "leanspace_orbits Resource - terraform-provider-leanspace"
subcategory: ""
description: |-
  
---

# leanspace_orbits (Resource)



## Example Usage

```terraform
variable "satellite_id" {
  type        = string
  description = "The ID of the satellite to which the orbit will be added."
}

resource "leanspace_orbits" "an_orbit" {
  name         = "Terraform Orbit"
  satellite_id = var.satellite_id
  ideal_orbit {
    type                              = "LEO"
    inclination                       = 97.5
    right_ascension_of_ascending_node = 50.0
    argument_of_perigee               = 0.8
    altitude_in_meters                = 150.0
    eccentricity                      = 0.7
  }
  tags {
    key   = "Mission"
    value = "Terraform"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String)
- `satellite_id` (String)

### Optional

- `ideal_orbit` (Block List, Max: 1) (see [below for nested schema](#nestedblock--ideal_orbit))
- `tags` (Block Set) (see [below for nested schema](#nestedblock--tags))

### Read-Only

- `created_at` (String) When it was created
- `created_by` (String) Who created it
- `id` (String) The ID of this resource.
- `last_modified_at` (String) When it was last modified
- `last_modified_by` (String) Who modified it the last

<a id="nestedblock--ideal_orbit"></a>
### Nested Schema for `ideal_orbit`

Required:

- `altitude_in_meters` (Number)
- `argument_of_perigee` (Number)
- `eccentricity` (Number)
- `inclination` (Number)
- `right_ascension_of_ascending_node` (Number)
- `type` (String) it must be one of these values: SSO, POLAR, LEO, GEO, MEO, OTHER

Read-Only:

- `apogee_altitude_in_meters` (Number)
- `perigee_altitude_in_meters` (Number)
- `semi_major_axis` (Number)


<a id="nestedblock--tags"></a>
### Nested Schema for `tags`

Required:

- `key` (String)

Optional:

- `value` (String)
