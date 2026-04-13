package generic_plugins

import (
	"github.com/leanspace/terraform-provider-leanspace/provider"
	"github.com/leanspace/terraform-provider-leanspace/services/plugins"
)

func (genericPlugin *GenericPlugin) PostCreateProcess(client *provider.Client, destGenericPluginRaw any) error {
	return plugins.DoPostCreateProcess[*GenericPlugin](client, genericPlugin, GenericPluginDataType.Path, destGenericPluginRaw)
}

func (plugin *GenericPlugin) PostUpdateProcess(client *provider.Client, destPluginRaw any) error {
	return plugin.PostCreateProcess(client, destPluginRaw)
}

func (genericPlugin *GenericPlugin) PostReadProcess(client *provider.Client, destPluginRaw any) error {
	createdGenericPlugin := destPluginRaw.(*GenericPlugin)
	genericPlugin.SourceCodeLink = createdGenericPlugin.SourceCodeLink
	return plugins.DoPostReadProcess[*GenericPlugin](client, genericPlugin, destPluginRaw)
}
