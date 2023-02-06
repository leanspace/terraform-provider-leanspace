---
page_title: "leanspace_command_definitions Resource - terraform-provider-leanspace"
subcategory: ""
description: |-
  
---

# leanspace_command_definitions (Resource)



## Example Usage

```terraform
variable "node_id" {
  type        = string
  description = "The ID of the node to which the command definitions will be added."
}

resource "leanspace_command_definitions" "test" {
  name        = "Terraform Command"
  description = "A complex command definition, entirely created under terraform."
  node_id     = var.node_id
  identifier  = "TERRA_CMD"

  metadata {
    name        = "TestMetadataNumeric"
    description = "A numeric metadata value"
    attributes {
      value = 2
      type  = "NUMERIC"
    }
  }
  metadata {
    name        = "TestMetadataText"
    description = "A text metadata value"
    attributes {
      value = "test"
      type  = "TEXT"
    }
  }
  metadata {
    name        = "TestMetadataBool"
    description = "A boolean metadata value"
    attributes {
      value = true
      type  = "BOOLEAN"
    }
  }
  metadata {
    name        = "TestMetadataTimestamp"
    description = "A timestamp metadata value"
    attributes {
      value = "2022-06-30T13:57:23Z"
      type  = "TIMESTAMP"
    }
  }
  metadata {
    name        = "TestMetadataDate"
    description = "A date metadata value"
    attributes {
      value = "2022-06-30"
      type  = "DATE"
    }
  }
  metadata {
    name        = "TestMetadataTime"
    description = "A time metadata value"
    attributes {
      value = "10:37:19"
      type  = "TIME"
    }
  }

  arguments {
    name        = "TestArgumentNumeric"
    identifier  = "NUMERIC"
    description = "A numeric input"
    attributes {
      default_value = 2
      type          = "NUMERIC"
      required      = true
    }
  }
  arguments {
    name        = "TestArgumentText"
    identifier  = "TEXT"
    description = "A text input"
    attributes {
      default_value = "test"
      type          = "TEXT"
    }
  }
  arguments {
    name        = "TestArgumentBool"
    identifier  = "BOOL"
    description = "A boolean input"
    attributes {
      default_value = true
      type          = "BOOLEAN"
      required      = true
    }
  }
  arguments {
    name        = "TestArgumentTimestamp"
    identifier  = "TIMESTAMP"
    description = "A timestamp input"
    attributes {
      default_value = "2022-06-30T13:57:23Z"
      type          = "TIMESTAMP"
      required      = true
    }
  }
  arguments {
    name        = "TestArgumentDate"
    identifier  = "DATE"
    description = "A date input"
    attributes {
      default_value = "2022-06-30"
      type          = "DATE"
      required      = true
    }
  }
  arguments {
    name        = "TestArgumentTime"
    identifier  = "TIME"
    description = "A time input"
    attributes {
      default_value = "10:37:19"
      type          = "TIME"
      required      = true
    }
  }
  arguments {
    name        = "TestArgumentEnum"
    identifier  = "ENUM"
    description = "An enum input"
    attributes {
      default_value = 1
      options       = { 1 = "test" }
      type          = "ENUM"
      required      = true
    }
  }
  arguments {
    name        = "TestArgumentNumericArray"
    identifier  = "Numeric ARRAY"
    description = "A numeric array"
    attributes {
      type          = "ARRAY"
      required      = true
      min_size      = 1
      max_size      = 4
      unique        = true
      default_value = "1,2,3"
      constraint    { 
        type        = "NUMERIC"
        min         = 1
        max         = 10
      }
    }
  }
  arguments {
    name        = "TestArgumentTextArray"
    identifier  = "Text ARRAY"
    description = "A text array"
    attributes {
      type          = "ARRAY"
      required      = true
      min_size      = 1
      max_size      = 4
      unique        = true
      default_value = "value1,value2,value3"
      constraint    { 
        type        = "TEXT"
        min_length  = 5
        max_length  = 10
      }
    }
  }
  arguments {
    name        = "TestArgumentBooleanArray"
    identifier  = "Boolean ARRAY"
    description = "A boolean array"
    attributes {
      type          = "ARRAY"
      required      = true
      min_size      = 1
      max_size      = 4
      unique        = false
      default_value = "true,false,true"
      constraint    { 
        type        = "BOOLEAN"
      }
    }
  }
  arguments {
    name        = "TestArgumentTimeArray"
    identifier  = "Time ARRAY"
    description = "A time array"
    attributes {
      type          = "ARRAY"
      required      = true
      min_size      = 1
      max_size      = 4
      unique        = true
      default_value = "08:37:19,10:37:19,15:37:19"
      constraint    { 
        type        = "TIME"
        before      = "20:00:00"
        after       = "07:00:00"
      }
    }
  }
  arguments {
    name        = "TestArgumentDateArray"
    identifier  = "Date ARRAY"
    description = "A date array"
    attributes {
      type          = "ARRAY"
      required      = true
      min_size      = 1
      max_size      = 4
      unique        = true
      default_value = "2023-03-30,2023-05-11,2023-07-02"
      constraint    { 
        type        = "DATE"
        before      = "2023-08-01"
        after       = "2023-02-01"
      }
    }
  }
  arguments {
    name        = "TestArgumentTimeStampArray"
    identifier  = "TimeStamp ARRAY"
    description = "A timeStamp array"
    attributes {
      type          = "ARRAY"
      required      = true
      min_size      = 1
      max_size      = 4
      unique        = true
      default_value = "2023-01-30T13:00:00Z,2023-01-29T01:00:00Z,2023-01-31T19:57:23Z"
      constraint    { 
        type        = "TIMESTAMP"
        before      = "2023-01-31T20:00:00Z"
        after       = "2023-01-29T00:00:00Z"
      }
    }
  }
    arguments {
    name        = "TestArgumentEnumArray"
    identifier  = "Enum ARRAY"
    description = "A enum array"
    attributes {
      type          = "ARRAY"
      required      = true
      min_size      = 1
      max_size      = 4
      unique        = false
      default_value = "1,2,3,1"
      constraint    { 
        type        = "ENUM"
        options       = { 1 = "value1", 2 = "value2", 3 = "value3" }
      }
    }
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String)
- `node_id` (String)

### Optional

- `arguments` (Block Set) (see [below for nested schema](#nestedblock--arguments))
- `description` (String)
- `identifier` (String)
- `metadata` (Block Set) (see [below for nested schema](#nestedblock--metadata))

### Read-Only

- `created_at` (String) When it was created
- `created_by` (String) Who created it
- `id` (String) The ID of this resource.
- `last_modified_at` (String) When it was last modified
- `last_modified_by` (String) Who modified it the last

<a id="nestedblock--arguments"></a>
### Nested Schema for `arguments`

Required:

- `attributes` (Block List, Min: 1, Max: 1) (see [below for nested schema](#nestedblock--arguments--attributes))
- `identifier` (String)
- `name` (String)

Optional:

- `description` (String)

Read-Only:

- `id` (String) The ID of this resource.

<a id="nestedblock--arguments--attributes"></a>
### Nested Schema for `arguments.attributes`

Required:

- `type` (String) it must be one of these values: NUMERIC, BOOLEAN, TEXT, DATE, TIME, TIMESTAMP, ENUM, BINARY, ARRAY

Optional:

- `after` (String) Time/date/timestamp only: Minimum date allowed
- `before` (String) Time/date/timestamp only: Maximum date allowed
- `constraint` (Block List, Max: 1) Array only: Constraint applied to all elements in the array (see [below for nested schema](#nestedblock--arguments--attributes--constraint))
- `default_value` (String) The default value can be of any type. In case of an array type, please surround the list values with double quotes and use the comma separator.
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

<a id="nestedblock--arguments--attributes--constraint"></a>
### Nested Schema for `arguments.attributes.constraint`

Required:

- `type` (String) it must be one of these values: NUMERIC, BOOLEAN, TEXT, DATE, TIME, TIMESTAMP, ENUM

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




<a id="nestedblock--metadata"></a>
### Nested Schema for `metadata`

Required:

- `attributes` (Block List, Min: 1, Max: 1) (see [below for nested schema](#nestedblock--metadata--attributes))
- `name` (String)

Optional:

- `description` (String)

Read-Only:

- `id` (String) The ID of this resource.

<a id="nestedblock--metadata--attributes"></a>
### Nested Schema for `metadata.attributes`

Required:

- `type` (String) it must be one of these values: NUMERIC, BOOLEAN, TEXT, DATE, TIME, TIMESTAMP, ENUM

Optional:

- `unit_id` (String)
- `value` (String)
