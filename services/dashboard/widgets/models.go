package widgets

import "github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

type Widget struct {
	general_objects.AuditModel
	Name                 string                     `json:"name"`
	Description          *string                    `json:"description,omitempty"`
	Type                 string                     `json:"type"`
	Granularity          string                     `json:"granularity"`
	QueryTimeDimension   string                     `json:"queryTimeDimension"`
	DisplayTimeDimension string                     `json:"displayTimeDimension"`
	Series               []Series                   `json:"series"`
	Metadata             *Metadata                  `json:"metadata,omitempty"`
	Dashboards           []DashboardInfo            `json:"dashboards"`
	Tags                 []general_objects.KeyValue `json:"tags,omitempty"`
}

type Series struct {
	ID          string   `json:"id"`
	Name        *string  `json:"name,omitempty"`
	Datasource  string   `json:"datasource"`
	Aggregation string   `json:"aggregation"`
	Filters     []Filter `json:"filters"`
}

type Filter struct {
	FilterBy string `json:"filterBy"`
	Operator string `json:"operator"`
	Value    string `json:"value"`
}

type Metadata struct {
	YAxisLabel *string     `json:"yAxisLabel,omitempty"`
	YAxisRange []*float64  `json:"yAxisRange,omitempty"`
	Thresholds []Threshold `json:"thresholds"`
}

type Threshold struct {
	From  *float64 `json:"from,omitempty"`
	To    *float64 `json:"to,omitempty"`
	Color string   `json:"color"`
}

type DashboardInfo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
