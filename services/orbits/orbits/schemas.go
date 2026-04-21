package orbits

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
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

var orbitSchema = general_objects.ResourceSchemaWith(map[string]resourceschema.Attribute{
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
})

var idealOrbitSchema = map[string]resourceschema.Attribute{
	"type": resourceschema.StringAttribute{
		Optional:    true,
		Description: helper.AllowedValuesToDescription(validIdealOrbitTypes),
		Validators:  []validator.String{stringvalidator.OneOf(validIdealOrbitTypes...), helper.RequiredIfParentConfigured()},
	},
	"inclination": resourceschema.Float64Attribute{
		Optional:   true,
		Validators: []validator.Float64{float64validator.Between(0.0, 180.0), helper.RequiredFloat64IfParentConfigured()},
	},
	"right_ascension_of_ascending_node": resourceschema.Float64Attribute{
		Optional:   true,
		Validators: []validator.Float64{float64validator.Between(0.0, 360.0), helper.RequiredFloat64IfParentConfigured()},
	},
	"argument_of_perigee": resourceschema.Float64Attribute{
		Optional:   true,
		Validators: []validator.Float64{float64validator.Between(0.0, 360.0), helper.RequiredFloat64IfParentConfigured()},
	},
	"altitude_in_meters": resourceschema.Float64Attribute{
		Optional:   true,
		Validators: []validator.Float64{float64validator.AtLeast(0.0), helper.RequiredFloat64IfParentConfigured()},
	},
	"eccentricity": resourceschema.Float64Attribute{
		Optional:   true,
		Validators: []validator.Float64{float64validator.AtLeast(0.0), helper.RequiredFloat64IfParentConfigured()},
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
		Optional:   true,
		Validators: append(helper.ValidUUID(), helper.RequiredIfParentConfigured()),
	},
	"metric_id_for_longitude": resourceschema.StringAttribute{
		Optional:   true,
		Validators: append(helper.ValidUUID(), helper.RequiredIfParentConfigured()),
	},
	"metric_id_for_altitude": resourceschema.StringAttribute{
		Optional:   true,
		Validators: append(helper.ValidUUID(), helper.RequiredIfParentConfigured()),
	},
	"metric_id_for_ground_speed": resourceschema.StringAttribute{
		Optional:   true,
		Validators: append(helper.ValidUUID(), helper.RequiredIfParentConfigured()),
	},
}

var standardDeviationsSchema = map[string]resourceschema.Attribute{
	"latitude": resourceschema.Float64Attribute{
		Optional:   true,
		Validators: []validator.Float64{float64validator.AtLeast(0.01), helper.RequiredFloat64IfParentConfigured()},
	},
	"longitude": resourceschema.Float64Attribute{
		Optional:   true,
		Validators: []validator.Float64{float64validator.AtLeast(0.01), helper.RequiredFloat64IfParentConfigured()},
	},
	"altitude": resourceschema.Float64Attribute{
		Optional:   true,
		Validators: []validator.Float64{float64validator.AtLeast(0.0), helper.RequiredFloat64IfParentConfigured()},
	},
	"ground_speed": resourceschema.Float64Attribute{
		Optional:   true,
		Validators: []validator.Float64{float64validator.AtLeast(0.0), helper.RequiredFloat64IfParentConfigured()},
	},
}

var satelliteConfigurationSchema = map[string]resourceschema.Attribute{
	"drag_cross_section": resourceschema.Float64Attribute{
		Optional:   true,
		Validators: []validator.Float64{float64validator.AtLeast(0.01), helper.RequiredFloat64IfParentConfigured()},
	},
	"radiation_cross_section": resourceschema.Float64Attribute{
		Optional:   true,
		Validators: []validator.Float64{float64validator.AtLeast(0.01), helper.RequiredFloat64IfParentConfigured()},
	},
}

var dataSourceFilterSchema = general_objects.AuditFilterFieldsWithTags(map[string]datasourceschema.Attribute{
	"satellite_ids": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators:  []validator.List{listvalidator.ValueStringsAre(helper.ValidUUID()...)},
	},
	"with_gps_metrics": datasourceschema.BoolAttribute{
		Optional: true,
	},
	"with_standard_deviations": datasourceschema.BoolAttribute{
		Optional: true,
	},
	"with_satellite_configuration": datasourceschema.BoolAttribute{
		Optional: true,
	},
})
