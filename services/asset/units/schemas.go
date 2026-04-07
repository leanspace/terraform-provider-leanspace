package units

import (
	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

var unitSchema = map[string]resourceschema.Attribute{
	"id": resourceschema.StringAttribute{
		Computed: true,
	},
	"display_name": resourceschema.StringAttribute{
		Required: true,
	},
	"symbol": resourceschema.StringAttribute{
		Required: true,
	},
}
