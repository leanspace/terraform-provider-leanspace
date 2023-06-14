package monitors

import (
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
	"github.com/leanspace/terraform-provider-leanspace/services/monitors/action_templates"

	"github.com/leanspace/terraform-provider-leanspace/helper"

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
	"metric_id": {
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validation.IsUUID,
	},
	"node_id": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"rule": {
		Type:     schema.TypeList,
		Required: true,
		MinItems: 1,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: ruleSchema,
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
	"type": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Represent the type of the monitor. This field is deprecated and it will be removed soon. Please use only this type: REALTIME.",
	},
}

var ruleSchema = map[string]*schema.Schema{ // ruleSchema
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
