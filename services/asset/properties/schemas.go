package properties

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

var validPropertyTypes = []string{"NUMERIC", "ENUM", "TEXT", "TIMESTAMP", "DATE", "TIME", "BOOLEAN", "GEOPOINT", "TLE"}

var PropertySchema = map[string]*schema.Schema{
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
		Optional:     true,
		ForceNew:     true,
		Description:  "This field is required when creating the property but optional when creating a node with properties.",
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
		ValidateFunc: helper.IsValidTimeDateOrTimestamp,
		Description:  "Time/date/timestamp only: Maximum date allowed",
	},
	"after": {
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: helper.IsValidTimeDateOrTimestamp,
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
	"built_in": {
		Type:        schema.TypeBool,
		Computed:    true,
		Description: "Indicates if it is a build-in property.",
	},
	"additional_properties": {
		Type:     schema.TypeMap,
		Optional: true,
	},
}

var geoPointFieldsSchema = map[string]*schema.Schema{
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
	"elevation": {
		Type:     schema.TypeList,
		MaxItems: 1,
		Required: true,
		Elem: &schema.Resource{
			Schema: propertyFieldSchema,
		},
	},
}

var propertyFieldSchema = map[string]*schema.Schema{
	"additional_properties": {
		Type:     schema.TypeMap,
		Optional: true,
	},
	"value": {
		Type:     schema.TypeString,
		Optional: true,
	},

	// Numeric only
	"scale": {
		Type:        schema.TypeInt,
		Optional:    true,
		Description: "Property field with numeric type only: the scale required.",
	},
	"unit_id": {
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: validation.IsUUID,
		Description:  "Property field with numeric type only",
	},
	"min": {
		Type:        schema.TypeFloat,
		Optional:    true,
		Description: "Property field with numeric type only: the minimum value allowed.",
	},
	"precision": {
		Type:        schema.TypeInt,
		Optional:    true,
		Description: "Property field with numeric type only: How many values after the comma should be accepted",
	},
	"max": {
		Type:        schema.TypeFloat,
		Optional:    true,
		Description: "Property field with numeric type only: the maximum value allowed.",
	},
}

var dataSourceFilterSchema = map[string]*schema.Schema{
	"built_in": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type: schema.TypeBool,
		},
	},
	"tags": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	},
	"names": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	},
	"node_ids": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.IsUUID,
		},
	},
}
