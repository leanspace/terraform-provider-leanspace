package dashboards

import (
	"terraform-provider-asset/asset/general_objects"
	"terraform-provider-asset/asset/widgets"
)

type Dashboard struct {
	ID             string                `json:"id" terra:"id"`
	Name           string                `json:"name" terra:"name"`
	Description    string                `json:"description,omitempty" terra:"description"`
	NodeIds        []string              `json:"nodeIds" terra:"node_ids"`
	WidgetInfo     []WidgetInfo          `json:"widgetInfo,omitempty" terra:"widget_info,omitempty"`
	Widgets        []DashboardWidget     `json:"widgets" terra:"widgets"`
	Tags           []general_objects.Tag `json:"tags,omitempty" terra:"tags"`
	CreatedAt      string                `json:"createdAt" terra:"created_at"`
	CreatedBy      string                `json:"createdBy" terra:"created_by"`
	LastModifiedAt string                `json:"lastModifiedAt" terra:"last_modified_at"`
	LastModifiedBy string                `json:"lastModifiedBy" terra:"last_modified_by"`
}

func (dashboard *Dashboard) GetID() string { return dashboard.ID }

type WidgetInfo struct {
	ID   string `json:"id" terra:"id"`
	Type string `json:"type" terra:"type"`
	W    int    `json:"w" terra:"w"`
	H    int    `json:"h" terra:"h"`
	X    int    `json:"x" terra:"x"`
	Y    int    `json:"y" terra:"y"`
	MinW int    `json:"minW" terra:"min_w"`
	MinH int    `json:"minH" terra:"min_h"`
}

type DashboardWidget struct {
	ID             string                  `json:"id" terra:"id"`
	Name           string                  `json:"name" terra:"name"`
	Description    string                  `json:"description,omitempty" terra:"description"`
	Type           string                  `json:"type" terra:"type"`
	Granularity    string                  `json:"granularity" terra:"granularity"`
	Series         []widgets.Series        `json:"series" terra:"series"`
	Metrics        []widgets.MetricInfo    `json:"metrics" terra:"metrics"`
	Metadata       widgets.Metadata        `json:"metadata" terra:"metadata"`
	Dashboards     []widgets.DashboardInfo `json:"dashboards" terra:"dashboards"`
	View           WidgetView              `json:"view" terra:"view"`
	Tags           []general_objects.Tag   `json:"tags,omitempty" terra:"tags"`
	CreatedAt      string                  `json:"createdAt" terra:"created_at"`
	CreatedBy      string                  `json:"createdBy" terra:"created_by"`
	LastModifiedAt string                  `json:"lastModifiedAt" terra:"last_modified_at"`
	LastModifiedBy string                  `json:"lastModifiedBy" terra:"last_modified_by"`
}

type WidgetView struct {
	WidgetType string `json:"widgetType" terra:"widget_type"`
	Grid       Grid   `json:"grid" terra:"grid"`
}

type Grid struct {
	Version int    `json:"version" terra:"version"`
	W       int    `json:"w" terra:"w"`
	H       int    `json:"h" terra:"h"`
	MinW    int    `json:"minW" terra:"min_w"`
	MinH    int    `json:"minH" terra:"min_h"`
	X       int    `json:"x" terra:"x"`
	Y       int    `json:"y" terra:"y"`
	I       string `json:"i" terra:"i"`
}
