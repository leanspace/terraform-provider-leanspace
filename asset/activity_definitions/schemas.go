package activity_definitions

import (
	"terraform-provider-asset/asset/general_objects"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var activityDefinitionSchema = map[string]*schema.Schema{
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
	"estimated_duration": {
		Type:         schema.TypeInt,
		Optional:     true,
		ValidateFunc: validation.IntAtLeast(0),
	},
	"metadata": {
		Type:     schema.TypeSet,
		Optional: true,
		Elem: &schema.Resource{
			Schema: metadataSchema,
		},
	},
	"argument_definitions": {
		Type:     schema.TypeSet,
		Optional: true,
		Elem: &schema.Resource{
			Schema: argumentDefinitionSchema,
		},
	},
	"command_mappings": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Resource{
			Schema: commandMappingSchema,
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
			Schema: general_objects.ValueAttributeSchema,
		},
	},
}

var argumentDefinitionSchema = map[string]*schema.Schema{
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
			Schema: general_objects.DefinitionAttributeSchema(nil, nil),
		},
	},
}

var commandMappingSchema = map[string]*schema.Schema{
	"command_definition_id": {
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validation.IsUUID,
	},
	"position": {
		Type:     schema.TypeInt,
		Computed: true,
	},
	"delay_in_milliseconds": {
		Type:         schema.TypeInt,
		Required:     true,
		ValidateFunc: validation.IntAtLeast(0),
	},
	"argument_mappings": {
		Type:     schema.TypeSet,
		Optional: true,
		Elem: &schema.Resource{
			Schema: argumentMappingSchema,
		},
	},
	"metadata_mappings": {
		Type:     schema.TypeSet,
		Optional: true,
		Elem: &schema.Resource{
			Schema: metadataMappingSchema,
		},
	},
}

var argumentMappingSchema = map[string]*schema.Schema{
	"activity_definition_argument_name": {
		Type:     schema.TypeString,
		Required: true,
	},
	"command_definition_argument_name": {
		Type:     schema.TypeString,
		Required: true,
	},
}

var metadataMappingSchema = map[string]*schema.Schema{
	"activity_definition_metadata_name": {
		Type:     schema.TypeString,
		Required: true,
	},
	"command_definition_argument_name": {
		Type:     schema.TypeString,
		Required: true,
	},
}
