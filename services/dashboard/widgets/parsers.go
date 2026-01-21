package widgets

import (
	"strconv"

	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func (widget *Widget) ToMap() map[string]any {
	widgetMap := make(map[string]any)
	widgetMap["id"] = widget.ID
	widgetMap["name"] = widget.Name
	widgetMap["description"] = widget.Description
	widgetMap["type"] = widget.Type
	widgetMap["granularity"] = widget.Granularity
	widgetMap["query_time_dimension"] = widget.QueryTimeDimension
	widgetMap["display_time_dimension"] = widget.DisplayTimeDimension
	widgetMap["series"] = helper.ParseToMaps(widget.Series)
	if metadataMap := widget.Metadata.ToMap(); metadataMap != nil {
		widgetMap["metadata"] = []any{metadataMap}
	}
	widgetMap["dashboards"] = helper.ParseToMaps(widget.Dashboards)
	widgetMap["tags"] = helper.ParseToMaps(widget.Tags)
	widgetMap["created_at"] = widget.CreatedAt
	widgetMap["created_by"] = widget.CreatedBy
	widgetMap["last_modified_at"] = widget.LastModifiedAt
	widgetMap["last_modified_by"] = widget.LastModifiedBy
	return widgetMap
}

func (series *Series) ToMap() map[string]any {
	seriesMap := make(map[string]any)
	seriesMap["id"] = series.ID
	seriesMap["name"] = series.Name
	seriesMap["datasource"] = series.Datasource
	seriesMap["aggregation"] = series.Aggregation
	seriesMap["filters"] = helper.ParseToMaps(series.Filters)
	return seriesMap
}

func (filter *Filter) ToMap() map[string]any {
	filterMap := make(map[string]any)
	filterMap["filter_by"] = filter.FilterBy
	filterMap["operator"] = filter.Operator
	filterMap["value"] = filter.Value
	return filterMap
}

func (metadata *Metadata) ToMap() map[string]any {
	min_set, max_set := false, false
	metadataMap := make(map[string]any)
	metadataMap["y_axis_label"] = metadata.YAxisLabel
	metadataMap["thresholds"] = helper.ParseToMaps(metadata.Thresholds)
	if metadata.YAxisRange != nil && len(metadata.YAxisRange) == 2 {
		if metadata.YAxisRange[0] != nil {
			metadataMap["y_axis_range_min"] = []any{metadata.YAxisRange[0]}
			min_set = true
		}
		if metadata.YAxisRange[1] != nil {
			metadataMap["y_axis_range_max"] = []any{metadata.YAxisRange[1]}
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
	thresoldMap["color"] = threshold.Color
	return thresoldMap
}

func (dashboardInfo *DashboardInfo) ToMap() map[string]any {
	dashboardInfoMap := make(map[string]any)
	dashboardInfoMap["id"] = dashboardInfo.ID
	dashboardInfoMap["name"] = dashboardInfo.Name
	return dashboardInfoMap
}

func (widget *Widget) FromMap(widgetMap map[string]any) error {
	widget.ID = widgetMap["id"].(string)
	widget.Name = widgetMap["name"].(string)
	widget.Description = widgetMap["description"].(string)
	widget.Type = widgetMap["type"].(string)
	widget.Granularity = widgetMap["granularity"].(string)
	widget.QueryTimeDimension = widgetMap["query_time_dimension"].(string)
	widget.DisplayTimeDimension = widgetMap["display_time_dimension"].(string)
	if series, err := helper.ParseFromMaps[Series](widgetMap["series"].([]any)); err != nil {
		return err
	} else {
		widget.Series = series
	}
	if len(widgetMap["metadata"].([]any)) > 0 && widgetMap["metadata"].([]any)[0] != nil {
		if err := widget.Metadata.FromMap(widgetMap["metadata"].([]any)[0].(map[string]any)); err != nil {
			return err
		}
	}
	if dashboards, err := helper.ParseFromMaps[DashboardInfo](widgetMap["dashboards"].(*schema.Set).List()); err != nil {
		return err
	} else {
		widget.Dashboards = dashboards
	}
	if tags, err := helper.ParseFromMaps[general_objects.KeyValue](widgetMap["tags"].(*schema.Set).List()); err != nil {
		return err
	} else {
		widget.Tags = tags
	}
	widget.CreatedAt = widgetMap["created_at"].(string)
	widget.CreatedBy = widgetMap["created_by"].(string)
	widget.LastModifiedAt = widgetMap["last_modified_at"].(string)
	widget.LastModifiedBy = widgetMap["last_modified_by"].(string)
	return nil
}

func (series *Series) FromMap(seriesMap map[string]any) error {
	series.ID = seriesMap["id"].(string)
	series.Name = seriesMap["name"].(string)
	series.Datasource = seriesMap["datasource"].(string)
	series.Aggregation = seriesMap["aggregation"].(string)
	if filters, err := helper.ParseFromMaps[Filter](seriesMap["filters"].(*schema.Set).List()); err != nil {
		return err
	} else {
		series.Filters = filters
	}
	return nil
}

func (filter *Filter) FromMap(filterMap map[string]any) error {
	filter.FilterBy = filterMap["filter_by"].(string)
	filter.Operator = filterMap["operator"].(string)
	filter.Value = filterMap["value"].(string)
	return nil
}

func (metadata *Metadata) FromMap(metadataMap map[string]any) error {
	metadata.YAxisLabel = metadataMap["y_axis_label"].(string)

	metadata.YAxisRange = make([]*float64, 2)
	if len(metadataMap["thresholds"].([]any)) > 0 {
		if thresholds, err := helper.ParseFromMaps[Threshold](metadataMap["thresholds"].([]any)); err != nil {
			return err
		} else {
			metadata.Thresholds = thresholds
		}
	}
	if min, exists := metadataMap["y_axis_range_min"]; exists {
		if len(min.([]any)) == 1 {
			min := min.([]any)[0].(float64)
			metadata.YAxisRange[0] = &min
		}
	}
	if max, exists := metadataMap["y_axis_range_max"]; exists {
		if len(max.([]any)) == 1 {
			max := max.([]any)[0].(float64)
			metadata.YAxisRange[1] = &max
		}
	}
	return nil
}

func (thresold *Threshold) FromMap(thresoldMap map[string]any) error {
	from := thresoldMap["from"].(string)
	fromInFloat, _ := strconv.ParseFloat(from, 64)
	to := thresoldMap["to"].(string)
	toInFloat, _ := strconv.ParseFloat(to, 64)
	if from != "" {
		thresold.From = &fromInFloat
	}
	if to != "" {
		thresold.To = &toInFloat
	}
	thresold.Color = thresoldMap["color"].(string)
	return nil
}

func (dashboardInfo *DashboardInfo) FromMap(dashboardInfoMap map[string]any) error {
	dashboardInfo.ID = dashboardInfoMap["id"].(string)
	dashboardInfo.Name = dashboardInfoMap["name"].(string)
	return nil
}
