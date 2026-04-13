package dashboards

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
	"github.com/leanspace/terraform-provider-leanspace/services/dashboard/widgets"
)

type DashboardTF struct {
	general_objects.AuditModelTF
	Name            types.String                 `tfsdk:"name"`
	Description     types.String                 `tfsdk:"description"`
	NodeIds         []types.String               `tfsdk:"node_ids"`
	WidgetInfo      []WidgetInfoTF               `tfsdk:"widget_info"`
	Widgets         []DashboardWidgetTF          `tfsdk:"widgets"`
	Tags            []general_objects.KeyValueTF `tfsdk:"tags"`
	TimestampFormat types.String                 `tfsdk:"timestamp_format"`
}

type WidgetInfoTF struct {
	ID   types.String `tfsdk:"id"`
	Type types.String `tfsdk:"type"`
	W    types.Int64  `tfsdk:"w"`
	H    types.Int64  `tfsdk:"h"`
	X    types.Int64  `tfsdk:"x"`
	Y    types.Int64  `tfsdk:"y"`
	MinW types.Int64  `tfsdk:"min_w"`
	MinH types.Int64  `tfsdk:"min_h"`
}

type DashboardWidgetTF struct {
	general_objects.AuditModelTF
	Name                 types.String                 `tfsdk:"name"`
	Description          types.String                 `tfsdk:"description"`
	Type                 types.String                 `tfsdk:"type"`
	Granularity          types.String                 `tfsdk:"granularity"`
	QueryTimeDimension   types.String                 `tfsdk:"query_time_dimension"`
	DisplayTimeDimension types.String                 `tfsdk:"display_time_dimension"`
	Series               []DashboardSeriesTF          `tfsdk:"series"`
	Metadata             []DashboardWidgetMetadataTF  `tfsdk:"metadata"`
	View                 []DashboardViewInfoTF        `tfsdk:"view"`
	Tags                 []general_objects.KeyValueTF `tfsdk:"tags"`
}

type DashboardSeriesTF struct {
	ID          types.String            `tfsdk:"id"`
	Name        types.String            `tfsdk:"name"`
	Datasource  types.String            `tfsdk:"datasource"`
	Aggregation types.String            `tfsdk:"aggregation"`
	Filters     []DashboardFilterTF     `tfsdk:"filters"`
}

type DashboardFilterTF struct {
	FilterBy types.String `tfsdk:"filter_by"`
	Operator types.String `tfsdk:"operator"`
	Value    types.String `tfsdk:"value"`
}

type DashboardWidgetMetadataTF struct {
	YAxisLabel    types.String    `tfsdk:"y_axis_label"`
	YAxisRangeMin []types.Float64 `tfsdk:"y_axis_range_min"`
	YAxisRangeMax []types.Float64 `tfsdk:"y_axis_range_max"`
	Thresholds    []DashboardThresholdTF `tfsdk:"thresholds"`
}

type DashboardThresholdTF struct {
	From  types.String `tfsdk:"from"`
	To    types.String `tfsdk:"to"`
	Color types.String `tfsdk:"color"`
}

type DashboardViewInfoTF struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

func (x *Dashboard) ToTF() any {
	widgetInfos := make([]WidgetInfoTF, len(x.WidgetInfo))
	for i, wi := range x.WidgetInfo {
		widgetInfos[i] = WidgetInfoTF{
			ID:   helper.TFStringValue(wi.ID),
			Type: helper.TFStringValue(wi.Type),
			W:    helper.TFInt64Value(wi.W),
			H:    helper.TFInt64Value(wi.H),
			X:    helper.TFInt64Value(wi.X),
			Y:    helper.TFInt64Value(wi.Y),
			MinW: helper.TFInt64Value(wi.MinW),
			MinH: helper.TFInt64Value(wi.MinH),
		}
	}

	dashWidgets := make([]DashboardWidgetTF, len(x.Widgets))
	for i, w := range x.Widgets {
		series := make([]DashboardSeriesTF, len(w.Series))
		for j, s := range w.Series {
			filters := make([]DashboardFilterTF, len(s.Filters))
			for k, f := range s.Filters {
				filters[k] = DashboardFilterTF{
					FilterBy: helper.TFStringValue(f.FilterBy),
					Operator: helper.TFStringValue(f.Operator),
					Value:    helper.TFStringValue(f.Value),
				}
			}
			series[j] = DashboardSeriesTF{
				ID:          helper.TFStringValue(s.ID),
				Name:        helper.TFStringValue(s.Name),
				Datasource:  helper.TFStringValue(s.Datasource),
				Aggregation: helper.TFStringValue(s.Aggregation),
				Filters:     filters,
			}
		}

		var metadataSlice []DashboardWidgetMetadataTF
		hasMetadata := w.Metadata.YAxisLabel != "" || len(w.Metadata.Thresholds) > 0 ||
			(w.Metadata.YAxisRange != nil && len(w.Metadata.YAxisRange) == 2)
		if hasMetadata {
			md := DashboardWidgetMetadataTF{
				YAxisLabel: helper.TFStringValue(w.Metadata.YAxisLabel),
			}
			if w.Metadata.YAxisRange != nil && len(w.Metadata.YAxisRange) == 2 {
				if w.Metadata.YAxisRange[0] != nil {
					md.YAxisRangeMin = []types.Float64{helper.TFFloat64PtrValue(w.Metadata.YAxisRange[0])}
				}
				if w.Metadata.YAxisRange[1] != nil {
					md.YAxisRangeMax = []types.Float64{helper.TFFloat64PtrValue(w.Metadata.YAxisRange[1])}
				}
			}
			thresholds := make([]DashboardThresholdTF, len(w.Metadata.Thresholds))
			for j, t := range w.Metadata.Thresholds {
				thresholds[j] = DashboardThresholdTF{Color: helper.TFStringValue(t.Color)}
				// from/to handled as strings in the original parsers
			}
			md.Thresholds = thresholds
			metadataSlice = []DashboardWidgetMetadataTF{md}
		}

		dashWidgets[i] = DashboardWidgetTF{
			AuditModelTF:         general_objects.AuditModelToTF(&w.AuditModel),
			Name:                 helper.TFStringValue(w.Name),
			Description:          helper.TFStringValue(w.Description),
			Type:                 helper.TFStringValue(w.Type),
			Granularity:          helper.TFStringValue(w.Granularity),
			QueryTimeDimension:   helper.TFStringValue(w.QueryTimeDimension),
			DisplayTimeDimension: helper.TFStringValue(w.DisplayTimeDimension),
			Series:               series,
			Metadata:             metadataSlice,
			Tags:                 general_objects.KeyValuesToTF(w.Tags),
		}
	}

	return &DashboardTF{
		AuditModelTF:    general_objects.AuditModelToTF(&x.AuditModel),
		Name:            helper.TFStringValue(x.Name),
		Description:     helper.TFStringValue(x.Description),
		NodeIds:         helper.TFStringsValue(x.NodeIds),
		WidgetInfo:      widgetInfos,
		Widgets:         dashWidgets,
		Tags:            general_objects.KeyValuesToTF(x.Tags),
		TimestampFormat: helper.TFStringValue(x.TimestampFormat),
	}
}

func (tf *DashboardTF) ToAPI() any {
	nodeIds := helper.FromTFStrings(tf.NodeIds)

	widgetInfos := make([]WidgetInfo, len(tf.WidgetInfo))
	for i, wi := range tf.WidgetInfo {
		widgetInfos[i] = WidgetInfo{
			ID:   helper.FromTFString(wi.ID),
			Type: helper.FromTFString(wi.Type),
			W:    helper.FromTFInt64(wi.W),
			H:    helper.FromTFInt64(wi.H),
			X:    helper.FromTFInt64(wi.X),
			Y:    helper.FromTFInt64(wi.Y),
			MinW: helper.FromTFInt64(wi.MinW),
			MinH: helper.FromTFInt64(wi.MinH),
		}
	}

	dashWidgets := make([]DashboardWidget, len(tf.Widgets))
	for i, w := range tf.Widgets {
		series := make([]widgets.Series, len(w.Series))
		for j, s := range w.Series {
			filters := make([]widgets.Filter, len(s.Filters))
			for k, f := range s.Filters {
				filters[k] = widgets.Filter{
					FilterBy: helper.FromTFString(f.FilterBy),
					Operator: helper.FromTFString(f.Operator),
					Value:    helper.FromTFString(f.Value),
				}
			}
			series[j] = widgets.Series{
				ID:          helper.FromTFString(s.ID),
				Name:        helper.FromTFString(s.Name),
				Datasource:  helper.FromTFString(s.Datasource),
				Aggregation: helper.FromTFString(s.Aggregation),
				Filters:     filters,
			}
		}

		dw := DashboardWidget{
			AuditModel:           general_objects.AuditModelFromTF(w.AuditModelTF),
			Name:                 helper.FromTFString(w.Name),
			Description:          helper.FromTFString(w.Description),
			Type:                 helper.FromTFString(w.Type),
			Granularity:          helper.FromTFString(w.Granularity),
			QueryTimeDimension:   helper.FromTFString(w.QueryTimeDimension),
			DisplayTimeDimension: helper.FromTFString(w.DisplayTimeDimension),
			Series:               series,
			Tags:                 general_objects.KeyValuesFromTF(w.Tags),
		}
		if len(w.Metadata) > 0 {
			md := w.Metadata[0]
			dw.Metadata = widgets.Metadata{
				YAxisLabel: helper.FromTFString(md.YAxisLabel),
				YAxisRange: make([]*float64, 2),
			}
			if len(md.YAxisRangeMin) > 0 {
				v := md.YAxisRangeMin[0].ValueFloat64()
				dw.Metadata.YAxisRange[0] = &v
			}
			if len(md.YAxisRangeMax) > 0 {
				v := md.YAxisRangeMax[0].ValueFloat64()
				dw.Metadata.YAxisRange[1] = &v
			}
		}
		dashWidgets[i] = dw
	}

	return &Dashboard{
		AuditModel:      general_objects.AuditModelFromTF(tf.AuditModelTF),
		Name:            helper.FromTFString(tf.Name),
		Description:     helper.FromTFString(tf.Description),
		NodeIds:         nodeIds,
		WidgetInfo:      widgetInfos,
		Widgets:         dashWidgets,
		Tags:            general_objects.KeyValuesFromTF(tf.Tags),
		TimestampFormat: helper.FromTFString(tf.TimestampFormat),
	}
}
