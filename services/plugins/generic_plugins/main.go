package generic_plugins

import (
	"github.com/leanspace/terraform-provider-leanspace/provider"
)

var GenericPluginDataType = provider.DataSourceType[GenericPlugin, *GenericPlugin]{
	ResourceIdentifier: "leanspace_generic_plugins",
	Path:               "plugins-repository/generic-plugins",
	Schema:             genericPluginSchema,
	FilterSchema:       dataSourceFilterSchema,
}
