package activity_definitions

import (
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

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
	"mapping_status": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Mapping status with Command definition arguments. Can be IN_SYNC or OUT_OF_SYNC",
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
	"tags": general_objects.KeyValuesSchema,
}

var metadataSchema = map[string]*schema.Schema{
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
			Schema: general_objects.ValueAttributeSchema([]string{"ENUM", "STRUCTURE", "TLE"}),
		},
	},
}

var argumentDefinitionSchema = map[string]*schema.Schema{
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
			Schema: general_objects.DefinitionAttributeSchema(
				[]string{"STRUCTURE", "TLE"}, // attribute types not allowed in command definition attributes
				nil,                          // All fields are used
				false,                        // Does not force recreation if the type changes
			),
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
		Description:  "Delay to execute this command",
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
	"mapping_status": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Mapping status with the Command definition argument. Can be IN_SYNC or OUT_OF_SYNC",
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
	"mapping_status": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Mapping status with the Command definition argument. Can be IN_SYNC or OUT_OF_SYNC",
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
}
