package widgets

import "leanspace-terraform-provider/helper/general_objects"

type Widget struct {
	ID             string                `json:"id" terra:"id"`
	Name           string                `json:"name" terra:"name"`
	Description    string                `json:"description,omitempty" terra:"description"`
	Type           string                `json:"type" terra:"type"`
	Granularity    string                `json:"granularity" terra:"granularity"`
	Series         []Series              `json:"series" terra:"series"`
	Metrics        []MetricInfo          `json:"metrics" terra:"metrics"`
	Metadata       Metadata              `json:"metadata" terra:"metadata"`
	Dashboards     []DashboardInfo       `json:"dashboards" terra:"dashboards"`
	Tags           []general_objects.Tag `json:"tags,omitempty" terra:"tags"`
	CreatedAt      string                `json:"createdAt" terra:"created_at"`
	CreatedBy      string                `json:"createdBy" terra:"created_by"`
	LastModifiedAt string                `json:"lastModifiedAt" terra:"last_modified_at"`
	LastModifiedBy string                `json:"lastModifiedBy" terra:"last_modified_by"`
}

func (widget *Widget) GetID() string { return widget.ID }

type Series struct {
	ID          string   `json:"id" terra:"id"`
	Datasource  string   `json:"datasource" terra:"datasource"`
	Aggregation string   `json:"aggregation" terra:"aggregation"`
	Filters     []Filter `json:"filters" terra:"filters"`
}

type Filter struct {
	FilterBy string `json:"filterBy" terra:"filter_by"`
	Operator string `json:"operator" terra:"operator"`
	Value    string `json:"value" terra:"value"`
}

type MetricInfo struct {
	ID          string `json:"id" terra:"id"`
	Aggregation string `json:"aggregation" terra:"aggregation"`
}

type Metadata struct {
	YAxisLabel string     `json:"yAxisLabel" terra:"y_axis_label"`
	YAxisRange []*float64 `json:"yAxisRange" terra:"y_axis_range"`
}

type DashboardInfo struct {
	ID   string `json:"id" terra:"id"`
	Name string `json:"name" terra:"name"`
}
