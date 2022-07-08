package dashboards

import (
	"terraform-provider-asset/asset"
)

var DashboardDataType = asset.DataSourceType[Dashboard, *Dashboard]{
	ResourceIdentifier: "leanspace_dashboards",
	Name:               "dashboard",
	Path:               "dashboard-repository/dashboards",
	Schema:             dashboardSchema,
}
