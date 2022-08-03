package dashboards

import (
	"leanspace-terraform-provider/helper/general_objects"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"leanspace-terraform-provider/helper"
)

var validWidgetTypes = []string{"TABLE", "LINE", "BAR", "AREA", "VALUE"}

var dashboardSchema = map[string]*schema.Schema{
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
	"node_ids": {
		Type:     schema.TypeSet,
		Optional: true,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.IsUUID,
		},
	},
	"widget_info": {
		Type:     schema.TypeSet,
		Optional: true,
		Elem: &schema.Resource{
			Schema: widgetInfoSchema,
		},
	},
	"widgets": {
		Type:     schema.TypeSet,
		Computed: true,
		Elem: &schema.Resource{
			Schema: dashboardWidgetSchema,
		},
	},
	"tags": general_objects.TagsSchema,
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

var widgetInfoSchema = map[string]*schema.Schema{
	"id": {
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validation.IsUUID,
	},
	"type": {
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validation.StringInSlice(validWidgetTypes, false),
		Description:  helper.AllowedValuesToDescription(validWidgetTypes),
	},
	"w": {
		Type:         schema.TypeInt,
		Required:     true,
		ValidateFunc: validation.IntAtLeast(1),
	},
	"h": {
		Type:         schema.TypeInt,
		Required:     true,
		ValidateFunc: validation.IntAtLeast(1),
	},
	"x": {
		Type:         schema.TypeInt,
		Required:     true,
		ValidateFunc: validation.IntAtLeast(0),
	},
	"y": {
		Type:         schema.TypeInt,
		Required:     true,
		ValidateFunc: validation.IntAtLeast(0),
	},
	"min_w": {
		Type:         schema.TypeInt,
		Optional:     true,
		ValidateFunc: validation.IntAtLeast(1),
	},
	"min_h": {
		Type:         schema.TypeInt,
		Optional:     true,
		ValidateFunc: validation.IntAtLeast(1),
	},
}

var dashboardWidgetSchema = map[string]*schema.Schema{
	"id": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"name": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"description": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"type": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"granularity": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"series": {
		Type:     schema.TypeList,
		Computed: true,
		Elem: &schema.Resource{
			Schema: seriesSchema,
		},
	},
	"metrics": {
		Type:     schema.TypeList,
		Computed: true,
		Elem: &schema.Resource{
			Schema: metricInfoSchema,
		},
	},
	"metadata": {
		Type:     schema.TypeList,
		Computed: true,
		Elem: &schema.Resource{
			Schema: metadataSchema,
		},
	},
	"view": {
		Type:     schema.TypeList,
		Computed: true,
		Elem: &schema.Resource{
			Schema: dashboardInfoSchema,
		},
	},
	"tags": general_objects.TagsSchema,
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
		Computed: true,
	},
	"datasource": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"aggregation": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"filters": {
		Type:     schema.TypeSet,
		Computed: true,
		Elem: &schema.Resource{
			Schema: filterSchema,
		},
	},
}

var filterSchema = map[string]*schema.Schema{
	"filter_by": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"operator": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"value": {
		Type:     schema.TypeString,
		Computed: true,
	},
}

var metricInfoSchema = map[string]*schema.Schema{
	"id": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"aggregation": {
		Type:     schema.TypeString,
		Computed: true,
	},
}

var metadataSchema = map[string]*schema.Schema{
	"y_axis_label": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"y_axis_range_min": {
		Type:     schema.TypeList,
		Computed: true,
		Elem: &schema.Schema{
			Type: schema.TypeFloat,
		},
	},
	"y_axis_range_max": {
		Type:     schema.TypeList,
		Computed: true,
		Elem: &schema.Schema{
			Type: schema.TypeFloat,
		},
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
