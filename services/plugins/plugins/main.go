package plugins

import (
	"fmt"

	"github.com/leanspace/terraform-provider-leanspace/provider"
)

var PluginDataType = provider.DataSourceType[Plugin, *Plugin]{
	ResourceIdentifier: "leanspace_plugins",
	Path:               "plugins-repository/plugins",
	Schema:             pluginSchema,
	FilterSchema:       dataSourceFilterSchema,
	ReadPath: func(id string) string {
		return fmt.Sprintf("plugins-repository/plugins/%s/metadata", id)
	},
}
