package properties

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"leanspace-terraform-provider/helper/general_objects"
)

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
		Type:     schema.TypeString,
		Computed: true,
	},
	"created_by": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"last_modified_at": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"last_modified_by": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"value": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"type": {
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validation.StringInSlice([]string{"NUMERIC", "ENUM", "TEXT", "TIMESTAMP", "DATE", "TIME", "BOOLEAN", "GEOPOINT"}, false),
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
		Type:     schema.TypeString,
		Computed: true,
	},
	"created_by": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"last_modified_at": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"last_modified_by": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"tags": general_objects.TagsSchema,
	"min_length": {
		Type:         schema.TypeInt,
		Optional:     true,
		ValidateFunc: validation.IntAtLeast(1),
	},
	"max_length": {
		Type:     schema.TypeInt,
		Optional: true,
	},
	"pattern": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"before": {
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: validation.IsRFC3339Time,
	},
	"after": {
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: validation.IsRFC3339Time,
	},
	"fields": {
		Type:     schema.TypeList,
		MaxItems: 1,
		Optional: true,
		Elem: &schema.Resource{
			Schema: geoPointFieldsSchema,
		},
	},
	"options": {
		Type:     schema.TypeMap,
		Optional: true,
	},
	"min": {
		Type:     schema.TypeFloat,
		Optional: true,
	},
	"max": {
		Type:     schema.TypeFloat,
		Optional: true,
	},
	"scale": {
		Type:     schema.TypeInt,
		Optional: true,
	},
	"precision": {
		Type:     schema.TypeInt,
		Optional: true,
	},
	"unit_id": {
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: validation.IsUUID,
	},
	"value": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"type": {
		Type:         schema.TypeString,
		Required:     true,
		ForceNew:     true,
		ValidateFunc: validation.StringInSlice([]string{"NUMERIC", "ENUM", "TEXT", "TIMESTAMP", "DATE", "TIME", "BOOLEAN", "GEOPOINT"}, false),
	},
}
