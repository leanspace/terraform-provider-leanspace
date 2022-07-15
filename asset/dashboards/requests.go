package dashboards

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"terraform-provider-asset/asset"
)

type apiValidWidgetInfo struct {
	View WidgetView `json:"view"`
}

func (widgetInfo *WidgetInfo) toAPIFormat() ([]byte, error) {
	widgetView := apiValidWidgetInfo{
		View: WidgetView{
			WidgetType: widgetInfo.Type,
			Grid: Grid{
				Version: 1,
				I:       widgetInfo.ID,
				W:       widgetInfo.W,
				H:       widgetInfo.H,
				MinW:    widgetInfo.MinW,
				MinH:    widgetInfo.MinH,
				X:       widgetInfo.X,
				Y:       widgetInfo.Y,
			},
		},
	}
	return json.Marshal(widgetView)
}

func nodeAction(action string, node string, dashboardId string, client *asset.Client) error {
	path := fmt.Sprintf("%s/%s/%s/nodes/%s", client.HostURL, DashboardDataType.Path, dashboardId, node)
	req, err := http.NewRequest(action, path, nil)
	if err != nil {
		return err
	}
	_, err, _ = client.DoRequest(req, &(client).Token)
	if err != nil {
		return err
	}
	return nil
}

func attachNode(node string, dashboardId string, client *asset.Client) error {
	return nodeAction("POST", node, dashboardId, client)
}

func detachNode(node string, dashboardId string, client *asset.Client) error {
	return nodeAction("DELETE", node, dashboardId, client)
}

func addWidget(widget WidgetInfo, dashboardId string, client *asset.Client) error {
	data, err := widget.toAPIFormat()
	if err != nil {
		return err
	}
	path := fmt.Sprintf("%s/%s/%s/widgets/%s", client.HostURL, DashboardDataType.Path, dashboardId, widget.ID)
	req, err := http.NewRequest("POST", path, strings.NewReader(string(data)))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	_, err, _ = client.DoRequest(req, &(client).Token)
	if err != nil {
		return err
	}
	return nil
}

func removeWidget(widget WidgetInfo, dashboardId string, client *asset.Client) error {
	path := fmt.Sprintf("%s/%s/%s/widgets/%s", client.HostURL, DashboardDataType.Path, dashboardId, widget.ID)
	req, err := http.NewRequest("DELETE", path, nil)
	if err != nil {
		return err
	}
	_, err, _ = client.DoRequest(req, &(client).Token)
	if err != nil {
		return err
	}
	return nil
}

// 0 = present, same value
// 1 = present, diff value
// 2 = absent
func diffWidgetView(widgets []WidgetInfo, target WidgetInfo) int {
	for _, w := range widgets {
		if w.ID == target.ID {
			if w == target {
				return 0
			} else {
				return 1
			}
		}
	}
	return 2
}

func (dashboard *Dashboard) PostCreateProcess(client *asset.Client, created any) error {
	createdDashboard := created.(*Dashboard)
	for _, node := range dashboard.NodeIds { // Attach all nodes
		if err := attachNode(node, createdDashboard.ID, client); err != nil {
			return err
		}
	}
	for _, widget := range dashboard.WidgetInfo { // Add all widgets
		if err := addWidget(widget, createdDashboard.ID, client); err != nil {
			return err
		}
	}
	return nil
}

func (dashboard *Dashboard) PostUpdateProcess(client *asset.Client, updated any) error {
	updatedDashboard := updated.(*Dashboard)
	for _, node := range dashboard.NodeIds { // Attach needed nodes
		if !asset.Contains(updatedDashboard.NodeIds, node) {
			if err := attachNode(node, updatedDashboard.ID, client); err != nil {
				return err
			}
		}
	}
	for _, node := range updatedDashboard.NodeIds { // Remove extra nodes
		if !asset.Contains(dashboard.NodeIds, node) {
			if err := detachNode(node, updatedDashboard.ID, client); err != nil {
				return err
			}
		}
	}

	for _, widget := range dashboard.WidgetInfo { // Add all missing widgets
		diff := diffWidgetView(updatedDashboard.WidgetInfo, widget)
		if diff == 2 { // Missing, add widget
			if err := addWidget(widget, updatedDashboard.ID, client); err != nil {
				return err
			}
		}
		if diff == 1 { // Present but changed, remove and re-add
			if err := removeWidget(widget, updatedDashboard.ID, client); err != nil {
				return err
			}
			if err := addWidget(widget, updatedDashboard.ID, client); err != nil {
				return err
			}
		}
	}
	for _, widget := range updatedDashboard.WidgetInfo { // Remove all extra widgets
		if diffWidgetView(dashboard.WidgetInfo, widget) == 2 { // Missing, remove
			if err := removeWidget(widget, updatedDashboard.ID, client); err != nil {
				return err
			}
		}
	}
	return nil
}
