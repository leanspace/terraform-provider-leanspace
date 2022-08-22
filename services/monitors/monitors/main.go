package monitors

import "leanspace-terraform-provider/provider"

var MonitorDataType = provider.DataSourceType[Monitor, *Monitor]{
	ResourceIdentifier: "leanspace_monitors",
	Path:               "monitors-repository/monitors",
	Schema:             monitorSchema,
}
