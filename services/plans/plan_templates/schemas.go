package plan_templates

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

var nameRegex = regexp.MustCompile(`^[a-zA-Z0-9_-]*$`)

var planTemplateSchema = map[string]*schema.Schema{
	"id": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"asset_id": {
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validation.IsUUID,
	},
	"name": {
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validation.StringMatch(nameRegex, "Must be a valid Plan Template name"),
	},
	"description": {
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: validation.StringLenBetween(0, 2000),
	},
	"integrity_status": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"activity_configs": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Resource{
			Schema: activityConfigResultSchema,
		},
	},
	"estimated_duration_in_seconds": {
		Type:     schema.TypeInt,
		Computed: true,
	},
	"invalid_plan_template_reasons": {
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
	"activity_definition_id": {
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validation.IsUUID,
	},
	"delay_reference_on_predecessor": {
		Type:     schema.TypeString,
		Optional: true,
	},

	"position": {
		Type:         schema.TypeInt,
		Required:     true,
		ValidateFunc: validation.IntBetween(0, 499),
	},

	"delay_in_seconds": {
		Type:         schema.TypeInt,
		Required:     true,
		ValidateFunc: validation.IntBetween(0, 86400),
	},

	"estimated_duration_in_seconds": {
		Type:         schema.TypeInt,
		Optional:     true,
		ValidateFunc: validation.IntBetween(0, 86400),
	},

	"name": {
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: validation.StringMatch(nameRegex, "Must be a valid name"),
	},

	"arguments": {
		Type:     schema.TypeSet,
		Optional: true,
		Elem: &schema.Resource{
			Schema: argumentSchema,
		},
	},

	"resource_function_formulas": {
		Type:     schema.TypeSet,
		Optional: true,
		Elem: &schema.Resource{
			Schema: resourceFunctionFormulaOverloadSchema,
		},
	},

	"tags": general_objects.KeyValuesSchema,

	"definition_link_status": {
		Type:     schema.TypeString,
		Computed: true,
	},

	"invalid_definition_link_reasons": {
		Type:     schema.TypeSet,
		Computed: true,
		Elem: &schema.Resource{
			Schema: invalidDefinitionLinkReasonSchema,
		},
	},
}

var invalidPlanTemplateReasonSchema = map[string]*schema.Schema{
	"code": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"message": {
		Type:     schema.TypeString,
		Computed: true,
	},
}

var argumentSchema = map[string]*schema.Schema{
	"name": {
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validation.StringMatch(nameRegex, "Must be a valid name"),
	},
	"attributes": {
		Type:     schema.TypeSet,
		Required: true,
		MinItems: 1,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: general_objects.ValueAttributeSchema([]string{"TLE", "STRUCTURE"}),
		},
	},
}

var resourceFunctionFormulaOverloadSchema = map[string]*schema.Schema{
	"resource_function_id": {
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validation.IsUUID,
	},
	"formula": {
		Type:     schema.TypeSet,
		Required: true,
		Elem: &schema.Resource{
			Schema: resourceFunctionFormulaSchema,
		},
	},
}

var resourceFunctionFormulaSchema = map[string]*schema.Schema{
	"type": {
		Type:     schema.TypeString,
		Required: true,
	},
	"constant": {
		Type:     schema.TypeFloat,
		Required: true,
	},
	"rate": {
		Type:     schema.TypeFloat,
		Required: true,
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
