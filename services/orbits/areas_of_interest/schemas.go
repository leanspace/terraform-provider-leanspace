package areas_of_interest

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	datasourceschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

var validShapeTypes = []string{
	"POINT", "CIRCLE", "POLYGON",
}

var areaOfInterestSchema = map[string]resourceschema.Attribute{
	"id": resourceschema.StringAttribute{
		Computed: true,
	},
	"name": resourceschema.StringAttribute{
		Required:   true,
		Validators: helper.ValidName(),
	},
	"shape": resourceschema.SingleNestedAttribute{
		Required:   true,
		Attributes: areaOfInterestShapeSchema,
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
		Required:   true,
		Validators: []validator.Float64{float64validator.Between(-90.0, 90.0)},
	},
	"longitude": resourceschema.Float64Attribute{
		Required:   true,
		Validators: []validator.Float64{float64validator.Between(-180.0, 180)},
	},
	"altitude": resourceschema.Float64Attribute{
		Optional:   true,
		Validators: []validator.Float64{float64validator.AtLeast(0)},
	},
}

var dataSourceFilterSchema = map[string]datasourceschema.Attribute{
	"types": datasourceschema.ListAttribute{
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
