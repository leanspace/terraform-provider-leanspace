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
		Type:     schema.TypeString,
		Required: true,
		ValidateFunc: validation.StringInSlice(validIdealOrbitTypes, false),
        Description:  helper.AllowedValuesToDescription(validIdealOrbitTypes),
	},
	"inclination": {
		Type:     schema.TypeFloat,
		Required: true,
	},
	"right_ascension_of_ascending_node": {
		Type:     schema.TypeFloat,
		Required: true,
	},
	"argument_of_perigee": {
		Type:     schema.TypeFloat,
		Required: true,
	},
	"altitude_in_meters": {
		Type:     schema.TypeFloat,
		Required: true,
	},
	"eccentricity": {
		Type:     schema.TypeFloat,
		Required: true,
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
