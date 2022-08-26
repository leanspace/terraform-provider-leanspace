package remote_agents

import (
	"leanspace-terraform-provider/helper"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var validConnectorTypes = []string{"INBOUND", "OUTBOUND"}
var validProtocolTypes = []string{"TCP", "UDP"}

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
		Type:         schema.TypeString,
		Computed:     true,
		Optional:     true,
		ValidateFunc: validation.IsUUID,
	},
	"connectors": {
		Type:     schema.TypeSet,
		Optional: true,
		Elem: &schema.Resource{
			Schema: connectorSchema,
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
}

var connectorSchema = map[string]*schema.Schema{
	"id": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"gateway_id": {
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validation.IsUUID,
		Description:  "Id of the node",
	},
	"type": {
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validation.StringInSlice(validConnectorTypes, false),
		Description:  helper.AllowedValuesToDescription(validConnectorTypes),
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

	// inbound only
	"stream_id": {
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: validation.IsUUID,
		Description:  "Only required for inbound connectors.",
	},
	"destination": {
		Type:        schema.TypeList,
		Computed:    true,
		Description: "Only used for inbound connectors.",
		Elem: &schema.Resource{
			Schema: connTargetSchema,
		},
	},
	// outbound only
	"command_queue_id": {
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: validation.IsUUID,
		Description:  "Only required for outbound connectors.",
	},
	"source": {
		Type:        schema.TypeList,
		Computed:    true,
		Description: "Only used for outbound connectors.",
		Elem: &schema.Resource{
			Schema: connTargetSchema,
		},
	},
}

var socketSchema = map[string]*schema.Schema{
	"type": {
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validation.StringInSlice(validProtocolTypes, false),
		Description:  helper.AllowedValuesToDescription(validProtocolTypes),
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
		Computed: true,
	},
}
