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
var validPropertyCategories = []string{"BUILT_IN_PROPERTIES_ONLY", "USER_PROPERTIES_ONLY", "ALL_PROPERTIES"}

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
		Attributes:  general_objects.CreateGeoPointFieldsSchema(true),
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

var dataSourceFilterSchema = general_objects.AuditFilterFieldsWithTagsAndSingularBy(map[string]datasourceschema.Attribute{
	"category": datasourceschema.StringAttribute{
		Optional:    true,
		Description: helper.AllowedValuesToDescription(validPropertyCategories),
	},
	"node_kinds": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Description: helper.AllowedValuesToDescription(validNodeKinds),
	},
	"node_ids": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Description: "Only returns node whose id matches one of the provided values",
	},
	"node_types": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Description: helper.AllowedValuesToDescription(validNodeTypes),
	},
})
