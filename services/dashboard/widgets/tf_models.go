package widgets

import (
	"sort"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

var widgetDashboardInfoAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"name": types.StringType,
}

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
	Dashboards           types.Set                    `tfsdk:"dashboards"`
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

func (x *Widget) ToTF() any {
	series := make([]SeriesTF, len(x.Series))
	for i, s := range x.Series {
		filters := make([]FilterTF, len(s.Filters))
		for j, f := range s.Filters {
			filters[j] = FilterTF{
				FilterBy: types.StringValue(f.FilterBy),
				Operator: types.StringValue(f.Operator),
				Value:    types.StringValue(f.Value),
			}
		}
		// Sort filters canonically so list order is stable regardless of API return order.
		sort.Slice(filters, func(a, b int) bool {
			fa, fb := filters[a], filters[b]
			if fa.FilterBy != fb.FilterBy {
				return fa.FilterBy.ValueString() < fb.FilterBy.ValueString()
			}
			if fa.Operator != fb.Operator {
				return fa.Operator.ValueString() < fb.Operator.ValueString()
			}
			return fa.Value.ValueString() < fb.Value.ValueString()
		})
		series[i] = SeriesTF{
			ID:          types.StringValue(s.ID),
			Name:        helper.TFStringPtrValue(s.Name),
			Datasource:  types.StringValue(s.Datasource),
			Aggregation: types.StringValue(s.Aggregation),
			Filters:     filters,
		}
	}

	var metadata *MetadataTF
	if x.Metadata != nil {
		md := &MetadataTF{
			YAxisLabel: helper.TFStringPtrValue(x.Metadata.YAxisLabel),
		}
		if len(x.Metadata.YAxisRange) == 2 {
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
				from = types.StringValue(strconv.FormatFloat(*t.From, 'g', -1, 64))
			}
			if t.To != nil {
				to = types.StringValue(strconv.FormatFloat(*t.To, 'g', -1, 64))
			}
			thresholds[j] = ThresholdTF{From: from, To: to, Color: types.StringValue(t.Color)}
		}
		md.Thresholds = thresholds

		// Only set metadata if there's actual content
		if x.Metadata.YAxisLabel != nil || (len(x.Metadata.YAxisRange) == 2 && (x.Metadata.YAxisRange[0] != nil || x.Metadata.YAxisRange[1] != nil)) || len(x.Metadata.Thresholds) > 0 {
			metadata = md
		}
	}

	dashElems := make([]attr.Value, len(x.Dashboards))
	for i, d := range x.Dashboards {
		dObj, _ := types.ObjectValue(widgetDashboardInfoAttrTypes, map[string]attr.Value{
			"id":   types.StringValue(d.ID),
			"name": types.StringValue(d.Name),
		})
		dashElems[i] = dObj
	}
	dashboards := types.SetValueMust(types.ObjectType{AttrTypes: widgetDashboardInfoAttrTypes}, dashElems)

	return &WidgetTF{
		AuditModelTF:         general_objects.AuditModelToTF(&x.AuditModel),
		Name:                 types.StringValue(x.Name),
		Description:          helper.TFStringPtrValue(x.Description),
		Type:                 types.StringValue(x.Type),
		Granularity:          types.StringValue(x.Granularity),
		QueryTimeDimension:   types.StringValue(x.QueryTimeDimension),
		DisplayTimeDimension: types.StringValue(x.DisplayTimeDimension),
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
			Name:        helper.FromTFStringPtr(s.Name),
			Datasource:  helper.FromTFString(s.Datasource),
			Aggregation: helper.FromTFString(s.Aggregation),
			Filters:     filters,
		}
	}

	var metadata *Metadata
	if tf.Metadata != nil {
		metadata = &Metadata{
			YAxisLabel: helper.FromTFStringPtr(tf.Metadata.YAxisLabel),
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

	return &Widget{
		AuditModel:           general_objects.AuditModelFromTF(tf.AuditModelTF),
		Name:                 helper.FromTFString(tf.Name),
		Description:          helper.FromTFStringPtr(tf.Description),
		Type:                 helper.FromTFString(tf.Type),
		Granularity:          helper.FromTFString(tf.Granularity),
		QueryTimeDimension:   helper.FromTFString(tf.QueryTimeDimension),
		DisplayTimeDimension: helper.FromTFString(tf.DisplayTimeDimension),
		Series:               series,
		Metadata:             metadata,
		Tags:                 general_objects.KeyValuesFromTF(tf.Tags),
	}
}
