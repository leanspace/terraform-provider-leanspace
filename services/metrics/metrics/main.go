package metrics

import "github.com/leanspace/terraform-provider-leanspace/provider"

var MetricDataType = provider.DataSourceType[Metric[any], *Metric[any]]{
	ResourceIdentifier: "leanspace_metrics",
	Path:               "metrics-repository/metrics",
	Schema:             metricSchema,
	FilterSchema:       dataSourceFilterSchema,
}
