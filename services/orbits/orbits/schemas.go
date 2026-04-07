package orbits

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	datasourceschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

var validIdealOrbitTypes = []string{"SSO", "POLAR", "LEO", "GEO", "MEO", "OTHER"}

var orbitSchema = map[string]resourceschema.Attribute{
	"id": resourceschema.StringAttribute{
		Computed: true,
	},
	"satellite_id": resourceschema.StringAttribute{
		Required:      true,
		Validators:    helper.ValidUUID(),
		PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
	},
	"name": resourceschema.StringAttribute{
		Required: true,
	},
	"ideal_orbit": resourceschema.SingleNestedAttribute{
		Optional:   true,
		Attributes: idealOrbitSchema,
	},
	"gps_configuration": resourceschema.SingleNestedAttribute{
		Optional:   true,
		Attributes: gpsConfigurationSchema,
	},
	"satellite_configuration": resourceschema.SingleNestedAttribute{
		Optional:   true,
		Attributes: satelliteConfigurationSchema,
	},
	"tags": general_objects.KeyValuesSchema,
	"created_at": resourceschema.StringAttribute{
		Computed:    true,
		Description: "When it was created",
	},
	"created_by": resourceschema.StringAttribute{
		Computed:    true,
		Description: "Who created it",
	},
	"last_modified_at": resourceschema.StringAttribute{
		Computed:    true,
		Description: "When it was last modified",
	},
	"last_modified_by": resourceschema.StringAttribute{
		Computed:    true,
		Description: "Who modified it the last",
	},
}

var idealOrbitSchema = map[string]resourceschema.Attribute{
	"type": resourceschema.StringAttribute{
		Required:    true,
		Description: helper.AllowedValuesToDescription(validIdealOrbitTypes),
		Validators:  []validator.String{stringvalidator.OneOf(validIdealOrbitTypes...)},
	},
	"inclination": resourceschema.Float64Attribute{
		Required:   true,
		Validators: []validator.Float64{float64validator.Between(0.0, 180.0)},
	},
	"right_ascension_of_ascending_node": resourceschema.Float64Attribute{
		Required:   true,
		Validators: []validator.Float64{float64validator.Between(0.0, 360.0)},
	},
	"argument_of_perigee": resourceschema.Float64Attribute{
		Required:   true,
		Validators: []validator.Float64{float64validator.Between(0.0, 360.0)},
	},
	"altitude_in_meters": resourceschema.Float64Attribute{
		Required:   true,
		Validators: []validator.Float64{float64validator.AtLeast(0.0)},
	},
	"eccentricity": resourceschema.Float64Attribute{
		Required:   true,
		Validators: []validator.Float64{float64validator.AtLeast(0.0)},
	},
	"perigee_altitude_in_meters": resourceschema.Float64Attribute{
		Computed: true,
	},
	"apogee_altitude_in_meters": resourceschema.Float64Attribute{
		Computed: true,
	},
	"semi_major_axis": resourceschema.Float64Attribute{
		Computed: true,
	},
}

var gpsConfigurationSchema = map[string]resourceschema.Attribute{
	"gps_metrics": resourceschema.SingleNestedAttribute{
		Optional:   true,
		Attributes: gpsMetricsSchema,
	},
	"standard_deviations": resourceschema.SingleNestedAttribute{
		Optional:   true,
		Attributes: standardDeviationsSchema,
	},
}

var gpsMetricsSchema = map[string]resourceschema.Attribute{
	"metric_id_for_latitude": resourceschema.StringAttribute{
		Required:   true,
		Validators: helper.ValidUUID(),
	},
	"metric_id_for_longitude": resourceschema.StringAttribute{
		Required:   true,
		Validators: helper.ValidUUID(),
	},
	"metric_id_for_altitude": resourceschema.StringAttribute{
		Required:   true,
		Validators: helper.ValidUUID(),
	},
	"metric_id_for_ground_speed": resourceschema.StringAttribute{
		Required:   true,
		Validators: helper.ValidUUID(),
	},
}

var standardDeviationsSchema = map[string]resourceschema.Attribute{
	"latitude": resourceschema.Float64Attribute{
		Required:   true,
		Validators: []validator.Float64{float64validator.AtLeast(0.01)},
	},
	"longitude": resourceschema.Float64Attribute{
		Required:   true,
		Validators: []validator.Float64{float64validator.AtLeast(0.01)},
	},
	"altitude": resourceschema.Float64Attribute{
		Required:   true,
		Validators: []validator.Float64{float64validator.AtLeast(0.0)},
	},
	"ground_speed": resourceschema.Float64Attribute{
		Required:   true,
		Validators: []validator.Float64{float64validator.AtLeast(0.0)},
	},
}

var satelliteConfigurationSchema = map[string]resourceschema.Attribute{
	"drag_cross_section": resourceschema.Float64Attribute{
		Required:   true,
		Validators: []validator.Float64{float64validator.AtLeast(0.01)},
	},
	"radiation_cross_section": resourceschema.Float64Attribute{
		Required:   true,
		Validators: []validator.Float64{float64validator.AtLeast(0.01)},
	},
}

var dataSourceFilterSchema = map[string]datasourceschema.Attribute{
	"satellite_ids": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
	},
	"tags": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
	},
	"created_bys": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Description: "Filter on the user who created the entry. If you have no wish to use this field as a filter, either provide a null value or remove the field.",
	},
	"last_modified_bys": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Description: "Filter on the user who last modified the entry. If you have no wish to use this field as a filter, either provide a null value or remove the field.",
	},
	"from_created_at": datasourceschema.StringAttribute{
		Optional:    true,
		Description: "Filter on the creation date. Entries with a creation date greater or equals than the filter value will be selected (if they are not excluded by other filters). If you have no wish to use this field as a filter, either provide a null value or remove the field.",
	},
	"from_last_modified_at": datasourceschema.StringAttribute{
		Optional:    true,
		Description: "Filter on the last modification date. Entries with a last modification date greater or equals than the filter value will be selected (if they are not excluded by other filters). If you have no wish to use this field as a filter, either provide a null value or remove the field.",
	},
	"to_created_at": datasourceschema.StringAttribute{
		Optional:    true,
		Description: "Filter on the creation date. Entries with a creation date lower or equals than the filter value will be selected (if they are not excluded by other filters). If you have no wish to use this field as a filter, either provide a null value or remove the field.",
	},
	"to_last_modified_at": datasourceschema.StringAttribute{
		Optional:    true,
		Description: "Filter on the last modification date. Entries with a last modification date lower or equals than the filter value will be selected (if they are not excluded by other filters). If you have no wish to use this field as a filter, either provide a null value or remove the field.",
	},
}
