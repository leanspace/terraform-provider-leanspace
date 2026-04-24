package generic_plugins

import (
	"github.com/leanspace/terraform-provider-leanspace/provider"
)

var GenericPluginDataType = provider.DataSourceType[GenericPlugin, *GenericPlugin]{
	ResourceIdentifier: "leanspace_generic_plugins",
	Path:               "plugins-repository/generic-plugins",
	Schema:             genericPluginSchema,
	FilterSchema:       dataSourceFilterSchema,
	TFModelFactory:     func() any { return &GenericPluginTF{} },
	SchemaVersion:      1,
	StateUpgraders:     provider.UnwrapSingleNestedUpgraderMap(genericPluginSchema),
}
