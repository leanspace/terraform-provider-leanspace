package action_templates

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var validTypes = []string{"WEBHOOK"}

var ActionTemplateSchema = map[string]*schema.Schema{
	"id": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"name": {
		Type:     schema.TypeString,
		Required: true,
	},
	"type": {
		Type:         schema.TypeString,
		Optional:     true,
		Default:      "WEBHOOK",
		ValidateFunc: validation.StringInSlice(validTypes, false),
	},
	"url": {
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validation.IsURLWithHTTPorHTTPS,
	},
	"payload": {
		Type:     schema.TypeString,
		Required: true,
	},
	"headers": {
		Type:     schema.TypeMap,
		Optional: true,
		Default:  make(map[string]string),
		Elem: &schema.Schema{
			Type: schema.TypeString,
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
