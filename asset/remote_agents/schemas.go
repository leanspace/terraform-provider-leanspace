package remote_agents

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var remoteAgentSchema = map[string]*schema.Schema{
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
	"service_account_id": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"connectors": {
		Type:     schema.TypeSet,
		Optional: true,
		Elem: &schema.Resource{
			Schema: connectorSchema,
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

var connectorSchema = map[string]*schema.Schema{
	"id": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"gateway_id": {
		Type:     schema.TypeString,
		Required: true,
	},
	"type": {
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validation.StringInSlice([]string{"INBOUND", "OUTBOUND"}, false),
	},
	"socket": {
		Type:     schema.TypeList,
		Required: true,
		MinItems: 1,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: socketSchema,
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

	// inbound only
	"stream_id": {
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: validation.IsUUID,
	},
	"destination": {
		Type:     schema.TypeList,
		Computed: true,
		Elem: &schema.Resource{
			Schema: connTargetSchema,
		},
	},
	// outbound only
	"command_queue_id": {
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: validation.IsUUID,
	},
	"source": {
		Type:     schema.TypeList,
		Computed: true,
		Elem: &schema.Resource{
			Schema: connTargetSchema,
		},
	},
}

var socketSchema = map[string]*schema.Schema{
	"type": {
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validation.StringInSlice([]string{"TCP", "UDP"}, false),
	},
	"host": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"port": {
		Type:     schema.TypeInt,
		Required: true,
	},
}

var connTargetSchema = map[string]*schema.Schema{
	"type": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"binding": {
		Type:     schema.TypeString,
		Optional: true,
	},
}
