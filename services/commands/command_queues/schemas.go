package command_queues

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var commandQueueSchema = map[string]*schema.Schema{
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
	"ground_station_ids": {
		Type:     schema.TypeSet,
		Optional: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	},
	"command_transformer_plugin_id": {
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: validation.IsUUID,
	},
	"protocol_transformer_plugin_id": {
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: validation.IsUUID,
	},
	"protocol_transformer_init_data": {
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
}
