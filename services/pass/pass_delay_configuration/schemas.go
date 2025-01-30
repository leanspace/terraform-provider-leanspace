package pass_delay_configuration

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var passDelayConfigurationSchema = map[string]*schema.Schema{
	"id": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"name": {
		Type:     schema.TypeString,
		Required: true,
	},
	"aos_delay_in_millisecond": {
		Type:     schema.TypeFloat,
		Required: true,
	},
	"los_delay_in_millisecond": {
		Type:     schema.TypeFloat,
		Required: true,
	},
}
