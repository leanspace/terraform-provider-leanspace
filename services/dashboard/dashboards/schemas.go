package dashboards

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	datasourceschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
	"github.com/leanspace/terraform-provider-leanspace/services/dashboard/widgets"
)

var dashboardSchema = general_objects.ResourceSchemaWith(map[string]resourceschema.Attribute{
	"name": resourceschema.StringAttribute{
		Required: true,
	},
	"description": resourceschema.StringAttribute{
		Optional: true,
	},
	"node_ids": resourceschema.SetAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators: []validator.Set{
			setvalidator.ValueStringsAre(helper.ValidUUID()...),
		},
	},
	"widget_info": resourceschema.SetNestedAttribute{
		Optional: true,
		NestedObject: resourceschema.NestedAttributeObject{
			Attributes: widgetInfoSchema,
		},
	},
	"widgets": resourceschema.SetNestedAttribute{
		Computed: true,
		NestedObject: resourceschema.NestedAttributeObject{
			Attributes: dashboardWidgetSchema,
		},
	},
	"tags": general_objects.KeyValuesSchema,
})

var widgetInfoSchema = map[string]resourceschema.Attribute{
	"id": resourceschema.StringAttribute{
		Required:   true,
		Validators: helper.ValidUUID(),
	},
	"type": resourceschema.StringAttribute{
		Computed:    true,
		Description: helper.AllowedValuesToDescription(widgets.ValidWidgetTypes),
	},
	"w": resourceschema.Int64Attribute{
		Required:   true,
		Validators: []validator.Int64{int64validator.AtLeast(1)},
	},
	"h": resourceschema.Int64Attribute{
		Required:   true,
		Validators: []validator.Int64{int64validator.AtLeast(1)},
	},
	"x": resourceschema.Int64Attribute{
		Required:   true,
		Validators: []validator.Int64{int64validator.AtLeast(0)},
	},
	"y": resourceschema.Int64Attribute{
		Required:   true,
		Validators: []validator.Int64{int64validator.AtLeast(0)},
	},
	"min_w": resourceschema.Int64Attribute{
		Optional:   true,
		Validators: []validator.Int64{int64validator.AtLeast(1)},
	},
	"min_h": resourceschema.Int64Attribute{
		Optional:   true,
		Validators: []validator.Int64{int64validator.AtLeast(1)},
	},
}

var dashboardWidgetSchema = general_objects.ResourceSchemaWith(map[string]resourceschema.Attribute{
	"name": resourceschema.StringAttribute{
		Computed: true,
	},
	"description": resourceschema.StringAttribute{
		Computed: true,
	},
	"type": resourceschema.StringAttribute{
		Computed: true,
	},
	"granularity": resourceschema.StringAttribute{
		Computed: true,
	},
	"query_time_dimension": resourceschema.StringAttribute{
		Computed: true,
	},
	"display_time_dimension": resourceschema.StringAttribute{
		Computed: true,
	},
	"series": resourceschema.ListNestedAttribute{
		Computed: true,
		NestedObject: resourceschema.NestedAttributeObject{
			Attributes: seriesSchema,
		},
	},
	"metadata": resourceschema.ListNestedAttribute{
		Computed: true,
		NestedObject: resourceschema.NestedAttributeObject{
			Attributes: metadataSchema,
		},
	},
	"view": resourceschema.ListNestedAttribute{
		Computed: true,
		NestedObject: resourceschema.NestedAttributeObject{
			Attributes: dashboardInfoSchema,
		},
	},
	"tags": general_objects.KeyValuesSchema,
})

var seriesSchema = map[string]resourceschema.Attribute{
	"id": resourceschema.StringAttribute{
		Computed: true,
	},
	"name": resourceschema.StringAttribute{
		Computed: true,
	},
	"datasource": resourceschema.StringAttribute{
		Computed: true,
	},
	"aggregation": resourceschema.StringAttribute{
		Computed: true,
	},
	"filters": resourceschema.SetNestedAttribute{
		Computed: true,
		NestedObject: resourceschema.NestedAttributeObject{
			Attributes: filterSchema,
		},
	},
}

var filterSchema = map[string]resourceschema.Attribute{
	"filter_by": resourceschema.StringAttribute{
		Computed: true,
	},
	"operator": resourceschema.StringAttribute{
		Computed: true,
	},
	"value": resourceschema.StringAttribute{
		Computed: true,
	},
}

var metadataSchema = map[string]resourceschema.Attribute{
	"y_axis_label": resourceschema.StringAttribute{
		Computed: true,
	},
	"y_axis_range_min": resourceschema.ListAttribute{
		ElementType: types.Float64Type,
		Computed:    true,
	},
	"y_axis_range_max": resourceschema.ListAttribute{
		ElementType: types.Float64Type,
		Computed:    true,
	},
	"thresholds": resourceschema.ListNestedAttribute{
		Computed:    true,
		Description: "The threshold applies only to the Gauge widget.",
		NestedObject: resourceschema.NestedAttributeObject{
			Attributes: thresholdSchema,
		},
	},
}

var thresholdSchema = map[string]resourceschema.Attribute{
	"from": resourceschema.StringAttribute{
		Computed: true,
	},
	"to": resourceschema.StringAttribute{
		Computed: true,
	},
	"color": resourceschema.StringAttribute{
		Computed: true,
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
	"node_ids": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
	},
	"widget_ids": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
	},
	"tags": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
	},
}
