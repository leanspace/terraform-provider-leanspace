package plugins

import (
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/provider"
	"github.com/leanspace/terraform-provider-leanspace/services/plugins"
)

func (plugin *Plugin) ToMap() map[string]any {
	pluginMap := plugin.ToAuditMap()
	pluginMap["type"] = helper.NilIfEmpty(plugin.Type)
	pluginMap["implementation_class_name"] = helper.NilIfEmpty(plugin.ImplementationClassName)
	pluginMap["name"] = helper.NilIfEmpty(plugin.Name)
	pluginMap["description"] = helper.NilIfEmpty(plugin.Description)
	pluginMap["source_code_file_download_authorized"] = helper.NilIfEmpty(plugin.SourceCodeFileDownloadAuthorized)
	pluginMap["file_path"] = helper.NilIfEmpty(plugin.FilePath)
	pluginMap["file_sha"] = helper.NilIfEmpty(plugin.FileSha)
	pluginMap["sdk_version"] = helper.NilIfEmpty(plugin.SdkVersion)
	pluginMap["sdk_version_family"] = helper.NilIfEmpty(plugin.SdkVersionFamily)
	pluginMap["status"] = helper.NilIfEmpty(plugin.Status)
	return pluginMap
}

func (plugin *Plugin) FromMap(pluginMap map[string]any) error {
	plugin.FromAuditMap(pluginMap)
	plugin.Type = helper.CastString(pluginMap, "type")
	plugin.ImplementationClassName = helper.CastString(pluginMap, "implementation_class_name")
	plugin.Name = helper.CastString(pluginMap, "name")
	plugin.Description = helper.CastString(pluginMap, "description")
	plugin.SourceCodeFileDownloadAuthorized = helper.CastBool(pluginMap, "source_code_file_download_authorized")
	plugin.FilePath = helper.CastString(pluginMap, "file_path")
	plugin.FileSha = helper.CastString(pluginMap, "file_sha")
	plugin.SdkVersion = helper.CastString(pluginMap, "sdk_version")
	plugin.SdkVersionFamily = helper.CastString(pluginMap, "sdk_version_family")
	plugin.Status = helper.CastString(pluginMap, "status")
	return nil
}

func (plugin *Plugin) PostCreateProcess(client *provider.Client, destPluginRaw any) error {
	return plugins.DoPostCreateProcess[*Plugin](client, plugin, PluginDataType.Path+"/metadata", destPluginRaw)
}

func (plugin *Plugin) PostUpdateProcess(client *provider.Client, destPluginRaw any) error {
	return plugin.PostCreateProcess(client, destPluginRaw)
}

func (plugin *Plugin) PostReadProcess(client *provider.Client, destPluginRaw any) error {
	return plugins.DoPostReadProcess[*Plugin](client, plugin, destPluginRaw)
}
