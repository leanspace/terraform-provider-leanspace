package contact_reservations

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var contactReservationSchema = map[string]*schema.Schema{
	"id": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"contact_state_id": {
		Type:     schema.TypeString,
		Required: true,
	},
	"leafspace_status": {
		Type:     schema.TypeString,
		Required: true,
	},
	"status": {
		Type:     schema.TypeString,
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
	"leafspace_statuses": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	},
}
