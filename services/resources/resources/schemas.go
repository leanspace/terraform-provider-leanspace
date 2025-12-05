package resources

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

var validResourceConstraintTypes = []string{"LIMIT", "THRESHOLD"}
var validResourceConstraintKinds = []string{"UPPER", "LOWER"}

var resourceSchema = map[string]*schema.Schema{
	"id": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"asset_id": {
		Type:         schema.TypeString,
		Required:     true,
		ForceNew:     true,
		ValidateFunc: validation.IsUUID,
	},
	"unit_id": {
		Type:         schema.TypeString,
		Optional:     true,
		ForceNew:     true,
		ValidateFunc: validation.IsUUID,
	},
	"metric_id": {
		Type:         schema.TypeString,
		Optional:     true,
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
	"default_level": {
		Type:     schema.TypeFloat,
		Optional: true,
	},
	"constraints": {
		Type:       schema.TypeSet,
		Deprecated: "Prefer using the lowerLimit, upperLimit and thresholds fields",
		Optional:   true,
		Elem: &schema.Resource{
			Schema: resourceConstraintsSchema,
		},
	},
	"lower_limit": {
		Type:     schema.TypeList,
		Optional: true,
		MaxItems: 1,
		Elem:     &schema.Schema{Type: schema.TypeFloat},
	},
	"upper_limit": {
		Type:     schema.TypeList,
		Optional: true,
		MaxItems: 1,
		Elem:     &schema.Schema{Type: schema.TypeFloat},
	},
	"thresholds": {
		Type:        schema.TypeSet,
		Optional:    true,
		Description: "Currently, at most three LOWER and three UPPER thresholds can be set",
		Elem: &schema.Resource{
			Schema: resourceThresholdSchema,
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

var resourceConstraintsSchema = map[string]*schema.Schema{
	"type": {
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validation.StringInSlice(validResourceConstraintTypes, false),
		Description:  helper.AllowedValuesToDescription(validResourceConstraintTypes),
	},
	"kind": {
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validation.StringInSlice(validResourceConstraintKinds, false),
		Description:  helper.AllowedValuesToDescription(validResourceConstraintKinds),
	},
	"value": {
		Type:     schema.TypeFloat,
		Required: true,
	},
	"name": {
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: helper.IsValidName,
	},
}

var resourceThresholdSchema = map[string]*schema.Schema{
	"kind": {
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validation.StringInSlice(validResourceConstraintKinds, false),
		Description:  helper.AllowedValuesToDescription(validResourceConstraintKinds),
	},
	"name": {
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: helper.IsValidName,
	},
	"violation_when_reached": {
		Type:     schema.TypeBool,
		Optional: true,
	},
	"value": {
		Type:     schema.TypeFloat,
		Required: true,
	},
}

var dataSourceFilterSchema = map[string]*schema.Schema{
	"ids": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.IsUUID,
		},
	},
	"asset_ids": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.IsUUID,
		},
	},
	"unit_ids": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.IsUUID,
		},
	},
	"metric_ids": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.IsUUID,
		},
	},
	"tags": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	},
	"created_bys": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.IsUUID,
		},
		Description: "Filter on the user who created the Resource. If you have no wish to use this field as a filter, either provide a null value or remove the field.",
	},
	"last_modified_bys": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.IsUUID,
		},
		Description: "Filter on the user who last modified the Resource. If you have no wish to use this field as a filter, either provide a null value or remove the field.",
	},
	"from_created_at": {
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: helper.IsValidTimeDateOrTimestamp,
		Description:  "Filter on the Resource creation date. Resources with a creation date greater or equals than the filter value will be selected (if they are not excluded by other filters). If you have no wish to use this field as a filter, either provide a null value or remove the field.",
	},
	"from_last_modified_at": {
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: helper.IsValidTimeDateOrTimestamp,
		Description:  "Filter on the Resource last modification date. Resources with a last modification date greater or equals than the filter value will be selected (if they are not excluded by other filters). If you have no wish to use this field as a filter, either provide a null value or remove the field.",
	},
	"to_created_at": {
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: helper.IsValidTimeDateOrTimestamp,
		Description:  "Filter on the Resource creation date. Resources with a creation date lower or equals than the filter value will be selected (if they are not excluded by other filters). If you have no wish to use this field as a filter, either provide a null value or remove the field.",
	},
	"to_last_modified_at": {
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: helper.IsValidTimeDateOrTimestamp,
		Description:  "Filter on the Resource last modification date. Resources with a last modification date lower or equals than the filter value will be selected (if they are not excluded by other filters). If you have no wish to use this field as a filter, either provide a null value or remove the field.",
	},
}
