package orbit_resources

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var validDataSources = []string{"TLE_CELESTRAK", "TLE_MANUAL", "GPS_METRIC", "PROVIDED_PREDICTED", "PROVIDED_MEASURED"}

var orbitResourceSchema = map[string]*schema.Schema{
	"id": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"satellite_id": {
		Type:         schema.TypeString,
		Required:     true,
		ForceNew:     true,
		ValidateFunc: validation.IsUUID,
	},
	"name": {
		Type:     schema.TypeString,
		Required: true,
	},
	"data_source": {
		Type:     schema.TypeString,
		Required: true,
	},
	"automatic_propagation": {
		Type:     schema.TypeBool,
		Optional: true,
	},
	"gps_metric_ids": {
		Type:     schema.TypeList,
		Optional: true,
		MinItems: 1,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: gpsMetricIdsSchema,
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

var gpsMetricIdsSchema = map[string]*schema.Schema{
	"metric_id_for_position_x": {
		Type:     schema.TypeString,
		Required: true,
	},
	"metric_id_for_position_y": {
		Type:     schema.TypeString,
		Required: true,
	},
	"metric_id_for_position_z": {
		Type:     schema.TypeString,
		Required: true,
	},
	"metric_id_for_velocity_x": {
		Type:     schema.TypeString,
		Required: true,
	},
	"metric_id_for_velocity_y": {
		Type:     schema.TypeString,
		Required: true,
	},
	"metric_id_for_velocity_z": {
		Type:     schema.TypeString,
		Required: true,
	},
}

var dataSourceFilterSchema = map[string]*schema.Schema{
	"satellite_ids": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.IsUUID,
		},
	},
	"ids": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.IsUUID,
		},
	},
	"data_sources": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.StringInSlice(validDataSources, false),
		},
	},
	"automatic_propagation": {
		Type:     schema.TypeBool,
		Optional: true,
	},
	"query": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Search by name",
	},
}
