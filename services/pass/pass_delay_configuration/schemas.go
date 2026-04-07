package pass_delay_configuration

import (
	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

var passDelayConfigurationSchema = map[string]resourceschema.Attribute{
	"id": resourceschema.StringAttribute{
		Computed: true,
	},
	"name": resourceschema.StringAttribute{
		Required: true,
	},
	"aos_delay_in_millisecond": resourceschema.Float64Attribute{
		Required: true,
	},
	"los_delay_in_millisecond": resourceschema.Float64Attribute{
		Required: true,
	},
}
