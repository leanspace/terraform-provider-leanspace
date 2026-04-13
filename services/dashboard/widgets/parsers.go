package widgets

import (
	"strconv"

	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

func (widget *Widget) ToMap() map[string]any {
	widgetMap := widget.ToAuditMap()
	widgetMap["name"] = helper.NilIfEmpty(widget.Name)
	widgetMap["description"] = helper.NilIfEmpty(widget.Description)
	widgetMap["type"] = helper.NilIfEmpty(widget.Type)
	widgetMap["granularity"] = helper.NilIfEmpty(widget.Granularity)
	widgetMap["query_time_dimension"] = helper.NilIfEmpty(widget.QueryTimeDimension)
	widgetMap["display_time_dimension"] = helper.NilIfEmpty(widget.DisplayTimeDimension)
	widgetMap["series"] = helper.ParseToMaps(widget.Series)
	if widget.Metadata != nil {
		widgetMap["metadata"] = widget.Metadata.ToMap()
	}
	widgetMap["dashboards"] = helper.ParseToMaps(widget.Dashboards)
	widgetMap["tags"] = helper.ParseToMaps(widget.Tags)
	return widgetMap
}

func (series *Series) ToMap() map[string]any {
	seriesMap := make(map[string]any)
	seriesMap["id"] = helper.NilIfEmpty(series.ID)
	seriesMap["name"] = helper.NilIfEmpty(series.Name)
	seriesMap["datasource"] = helper.NilIfEmpty(series.Datasource)
	seriesMap["aggregation"] = helper.NilIfEmpty(series.Aggregation)
	seriesMap["filters"] = helper.ParseToMaps(series.Filters)
	return seriesMap
}

func (filter *Filter) ToMap() map[string]any {
	filterMap := make(map[string]any)
	filterMap["filter_by"] = helper.NilIfEmpty(filter.FilterBy)
	filterMap["operator"] = helper.NilIfEmpty(filter.Operator)
	filterMap["value"] = helper.NilIfEmpty(filter.Value)
	return filterMap
}

func (metadata *Metadata) ToMap() map[string]any {
	min_set, max_set := false, false
	metadataMap := make(map[string]any)
	metadataMap["y_axis_label"] = helper.NilIfEmpty(metadata.YAxisLabel)
	metadataMap["thresholds"] = helper.ParseToMaps(metadata.Thresholds)
	if metadata.YAxisRange != nil && len(metadata.YAxisRange) == 2 {
		if metadata.YAxisRange[0] != nil {
			minPointer := metadata.YAxisRange[0]
			metadataMap["y_axis_range_min"] = *minPointer
			min_set = true
		}
		if metadata.YAxisRange[1] != nil {
			maxPointer := metadata.YAxisRange[1]
			metadataMap["y_axis_range_max"] = *maxPointer
			max_set = true
		}
	}
	if !min_set && !max_set && metadata.YAxisLabel == "" {
		return nil
	}
	return metadataMap
}

func (threshold *Threshold) ToMap() map[string]any {
	thresoldMap := make(map[string]any)
	if threshold.From != nil {
		from := *threshold.From
		fromInString := strconv.FormatFloat(from, 'g', -1, 64)
		thresoldMap["from"] = fromInString
	}
	if threshold.To != nil {
		to := *threshold.To
		toInString := strconv.FormatFloat(to, 'g', -1, 64)
		thresoldMap["to"] = toInString
	}
	thresoldMap["color"] = helper.NilIfEmpty(threshold.Color)
	return thresoldMap
}

func (dashboardInfo *DashboardInfo) ToMap() map[string]any {
	dashboardInfoMap := make(map[string]any)
	dashboardInfoMap["id"] = helper.NilIfEmpty(dashboardInfo.ID)
	dashboardInfoMap["name"] = helper.NilIfEmpty(dashboardInfo.Name)
	return dashboardInfoMap
}

func (widget *Widget) FromMap(widgetMap map[string]any) error {
	widget.FromAuditMap(widgetMap)
	widget.Name = helper.CastString(widgetMap, "name")
	widget.Description = helper.CastString(widgetMap, "description")
	widget.Type = helper.CastString(widgetMap, "type")
	widget.Granularity = helper.CastString(widgetMap, "granularity")
	widget.QueryTimeDimension = helper.CastString(widgetMap, "query_time_dimension")
	widget.DisplayTimeDimension = helper.CastString(widgetMap, "display_time_dimension")
	if series, err := helper.ParseFromMaps[Series](helper.CastSlice(widgetMap, "series")); err != nil {
		return err
	} else {
		widget.Series = series
	}
	if widgetMap["metadata"] != nil {
		if widget.Metadata == nil {
			widget.Metadata = &Metadata{}
		}
		if err := widget.Metadata.FromMap(helper.CastMapAny(widgetMap, "metadata")); err != nil {
			return err
		}
	}
	if dashboards, err := helper.ParseFromMaps[DashboardInfo](helper.CastSlice(widgetMap, "dashboards")); err != nil {
		return err
	} else {
		widget.Dashboards = dashboards
	}
	if tags, err := helper.ParseFromMaps[general_objects.KeyValue](helper.CastSlice(widgetMap, "tags")); err != nil {
		return err
	} else {
		widget.Tags = tags
	}
	return nil
}

func (series *Series) FromMap(seriesMap map[string]any) error {
	series.ID = helper.CastString(seriesMap, "id")
	series.Name = helper.CastString(seriesMap, "name")
	series.Datasource = helper.CastString(seriesMap, "datasource")
	series.Aggregation = helper.CastString(seriesMap, "aggregation")
	if filters, err := helper.ParseFromMaps[Filter](helper.CastSlice(seriesMap, "filters")); err != nil {
		return err
	} else {
		series.Filters = filters
	}
	return nil
}

func (filter *Filter) FromMap(filterMap map[string]any) error {
	filter.FilterBy = helper.CastString(filterMap, "filter_by")
	filter.Operator = helper.CastString(filterMap, "operator")
	filter.Value = helper.CastString(filterMap, "value")
	return nil
}

func (metadata *Metadata) FromMap(metadataMap map[string]any) error {
	metadata.YAxisLabel = helper.CastString(metadataMap, "y_axis_label")

	metadata.YAxisRange = make([]*float64, 2)
	if len(helper.CastSlice(metadataMap, "thresholds")) > 0 {
		if thresholds, err := helper.ParseFromMaps[Threshold](helper.CastSlice(metadataMap, "thresholds")); err != nil {
			return err
		} else {
			metadata.Thresholds = thresholds
		}
	}
	if min, exists := metadataMap["y_axis_range_min"]; exists {
		if min != nil {
			min := min.(float64)
			metadata.YAxisRange[0] = &min
		}
	}
	if max, exists := metadataMap["y_axis_range_max"]; exists {
		if max != nil {
			max := max.(float64)
			metadata.YAxisRange[1] = &max
		}
	}
	return nil
}

func (thresold *Threshold) FromMap(thresoldMap map[string]any) error {
	from := helper.CastString(thresoldMap, "from")
	fromInFloat, _ := strconv.ParseFloat(from, 64)
	to := helper.CastString(thresoldMap, "to")
	toInFloat, _ := strconv.ParseFloat(to, 64)
	if from != "" {
		thresold.From = &fromInFloat
	}
	if to != "" {
		thresold.To = &toInFloat
	}
	thresold.Color = helper.CastString(thresoldMap, "color")
	return nil
}

func (dashboardInfo *DashboardInfo) FromMap(dashboardInfoMap map[string]any) error {
	dashboardInfo.ID = helper.CastString(dashboardInfoMap, "id")
	dashboardInfo.Name = helper.CastString(dashboardInfoMap, "name")
	return nil
}
