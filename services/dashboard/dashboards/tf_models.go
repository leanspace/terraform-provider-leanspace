package dashboards

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

var dashFilterAttrTypes = map[string]attr.Type{
	"filter_by": types.StringType,
	"operator":  types.StringType,
	"value":     types.StringType,
}

var dashSeriesAttrTypes = map[string]attr.Type{
	"id":          types.StringType,
	"name":        types.StringType,
	"datasource":  types.StringType,
	"aggregation": types.StringType,
	"filters":     types.SetType{ElemType: types.ObjectType{AttrTypes: dashFilterAttrTypes}},
}

var dashThresholdAttrTypes = map[string]attr.Type{
	"from":  types.StringType,
	"to":    types.StringType,
	"color": types.StringType,
}

var dashMetadataAttrTypes = map[string]attr.Type{
	"y_axis_label":     types.StringType,
	"y_axis_range_min": types.ListType{ElemType: types.Float64Type},
	"y_axis_range_max": types.ListType{ElemType: types.Float64Type},
	"thresholds":       types.ListType{ElemType: types.ObjectType{AttrTypes: dashThresholdAttrTypes}},
}

var dashViewInfoAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"name": types.StringType,
}

var dashTagAttrTypes = map[string]attr.Type{
	"key":   types.StringType,
	"value": types.StringType,
}

var dashboardWidgetAttrTypes = map[string]attr.Type{
	"id":                     types.StringType,
	"created_at":             types.StringType,
	"created_by":             types.StringType,
	"last_modified_at":       types.StringType,
	"last_modified_by":       types.StringType,
	"name":                   types.StringType,
	"description":            types.StringType,
	"type":                   types.StringType,
	"granularity":            types.StringType,
	"query_time_dimension":   types.StringType,
	"display_time_dimension": types.StringType,
	"series":                 types.ListType{ElemType: types.ObjectType{AttrTypes: dashSeriesAttrTypes}},
	"metadata":               types.ListType{ElemType: types.ObjectType{AttrTypes: dashMetadataAttrTypes}},
	"view":                   types.ListType{ElemType: types.ObjectType{AttrTypes: dashViewInfoAttrTypes}},
	"tags":                   types.SetType{ElemType: types.ObjectType{AttrTypes: dashTagAttrTypes}},
}

type DashboardTF struct {
	general_objects.AuditModelTF
	Name            types.String                 `tfsdk:"name"`
	Description     types.String                 `tfsdk:"description"`
	NodeIds         []types.String               `tfsdk:"node_ids"`
	WidgetInfo      []WidgetInfoTF               `tfsdk:"widget_info"`
	Widgets         types.Set                    `tfsdk:"widgets"`
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

	widgetElems := make([]attr.Value, len(x.Widgets))
	for i, w := range x.Widgets {
		// Build series list
		seriesElems := make([]attr.Value, len(w.Series))
		for j, s := range w.Series {
			filterElems := make([]attr.Value, len(s.Filters))
			for k, f := range s.Filters {
				fObj, _ := types.ObjectValue(dashFilterAttrTypes, map[string]attr.Value{
					"filter_by": helper.TFStringValue(f.FilterBy),
					"operator":  helper.TFStringValue(f.Operator),
					"value":     helper.TFStringValue(f.Value),
				})
				filterElems[k] = fObj
			}
			filtersSet, _ := types.SetValue(types.ObjectType{AttrTypes: dashFilterAttrTypes}, filterElems)
			sObj, _ := types.ObjectValue(dashSeriesAttrTypes, map[string]attr.Value{
				"id":          helper.TFStringValue(s.ID),
				"name":        helper.TFStringPtrValue(s.Name),
				"datasource":  helper.TFStringValue(s.Datasource),
				"aggregation": helper.TFStringValue(s.Aggregation),
				"filters":     filtersSet,
			})
			seriesElems[j] = sObj
		}
		seriesList, _ := types.ListValue(types.ObjectType{AttrTypes: dashSeriesAttrTypes}, seriesElems)

		// Build metadata list
		var metadataElems []attr.Value
		hasMetadata := w.Metadata.YAxisLabel != nil || len(w.Metadata.Thresholds) > 0 ||
			(w.Metadata.YAxisRange != nil && len(w.Metadata.YAxisRange) == 2)
		if hasMetadata {
			thresholdElems := make([]attr.Value, len(w.Metadata.Thresholds))
			for j, t := range w.Metadata.Thresholds {
				tObj, _ := types.ObjectValue(dashThresholdAttrTypes, map[string]attr.Value{
					"from":  types.StringNull(),
					"to":    types.StringNull(),
					"color": helper.TFStringValue(t.Color),
				})
				thresholdElems[j] = tObj
			}
			thresholdList, _ := types.ListValue(types.ObjectType{AttrTypes: dashThresholdAttrTypes}, thresholdElems)

			var minElems []attr.Value
			if w.Metadata.YAxisRange != nil && len(w.Metadata.YAxisRange) == 2 && w.Metadata.YAxisRange[0] != nil {
				minElems = []attr.Value{helper.TFFloat64PtrValue(w.Metadata.YAxisRange[0])}
			}
			minList, _ := types.ListValue(types.Float64Type, minElems)

			var maxElems []attr.Value
			if w.Metadata.YAxisRange != nil && len(w.Metadata.YAxisRange) == 2 && w.Metadata.YAxisRange[1] != nil {
				maxElems = []attr.Value{helper.TFFloat64PtrValue(w.Metadata.YAxisRange[1])}
			}
			maxList, _ := types.ListValue(types.Float64Type, maxElems)

			mdObj, _ := types.ObjectValue(dashMetadataAttrTypes, map[string]attr.Value{
				"y_axis_label":     helper.TFStringPtrValue(w.Metadata.YAxisLabel),
				"y_axis_range_min": minList,
				"y_axis_range_max": maxList,
				"thresholds":       thresholdList,
			})
			metadataElems = []attr.Value{mdObj}
		}
		metadataList, _ := types.ListValue(types.ObjectType{AttrTypes: dashMetadataAttrTypes}, metadataElems)

		// Build view list (always empty — API view is complex struct not mapped here)
		viewList, _ := types.ListValue(types.ObjectType{AttrTypes: dashViewInfoAttrTypes}, nil)

		// Build tags set
		tagElems := make([]attr.Value, len(w.Tags))
		for j, t := range w.Tags {
			tObj, _ := types.ObjectValue(dashTagAttrTypes, map[string]attr.Value{
				"key":   helper.TFStringValue(t.Key),
				"value": helper.TFStringPtrValue(t.Value),
			})
			tagElems[j] = tObj
		}
		tagsSet, _ := types.SetValue(types.ObjectType{AttrTypes: dashTagAttrTypes}, tagElems)

		am := general_objects.AuditModelToTF(&w.AuditModel)
		widgetObj, _ := types.ObjectValue(dashboardWidgetAttrTypes, map[string]attr.Value{
			"id":                     am.ID,
			"created_at":             am.CreatedAt,
			"created_by":             am.CreatedBy,
			"last_modified_at":       am.LastModifiedAt,
			"last_modified_by":       am.LastModifiedBy,
			"name":                   helper.TFStringValue(w.Name),
			"description":            helper.TFStringValue(w.Description),
			"type":                   helper.TFStringValue(w.Type),
			"granularity":            helper.TFStringValue(w.Granularity),
			"query_time_dimension":   helper.TFStringValue(w.QueryTimeDimension),
			"display_time_dimension": helper.TFStringValue(w.DisplayTimeDimension),
			"series":                 seriesList,
			"metadata":               metadataList,
			"view":                   viewList,
			"tags":                   tagsSet,
		})
		widgetElems[i] = widgetObj
	}
	dashWidgets := types.SetValueMust(types.ObjectType{AttrTypes: dashboardWidgetAttrTypes}, widgetElems)

	return &DashboardTF{
		AuditModelTF:    general_objects.AuditModelToTF(&x.AuditModel),
		Name:            helper.TFStringValue(x.Name),
		Description:     helper.TFStringPtrValue(x.Description),
		NodeIds:         helper.TFStringsValue(x.NodeIds),
		WidgetInfo:      widgetInfos,
		Widgets:         dashWidgets,
		Tags:            general_objects.KeyValuesToTF(x.Tags),
		TimestampFormat: helper.TFStringPtrValue(x.TimestampFormat),
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

	return &Dashboard{
		AuditModel:      general_objects.AuditModelFromTF(tf.AuditModelTF),
		Name:            helper.FromTFString(tf.Name),
		Description:     helper.FromTFStringPtr(tf.Description),
		NodeIds:         nodeIds,
		WidgetInfo:      widgetInfos,
		Tags:            general_objects.KeyValuesFromTF(tf.Tags),
		TimestampFormat: helper.FromTFStringPtr(tf.TimestampFormat),
	}
}
