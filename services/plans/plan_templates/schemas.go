package plan_templates

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

var nameRegex = regexp.MustCompile(`^[A-Z](?:[A-Z_]*[A-Z])?$`)

var planTemplateSchema = map[string]*schema.Schema{
	"id": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"assetId": {
		Type:     schema.TypeString,
		Required: true,
	},
	"name": {
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validation.StringMatch(nameRegex, "Must be a valid Plan Template name"),
	},
	"description": {
		Type:         schema.TypeString,
		Required:     false,
		ValidateFunc: validation.StringLenBetween(1, 2000),
	},
	"integrityStatus": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"activityConfigs": {
		Type:     schema.TypeList,
		Required: true,
		Elem: &schema.Resource{
			Schema: activityConfigResultSchema,
		},
	},
	"estimatedDurationInSeconds": {
		Type:     schema.TypeInt,
		Computed: true,
	},
	"invalidPlanTemplateReasons": {
		Type:     schema.TypeList,
		Computed: true,
		Elem: &schema.Resource{
			Schema: invalidPlanTemplateReasonSchema,
		},
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
}

var activityConfigResultSchema = map[string]*schema.Schema{
	"activityDefinitionId": {
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validation.IsUUID,
	},
	"delayReferenceOnPredecessor": {
		Type: schema.TypeString,
	},
	"position": {
		Type:         schema.TypeInt,
		ValidateFunc: validation.IntBetween(0, 499),
	},
	"delayInSeconds": {
		Type:         schema.TypeInt,
		Required:     true,
		ValidateFunc: validation.IntBetween(0, 499),
	},
	"estimatedDurationInSeconds": {
		Type:         schema.TypeInt,
		Required:     true,
		ValidateFunc: validation.IntBetween(0, 86400),
	},
	"name": {
		Type:         schema.TypeString,
		ValidateFunc: validation.StringMatch(nameRegex, "Must be a valid name"),
	},
	"arguments": {
		Type: schema.TypeList,
		Elem: &schema.Resource{
			Schema: argumentSchema,
		},
	},
	"resourceFunctionFormulas": {
		Type: schema.TypeList,
		Elem: &schema.Resource{
			Schema: resourceFunctionFormulaOverloadSchema,
		},
	},
	"tags": general_objects.KeyValuesSchema,
	"definitionLinkStatus": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"invalidDefinitionLinkReasons": {
		Type:     schema.TypeList,
		Computed: true,
		Elem: &schema.Resource{
			Schema: invalidDefinitionLinkReasonSchema,
		},
	},
}

var invalidPlanTemplateReasonSchema = map[string]*schema.Schema{
	"code": {
		Type:     schema.TypeString,
		Required: true,
	},
	"message": {
		Type:     schema.TypeString,
		Required: true,
	},
}

var argumentSchema = map[string]*schema.Schema{
	"name": {
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validation.StringMatch(nameRegex, "Must be a valid name"),
	},
	"attributes": {
		Type:     schema.TypeList,
		Required: true,
		MinItems: 1,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: general_objects.ValueAttributeSchema([]string{"NUMERIC", "TEXT", "BOOLEAN", "ENUM", "TIMESTAMP", "DATE", "TIME", "GEOPOINT", "BINARY", "ARRAY"}),
		},
	},
}

var resourceFunctionFormulaOverloadSchema = map[string]*schema.Schema{
	"resourceFunctionId": {
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validation.IsUUID,
	},
	"formula": {
		Elem: &schema.Resource{
			Schema: resourceFunctionFormulaSchema,
		},
	},
}

var resourceFunctionFormulaSchema = map[string]*schema.Schema{
	"type": {
		Type: schema.TypeString,
	},
}

var invalidDefinitionLinkReasonSchema = map[string]*schema.Schema{
	"code": {
		Type:     schema.TypeString,
		Required: true,
	},
	"message": {
		Type:     schema.TypeString,
		Required: true,
	},
}
