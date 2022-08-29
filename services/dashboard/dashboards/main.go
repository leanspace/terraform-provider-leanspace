package dashboards

import "github.com/leanspace/terraform-provider-leanspace/provider"

var DashboardDataType = provider.DataSourceType[Dashboard, *Dashboard]{
	ResourceIdentifier: "leanspace_dashboards",
	Path:               "dashboard-repository/dashboards",
	Schema:             dashboardSchema,
	FilterSchema:       dataSourceFilterSchema,
}
