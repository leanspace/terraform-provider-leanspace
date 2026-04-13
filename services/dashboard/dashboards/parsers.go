package dashboards

import (
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
	"github.com/leanspace/terraform-provider-leanspace/provider"
	"github.com/leanspace/terraform-provider-leanspace/services/dashboard/widgets"
)

func (dashboard *Dashboard) ToMap() map[string]any {
	dashboardMap := dashboard.ToAuditMap()
	dashboardMap["name"] = helper.NilIfEmpty(dashboard.Name)
	dashboardMap["description"] = helper.NilIfEmpty(dashboard.Description)
	dashboardMap["node_ids"] = helper.NilIfEmpty(dashboard.NodeIds)
	dashboardMap["widget_info"] = helper.ParseToMaps(dashboard.WidgetInfo)
	dashboardMap["widgets"] = helper.ParseToMaps(dashboard.Widgets)
	dashboardMap["tags"] = helper.ParseToMaps(dashboard.Tags)
	dashboardMap["timestamp_format"] = helper.NilIfEmpty(dashboard.TimestampFormat)
	return dashboardMap
}

func (widgetInfo *WidgetInfo) ToMap() map[string]any {
	widgetInfoMap := make(map[string]any)
	widgetInfoMap["id"] = helper.NilIfEmpty(widgetInfo.ID)
	widgetInfoMap["type"] = helper.NilIfEmpty(widgetInfo.Type)
	widgetInfoMap["x"] = helper.NilIfEmpty(widgetInfo.X)
	widgetInfoMap["y"] = helper.NilIfEmpty(widgetInfo.Y)
	widgetInfoMap["w"] = helper.NilIfEmpty(widgetInfo.W)
	widgetInfoMap["h"] = helper.NilIfEmpty(widgetInfo.H)
	widgetInfoMap["min_w"] = helper.NilIfEmpty(widgetInfo.MinW)
	widgetInfoMap["min_h"] = helper.NilIfEmpty(widgetInfo.MinH)
	return widgetInfoMap
}

func (widget *DashboardWidget) ToMap() map[string]any {
	widgetMap := widget.ToAuditMap()
	widgetMap["name"] = helper.NilIfEmpty(widget.Name)
	widgetMap["description"] = helper.NilIfEmpty(widget.Description)
	widgetMap["type"] = helper.NilIfEmpty(widget.Type)
	widgetMap["granularity"] = helper.NilIfEmpty(widget.Granularity)
	widgetMap["query_time_dimension"] = helper.NilIfEmpty(widget.QueryTimeDimension)
	widgetMap["display_time_dimension"] = helper.NilIfEmpty(widget.DisplayTimeDimension)
	widgetMap["series"] = helper.ParseToMaps(widget.Series)
	if metadataMap := widget.Metadata.ToMap(); metadataMap != nil {
		widgetMap["metadata"] = []any{metadataMap}
	}
	widgetMap["tags"] = helper.ParseToMaps(widget.Tags)
	return widgetMap
}

func (view *WidgetView) ToMap() map[string]any {
	viewMap := make(map[string]any)
	viewMap["type"] = helper.NilIfEmpty(view.WidgetType)
	viewMap["grid"] = []any{view.Grid.ToMap()}
	return viewMap
}

func (grid *Grid) ToMap() map[string]any {
	gridMap := make(map[string]any)
	gridMap["version"] = helper.NilIfEmpty(grid.Version)
	gridMap["w"] = helper.NilIfEmpty(grid.W)
	gridMap["h"] = helper.NilIfEmpty(grid.H)
	gridMap["min_w"] = helper.NilIfEmpty(grid.MinW)
	gridMap["min_h"] = helper.NilIfEmpty(grid.MinH)
	gridMap["x"] = helper.NilIfEmpty(grid.X)
	gridMap["y"] = helper.NilIfEmpty(grid.Y)
	gridMap["i"] = helper.NilIfEmpty(grid.I)
	return gridMap
}

func (dashboard *Dashboard) FromMap(dashboardMap map[string]any) error {
	dashboard.FromAuditMap(dashboardMap)
	dashboard.Name = helper.CastString(dashboardMap, "name")
	dashboard.Description = helper.CastString(dashboardMap, "description")
	dashboard.NodeIds = make([]string, len(helper.CastSlice(dashboardMap, "node_ids")))
	for index, node := range helper.CastSlice(dashboardMap, "node_ids") {
		dashboard.NodeIds[index] = node.(string)
	}
	if widgetInfo, err := helper.ParseFromMaps[WidgetInfo](helper.CastSlice(dashboardMap, "widget_info")); err != nil {
		return err
	} else {
		dashboard.WidgetInfo = widgetInfo
	}
	if widgets, err := helper.ParseFromMaps[DashboardWidget](helper.CastSlice(dashboardMap, "widgets")); err != nil {
		return err
	} else {
		dashboard.Widgets = widgets
	}
	if tags, err := helper.ParseFromMaps[general_objects.KeyValue](helper.CastSlice(dashboardMap, "tags")); err != nil {
		return err
	} else {
		dashboard.Tags = tags
	}
	dashboard.TimestampFormat = helper.CastString(dashboardMap, "timestamp_format")
	return nil
}

func (widgetInfo *WidgetInfo) FromMap(widgetInfoMap map[string]any) error {
	widgetInfo.ID = helper.CastString(widgetInfoMap, "id")
	widgetInfo.Type = helper.CastString(widgetInfoMap, "type")
	widgetInfo.X = helper.CastInt(widgetInfoMap, "x")
	widgetInfo.Y = helper.CastInt(widgetInfoMap, "y")
	widgetInfo.W = helper.CastInt(widgetInfoMap, "w")
	widgetInfo.H = helper.CastInt(widgetInfoMap, "h")
	widgetInfo.MinW = helper.CastInt(widgetInfoMap, "min_w")
	widgetInfo.MinH = helper.CastInt(widgetInfoMap, "min_h")
	return nil
}

func (widget *DashboardWidget) FromMap(widgetMap map[string]any) error {
	widget.FromAuditMap(widgetMap)
	widget.Name = helper.CastString(widgetMap, "name")
	widget.Description = helper.CastString(widgetMap, "description")
	widget.Type = helper.CastString(widgetMap, "type")
	widget.Granularity = helper.CastString(widgetMap, "granularity")
	widget.QueryTimeDimension = helper.CastString(widgetMap, "query_time_dimension")
	widget.DisplayTimeDimension = helper.CastString(widgetMap, "display_time_dimension")
	if series, err := helper.ParseFromMaps[widgets.Series](helper.CastSlice(widgetMap, "series")); err != nil {
		return err
	} else {
		widget.Series = series
	}
	if len(helper.CastSlice(widgetMap, "metadata")) > 0 {
		if err := widget.Metadata.FromMap(helper.CastSlice(widgetMap, "metadata")[0].(map[string]any)); err != nil {
			return err
		}
	}
	if len(helper.CastSlice(widgetMap, "view")) > 0 {
		if err := widget.View.FromMap(helper.CastSlice(widgetMap, "view")[0].(map[string]any)); err != nil {
			return err
		}
	}
	if tags, err := helper.ParseFromMaps[general_objects.KeyValue](helper.CastSlice(widgetMap, "tags")); err != nil {
		return err
	} else {
		widget.Tags = tags
	}
	return nil
}

func (view *WidgetView) FromMap(viewMap map[string]any) error {
	view.WidgetType = helper.CastString(viewMap, "type")
	if len(helper.CastSlice(viewMap, "grid")) > 0 {
		if err := view.Grid.FromMap(helper.CastSlice(viewMap, "grid")[0].(map[string]any)); err != nil {
			return err
		}
	}
	return nil
}

func (grid *Grid) FromMap(gridMap map[string]any) error {
	grid.Version = helper.CastInt(gridMap, "version")
	grid.W = helper.CastInt(gridMap, "w")
	grid.H = helper.CastInt(gridMap, "h")
	grid.MinW = helper.CastInt(gridMap, "min_w")
	grid.MinH = helper.CastInt(gridMap, "min_h")
	grid.X = helper.CastInt(gridMap, "x")
	grid.Y = helper.CastInt(gridMap, "y")
	grid.I = helper.CastString(gridMap, "i")
	return nil
}

// PostReadProcess reorders widget_info in the API response to match the prior
// state order (matched by widget ID), preventing perpetual diffs when the API
// returns widgets in a different order than they were configured.
// Note: PostUnmarshallProcess has already rebuilt WidgetInfo from Widgets by this point.
func (dashboard *Dashboard) PostReadProcess(_ *provider.Client, newValue any) error {
	newDashboard, ok := newValue.(*Dashboard)
	if !ok || newDashboard == nil {
		return nil
	}
	newDashboard.WidgetInfo = helper.ReorderByKey(dashboard.WidgetInfo, newDashboard.WidgetInfo, func(w WidgetInfo) string { return w.ID })
	return nil
}

func (dashboard *Dashboard) PostUnmarshallProcess() error {
	// Update WidgetInfo array with "fresh" state
	// Because dashboard.widgets is marked as computed and not handled by terraform,
	// the only place where terraform will notice a change is in the widget_info
	// array (that is handled by the provider and unrelated to the API)
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
