package properties

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"leanspace-terraform-provider/helper"
	"leanspace-terraform-provider/helper/general_objects"
)

var validPropertyTypes = []string{"NUMERIC", "ENUM", "TEXT", "TIMESTAMP", "DATE", "TIME", "BOOLEAN", "GEOPOINT"}

var propertyFieldSchema = map[string]*schema.Schema{
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
	"value": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"type": {
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validation.StringInSlice(validPropertyTypes, false),
		Description:  helper.AllowedValuesToDescription(validPropertyTypes),
	},
}

var geoPointFieldsSchema = map[string]*schema.Schema{
	"elevation": {
		Type:     schema.TypeList,
		MaxItems: 1,
		Required: true,
		Elem: &schema.Resource{
			Schema: propertyFieldSchema,
		},
	},
	"latitude": {
		Type:     schema.TypeList,
		MaxItems: 1,
		Required: true,
		Elem: &schema.Resource{
			Schema: propertyFieldSchema,
		},
	},
	"longitude": {
		Type:     schema.TypeList,
		MaxItems: 1,
		Required: true,
		Elem: &schema.Resource{
			Schema: propertyFieldSchema,
		},
	},
}

var propertySchema = map[string]*schema.Schema{
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
	"node_id": {
		Type:         schema.TypeString,
		Required:     true,
		ForceNew:     true,
		ValidateFunc: validation.IsUUID,
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
	"tags": general_objects.TagsSchema,
	"min_length": {
		Type:         schema.TypeInt,
		Optional:     true,
		ValidateFunc: validation.IntAtLeast(1),
		Description:  "Text only: Minimum length of this text (at least 1)",
	},
	"max_length": {
		Type:         schema.TypeInt,
		Optional:     true,
		ValidateFunc: validation.IntAtLeast(1),
		Description:  "Text only: Maximum length of this text (at least 1)",
	},
	"pattern": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Text only: Regex defined the allowed pattern of this text",
	},
	"before": {
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: validation.IsRFC3339Time,
		Description:  "Time/date/timestamp only: Maximum date allowed",
	},
	"after": {
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: validation.IsRFC3339Time,
		Description:  "Time/date/timestamp only: Minimum date allowed",
	},
	"fields": {
		Type:     schema.TypeList,
		MaxItems: 1,
		Optional: true,
		Elem: &schema.Resource{
			Schema: geoPointFieldsSchema,
		},
		Description: "Geopoint only",
	},
	"options": {
		Type:        schema.TypeMap,
		Optional:    true,
		Description: "Enum only: The allowed values for the enum in the format 1 = \"value\"",
	},
	"min": {
		Type:        schema.TypeFloat,
		Optional:    true,
		Description: "Numeric only",
	},
	"max": {
		Type:        schema.TypeFloat,
		Optional:    true,
		Description: "Numeric only",
	},
	"scale": {
		Type:        schema.TypeInt,
		Optional:    true,
		Description: "Numeric only",
	},
	"precision": {
		Type:        schema.TypeInt,
		Optional:    true,
		Description: "Numeric only: How many values after the comma should be accepted",
	},
	"unit_id": {
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: validation.IsUUID,
		Description:  "Numeric only",
	},
	"value": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"type": {
		Type:         schema.TypeString,
		Required:     true,
		ForceNew:     true,
		ValidateFunc: validation.StringInSlice(validPropertyTypes, false),
		Description:  helper.AllowedValuesToDescription(validPropertyTypes),
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
	"node_types": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	},
	"node_kinds": {
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
}
