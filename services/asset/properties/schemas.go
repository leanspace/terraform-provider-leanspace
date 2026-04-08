package properties

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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

var validPropertyTypes = []string{"NUMERIC", "ENUM", "TEXT", "TIMESTAMP", "DATE", "TIME", "BOOLEAN", "GEOPOINT", "TLE"}
var validNodeTypes = []string{"ASSET", "GROUP", "COMPONENT"}
var validNodeKinds = []string{"GENERIC", "SATELLITE", "GROUND_STATION"}

var propertySchema = general_objects.ResourceSchemaWith(map[string]resourceschema.Attribute{
	"name": resourceschema.StringAttribute{
		Required: true,
	},
	"description": resourceschema.StringAttribute{
		Optional: true,
	},
	"node_id": resourceschema.StringAttribute{
		Required:      true,
		Validators:    helper.ValidUUID(),
		PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
	},
	"tags": general_objects.KeyValuesSchema,
	"min_length": resourceschema.Int64Attribute{
		Optional:    true,
		Description: "Text only: Minimum length of this text (at least 1)",
		Validators:  []validator.Int64{int64validator.AtLeast(1)},
	},
	"max_length": resourceschema.Int64Attribute{
		Optional:    true,
		Description: "Text only: Maximum length of this text (at least 1)",
		Validators:  []validator.Int64{int64validator.AtLeast(1)},
	},
	"pattern": resourceschema.StringAttribute{
		Optional:    true,
		Description: "Text only: Regex defined the allowed pattern of this text",
	},
	"before": resourceschema.StringAttribute{
		Optional:    true,
		Description: "Time/date/timestamp only: Maximum date allowed",
		Validators:  helper.IsValidTimeDateOrTimestamp(),
	},
	"after": resourceschema.StringAttribute{
		Optional:    true,
		Description: "Time/date/timestamp only: Minimum date allowed",
		Validators:  helper.IsValidTimeDateOrTimestamp(),
	},
	"fields": resourceschema.SingleNestedAttribute{
		Optional:    true,
		Description: "Geopoint only",
		Attributes:  geoPointFieldsSchema,
	},
	"options": resourceschema.MapAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Description: "Enum only: The allowed values for the enum in the format 1 = \"value\"",
	},
	"min": resourceschema.Float64Attribute{
		Optional:    true,
		Description: "Numeric only",
	},
	"max": resourceschema.Float64Attribute{
		Optional:    true,
		Description: "Numeric only",
	},
	"scale": resourceschema.Int64Attribute{
		Optional:    true,
		Description: "Numeric only",
	},
	"precision": resourceschema.Int64Attribute{
		Optional:    true,
		Description: "Numeric only: How many values after the comma should be accepted",
	},
	"unit_id": resourceschema.StringAttribute{
		Optional:    true,
		Description: "Numeric only",
		Validators:  helper.ValidUUID(),
	},
	"value": resourceschema.StringAttribute{
		Optional: true,
	},
	"type": resourceschema.StringAttribute{
		Required:      true,
		Description:   helper.AllowedValuesToDescription(validPropertyTypes),
		Validators:    []validator.String{stringvalidator.OneOf(validPropertyTypes...)},
		PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
	},
	"built_in": resourceschema.BoolAttribute{
		Computed:    true,
		Description: "Indicates if it is a build-in property.",
	},
})

var geoPointFieldsSchema = map[string]resourceschema.Attribute{
	"latitude": resourceschema.SingleNestedAttribute{
		Required:   true,
		Attributes: propertyFieldSchema(true),
	},
	"longitude": resourceschema.SingleNestedAttribute{
		Required:   true,
		Attributes: propertyFieldSchema(true),
	},
	"elevation": resourceschema.SingleNestedAttribute{
		Required:   true,
		Attributes: propertyFieldSchema(false),
	},
}

func propertyFieldSchema(computedMinMax bool) map[string]resourceschema.Attribute {
	return map[string]resourceschema.Attribute{
		"value": resourceschema.StringAttribute{
			Optional: true,
		},

		// Numeric only
		"scale": resourceschema.Int64Attribute{
			Optional:    true,
			Description: "Property field with numeric type only: the scale required.",
		},
		"unit_id": resourceschema.StringAttribute{
			Optional:    true,
			Description: "Property field with numeric type only",
			Validators:  helper.ValidUUID(),
		},
		"min": resourceschema.Float64Attribute{
			Computed:    computedMinMax,
			Optional:    !computedMinMax,
			Description: "Property field with numeric type only: the minimum value allowed.",
		},
		"precision": resourceschema.Int64Attribute{
			Optional:    true,
			Description: "Property field with numeric type only: How many values after the comma should be accepted",
		},
		"max": resourceschema.Float64Attribute{
			Computed:    computedMinMax,
			Optional:    !computedMinMax,
			Description: "Property field with numeric type only: the maximum value allowed.",
		},
	}
}

var dataSourceFilterSchema = map[string]datasourceschema.Attribute{
	"category": datasourceschema.StringAttribute{
		Optional:    true,
		Description: "Allowed values : BUILT_IN_PROPERTIES_ONLY, USER_PROPERTIES_ONLY, ALL_PROPERTIES",
	},
	"created_by": datasourceschema.StringAttribute{
		Optional:    true,
		Description: "Filter on the user who created the Property. If you have no wish to use this field as a filter, either provide a null value or remove the field.",
	},
	"from_created_at": datasourceschema.StringAttribute{
		Optional:    true,
		Description: "Filter on the Property creation date. Properties with a creation date greater or equals than the filter value will be selected (if they are not excluded by other filters). If you have no wish to use this field as a filter, either provide a null value or remove the field.",
	},
	"from_last_modified_at": datasourceschema.StringAttribute{
		Optional:    true,
		Description: "Filter on the Property last modification date. Properties with a last modification date greater or equals than the filter value will be selected (if they are not excluded by other filters). If you have no wish to use this field as a filter, either provide a null value or remove the field.",
	},
	"last_modified_by": datasourceschema.StringAttribute{
		Optional:    true,
		Description: "Filter on the user who modified last the Property. If you have no wish to use this field as a filter, either provide a null value or remove the field.",
	},
	"to_created_at": datasourceschema.StringAttribute{
		Optional:    true,
		Description: "Filter on the Property creation date. Properties with a creation date lower or equals than the filter value will be selected (if they are not excluded by other filters). If you have no wish to use this field as a filter, either provide a null value or remove the field.",
	},
	"to_last_modified_at": datasourceschema.StringAttribute{
		Optional:    true,
		Description: "Filter on the Property last modification date. Properties with a last modification date lower or equals than the filter value will be selected (if they are not excluded by other filters). If you have no wish to use this field as a filter, either provide a null value or remove the field.",
	},
	"kinds": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Description: "Allowed values : GENERIC, SATELLITE, GROUND_STATION",
	},
	"node_ids": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Description: "Only returns node whose id matches one of the provided values",
	},
	"node_types": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
	},
	"tags": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
	},
}
