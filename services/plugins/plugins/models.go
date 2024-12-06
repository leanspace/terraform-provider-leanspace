package plugins

import (
	"fmt"
	"net/http"

	"github.com/leanspace/terraform-provider-leanspace/provider"
	"github.com/leanspace/terraform-provider-leanspace/services/plugins"
)

type Plugin struct {
	ID                               string `json:"id"`
	Type                             string `json:"type"`
	ImplementationClassName          string `json:"implementationClassName"`
	Name                             string `json:"name"`
	Description                      string `json:"description"`
	SourceCodeFileDownloadAuthorized bool   `json:"sourceCodeFileDownloadAuthorized,omitempty"`
	FilePath                         string `json:"filePath"`
	CreatedAt                        string `json:"createdAt"`
	CreatedBy                        string `json:"createdBy"`
	LastModifiedAt                   string `json:"lastModifiedAt"`
	LastModifiedBy                   string `json:"lastModifiedBy"`
	SdkVersion                       string `json:"sdkVersion,omitempty"`
	SdkVersionFamily                 string `json:"sdkVersionFamily"`
	Status                           string `json:"status"`
	FileSha                          string `json:"fileSha"`
}

func (plugin Plugin) GetID() string     { return plugin.ID }
func (plugin Plugin) GetStatus() string { return plugin.Status }
func (plugin *Plugin) SetStatus(status string) {
	plugin.Status = status
}
func (plugin Plugin) GetFilePath() string { return plugin.FilePath }
func (plugin *Plugin) SetFilePath(filePath string) {
	plugin.FilePath = filePath
}
func (plugin Plugin) GetFileSha() string { return plugin.FileSha }
func (plugin *Plugin) SetFileSha(fileSha string) {
	plugin.FileSha = fileSha
}

func (plugin Plugin) PersistFilePath(destPlugin plugins.AbstractPlugin) error {
	destPlugin.SetFilePath(plugin.FilePath)
	return nil
}

func (plugin Plugin) PersistFileSha(destPlugin plugins.AbstractPlugin) error {
	sourceCodeSha, _ := plugins.CalculateFileSha(plugin.FilePath)
	plugin.FileSha = sourceCodeSha
	destPlugin.SetFileSha(plugin.FileSha)
	return nil
}

func (plugin Plugin) CallGetPlugin(client *provider.Client) (plugins.AbstractPlugin, error) {
	return plugins.GetPlugin[Plugin](plugin.ID, PluginDataType.Path, "/metadata", client)
}

func (plugin Plugin) CallReadProcess(client *provider.Client) ([]byte, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/plugins-repository/plugins/%s/source-code-file", client.HostURL, plugin.GetID()), nil)
	if err != nil {
		return nil, err
	}

	body, err, _ := client.DoRequest(req, &(client).Token)
	if err != nil {
		return nil, err
	}

	return body, err
}
