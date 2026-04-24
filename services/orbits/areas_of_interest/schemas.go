package areas_of_interest

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	datasourceschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64default"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

var validShapeTypes = []string{
	"POINT", "CIRCLE", "POLYGON",
}

var areaOfInterestSchema = general_objects.ResourceSchemaWith(map[string]resourceschema.Attribute{
	"name": resourceschema.StringAttribute{
		Required:   true,
		Validators: helper.ValidName(),
	},
	"shape": resourceschema.SingleNestedAttribute{
		Required:   true,
		Attributes: areaOfInterestShapeSchema,
	},
	"tags": general_objects.KeyValuesSchema,
})

var areaOfInterestShapeSchema = map[string]resourceschema.Attribute{
	"type": resourceschema.StringAttribute{
		Required:   true,
		Validators: []validator.String{stringvalidator.OneOf(validShapeTypes...)},
	},
	"geolocation": resourceschema.SingleNestedAttribute{
		Optional:   true,
		Attributes: geoPointSchema,
	},
	"center_geolocation": resourceschema.SingleNestedAttribute{
		Optional:   true,
		Attributes: geoPointSchema,
	},
	"radius_in_meters": resourceschema.Float64Attribute{
		Optional: true,
	},
	"vertices_geolocation": resourceschema.SetNestedAttribute{
		Optional: true,
		NestedObject: resourceschema.NestedAttributeObject{
			Attributes: geoPointSchema,
		},
	},
}

var geoPointSchema = map[string]resourceschema.Attribute{
	"latitude": resourceschema.Float64Attribute{
		Optional: true,
		Validators: []validator.Float64{
			float64validator.Between(-90.0, 90.0),
			helper.RequiredFloat64IfParentConfigured(),
		},
	},
	"longitude": resourceschema.Float64Attribute{
		Optional: true,
		Validators: []validator.Float64{
			float64validator.Between(-180.0, 180),
			helper.RequiredFloat64IfParentConfigured(),
		},
	},
	"altitude": resourceschema.Float64Attribute{
		Optional:   true,
		Computed:   true,
		Default:    float64default.StaticFloat64(0.0),
		Validators: []validator.Float64{float64validator.AtLeast(0)},
	},
}

var dataSourceFilterSchema = general_objects.AuditFilterFieldsWithTags(map[string]datasourceschema.Attribute{
	"types": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators:  []validator.List{listvalidator.ValueStringsAre(stringvalidator.OneOf(validShapeTypes...))},
	},
})
