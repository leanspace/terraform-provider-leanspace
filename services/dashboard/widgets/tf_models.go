package widgets

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

type WidgetTF struct {
	general_objects.AuditModelTF
	Name                 types.String                 `tfsdk:"name"`
	Description          types.String                 `tfsdk:"description"`
	Type                 types.String                 `tfsdk:"type"`
	Granularity          types.String                 `tfsdk:"granularity"`
	QueryTimeDimension   types.String                 `tfsdk:"query_time_dimension"`
	DisplayTimeDimension types.String                 `tfsdk:"display_time_dimension"`
	Series               []SeriesTF                   `tfsdk:"series"`
	Metadata             *MetadataTF                  `tfsdk:"metadata"`
	Dashboards           []DashboardInfoTF            `tfsdk:"dashboards"`
	Tags                 []general_objects.KeyValueTF `tfsdk:"tags"`
}

type SeriesTF struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Datasource  types.String `tfsdk:"datasource"`
	Aggregation types.String `tfsdk:"aggregation"`
	Filters     []FilterTF   `tfsdk:"filters"`
}

type FilterTF struct {
	FilterBy types.String `tfsdk:"filter_by"`
	Operator types.String `tfsdk:"operator"`
	Value    types.String `tfsdk:"value"`
}

type MetadataTF struct {
	YAxisLabel    types.String  `tfsdk:"y_axis_label"`
	YAxisRangeMin types.Float64 `tfsdk:"y_axis_range_min"`
	YAxisRangeMax types.Float64 `tfsdk:"y_axis_range_max"`
	Thresholds    []ThresholdTF `tfsdk:"thresholds"`
}

type ThresholdTF struct {
	From  types.String `tfsdk:"from"`
	To    types.String `tfsdk:"to"`
	Color types.String `tfsdk:"color"`
}

type DashboardInfoTF struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

func (x *Widget) ToTF() any {
	series := make([]SeriesTF, len(x.Series))
	for i, s := range x.Series {
		filters := make([]FilterTF, len(s.Filters))
		for j, f := range s.Filters {
			filters[j] = FilterTF{
				FilterBy: helper.TFStringValue(f.FilterBy),
				Operator: helper.TFStringValue(f.Operator),
				Value:    helper.TFStringValue(f.Value),
			}
		}
		series[i] = SeriesTF{
			ID:          helper.TFStringValue(s.ID),
			Name:        helper.TFStringValue(s.Name),
			Datasource:  helper.TFStringValue(s.Datasource),
			Aggregation: helper.TFStringValue(s.Aggregation),
			Filters:     filters,
		}
	}

	var metadata *MetadataTF
	if x.Metadata != nil {
		md := &MetadataTF{
			YAxisLabel: helper.TFStringValue(x.Metadata.YAxisLabel),
		}
		if x.Metadata.YAxisRange != nil && len(x.Metadata.YAxisRange) == 2 {
			if x.Metadata.YAxisRange[0] != nil {
				md.YAxisRangeMin = helper.TFFloat64PtrValue(x.Metadata.YAxisRange[0])
			}
			if x.Metadata.YAxisRange[1] != nil {
				md.YAxisRangeMax = helper.TFFloat64PtrValue(x.Metadata.YAxisRange[1])
			}
		}
		thresholds := make([]ThresholdTF, len(x.Metadata.Thresholds))
		for j, t := range x.Metadata.Thresholds {
			var from, to types.String
			if t.From != nil {
				from = helper.TFStringValue(strconv.FormatFloat(*t.From, 'g', -1, 64))
			}
			if t.To != nil {
				to = helper.TFStringValue(strconv.FormatFloat(*t.To, 'g', -1, 64))
			}
			thresholds[j] = ThresholdTF{From: from, To: to, Color: helper.TFStringValue(t.Color)}
		}
		md.Thresholds = thresholds

		// Only set metadata if there's actual content
		if x.Metadata.YAxisLabel != "" || (x.Metadata.YAxisRange != nil && len(x.Metadata.YAxisRange) == 2 && (x.Metadata.YAxisRange[0] != nil || x.Metadata.YAxisRange[1] != nil)) || len(x.Metadata.Thresholds) > 0 {
			metadata = md
		}
	}

	dashboards := make([]DashboardInfoTF, len(x.Dashboards))
	for i, d := range x.Dashboards {
		dashboards[i] = DashboardInfoTF{
			ID:   helper.TFStringValue(d.ID),
			Name: helper.TFStringValue(d.Name),
		}
	}

	return &WidgetTF{
		AuditModelTF:         general_objects.AuditModelToTF(&x.AuditModel),
		Name:                 helper.TFStringValue(x.Name),
		Description:          helper.TFStringValue(x.Description),
		Type:                 helper.TFStringValue(x.Type),
		Granularity:          helper.TFStringValue(x.Granularity),
		QueryTimeDimension:   helper.TFStringValue(x.QueryTimeDimension),
		DisplayTimeDimension: helper.TFStringValue(x.DisplayTimeDimension),
		Series:               series,
		Metadata:             metadata,
		Dashboards:           dashboards,
		Tags:                 general_objects.KeyValuesToTF(x.Tags),
	}
}

func (tf *WidgetTF) ToAPI() any {
	series := make([]Series, len(tf.Series))
	for i, s := range tf.Series {
		filters := make([]Filter, len(s.Filters))
		for j, f := range s.Filters {
			filters[j] = Filter{
				FilterBy: helper.FromTFString(f.FilterBy),
				Operator: helper.FromTFString(f.Operator),
				Value:    helper.FromTFString(f.Value),
			}
		}
		series[i] = Series{
			ID:          helper.FromTFString(s.ID),
			Name:        helper.FromTFString(s.Name),
			Datasource:  helper.FromTFString(s.Datasource),
			Aggregation: helper.FromTFString(s.Aggregation),
			Filters:     filters,
		}
	}

	var metadata *Metadata
	if tf.Metadata != nil {
		metadata = &Metadata{
			YAxisLabel: helper.FromTFString(tf.Metadata.YAxisLabel),
			YAxisRange: make([]*float64, 2),
		}
		if !tf.Metadata.YAxisRangeMin.IsNull() && !tf.Metadata.YAxisRangeMin.IsUnknown() {
			v := tf.Metadata.YAxisRangeMin.ValueFloat64()
			metadata.YAxisRange[0] = &v
		}
		if !tf.Metadata.YAxisRangeMax.IsNull() && !tf.Metadata.YAxisRangeMax.IsUnknown() {
			v := tf.Metadata.YAxisRangeMax.ValueFloat64()
			metadata.YAxisRange[1] = &v
		}
		thresholds := make([]Threshold, len(tf.Metadata.Thresholds))
		for j, t := range tf.Metadata.Thresholds {
			from := helper.FromTFString(t.From)
			to := helper.FromTFString(t.To)
			th := Threshold{Color: helper.FromTFString(t.Color)}
			if from != "" {
				f, _ := strconv.ParseFloat(from, 64)
				th.From = &f
			}
			if to != "" {
				f, _ := strconv.ParseFloat(to, 64)
				th.To = &f
			}
			thresholds[j] = th
		}
		metadata.Thresholds = thresholds
	}

	dashboards := make([]DashboardInfo, len(tf.Dashboards))
	for i, d := range tf.Dashboards {
		dashboards[i] = DashboardInfo{
			ID:   helper.FromTFString(d.ID),
			Name: helper.FromTFString(d.Name),
		}
	}

	return &Widget{
		AuditModel:           general_objects.AuditModelFromTF(tf.AuditModelTF),
		Name:                 helper.FromTFString(tf.Name),
		Description:          helper.FromTFString(tf.Description),
		Type:                 helper.FromTFString(tf.Type),
		Granularity:          helper.FromTFString(tf.Granularity),
		QueryTimeDimension:   helper.FromTFString(tf.QueryTimeDimension),
		DisplayTimeDimension: helper.FromTFString(tf.DisplayTimeDimension),
		Series:               series,
		Metadata:             metadata,
		Dashboards:           dashboards,
		Tags:                 general_objects.KeyValuesFromTF(tf.Tags),
	}
}
