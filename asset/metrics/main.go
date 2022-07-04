package metrics

import (
	"terraform-provider-asset/asset"
)

var MetricDataType = asset.DataSourceType[Metric[any], *Metric[any]]{
	ResourceIdentifier: "leanspace_metrics",
	Name:               "metric",
	Path:               "metrics-repository/metrics",
	Schema:             metricSchema,
}
