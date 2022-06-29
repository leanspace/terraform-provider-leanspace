package command_definitions

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var commandDefinitionSchema = map[string]*schema.Schema{
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
	"identifier": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"metadata": {
		Type:     schema.TypeSet,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
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
					ValidateFunc: validation.StringInSlice([]string{"NUMERIC", "TEXT", "TIMESTAMP", "DATE", "TIME", "BOOLEAN"}, false),
				},
			},
		},
	},
	"arguments": {
		Type:     schema.TypeSet,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"id": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"name": {
					Type:     schema.TypeString,
					Required: true,
				},
				"identifier": {
					Type:     schema.TypeString,
					Required: true,
				},
				"description": {
					Type:     schema.TypeString,
					Optional: true,
				},
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
				"default_value": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"type": {
					Type:         schema.TypeString,
					Required:     true,
					ValidateFunc: validation.StringInSlice([]string{"NUMERIC", "ENUM", "TEXT", "TIMESTAMP", "DATE", "TIME", "BOOLEAN"}, false),
				},
				"required": {
					Type:     schema.TypeBool,
					Optional: true,
				},
			},
		},
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
}
