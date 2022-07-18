package metrics

import "leanspace-terraform-provider/provider"

var MetricDataType = provider.DataSourceType[Metric[any], *Metric[any]]{
	ResourceIdentifier: "leanspace_metrics",
	Name:               "metric",
	Path:               "metrics-repository/metrics",
	Schema:             metricSchema,
}
