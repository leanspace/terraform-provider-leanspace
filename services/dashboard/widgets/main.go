package widgets

import "leanspace-terraform-provider/provider"

var WidgetDataType = provider.DataSourceType[Widget, *Widget]{
	ResourceIdentifier: "leanspace_widgets",
	Name:               "widget",
	Path:               "dashboard-repository/widgets",
	Schema:             widgetSchema,
}
