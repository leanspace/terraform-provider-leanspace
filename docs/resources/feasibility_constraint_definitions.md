---
page_title: "leanspace_feasibility_constraint_definitions Resource - terraform-provider-leanspace"
subcategory: ""
description: |-
  
---

# leanspace_feasibility_constraint_definitions (Resource)



## Example Usage

```terraform
resource "leanspace_feasibility_constraint_definitions" "a_feasibility_constraint_definition" {
  name        = "feasibilityConstraintDefinitionFromTerraform"
  description = "feasibilityConstraintDefinitionTerraformDescription"
  argument_definitions {
    name        = "NumericArgumentDefinition"
    description = "A numeric input"
    attributes {
      default_value = 2
      type          = "NUMERIC"
      required      = true
    }
  }
  argument_definitions {
    name        = "TimeArgumentDefinition"
    description = "A time input"
    attributes {
      default_value = "10:37:19"
      type          = "TIME"
      required      = true
    }
  }
  argument_definitions {
    name        = "GeoPointArgumentDefinition"
    description = "A geopoint input"
    attributes {
      type = "GEOPOINT"
      fields {
        elevation {
          default_value = 141.0
        }
        latitude {
          default_value = 48.5
        }
        longitude {
          default_value = 7.7
        }
      }
      required = true
    }
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String)

### Optional

- `argument_definitions` (Block Set, Max: 499) (see [below for nested schema](#nestedblock--argument_definitions))
- `description` (String)

### Read-Only

- `created_at` (String)
- `created_by` (String)
- `id` (String) The ID of this resource.
- `last_modified_at` (String)
- `last_modified_by` (String)

<a id="nestedblock--argument_definitions"></a>
### Nested Schema for `argument_definitions`

Required:

- `attributes` (Block List, Min: 1, Max: 1) (see [below for nested schema](#nestedblock--argument_definitions--attributes))
- `name` (String)

Optional:

- `description` (String)

<a id="nestedblock--argument_definitions--attributes"></a>
### Nested Schema for `argument_definitions.attributes`

Required:

- `type` (String) it must be one of these values: NUMERIC, TIME, GEOPOINT

Optional:

- `after` (String) Time/date/timestamp only: Minimum date allowed
- `before` (String) Time/date/timestamp only: Maximum date allowed
- `constraint` (Block List, Max: 1) Array only: Constraint applied to all elements in the array (see [below for nested schema](#nestedblock--argument_definitions--attributes--constraint))
- `default_value` (String) The default value can be of any type. In case of an array type, please surround the list values with double quotes and use the comma separator.
- `fields` (Block List, Max: 1) Geopoint only (see [below for nested schema](#nestedblock--argument_definitions--attributes--fields))
- `max` (Number) Numeric only
- `max_length` (Number) Text only: Maximum length of this text (at least 1)
- `max_size` (Number) Array only: The maximum number of elements allowed
- `min` (Number) Numeric only
- `min_length` (Number) Text only: Minimum length of this text (at least 1)
- `min_size` (Number) Array only: The minimum number of elements allowed
- `options` (Map of String) Enum only: The allowed values for the enum in the format 1 = "value"
- `pattern` (String) Text only: Regex defined the allowed pattern of this text
- `precision` (Number) Numeric only: How many values after the comma should be accepted
- `required` (Boolean)
- `scale` (Number) Numeric only
- `unique` (Boolean) Array only: No duplicated elements are allowed
- `unit_id` (String) Numeric only

<a id="nestedblock--argument_definitions--attributes--constraint"></a>
### Nested Schema for `argument_definitions.attributes.constraint`

Required:

- `type` (String) it must be one of these values: NUMERIC, BOOLEAN, TEXT, DATE, TIME, TIMESTAMP, ENUM, BINARY

Optional:

- `after` (String) Only array elements with time/date/timestamp type : Minimum date allowed
- `before` (String) Only array elements with time/date/timestamp type : Maximum date allowed
- `max` (Number) Only array elements with numeric type : maximum value allowed
- `max_length` (Number) Only array elements with text type: Maximum length of this text (at least 1)
- `min` (Number) Only array elements with numeric type : minimum value allowed
- `min_length` (Number) Only array elements with text type: Minimum length of this text (at least 1)
- `options` (Map of String) Only array elements with enum type : The allowed values for the enum in the format 1 = "value"
- `pattern` (String) Only array elements with text type: Regex defined the allowed pattern of this text
- `precision` (Number) Only array elements with numeric type : how many values after the comma should be accepted
- `required` (Boolean)
- `scale` (Number) Only array elements with numeric type
- `unit_id` (String) Only array elements with numeric type


<a id="nestedblock--argument_definitions--attributes--fields"></a>
### Nested Schema for `argument_definitions.attributes.fields`

Required:

- `elevation` (Block List, Min: 1, Max: 1) (see [below for nested schema](#nestedblock--argument_definitions--attributes--fields--elevation))
- `latitude` (Block List, Min: 1, Max: 1) (see [below for nested schema](#nestedblock--argument_definitions--attributes--fields--latitude))
- `longitude` (Block List, Min: 1, Max: 1) (see [below for nested schema](#nestedblock--argument_definitions--attributes--fields--longitude))

<a id="nestedblock--argument_definitions--attributes--fields--elevation"></a>
### Nested Schema for `argument_definitions.attributes.fields.elevation`

Optional:

- `default_value` (String)
- `max` (Number) Property field with numeric type only: the maximum value allowed.
- `min` (Number) Property field with numeric type only: the minimum value allowed.
- `precision` (Number) Property field with numeric type only: How many values after the comma should be accepted
- `scale` (Number) Property field with numeric type only: the scale required.
- `unit_id` (String) Property field with numeric type only


<a id="nestedblock--argument_definitions--attributes--fields--latitude"></a>
### Nested Schema for `argument_definitions.attributes.fields.latitude`

Optional:

- `default_value` (String)
- `max` (Number) Property field with numeric type only: the maximum value allowed.
- `min` (Number) Property field with numeric type only: the minimum value allowed.
- `precision` (Number) Property field with numeric type only: How many values after the comma should be accepted
- `scale` (Number) Property field with numeric type only: the scale required.
- `unit_id` (String) Property field with numeric type only


<a id="nestedblock--argument_definitions--attributes--fields--longitude"></a>
### Nested Schema for `argument_definitions.attributes.fields.longitude`

Optional:

- `default_value` (String)
- `max` (Number) Property field with numeric type only: the maximum value allowed.
- `min` (Number) Property field with numeric type only: the minimum value allowed.
- `precision` (Number) Property field with numeric type only: How many values after the comma should be accepted
- `scale` (Number) Property field with numeric type only: the scale required.
- `unit_id` (String) Property field with numeric type only