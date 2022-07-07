package widgets

import (
	"terraform-provider-asset/asset"
)

var WidgetDataType = asset.DataSourceType[Widget, *Widget]{
	ResourceIdentifier: "leanspace_widgets",
	Name:               "widget",
	Path:               "dashboard-repository/widgets",
	Schema:             widgetSchema,
}
