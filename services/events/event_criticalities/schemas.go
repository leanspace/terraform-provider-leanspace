package event_criticalities

import (
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var EventsCriticalitiesSchema = map[string]*schema.Schema{
	"id": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"name": {
		Type:     schema.TypeString,
		Required: true,
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

var dataSourceFilterSchema = map[string]*schema.Schema{
	"tags": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	},
}
