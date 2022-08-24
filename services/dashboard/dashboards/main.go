package dashboards

import "leanspace-terraform-provider/provider"

var DashboardDataType = provider.DataSourceType[Dashboard, *Dashboard]{
	ResourceIdentifier: "leanspace_dashboards",
	Path:               "dashboard-repository/dashboards",
	Schema:             dashboardSchema,
}
