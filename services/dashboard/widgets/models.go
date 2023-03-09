package widgets

import "github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

type Widget struct {
	ID             string                     `json:"id"`
	Name           string                     `json:"name"`
	Description    string                     `json:"description,omitempty"`
	Type           string                     `json:"type"`
	Granularity    string                     `json:"granularity"`
	Series         []Series                   `json:"series"`
	Metrics        []MetricInfo               `json:"metrics"`
	Metadata       Metadata                   `json:"metadata"`
	Dashboards     []DashboardInfo            `json:"dashboards"`
	Tags           []general_objects.KeyValue `json:"tags,omitempty"`
	CreatedAt      string                     `json:"createdAt"`
	CreatedBy      string                     `json:"createdBy"`
	LastModifiedAt string                     `json:"lastModifiedAt"`
	LastModifiedBy string                     `json:"lastModifiedBy"`
}

func (widget *Widget) GetID() string { return widget.ID }

type Series struct {
	ID          string   `json:"id"`
	Datasource  string   `json:"datasource"`
	Aggregation string   `json:"aggregation"`
	Filters     []Filter `json:"filters"`
}

type Filter struct {
	FilterBy string `json:"filterBy"`
	Operator string `json:"operator"`
	Value    string `json:"value"`
}

type MetricInfo struct {
	ID          string `json:"id"`
	Aggregation string `json:"aggregation"`
}

type Metadata struct {
	YAxisLabel string     `json:"yAxisLabel"`
	YAxisRange []*float64 `json:"yAxisRange"`
}

type DashboardInfo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
