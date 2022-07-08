package dashboards

import (
	"terraform-provider-asset/asset"
	"terraform-provider-asset/asset/general_objects"
	"terraform-provider-asset/asset/widgets"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func (dashboard *Dashboard) ToMap() map[string]any {
	dashboardMap := make(map[string]any)
	dashboardMap["id"] = dashboard.ID
	dashboardMap["name"] = dashboard.Name
	dashboardMap["description"] = dashboard.Description
	dashboardMap["node_ids"] = dashboard.NodeIds
	dashboardMap["widget_info"] = asset.ParseToMaps(dashboard.WidgetInfo)
	dashboardMap["widgets"] = asset.ParseToMaps(dashboard.Widgets)
	dashboardMap["tags"] = general_objects.TagsStructToMap(dashboard.Tags)
	dashboardMap["created_at"] = dashboard.CreatedAt
	dashboardMap["created_by"] = dashboard.CreatedBy
	dashboardMap["last_modified_at"] = dashboard.LastModifiedAt
	dashboardMap["last_modified_by"] = dashboard.LastModifiedBy
	return dashboardMap
}

func (widgetInfo *WidgetInfo) ToMap() map[string]any {
	widgetInfoMap := make(map[string]any)
	widgetInfoMap["id"] = widgetInfo.ID
	widgetInfoMap["type"] = widgetInfo.Type
	widgetInfoMap["x"] = widgetInfo.X
	widgetInfoMap["y"] = widgetInfo.Y
	widgetInfoMap["w"] = widgetInfo.W
	widgetInfoMap["h"] = widgetInfo.H
	widgetInfoMap["min_w"] = widgetInfo.MinW
	widgetInfoMap["min_h"] = widgetInfo.MinH
	return widgetInfoMap
}

func (widget *DashboardWidget) ToMap() map[string]any {
	widgetMap := make(map[string]any)
	widgetMap["id"] = widget.ID
	widgetMap["name"] = widget.Name
	widgetMap["description"] = widget.Description
	widgetMap["type"] = widget.Type
	widgetMap["granularity"] = widget.Granularity
	widgetMap["series"] = asset.ParseToMaps(widget.Series)
	widgetMap["metrics"] = asset.ParseToMaps(widget.Metrics)
	if metadataMap := widget.Metadata.ToMap(); metadataMap != nil {
		widgetMap["metadata"] = []any{metadataMap}
	}
	widgetMap["dashboards"] = asset.ParseToMaps(widget.Dashboards)
	widgetMap["tags"] = general_objects.TagsStructToMap(widget.Tags)
	widgetMap["created_at"] = widget.CreatedAt
	widgetMap["created_by"] = widget.CreatedBy
	widgetMap["last_modified_at"] = widget.LastModifiedAt
	widgetMap["last_modified_by"] = widget.LastModifiedBy
	return widgetMap
}

func (view *WidgetView) ToMap() map[string]any {
	viewMap := make(map[string]any)
	viewMap["type"] = view.WidgetType
	viewMap["grid"] = []any{view.Grid.ToMap()}
	return viewMap
}

func (grid *Grid) ToMap() map[string]any {
	gridMap := make(map[string]any)
	gridMap["version"] = grid.Version
	gridMap["w"] = grid.W
	gridMap["h"] = grid.H
	gridMap["min_w"] = grid.MinW
	gridMap["min_h"] = grid.MinH
	gridMap["x"] = grid.X
	gridMap["y"] = grid.Y
	gridMap["i"] = grid.I
	return gridMap
}

func (dashboard *Dashboard) FromMap(dashboardMap map[string]any) error {
	dashboard.ID = dashboardMap["id"].(string)
	dashboard.Name = dashboardMap["name"].(string)
	dashboard.Description = dashboardMap["description"].(string)
	dashboard.NodeIds = make([]string, dashboardMap["node_ids"].(*schema.Set).Len())
	for index, node := range dashboardMap["node_ids"].(*schema.Set).List() {
		dashboard.NodeIds[index] = node.(string)
	}
	if widgetInfo, err := asset.ParseFromMaps[WidgetInfo](dashboardMap["widget_info"].(*schema.Set).List()); err != nil {
		return err
	} else {
		dashboard.WidgetInfo = widgetInfo
	}
	if widgets, err := asset.ParseFromMaps[DashboardWidget](dashboardMap["widgets"].(*schema.Set).List()); err != nil {
		return err
	} else {
		dashboard.Widgets = widgets
	}
	dashboard.Tags = general_objects.TagsInterfaceToStruct(dashboardMap["tags"])
	dashboard.CreatedAt = dashboardMap["created_at"].(string)
	dashboard.CreatedBy = dashboardMap["created_by"].(string)
	dashboard.LastModifiedAt = dashboardMap["last_modified_at"].(string)
	dashboard.LastModifiedBy = dashboardMap["last_modified_by"].(string)
	return nil
}

func (widgetInfo *WidgetInfo) FromMap(widgetInfoMap map[string]any) error {
	widgetInfo.ID = widgetInfoMap["id"].(string)
	widgetInfo.Type = widgetInfoMap["type"].(string)
	widgetInfo.X = widgetInfoMap["x"].(int)
	widgetInfo.Y = widgetInfoMap["y"].(int)
	widgetInfo.W = widgetInfoMap["w"].(int)
	widgetInfo.H = widgetInfoMap["h"].(int)
	widgetInfo.MinW = widgetInfoMap["min_w"].(int)
	widgetInfo.MinH = widgetInfoMap["min_h"].(int)
	return nil
}

func (widget *DashboardWidget) FromMap(widgetMap map[string]any) error {
	widget.ID = widgetMap["id"].(string)
	widget.Name = widgetMap["name"].(string)
	widget.Description = widgetMap["description"].(string)
	widget.Type = widgetMap["type"].(string)
	widget.Granularity = widgetMap["granularity"].(string)
	if series, err := asset.ParseFromMaps[widgets.Series](widgetMap["series"].([]any)); err != nil {
		return err
	} else {
		widget.Series = series
	}
	if metrics, err := asset.ParseFromMaps[widgets.MetricInfo](widgetMap["metrics"].([]any)); err != nil {
		return err
	} else {
		widget.Metrics = metrics
	}
	if len(widgetMap["metadata"].([]any)) > 0 {
		if err := widget.Metadata.FromMap(widgetMap["metadata"].([]any)[0].(map[string]any)); err != nil {
			return err
		}
	}
	if err := widget.View.FromMap(widgetMap["view"].([]any)[0].(map[string]any)); err != nil {
		return err
	}
	widget.Tags = general_objects.TagsInterfaceToStruct(widgetMap["tags"])
	widget.CreatedAt = widgetMap["created_at"].(string)
	widget.CreatedBy = widgetMap["created_by"].(string)
	widget.LastModifiedAt = widgetMap["last_modified_at"].(string)
	widget.LastModifiedBy = widgetMap["last_modified_by"].(string)
	return nil
}

func (view *WidgetView) FromMap(viewMap map[string]any) error {
	view.WidgetType = viewMap["type"].(string)
	if err := view.Grid.FromMap(viewMap["grid"].([]any)[0].(map[string]any)); err != nil {
		return err
	}
	return nil
}

func (grid *Grid) FromMap(gridMap map[string]any) error {
	grid.Version = gridMap["version"].(int)
	grid.W = gridMap["w"].(int)
	grid.H = gridMap["h"].(int)
	grid.MinW = gridMap["min_w"].(int)
	grid.MinH = gridMap["min_h"].(int)
	grid.X = gridMap["x"].(int)
	grid.Y = gridMap["y"].(int)
	grid.I = gridMap["i"].(string)
	return nil
}

func (dashboard *Dashboard) PreMarshallProcess() error {
	// Transfer all data from the widgets array to the widget_info array,
	// to make sure the API call will ignore widgets (handled separately)
	// dashboard.WidgetInfo = make([]WidgetInfo, len(dashboard.Widgets))
	// for index, widget := range dashboard.Widgets {
	// 	dashboard.WidgetInfo[index] = WidgetInfo{
	// 		ID:   widget.ID,
	// 		Type: widget.Type,
	// 		X:    widget.View.Grid.X,
	// 		Y:    widget.View.Grid.Y,
	// 		W:    widget.View.Grid.W,
	// 		H:    widget.View.Grid.H,
	// 		MinW: widget.View.Grid.MinW,
	// 		MinH: widget.View.Grid.MinH,
	// 	}
	// }
	dashboard.Widgets = nil

	return nil
}

func (dashboard *Dashboard) PostUnmarshallProcess() error {
	// Update WidgetInfo array with "fresh" date
	dashboard.WidgetInfo = make([]WidgetInfo, len(dashboard.Widgets))
	for index, widget := range dashboard.Widgets {
		dashboard.WidgetInfo[index] = WidgetInfo{
			ID:   widget.ID,
			Type: widget.Type,
			X:    widget.View.Grid.X,
			Y:    widget.View.Grid.Y,
			W:    widget.View.Grid.W,
			H:    widget.View.Grid.H,
			MinW: widget.View.Grid.MinW,
			MinH: widget.View.Grid.MinH,
		}
	}
	return nil
}
