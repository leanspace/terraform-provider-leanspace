package plugins

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/leanspace/terraform-provider-leanspace/provider"
)

func (plugin *Plugin) ToMap() map[string]any {
	pluginMap := make(map[string]any)
	pluginMap["id"] = plugin.ID
	pluginMap["type"] = plugin.Type
	pluginMap["implementation_class_name"] = plugin.ImplementationClassName
	pluginMap["name"] = plugin.Name
	pluginMap["description"] = plugin.Description
	pluginMap["source_code_file_download_authorized"] = plugin.SourceCodeFileDownloadAuthorized
	pluginMap["file_path"] = plugin.FilePath
	pluginMap["file_sha"] = plugin.FileSha
	pluginMap["created_at"] = plugin.CreatedAt
	pluginMap["created_by"] = plugin.CreatedBy
	pluginMap["last_modified_at"] = plugin.LastModifiedAt
	pluginMap["last_modified_by"] = plugin.LastModifiedBy
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
	plugin.FileSha = pluginMap["file_sha"].(string)
	plugin.CreatedAt = pluginMap["created_at"].(string)
	plugin.CreatedBy = pluginMap["created_by"].(string)
	plugin.LastModifiedAt = pluginMap["last_modified_at"].(string)
	plugin.LastModifiedBy = pluginMap["last_modified_by"].(string)
	plugin.SdkVersion = pluginMap["sdk_version"].(string)
	plugin.SdkVersionFamily = pluginMap["sdk_version_family"].(string)
	plugin.Status = pluginMap["status"].(string)
	return nil
}

// Persist the file sha - this data is not returned from the backend, so when the resource
// is loaded (from create/read/update) the path is empty, and so terraform thinks the field was
// changed. This workaround prevents the value from changing - it's processed by terraform
// when reading the config and never changes again (except if the file changes).
func (plugin *Plugin) SetFileSha() error {
	if plugin.FilePath == "" {
		return nil
	}
	fileData, err := os.ReadFile(plugin.FilePath)
	if err != nil {
		return err
	}
	hasher := sha256.New()
	hasher.Write(fileData)
	plugin.FileSha = base64.URLEncoding.EncodeToString(hasher.Sum(nil))
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
func (plugin *Plugin) persistFileSha(destPlugin *Plugin) error {
	plugin.SetFileSha()
	destPlugin.FileSha = plugin.FileSha
	return nil
}

func (plugin *Plugin) PostCreateProcess(client *provider.Client, destPluginRaw any) error {
	createdPlugin := destPluginRaw.(*Plugin)

	metadata, err := GetPluginMetadata(createdPlugin.ID, client)
	plugin.persistFileSha(createdPlugin)
	plugin.persistFilePath(createdPlugin)
	if err != nil {
		return nil
	}

	startTime := time.Now()
	for metadata.Status != "ACTIVE" && metadata.Status != "FAILED" && time.Since(startTime).Seconds() < client.RetryTimeout.Seconds() {
		metadata, err = GetPluginMetadata(createdPlugin.ID, client)
		if err != nil {
			break
		}
		time.Sleep(1 * time.Second)
	}

	plugin.Status = metadata.Status
	return nil
}

func (plugin *Plugin) PostUpdateProcess(client *provider.Client, destPluginRaw any) error {
	return plugin.PostCreateProcess(client, destPluginRaw)
}
func (plugin *Plugin) PostReadProcess(client *provider.Client, destPluginRaw any) error {
	createdPlugin := destPluginRaw.(*Plugin)
	plugin.persistFilePath(createdPlugin)
	plugin.SetFileSha()

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/plugins-repository/plugins/%s/source-code-file", client.HostURL, createdPlugin.ID), nil)
	if err != nil {
		return err
	}

	body, err, _ := client.DoRequest(req, &(client).Token)
	if err != nil {
		return err
	}
	hasher := sha256.New()
	hasher.Write(body)
	createdPlugin.FileSha = base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	if createdPlugin.FileSha != plugin.FileSha && plugin.FilePath != "" {
		createdPlugin.FilePath = "file_changed" // this will cause the resource to be considered as changed
	}

	return nil
}

func GetPluginMetadata(pluginId string, client *provider.Client) (*Plugin, error) {
	path := fmt.Sprintf("%s/%s/%s/metadata", client.HostURL, PluginDataType.Path, pluginId)
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}
	data, err, code := client.DoRequest(req, &(client).Token)
	if code == http.StatusNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	var element Plugin
	err = json.Unmarshal(data, &element)
	if err != nil {
		return nil, err
	}
	return &element, nil
}
