package events_definitions

import (
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var validOperator = []string{"EQUAL_TO"}
var source = []string{"COMMAND_STATE_CHANGED", "MONITOR_TRIGGERED", "PASS_AOS", "PASS_LOS", "STREAM_DECODED"}
var state = []string{"ACTIVE", "INACTIVE"}
var validMetadataTypes = []string{
	"NUMERIC", "BOOLEAN", "TEXT",
}

var eventsDefinitions = map[string]*schema.Schema{
	"id": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"description": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"mappings": {
		Type:     schema.TypeSet,
		Optional: true,
		Elem: &schema.Resource{
			Schema: mappingSchema,
		},
	},
	"name": {
		Type:     schema.TypeString,
		Required: true,
	},
	"rules": {
		Type:     schema.TypeSet,
		Optional: true,
		Elem: &schema.Resource{
			Schema: ruleSchema,
		},
	},
	"source": {
		Type:         schema.TypeString,
		Required:     true,
		Description:  helper.AllowedValuesToDescription(source),
		ValidateFunc: validation.StringInSlice(source, false),
	},
	"state": {
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validation.StringInSlice(state, false),
		Description:  helper.AllowedValuesToDescription(state),
	},
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
	"tags": general_objects.KeyValuesSchema,
}

var mappingSchema = map[string]*schema.Schema{
	"default_value": {
		Type: schema.TypeMap,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional:    true,
		Description: "The allowed values are in the format \"key\" = \"value\"",
	},
	"origin": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"target": {
		Type:     schema.TypeString,
		Optional: true,
	},
}

var ruleSchema = map[string]*schema.Schema{
	"path": {
		Type:     schema.TypeString,
		Required: true,
	},
	"operator": {
		Type:         schema.TypeString,
		Required:     true,
		Description:  helper.AllowedValuesToDescription(validOperator),
		ValidateFunc: validation.StringInSlice(validOperator, false),
	},
	"comparison_value": {
		Type:     schema.TypeList,
		MinItems: 0,
		MaxItems: 1,
		Optional: true,
		Elem: &schema.Resource{
			Schema: ComparisonValueAttributeSchema,
		},
	},
}

var ComparisonValueAttributeSchema = map[string]*schema.Schema{
	"value": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"type": {
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validation.StringInSlice(validMetadataTypes, false),
		Description:  helper.AllowedValuesToDescription(validMetadataTypes),
	},
}

var dataSourceFilterSchema = map[string]*schema.Schema{
	"tags": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	},
}
