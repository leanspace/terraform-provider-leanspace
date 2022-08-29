package widgets

import "github.com/leanspace/terraform-provider-leanspace/provider"

var WidgetDataType = provider.DataSourceType[Widget, *Widget]{
	ResourceIdentifier: "leanspace_widgets",
	Path:               "dashboard-repository/widgets",
	Schema:             widgetSchema,
	FilterSchema:       dataSourceFilterSchema,
}
