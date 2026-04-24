package plugins

import (
	"io"

	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/provider"
	pluginscommon "github.com/leanspace/terraform-provider-leanspace/services/plugins"
)

func (plugin *Plugin) CustomEncoding(data []byte, isUpdating bool) (io.Reader, string, error) {
	return helper.FileAndDataToMultipart(plugin.FilePath, data)
}

func (plugin *Plugin) PostCreateProcess(client *provider.Client, destPluginRaw any) error {
	return pluginscommon.DoPostCreateProcess[*Plugin](client, plugin, PluginDataType.Path+"/metadata", destPluginRaw)
}

func (plugin *Plugin) PostUpdateProcess(client *provider.Client, destPluginRaw any) error {
	return plugin.PostCreateProcess(client, destPluginRaw)
}

func (plugin *Plugin) PostReadProcess(client *provider.Client, destPluginRaw any) error {
	return pluginscommon.DoPostReadProcess[*Plugin](client, plugin, destPluginRaw)
}
