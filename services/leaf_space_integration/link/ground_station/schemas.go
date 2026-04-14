package groundstation_links

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	datasourceschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/leanspace/terraform-provider-leanspace/helper"
)

var leafSpaceGroundStationLink = map[string]resourceschema.Attribute{
	"id": resourceschema.StringAttribute{
		Computed: true,
	},
	"leafspace_ground_station_id": resourceschema.StringAttribute{
		Required:      true,
		PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
	},
	"leafspace_ground_station_name": resourceschema.StringAttribute{
		Optional: true,
		Computed: true,
	},
	"leanspace_ground_station_id": resourceschema.StringAttribute{
		Required:      true,
		PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
	},
	"leanspace_ground_station_name": resourceschema.StringAttribute{
		Optional: true,
		Computed: true,
	},
}

var dataSourceFilterSchema = map[string]datasourceschema.Attribute{
	"leafspace_ground_station_ids": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Description: "list of the leafspace ground station ids",
	},

	"leanspace_ground_station_ids": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Description: "list of the leanspace ground station ids",
		Validators:  []validator.List{listvalidator.ValueStringsAre(helper.ValidUUID()...)},
	},
}
