package action_templates

import (
	"github.com/leanspace/terraform-provider-leanspace/helper"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var validTypes = []string{"WEBHOOK", "LEANSPACE_EVENT"}

var ValidTriggeredOn = []string{
	"TRIGGERED",
	"OK",
}

var baseActionTemplateSchema = MakeActionTemplateSchema(false)

func MakeActionTemplateSchema(includeTriggeredOn bool) map[string]*schema.Schema {
	baseSchema := map[string]*schema.Schema{
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
			Description:  helper.AllowedValuesToDescription(validTypes),
		},
		"url": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.IsURLWithHTTPorHTTPS,
		},
		"payload": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"content": {
			Type:     schema.TypeString,
			Computed: true,
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

	if includeTriggeredOn {
		baseSchema["triggered_on"] = &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice(ValidTriggeredOn, false),
				Description:  helper.AllowedValuesToDescription(ValidTriggeredOn),
			},
		}
	}

	return baseSchema
}

var dataSourceFilterSchema = map[string]*schema.Schema{
	"types": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.StringInSlice(validTypes, false),
			Description:  helper.AllowedValuesToDescription(validTypes),
		},
	},
	"monitor_ids": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.IsUUID,
		},
	},
}
