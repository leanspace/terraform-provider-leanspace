package metrics

import (
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var metricSchema = map[string]*schema.Schema{
	"id": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"node_id": {
		Type:         schema.TypeString,
		Required:     true,
		ForceNew:     true,
		ValidateFunc: validation.IsUUID,
	},
	"name": {
		Type:     schema.TypeString,
		Required: true,
	},
	"description": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"attributes": {
		Type:     schema.TypeList,
		MinItems: 1,
		MaxItems: 1,
		Required: true,
		Elem: &schema.Resource{
			Schema: general_objects.DefinitionAttributeSchema(
				[]string{"TIME", "ARRAY"},             // TIME and ARRAY type not allowed
				[]string{"required", "default_value"}, // Fields unused
				true,                                  // Force recreation if the type changes
			),
		},
	},
	"calibration": {
		Type:        schema.TypeList,
		MaxItems:    1,
		Optional:    true,
		Description: "Calibration applied to the metric values. Removing this block removes the calibration of the metric.",
		Elem: &schema.Resource{
			Schema: calibrationSchema,
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

var dataSourceFilterSchema = map[string]*schema.Schema{
	"node_ids": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.IsUUID,
		},
	},
	"attribute_types": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.StringInSlice(general_objects.ValidAttributeSchemaTypes, false),
			Description:  helper.AllowedValuesToDescription(general_objects.ValidAttributeSchemaTypes),
		},
	},
	"tags": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	},
}

func calibrationBoundSchema(description string) *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Optional:    true,
		MaxItems:    1,
		Elem:        &schema.Schema{Type: schema.TypeFloat},
		Description: description,
	}
}

var calibrationSchema = map[string]*schema.Schema{
	"type": {
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validation.StringInSlice(ValidCalibrationTypes, false),
		Description: helper.AllowedValuesToDescription(ValidCalibrationTypes) +
			". TEXT_MAPPING is only allowed for TEXT metrics, the other types only for NUMERIC metrics.",
	},
	"description": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"coefficients": {
		Type:        schema.TypeList,
		Optional:    true,
		MaxItems:    6,
		Elem:        &schema.Schema{Type: schema.TypeFloat},
		Description: "POLYNOMIAL, LOG and INVERSE_LOG only (required): between 1 and 6 coefficients.",
	},
	"base": {
		Type:         schema.TypeFloat,
		Optional:     true,
		ValidateFunc: validation.FloatAtLeast(1),
		Description:  "LOG and INVERSE_LOG only: the base of the logarithm, at least 1.",
	},
	"capping": {
		Type:        schema.TypeList,
		Optional:    true,
		MaxItems:    1,
		Description: "POLYNOMIAL, LOG, INVERSE_LOG and RANGE only: bounds applied to the calibrated value.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"min": calibrationBoundSchema("Lower bound of the calibrated value. Leave empty for no lower bound."),
				"max": calibrationBoundSchema("Upper bound of the calibrated value. Leave empty for no upper bound."),
			},
		},
	},
	"min": calibrationBoundSchema("CAPPING only: lower bound of the value. Leave empty for no lower bound."),
	"max": calibrationBoundSchema("CAPPING only: upper bound of the value. Leave empty for no upper bound."),
	"ranges": {
		Type:        schema.TypeList,
		Optional:    true,
		MaxItems:    10,
		Description: "RANGE only (required): between 1 and 10 non-overlapping ranges, each mapped to a value.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"min": calibrationBoundSchema("Lower bound of the range. Leave empty for no lower bound."),
				"max": calibrationBoundSchema("Upper bound of the range. Leave empty for no upper bound."),
				"value": {
					Type:        schema.TypeFloat,
					Required:    true,
					Description: "The calibrated value for values within this range.",
				},
			},
		},
	},
	"mappings": {
		Type:        schema.TypeMap,
		Optional:    true,
		Elem:        &schema.Schema{Type: schema.TypeString},
		Description: "TEXT_MAPPING only (required): maps raw text values to calibrated text values.",
	},
}
