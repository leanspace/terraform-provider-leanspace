package widgets

import (
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
	if series, err := helper.ParseFromMaps[Series](widgetMap["series"].(*schema.Set).List()); err != nil {
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

func (dashboardInfo *DashboardInfo) FromMap(dashboardInfoMap map[string]any) error {
	dashboardInfo.ID = dashboardInfoMap["id"].(string)
	dashboardInfo.Name = dashboardInfoMap["name"].(string)
	return nil
}
