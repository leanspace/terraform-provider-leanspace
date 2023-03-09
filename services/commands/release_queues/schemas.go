package release_queues

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

var validCommandTransformationStrategies = []string{"TEST", "NO_TRANSFORMATION", "USE_PLUGIN"}

var releaseQueueSchema = map[string]*schema.Schema{
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
	"name": {
		Type:     schema.TypeString,
		Required: true,
	},
	"description": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"command_transformer_plugin_id": {
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: validation.IsUUID,
		Description:  "The Id of the Command Transformer Plugin",
	},
	"command_transformation_strategy": {
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: validation.StringInSlice(validCommandTransformationStrategies, false),
		Description:  "What transformation strategy shall be applied on created and updated Commands",
	},
	"command_transformer_plugin_configuration_data": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Configuration data used by the Command Transformer Plugin (coming soon)",
	},
	"global_transmission_metadata": general_objects.KeyValuesSchema,
	"logical_lock": {
		Type:     schema.TypeBool,
		Computed: true,
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

var dataSourceFilterSchema = map[string]*schema.Schema{
	"asset_ids": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.IsUUID,
		},
	},
	"command_transformer_plugin_ids": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.IsUUID,
		},
	},
	"ids": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.IsUUID,
		},
		Description: "Only returns release queues whose id matches one of the provided values.",
	},
	"logical_lock": {
		Type:     schema.TypeBool,
		Optional: true,
	},
	"query": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Search by name or description",
	},
}
