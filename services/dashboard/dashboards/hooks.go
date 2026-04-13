package dashboards

import (
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/provider"
)

func (dashboard *Dashboard) PostReadProcess(_ *provider.Client, newValue any) error {
	newDashboard, ok := newValue.(*Dashboard)
	if !ok || newDashboard == nil {
		return nil
	}
	newDashboard.WidgetInfo = helper.ReorderByKey(dashboard.WidgetInfo, newDashboard.WidgetInfo, func(w WidgetInfo) string { return w.ID })
	return nil
}

func (dashboard *Dashboard) PostUnmarshallProcess() error {
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
