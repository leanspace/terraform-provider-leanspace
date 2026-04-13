package widgets

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	datasourceschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

var ValidWidgetTypes = []string{"TABLE", "LINE", "BAR", "AREA", "VALUE", "RESOURCES", "EARTH", "GAUGE", "ENUM", "ORBITAL_VIEW"}
var validGranularities = []string{"second", "minute", "hour", "day", "week", "month", "raw"}
var validDatasources = []string{"metric", "raw_stream", "resources", "topology", "orbits", "ground_stations", "areas_of_interest"}
var validAggregations = []string{"avg", "count", "sum", "min", "max", "none"}
var validFilterOperators = []string{"gt", "lt", "equals", "notEquals"}
var validTimeDimensions = []string{"timestamp", "received_at", "ingested_at"}

var colorRegex = regexp.MustCompile(`^#(?:[0-9a-fA-F]{3}){1,2}$`)

var widgetSchema = general_objects.ResourceSchemaWith(map[string]resourceschema.Attribute{
	"name": resourceschema.StringAttribute{
		Required: true,
	},
	"description": resourceschema.StringAttribute{
		Optional: true,
	},
	"type": resourceschema.StringAttribute{
		Required:    true,
		Description: helper.AllowedValuesToDescription(ValidWidgetTypes),
		Validators:  []validator.String{stringvalidator.OneOf(ValidWidgetTypes...)},
	},
	"granularity": resourceschema.StringAttribute{
		Required:    true,
		Description: helper.AllowedValuesToDescription(validGranularities),
		Validators:  []validator.String{stringvalidator.OneOf(validGranularities...)},
	},
	"query_time_dimension": resourceschema.StringAttribute{
		Optional:    true,
		Computed:    true,
		Default:     stringdefault.StaticString("timestamp"),
		Description: helper.AllowedValuesToDescription(validTimeDimensions),
		Validators:  []validator.String{stringvalidator.OneOf(validTimeDimensions...)},
	},
	"display_time_dimension": resourceschema.StringAttribute{
		Optional:    true,
		Computed:    true,
		Default:     stringdefault.StaticString("timestamp"),
		Description: helper.AllowedValuesToDescription(validTimeDimensions),
		Validators:  []validator.String{stringvalidator.OneOf(validTimeDimensions...)},
	},
	"series": resourceschema.ListNestedAttribute{
		Required: true,
		NestedObject: resourceschema.NestedAttributeObject{
			Attributes: seriesSchema,
		},
	},
	"metadata": resourceschema.SingleNestedAttribute{
		Optional:   true,
		Attributes: metadataSchema,
	},
	"dashboards": resourceschema.SetNestedAttribute{
		Computed: true,
		NestedObject: resourceschema.NestedAttributeObject{
			Attributes: dashboardInfoSchema,
		},
	},
	"tags": general_objects.KeyValuesSchema,
})

var seriesSchema = map[string]resourceschema.Attribute{
	"id": resourceschema.StringAttribute{
		Required: true,
	},
	"name": resourceschema.StringAttribute{
		Optional:    true,
		Description: "The datasource's name",
	},
	"datasource": resourceschema.StringAttribute{
		Required:    true,
		Description: helper.AllowedValuesToDescription(validDatasources),
		Validators:  []validator.String{stringvalidator.OneOf(validDatasources...)},
	},
	"aggregation": resourceschema.StringAttribute{
		Required:    true,
		Description: helper.AllowedValuesToDescription(validAggregations),
		Validators:  []validator.String{stringvalidator.OneOf(validAggregations...)},
	},
	"filters": resourceschema.ListNestedAttribute{ // workaround: we need a list instead of a set
		Optional: true,
		NestedObject: resourceschema.NestedAttributeObject{
			Attributes: filterSchema,
		},
		Validators: []validator.List{listvalidator.SizeAtMost(3)},
	},
}

var filterSchema = map[string]resourceschema.Attribute{
	"filter_by": resourceschema.StringAttribute{
		Required: true,
	},
	"operator": resourceschema.StringAttribute{
		Required:    true,
		Description: helper.AllowedValuesToDescription(validFilterOperators),
	},
	"value": resourceschema.StringAttribute{
		Required: true,
	},
}

var metadataSchema = map[string]resourceschema.Attribute{
	"y_axis_label": resourceschema.StringAttribute{
		Optional: true,
	},
	"y_axis_range_min": resourceschema.Float64Attribute{
		Optional:    true,
		Description: "The minimum value for the widget's Y axis.",
	},
	"y_axis_range_max": resourceschema.Float64Attribute{
		Optional:    true,
		Description: "The maximum value for the widget's Y axis.",
	},
	"thresholds": resourceschema.ListNestedAttribute{
		Optional:    true,
		Description: "The threshold applies only to the Gauge widget.",
		NestedObject: resourceschema.NestedAttributeObject{
			Attributes: thresholdSchema,
		},
		Validators: []validator.List{listvalidator.SizeBetween(1, 10)},
	},
}

/*
From and to are strings so that they can be nil.
*/
var thresholdSchema = map[string]resourceschema.Attribute{
	"from": resourceschema.StringAttribute{
		Optional: true,
	},
	"to": resourceschema.StringAttribute{
		Optional: true,
	},
	"color": resourceschema.StringAttribute{
		Required:   true,
		Validators: []validator.String{stringvalidator.RegexMatches(colorRegex, "Must be a valid hex color")},
	},
}

var dashboardInfoSchema = map[string]resourceschema.Attribute{
	"id": resourceschema.StringAttribute{
		Computed: true,
	},
	"name": resourceschema.StringAttribute{
		Computed: true,
	},
}

var dataSourceFilterSchema = map[string]datasourceschema.Attribute{
	"types": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators:  []validator.List{listvalidator.ValueStringsAre(stringvalidator.OneOf(ValidWidgetTypes...))},
	},
	"tags": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
	},
	"dashboard_ids": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
	},
	"datasource_ids": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
	},
	"datasources": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
	},
}
