package units

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var unitSchema = map[string]*schema.Schema{
	"id": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"display_name": {
		Type:     schema.TypeString,
		Required: true,
	},
	"symbol": {
		Type:     schema.TypeString,
		Required: true,
	},
}
