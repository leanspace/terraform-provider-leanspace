package plugins

import "github.com/leanspace/terraform-provider-leanspace/provider"

func (plugin *Plugin) ToMap() map[string]any {
	pluginMap := make(map[string]any)
	pluginMap["id"] = plugin.ID
	pluginMap["type"] = plugin.Type
	pluginMap["implementation_class_name"] = plugin.ImplementationClassName
	pluginMap["name"] = plugin.Name
	pluginMap["description"] = plugin.Description
	pluginMap["source_code_file_download_authorized"] = plugin.SourceCodeFileDownloadAuthorized
	pluginMap["file_path"] = plugin.FilePath
	pluginMap["created_at"] = plugin.CreatedAt
	pluginMap["created_by"] = plugin.CreatedBy
	pluginMap["last_modified_at"] = plugin.LastModifiedAt
	pluginMap["last_modified_by"] = plugin.LastModifiedBy
	pluginMap["function_name"] = plugin.FunctionName
	pluginMap["source_code_file_id"] = plugin.SourceCodeFileId
	pluginMap["sdk_version"] = plugin.SdkVersion
	pluginMap["sdk_version_family"] = plugin.SdkVersionFamily
	pluginMap["status"] = plugin.Status
	return pluginMap
}

func (plugin *Plugin) FromMap(pluginMap map[string]any) error {
	plugin.ID = pluginMap["id"].(string)
	plugin.Type = pluginMap["type"].(string)
	plugin.ImplementationClassName = pluginMap["implementation_class_name"].(string)
	plugin.Name = pluginMap["name"].(string)
	plugin.Description = pluginMap["description"].(string)
	plugin.SourceCodeFileDownloadAuthorized = pluginMap["source_code_file_download_authorized"].(bool)
	plugin.FilePath = pluginMap["file_path"].(string)
	plugin.CreatedAt = pluginMap["created_at"].(string)
	plugin.CreatedBy = pluginMap["created_by"].(string)
	plugin.LastModifiedAt = pluginMap["last_modified_at"].(string)
	plugin.LastModifiedBy = pluginMap["last_modified_by"].(string)
	plugin.FunctionName = pluginMap["function_name"].(string)
	plugin.SourceCodeFileId = pluginMap["source_code_file_id"].(string)
	plugin.SdkVersion = pluginMap["sdk_version"].(string)
	plugin.SdkVersionFamily = pluginMap["sdk_version_family"].(string)
	plugin.Status = pluginMap["status"].(string)
	return nil
}

// Persist the file path - this data is not returned from the backend, so when the resource
// is loaded (from create/read/update) the path is empty, and so terraform thinks the field was
// changed. This workaround prevents the value from changing - it's loaded by terraform
// when reading the config and never changes again (except if the config changes).
func (plugin *Plugin) persistFilePath(destPlugin *Plugin) error {
	destPlugin.FilePath = plugin.FilePath
	return nil
}

func (plugin *Plugin) PostCreateProcess(_ *provider.Client, destPluginRaw any) error {
	return plugin.persistFilePath(destPluginRaw.(*Plugin))
}
func (plugin *Plugin) PostUpdateProcess(_ *provider.Client, destPluginRaw any) error {
	return plugin.persistFilePath(destPluginRaw.(*Plugin))
}
func (plugin *Plugin) PostReadProcess(_ *provider.Client, destPluginRaw any) error {
	return plugin.persistFilePath(destPluginRaw.(*Plugin))
}
