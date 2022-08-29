package monitors

import "github.com/leanspace/terraform-provider-leanspace/provider"

var MonitorDataType = provider.DataSourceType[Monitor, *Monitor]{
	ResourceIdentifier: "leanspace_monitors",
	Path:               "monitors-repository/monitors",
	Schema:             monitorSchema,
	FilterSchema:       dataSourceFilterSchema,
}
