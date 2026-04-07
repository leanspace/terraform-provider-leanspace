package satellite_links

import (
	datasourceschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var leafSpaceSatelliteLink = map[string]resourceschema.Attribute{
	"id": resourceschema.StringAttribute{
		Computed: true,
	},
	"leafspace_satellite_id": resourceschema.StringAttribute{
		Required:      true,
		PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
	},
	"leanspace_satellite_id": resourceschema.StringAttribute{
		Required:      true,
		PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
	},
	"leafspace_satellite_name": resourceschema.StringAttribute{
		Optional: true,
		Computed: true,
	},
	"leanspace_satellite_name": resourceschema.StringAttribute{
		Optional: true,
		Computed: true,
	},
}
var dataSourceFilterSchema = map[string]datasourceschema.Attribute{
	"leafspace_satellite_ids": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Description: "list of the leafspace ground station ids",
	},

	"leanspace_satellite_ids": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Description: "list of the leanspace ground station ids",
	},
}
