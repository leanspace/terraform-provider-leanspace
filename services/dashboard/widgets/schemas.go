package widgets

import (
	"regexp"

	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

	"github.com/leanspace/terraform-provider-leanspace/helper"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var ValidWidgetTypes = []string{"TABLE", "LINE", "BAR", "AREA", "VALUE", "RESOURCES", "EARTH", "GAUGE", "ENUM", "ORBITAL_VIEW"}
var validGranularities = []string{"second", "minute", "hour", "day", "week", "month", "raw"}
var validDatasources = []string{"metric", "raw_stream", "resources", "topology", "orbits", "ground_stations", "areas_of_interest"}
var validAggregations = []string{"avg", "count", "sum", "min", "max", "none"}
var validFilterOperators = []string{"gt", "lt", "equals", "notEquals"}

var colorRegex = regexp.MustCompile(`^#(?:[0-9a-fA-F]{3}){1,2}$`)

var widgetSchema = map[string]*schema.Schema{
	"id": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"name": {
		Type:     schema.TypeString,
		Required: true,
	},
	"description": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"type": {
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validation.StringInSlice(ValidWidgetTypes, false),
		Description:  helper.AllowedValuesToDescription(ValidWidgetTypes),
	},
	"granularity": {
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validation.StringInSlice(validGranularities, false),
		Description:  helper.AllowedValuesToDescription(validGranularities),
	},
	"series": {
		Type:     schema.TypeSet,
		Required: true,
		Elem: &schema.Resource{
			Schema: seriesSchema,
		},
	},
	"metadata": {
		Type:     schema.TypeList,
		Optional: true,
		MinItems: 1,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: metadataSchema,
		},
	},
	"dashboards": {
		Type:     schema.TypeSet,
		Computed: true,
		Elem: &schema.Resource{
			Schema: dashboardInfoSchema,
		},
	},
	"tags": general_objects.KeyValuesSchema,
	"created_at": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "When it was created",
	},
	"created_by": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Who created it",
	},
	"last_modified_at": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "When it was last modified",
	},
	"last_modified_by": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Who modified it the last",
	},
}

var seriesSchema = map[string]*schema.Schema{
	"id": {
		Type:     schema.TypeString,
		Required: true,
	},
	"datasource": {
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validation.StringInSlice(validDatasources, false),
		Description:  helper.AllowedValuesToDescription(validDatasources),
	},
	"aggregation": {
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validation.StringInSlice(validAggregations, false),
		Description:  helper.AllowedValuesToDescription(validAggregations),
	},
	"filters": {
		Type:     schema.TypeSet,
		Optional: true,
		MaxItems: 3,
		Elem: &schema.Resource{
			Schema: filterSchema,
		},
	},
}

var filterSchema = map[string]*schema.Schema{
	"filter_by": {
		Type:     schema.TypeString,
		Required: true,
	},
	"operator": {
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validation.StringInSlice(validFilterOperators, false),
		Description:  helper.AllowedValuesToDescription(validFilterOperators),
	},
	"value": {
		Type:     schema.TypeString,
		Required: true,
	},
}

var metadataSchema = map[string]*schema.Schema{
	"y_axis_label": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"y_axis_range_min": {
		Type:     schema.TypeList,
		Optional: true,
		MinItems: 1,
		MaxItems: 1,
		Description: "The minimum value for the widget's Y axis. Set to an array with the value " +
			"inside (an empty array is treated as unset). This is due to Terraform limitations.",
		Elem: &schema.Schema{
			Type: schema.TypeFloat,
		},
	},
	"y_axis_range_max": {
		Type:     schema.TypeList,
		Optional: true,
		MinItems: 1,
		MaxItems: 1,
		Description: "The maximum value for the widget's Y axis. Set to an array with the value " +
			"inside (an empty array is treated as unset). This is due to Terraform limitations.",
		Elem: &schema.Schema{
			Type: schema.TypeFloat,
		},
	},
	"thresholds": {
		Type:        schema.TypeList,
		Optional:    true,
		MinItems:    1,
		MaxItems:    10,
		Description: "The threshold applies only to the Gauge widget.",
		Elem: &schema.Resource{
			Schema: thresholdSchema,
		},
	},
}

/*
From and to are strings so that they can be nil.
*/
var thresholdSchema = map[string]*schema.Schema{
	"from": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"to": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"color": {
		Type:         schema.TypeString,
		ValidateFunc: validation.StringMatch(colorRegex, "Must be a valid hex color"),
		Required:     true,
	},
}

var dashboardInfoSchema = map[string]*schema.Schema{
	"id": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"name": {
		Type:     schema.TypeString,
		Computed: true,
	},
}

var dataSourceFilterSchema = map[string]*schema.Schema{
	"types": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.StringInSlice(ValidWidgetTypes, false),
			Description:  helper.AllowedValuesToDescription(ValidWidgetTypes),
		},
	},
	"tags": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	},
	"dashboard_ids": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.IsUUID,
		},
	},
	"datasource_ids": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.IsUUID,
		},
	},
	"datasources": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.StringInSlice(validDatasources, false),
			Description:  helper.AllowedValuesToDescription(validDatasources),
		},
	},
}
