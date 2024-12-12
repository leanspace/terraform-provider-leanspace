package orbits

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

var validIdealOrbitTypes = []string{"SSO", "POLAR", "LEO", "GEO", "MEO", "OTHER"}

var orbitSchema = map[string]*schema.Schema{
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
	"ideal_orbit": {
		Type:     schema.TypeList,
		Optional: true,
		MinItems: 1,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: idealOrbitSchema,
		},
	},
	"gps_configuration": {
		Type:     schema.TypeList,
		Optional: true,
		MinItems: 1,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: gpsConfigurationSchema,
		},
	},
	"satellite_configuration": {
		Type:     schema.TypeList,
		Optional: true,
		MinItems: 1,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: satelliteConfigurationSchema,
		},
	},
	"tags": general_objects.KeyValuesSchema,
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

var idealOrbitSchema = map[string]*schema.Schema{
	"type": {
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validation.StringInSlice(validIdealOrbitTypes, false),
		Description:  helper.AllowedValuesToDescription(validIdealOrbitTypes),
	},
	"inclination": {
		Type:         schema.TypeFloat,
		Required:     true,
		ValidateFunc: validation.FloatBetween(0.0, 180.0),
	},
	"right_ascension_of_ascending_node": {
		Type:         schema.TypeFloat,
		Required:     true,
		ValidateFunc: validation.FloatBetween(0.0, 360.0),
	},
	"argument_of_perigee": {
		Type:         schema.TypeFloat,
		Required:     true,
		ValidateFunc: validation.FloatBetween(0.0, 360.0),
	},
	"altitude_in_meters": {
		Type:         schema.TypeFloat,
		Required:     true,
		ValidateFunc: validation.FloatAtLeast(0.0),
	},
	"eccentricity": {
		Type:         schema.TypeFloat,
		Required:     true,
		ValidateFunc: helper.FloatAtLeastAndLessThan(0.0, 1.0),
	},
	"perigee_altitude_in_meters": {
		Type:     schema.TypeFloat,
		Computed: true,
	},
	"apogee_altitude_in_meters": {
		Type:     schema.TypeFloat,
		Computed: true,
	},
	"semi_major_axis": {
		Type:     schema.TypeFloat,
		Computed: true,
	},
}

var gpsConfigurationSchema = map[string]*schema.Schema{
	"gps_metrics": {
		Type:     schema.TypeList,
		Optional: true,
		MinItems: 1,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: gpsMetricsSchema,
		},
	},
	"standard_deviations": {
		Type:     schema.TypeList,
		Optional: true,
		MinItems: 1,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: standardDeviationsSchema,
		},
	},
}

var gpsMetricsSchema = map[string]*schema.Schema{
	"metric_id_for_latitude": {
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validation.IsUUID,
	},
	"metric_id_for_longitude": {
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validation.IsUUID,
	},
	"metric_id_for_altitude": {
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validation.IsUUID,
	},
	"metric_id_for_ground_speed": {
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validation.IsUUID,
	},
}

var standardDeviationsSchema = map[string]*schema.Schema{
	"latitude": {
		Type:         schema.TypeFloat,
		Required:     true,
		ValidateFunc: helper.FloatAtLeastAndLessThan(0.1, 1.0),
	},
	"longitude": {
		Type:         schema.TypeFloat,
		Required:     true,
		ValidateFunc: helper.FloatAtLeastAndLessThan(0.1, 1.0),
	},
	"altitude": {
		Type:         schema.TypeFloat,
		Required:     true,
		ValidateFunc: helper.FloatAtLeastAndLessThan(0.0, 5000.1),
	},
	"ground_speed": {
		Type:         schema.TypeFloat,
		Required:     true,
		ValidateFunc: helper.FloatAtLeastAndLessThan(0.0, 5000.1),
	},
}

var satelliteConfigurationSchema = map[string]*schema.Schema{
	"drag_cross_section": {
		Type:         schema.TypeFloat,
		Required:     true,
		ValidateFunc: validation.FloatAtLeast(0.01),
	},
	"radiation_cross_section": {
		Type:         schema.TypeFloat,
		Required:     true,
		ValidateFunc: validation.FloatAtLeast(0.01),
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
	"tags": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	},
	"created_bys": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.IsUUID,
		},
		Description: "Filter on the user who created the Orbit. If you have no wish to use this field as a filter, either provide a null value or remove the field.",
	},
	"last_modified_bys": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.IsUUID,
		},
		Description: "Filter on the user who last modified the Orbit. If you have no wish to use this field as a filter, either provide a null value or remove the field.",
	},
	"from_created_at": {
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: helper.IsValidTimeDateOrTimestamp,
		Description:  "Filter on the Orbit creation date. Orbits with a creation date greater or equals than the filter value will be selected (if they are not excluded by other filters). If you have no wish to use this field as a filter, either provide a null value or remove the field.",
	},
	"from_last_modified_at": {
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: helper.IsValidTimeDateOrTimestamp,
		Description:  "Filter on the Orbit last modification date. Orbits with a last modification date greater or equals than the filter value will be selected (if they are not excluded by other filters). If you have no wish to use this field as a filter, either provide a null value or remove the field.",
	},
	"to_created_at": {
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: helper.IsValidTimeDateOrTimestamp,
		Description:  "Filter on the Orbit creation date. Orbits with a creation date lower or equals than the filter value will be selected (if they are not excluded by other filters). If you have no wish to use this field as a filter, either provide a null value or remove the field.",
	},
	"to_last_modified_at": {
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: helper.IsValidTimeDateOrTimestamp,
		Description:  "Filter on the Orbit last modification date. Orbits with a last modification date lower or equals than the filter value will be selected (if they are not excluded by other filters). If you have no wish to use this field as a filter, either provide a null value or remove the field.",
	},
}
