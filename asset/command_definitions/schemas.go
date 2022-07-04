package command_definitions

import (
	"terraform-provider-asset/asset/general_objects"

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
			Schema: metadataSchema,
		},
	},
	"arguments": {
		Type:     schema.TypeSet,
		Optional: true,
		Elem: &schema.Resource{
			Schema: argumentSchema,
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

var metadataSchema = map[string]*schema.Schema{
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
	"attributes": {
		Type:     schema.TypeList,
		Required: true,
		MinItems: 1,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: valueAttributeSchema,
		},
	},
}

var validMetadataTypes = []string{
	"NUMERIC", "BOOLEAN", "TEXT", "DATE", "TIME", "TIMESTAMP", "ENUM",
}

var valueAttributeSchema = map[string]*schema.Schema{
	"value": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"type": {
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validation.StringInSlice(validMetadataTypes, false),
	},
	// Numeric only
	"unit_id": {
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: validation.IsUUID,
	},
}

var argumentSchema = map[string]*schema.Schema{
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
	"attributes": {
		Type:     schema.TypeList,
		Required: true,
		MinItems: 1,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: general_objects.DefinitionAttributeSchema(nil, nil),
		},
	},
}
