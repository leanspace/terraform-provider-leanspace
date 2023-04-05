package plugins

func (plugin *Plugin) ToMap() map[string]any {
	pluginMap := make(map[string]any)
	pluginMap["id"] = plugin.ID
	pluginMap["type"] = plugin.Type
	pluginMap["implementation_class_name"] = plugin.ImplementationClassName
	pluginMap["name"] = plugin.Name
	pluginMap["description"] = plugin.Description
	pluginMap["source_code_file_name"] = plugin.SourceCodeFileName
	pluginMap["source_code_file_download_authorized"] = plugin.SourceCodeFileDownloadAuthorized
	pluginMap["file_path"] = plugin.FilePath
	pluginMap["created_at"] = plugin.CreatedAt
	pluginMap["created_by"] = plugin.CreatedBy
	pluginMap["last_modified_at"] = plugin.LastModifiedAt
	pluginMap["last_modified_by"] = plugin.LastModifiedBy
	pluginMap["sdk_version"] = plugin.SdkVersion
	pluginMap["sdk_version_family"] = plugin.SdkVersionFamily
	pluginMap["status"] = plugin.Status
	pluginMap["url"] = plugin.Url
	pluginMap["expires"] = plugin.Expires
	return pluginMap
}

func (plugin *Plugin) FromMap(pluginMap map[string]any) error {
	plugin.ID = pluginMap["id"].(string)
	plugin.Type = pluginMap["type"].(string)
	plugin.ImplementationClassName = pluginMap["implementation_class_name"].(string)
	plugin.Name = pluginMap["name"].(string)
	plugin.Description = pluginMap["description"].(string)
	plugin.SourceCodeFileName = pluginMap["source_code_file_name"].(string)
	plugin.SourceCodeFileDownloadAuthorized = pluginMap["source_code_file_download_authorized"].(bool)
	plugin.FilePath = pluginMap["file_path"].(string)
	plugin.CreatedAt = pluginMap["created_at"].(string)
	plugin.CreatedBy = pluginMap["created_by"].(string)
	plugin.LastModifiedAt = pluginMap["last_modified_at"].(string)
	plugin.LastModifiedBy = pluginMap["last_modified_by"].(string)
	plugin.SdkVersion = pluginMap["sdk_version"].(string)
	plugin.SdkVersionFamily = pluginMap["sdk_version_family"].(string)
	plugin.Status = pluginMap["status"].(string)
	plugin.Url = pluginMap["url"].(string)
	plugin.Expires = pluginMap["expires"].(string)
	return nil
}
