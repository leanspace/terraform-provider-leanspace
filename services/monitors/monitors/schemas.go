package monitors

import (
	"leanspace-terraform-provider/helper/general_objects"
	"leanspace-terraform-provider/services/monitors/action_templates"

	"leanspace-terraform-provider/helper"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var validComparisonOperators = []string{
	"GREATER_THAN",
	"LESSER_THAN",
	"GREATER_THAN_OR_EQUAL_TO",
	"LESSER_THAN_OR_EQUAL_TO",
	"EQUAL_TO",
	"NOT_EQUAL_TO",
}

var validAggregationFunctions = []string{
	"AVERAGE_VALUE",
	"HIGHEST_VALUE",
	"LOWEST_VALUE",
	"SUM_VALUE",
	"COUNT_VALUE",
}

var validPollingFrequencies = []int{1, 60, 1440}

var monitorSchema = map[string]*schema.Schema{
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
	"status": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"polling_frequency_in_minutes": {
		Type:         schema.TypeInt,
		Required:     true,
		ValidateFunc: validation.IntInSlice(validPollingFrequencies),
		Description:  helper.AllowedIntValuesToDescription(validPollingFrequencies),
	},
	"metric_id": {
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validation.IsUUID,
	},
	"node_id": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"statistics": {
		Type:     schema.TypeList,
		Computed: true,
		Elem: &schema.Resource{
			Schema: statisticsSchema,
		},
	},
	"expression": {
		Type:     schema.TypeList,
		Required: true,
		MinItems: 1,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: expressionSchema,
		},
	},
	"action_templates": {
		Type:     schema.TypeSet,
		Computed: true,
		Elem: &schema.Resource{
			Schema: action_templates.ActionTemplateSchema,
		},
	},
	"action_template_ids": {
		Type:     schema.TypeSet,
		Optional: true,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.IsUUID,
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

var statisticsSchema = map[string]*schema.Schema{
	"last_evaluation": {
		Type:     schema.TypeList,
		Computed: true,
		Elem: &schema.Resource{
			Schema: evaluationSchema,
		},
	},
}

var evaluationSchema = map[string]*schema.Schema{
	"timestamp": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"value": {
		Type:     schema.TypeFloat,
		Computed: true,
	},
	"status": {
		Type:     schema.TypeString,
		Computed: true,
	},
}

var expressionSchema = map[string]*schema.Schema{
	"comparison_operator": {
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validation.StringInSlice(validComparisonOperators, false),
		Description:  helper.AllowedValuesToDescription(validComparisonOperators),
	},
	"comparison_value": {
		Type:     schema.TypeFloat,
		Required: true,
	},
	"aggregation_function": {
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validation.StringInSlice(validAggregationFunctions, false),
		Description:  helper.AllowedValuesToDescription(validAggregationFunctions),
	},
	"tolerance": {
		Type:         schema.TypeFloat,
		Optional:     true,
		ValidateFunc: validation.FloatAtLeast(0),
		Description:  "Only valid for EQUAL_TO or NOT_EQUAL_TO comparison operator",
	},
}

var dataSourceFilterSchema = map[string]*schema.Schema{
	"metric_ids": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.IsUUID,
		},
	},
	"name": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"node_ids": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.IsUUID,
		},
	},
	"statuses": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	},
	"tags": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	},
	"action_template_ids": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.IsUUID,
		},
	},
}
