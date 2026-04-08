package sensors

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

var validShapeTypes = []string{
	"CIRCULAR", "RECTANGULAR",
}

var sensorSchema = general_objects.ResourceSchemaWith(map[string]resourceschema.Attribute{
	"satellite_id": resourceschema.StringAttribute{
		Required:      true,
		Validators:    helper.ValidUUID(),
		PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
	},
	"name": resourceschema.StringAttribute{
		Required:   true,
		Validators: helper.ValidName(),
	},
	"aperture_shape": resourceschema.SingleNestedAttribute{
		Required:   true,
		Attributes: apertureShapeSchema,
	},
	"tags": general_objects.KeyValuesSchema,
})

var apertureShapeSchema = map[string]resourceschema.Attribute{
	"type": resourceschema.StringAttribute{
		Required:   true,
		Validators: []validator.String{stringvalidator.OneOf(validShapeTypes...)},
	},
	"aperture_center": resourceschema.SingleNestedAttribute{
		Optional:   true,
		Attributes: vector3DSchema,
	},
	"half_aperture_angle": resourceschema.SingleNestedAttribute{
		Optional:   true,
		Attributes: halfApertureAngleSchema(180),
	},
	"first_axis_vector": resourceschema.SingleNestedAttribute{
		Optional:   true,
		Attributes: vector3DSchema,
	},
	"first_axis_half_aperture_angle": resourceschema.SingleNestedAttribute{
		Optional:   true,
		Attributes: halfApertureAngleSchema(90),
	},
	"second_axis_vector": resourceschema.SingleNestedAttribute{
		Optional:   true,
		Attributes: vector3DSchema,
	},
	"second_axis_half_aperture_angle": resourceschema.SingleNestedAttribute{
		Optional:   true,
		Attributes: halfApertureAngleSchema(90),
	},
}

var vector3DSchema = map[string]resourceschema.Attribute{
	"x": resourceschema.Float64Attribute{
		Optional:   true,
		Validators: []validator.Float64{helper.RequiredFloat64IfParentConfigured()},
	},
	"y": resourceschema.Float64Attribute{
		Optional:   true,
		Validators: []validator.Float64{helper.RequiredFloat64IfParentConfigured()},
	},
	"z": resourceschema.Float64Attribute{
		Optional:   true,
		Validators: []validator.Float64{helper.RequiredFloat64IfParentConfigured()},
	},
}

func halfApertureAngleSchema(maximum float64) map[string]resourceschema.Attribute {
	return map[string]resourceschema.Attribute{
		"degrees": resourceschema.Float64Attribute{
			Optional: true,
			Validators: []validator.Float64{
				float64validator.Between(0.0, maximum),
				helper.RequiredFloat64IfParentConfigured(),
			},
		},
	}
}

var dataSourceFilterSchema = map[string]datasourceschema.Attribute{
	"satellite_ids": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators:  []validator.List{listvalidator.ValueStringsAre(helper.ValidUUID()...)},
	},
	"aperture_shape_types": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators:  []validator.List{listvalidator.ValueStringsAre(stringvalidator.OneOf(validShapeTypes...))},
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
