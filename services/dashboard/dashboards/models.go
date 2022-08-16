package dashboards

import (
	"leanspace-terraform-provider/helper/general_objects"
	"leanspace-terraform-provider/services/dashboard/widgets"
)

type Dashboard struct {
	ID             string                `json:"id"`
	Name           string                `json:"name"`
	Description    string                `json:"description,omitempty"`
	NodeIds        []string              `json:"nodeIds"`
	WidgetInfo     []WidgetInfo          `json:"widgetInfo,omitempty"`
	Widgets        []DashboardWidget     `json:"widgets"`
	Tags           []general_objects.Tag `json:"tags,omitempty"`
	CreatedAt      string                `json:"createdAt"`
	CreatedBy      string                `json:"createdBy"`
	LastModifiedAt string                `json:"lastModifiedAt"`
	LastModifiedBy string                `json:"lastModifiedBy"`
}

func (dashboard *Dashboard) GetID() string { return dashboard.ID }

type WidgetInfo struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	W    int    `json:"w"`
	H    int    `json:"h"`
	X    int    `json:"x"`
	Y    int    `json:"y"`
	MinW int    `json:"minW"`
	MinH int    `json:"minH"`
}

type DashboardWidget struct {
	ID             string                `json:"id"`
	Name           string                `json:"name"`
	Description    string                `json:"description,omitempty"`
	Type           string                `json:"type"`
	Granularity    string                `json:"granularity"`
	Series         []widgets.Series      `json:"series"`
	Metrics        []widgets.MetricInfo  `json:"metrics"`
	Metadata       widgets.Metadata      `json:"metadata"`
	View           WidgetView            `json:"view"`
	Tags           []general_objects.Tag `json:"tags,omitempty"`
	CreatedAt      string                `json:"createdAt"`
	CreatedBy      string                `json:"createdBy"`
	LastModifiedAt string                `json:"lastModifiedAt"`
	LastModifiedBy string                `json:"lastModifiedBy"`
}

type WidgetView struct {
	WidgetType string `json:"widgetType"`
	Grid       Grid   `json:"grid"`
}

type Grid struct {
	Version int    `json:"version"`
	W       int    `json:"w"`
	H       int    `json:"h"`
	MinW    int    `json:"minW"`
	MinH    int    `json:"minH"`
	X       int    `json:"x"`
	Y       int    `json:"y"`
	I       string `json:"i"`
}
