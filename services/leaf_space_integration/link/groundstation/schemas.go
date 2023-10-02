package leaf_space_groundstation_links

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var leafSpaceGroundStationLink = map[string]*schema.Schema{
	"id": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"leafspace_ground_station_id": {
		Type:     schema.TypeString,
		Required: true,
		ForceNew: true,
	},
	"leafspace_ground_station_name": {
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},
	"leanspace_ground_station_id": {
		Type:     schema.TypeString,
		Required: true,
		ForceNew: true,
	},
	"leanspace_ground_station_name": {
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},
}

var dataSourceFilterSchema = map[string]*schema.Schema{
	"leafspace_ground_station_ids": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Description: "list of the leafspace ground station ids",
	},

	"leanspace_ground_station_ids": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.IsUUID,
		},
		Description: "list of the leanspace ground station ids",
	},
}
