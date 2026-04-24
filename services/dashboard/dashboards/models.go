package dashboards

import (
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
	"github.com/leanspace/terraform-provider-leanspace/services/dashboard/widgets"
)

type Dashboard struct {
	general_objects.AuditModel
	Name            string                     `json:"name"`
	Description     *string                    `json:"description,omitempty"`
	NodeIds         []string                   `json:"nodeIds"`
	WidgetInfo      []WidgetInfo               `json:"widgetInfo,omitempty"`
	Widgets         []DashboardWidget          `json:"widgets,omitempty"`
	Tags            []general_objects.KeyValue `json:"tags,omitempty"`
	TimestampFormat *string                    `json:"timestampFormat,omitempty"`
}

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
	general_objects.AuditModel
	Name                 string                     `json:"name"`
	Description          *string                    `json:"description,omitempty"`
	Type                 string                     `json:"type"`
	Granularity          string                     `json:"granularity"`
	QueryTimeDimension   string                     `json:"queryTimeDimension"`
	DisplayTimeDimension string                     `json:"displayTimeDimension"`
	Series               []widgets.Series           `json:"series"`
	Metadata             widgets.Metadata           `json:"metadata"`
	View                 WidgetView                 `json:"view"`
	Tags                 []general_objects.KeyValue `json:"tags,omitempty"`
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
