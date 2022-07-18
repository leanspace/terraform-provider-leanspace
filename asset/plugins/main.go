package plugins

import (
	"fmt"
	"terraform-provider-asset/asset"
)

var PluginDataType = asset.DataSourceType[Plugin, *Plugin]{
	ResourceIdentifier: "leanspace_plugins",
	Name:               "plugin",
	Path:               "plugins-repository/plugins",
	Schema:             pluginSchema,
	ReadPath: func(id string) string {
		return fmt.Sprintf("plugins-repository/plugins/%s/metadata", id)
	},
}
