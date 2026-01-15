---
page_title: "leanspace_widgets Resource - terraform-provider-leanspace"
subcategory: ""
description: |-
  
---

# leanspace_widgets (Resource)



## Example Usage

```terraform
variable "text_metric_id" {
  type        = string
  description = "The ID of the text metric to create widgets for."
}

variable "numeric_metric_id" {
  type        = string
  description = "The ID of the numeric metric to create widgets for."
}

variable "enum_metric_id" {
  type        = string
  description = "The ID of the enum metric to create widgets for."
}

variable "resource_id" {
  type        = string
  description = "The ID of the text resources to create widgets for."
}

variable "topology_id" {
  type        = string
  description = "The ID of the asset to create widgets for."
}

resource "leanspace_widgets" "test_table" {
  name        = "Terraform Table Widget"
  description = "A table widget created with Terraform"
  type        = "TABLE"
  granularity = "raw"
  query_time_dimension = "received_at"
  display_time_dimension = "timestamp"
  series {
    id          = var.text_metric_id
    name        = "Text Metric Series"
    datasource  = "metric"
    aggregation = "none"
  }
  tags {
    key   = "Mission"
    value = "Terraform"
  }
}

resource "leanspace_widgets" "test_line" {
  name        = "Terraform Line Widget"
  description = "A line widget created with Terraform"
  type        = "LINE"
  granularity = "second"
  query_time_dimension = "timestamp"
  display_time_dimension = "timestamp"
  series {
    id          = var.numeric_metric_id
    name        = "Numeric Metric Series"
    datasource  = "metric"
    aggregation = "avg"
    filters {
      filter_by = var.numeric_metric_id
      operator  = "gt"
      value     = 3
    }
  }
  metadata {
    y_axis_range_max = [100]
    y_axis_label     = "This is a label"
  }
  tags {
    key   = "Mission"
    value = "Terraform"
  }
}

resource "leanspace_widgets" "test_enum" {
  name        = "Terraform Enum Widget"
  description = "An enum widget created with Terraform"
  type        = "ENUM"
  granularity = "second"
  query_time_dimension = "received_at"
  display_time_dimension = "ingested_at"
  series {
    id          = var.enum_metric_id
    name        = "Enum Metric Series"
    datasource  = "metric"
    aggregation = "avg"
    filters {
      filter_by = var.enum_metric_id
      operator  = "gt"
      value     = 3
    }
  }
  metadata {
    y_axis_range_max = [100]
    y_axis_label     = "This is a label"
  }
  tags {
    key   = "Mission"
    value = "Terraform"
  }
}

resource "leanspace_widgets" "test_earth" {
  name        = "Terraform Earth Widget"
  description = "An earth widget created with Terraform"
  type        = "EARTH"
  granularity = "second"
  series {
    id          = var.topology_id
    datasource  = "topology"
    aggregation = "count"
  }
  metadata {
    y_axis_range_max = [100]
    y_axis_label     = "This is a label"
  }
  tags {
    key   = "Mission"
    value = "Terraform"
  }
}

resource "leanspace_widgets" "test_gauge" {
  name        = "Terraform Gauge Widget"
  description = "A gauge widget created with Terraform"
  type        = "GAUGE"
  granularity = "second"
  query_time_dimension = "ingested_at"
  display_time_dimension = "ingested_at"
  series {
    id          = var.numeric_metric_id
    name        = "Numeric Metric Series"
    datasource  = "metric"
    aggregation = "avg"
  }
  metadata {
    y_axis_range_max = [100]
    y_axis_label     = "This is a label"
    thresholds {
      to    = 49
      color = "#52C31A"
    }
    thresholds {
      from  = 49
      to    = 500
      color = "#FAAD14"
    }
    thresholds {
      from  = 500
      color = "#FF4D4F"
    }
  }
  tags {
    key   = "Mission"
    value = "Terraform"
  }
}

resource "leanspace_widgets" "test_bar" {
  name        = "Terraform Bar Widget"
  description = "A bar widget created with Terraform"
  type        = "BAR"
  granularity = "hour"
  query_time_dimension = "received_at"
  display_time_dimension = "timestamp"
  series {
    id          = "error_code"
    name        = "Error Code Series"
    datasource  = "raw_stream"
    aggregation = "count"
    filters {
      filter_by = "error_code"
      operator  = "notEquals"
      value     = 500
    }
  }
  metadata {
    y_axis_range_min = [200]
    y_axis_range_max = [600]
  }
  tags {
    key   = "Mission"
    value = "Terraform"
  }
}

resource "leanspace_widgets" "test_area" {
  name        = "Terraform Area Widget"
  description = "An area widget created with Terraform"
  type        = "AREA"
  granularity = "day"
  query_time_dimension = "timestamp"
  display_time_dimension = "timestamp"
  series {
    id          = var.numeric_metric_id
    name        = "Numeric Metric Series"
    datasource  = "metric"
    aggregation = "max"
    filters {
      filter_by = var.numeric_metric_id
      operator  = "lt"
      value     = 1000
    }
    filters {
      filter_by = var.numeric_metric_id
      operator  = "gt"
      value     = 0
    }
  }
  tags {
    key   = "Mission"
    value = "Terraform"
  }
}

resource "leanspace_widgets" "test_value" {
  name        = "Terraform Value Widget"
  description = "A value widget created with Terraform"
  type        = "VALUE"
  granularity = "minute"
  query_time_dimension = "received_at"
  display_time_dimension = "timestamp"
  series {
    id          = var.text_metric_id
    name        = "Text Metric Series"
    datasource  = "metric"
    aggregation = "max"
  }
  tags {
    key   = "Mission"
    value = "Terraform"
  }
}

resource "leanspace_widgets" "test_resources" {
  name        = "Terraform Resources Widget"
  description = "A resources widget created with Terraform"
  type        = "RESOURCES"
  granularity = "raw"
  series {
    id          = var.resource_id
    datasource  = "resources"
    aggregation = "none"
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

- `granularity` (String) it must be one of these values: second, minute, hour, day, week, month, raw
- `name` (String)
- `series` (Block Set, Min: 1) (see [below for nested schema](#nestedblock--series))
- `type` (String) it must be one of these values: TABLE, LINE, BAR, AREA, VALUE, RESOURCES, EARTH, GAUGE, ENUM, ORBITAL_VIEW

### Optional

- `description` (String)
- `display_time_dimension` (String) it must be one of these values: timestamp, received_at, ingested_at
- `metadata` (Block List, Max: 1) (see [below for nested schema](#nestedblock--metadata))
- `query_time_dimension` (String) it must be one of these values: timestamp, received_at, ingested_at
- `tags` (Block Set) (see [below for nested schema](#nestedblock--tags))

### Read-Only

- `created_at` (String) When it was created
- `created_by` (String) Who created it
- `dashboards` (Set of Object) (see [below for nested schema](#nestedatt--dashboards))
- `id` (String) The ID of this resource.
- `last_modified_at` (String) When it was last modified
- `last_modified_by` (String) Who modified it the last

<a id="nestedblock--series"></a>
### Nested Schema for `series`

Required:

- `aggregation` (String) it must be one of these values: avg, count, sum, min, max, none
- `datasource` (String) it must be one of these values: metric, raw_stream, resources, topology, orbits, ground_stations, areas_of_interest

Optional:

- `filters` (Block Set, Max: 3) (see [below for nested schema](#nestedblock--series--filters))
- `name` (String)

Read-Only:

- `id` (String) The ID of this resource.

<a id="nestedblock--series--filters"></a>
### Nested Schema for `series.filters`

Required:

- `filter_by` (String)
- `operator` (String) it must be one of these values: gt, lt, equals, notEquals
- `value` (String)



<a id="nestedblock--metadata"></a>
### Nested Schema for `metadata`

Optional:

- `thresholds` (Block List, Max: 10) The threshold applies only to the Gauge widget. (see [below for nested schema](#nestedblock--metadata--thresholds))
- `y_axis_label` (String)
- `y_axis_range_max` (List of Number) The maximum value for the widget's Y axis. Set to an array with the value inside (an empty array is treated as unset). This is due to Terraform limitations.
- `y_axis_range_min` (List of Number) The minimum value for the widget's Y axis. Set to an array with the value inside (an empty array is treated as unset). This is due to Terraform limitations.

<a id="nestedblock--metadata--thresholds"></a>
### Nested Schema for `metadata.thresholds`

Required:

- `color` (String)

Optional:

- `from` (String)
- `to` (String)



<a id="nestedblock--tags"></a>
### Nested Schema for `tags`

Required:

- `key` (String)

Optional:

- `value` (String)


<a id="nestedatt--dashboards"></a>
### Nested Schema for `dashboards`

Read-Only:

- `id` (String)
- `name` (String)
