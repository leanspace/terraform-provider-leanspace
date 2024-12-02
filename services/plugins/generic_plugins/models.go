package generic_plugins

import (
	"io"
	"net/http"

	"github.com/leanspace/terraform-provider-leanspace/provider"
	"github.com/leanspace/terraform-provider-leanspace/services/plugins"
)

type GenericPlugin struct {
	ID             string         `json:"id"`
	Name           string         `json:"name"`
	Description    string         `json:"description"`
	Type           string         `json:"type"`
	Language       string         `json:"language"`
	SourceCodeLink SourceCodeLink `json:"sourceCodeLink"`
	CreatedAt      string         `json:"createdAt"`
	CreatedBy      string         `json:"createdBy"`
	LastModifiedAt string         `json:"lastModifiedAt"`
	LastModifiedBy string         `json:"lastModifiedBy"`
	Status         string         `json:"status"`
	FilePath       string         `json:"source_code_path"`
	FileSha        string         `json:"source_code_sha"`
}

func (genericPlugin GenericPlugin) GetID() string     { return genericPlugin.ID }
func (genericPlugin GenericPlugin) GetStatus() string { return genericPlugin.Status }
func (genericPlugin *GenericPlugin) SetStatus(status string) {
	genericPlugin.Status = status
}
func (genericPlugin GenericPlugin) GetFilePath() string { return genericPlugin.FilePath }
func (genericPlugin *GenericPlugin) SetFilePath(filePath string) {
	genericPlugin.FilePath = filePath
}
func (genericPlugin GenericPlugin) GetFileSha() string { return genericPlugin.FileSha }
func (genericPlugin *GenericPlugin) SetFileSha(fileSha string) {
	genericPlugin.FileSha = fileSha
}

type SourceCodeLink struct {
	ExpirationTime string `json:"expirationTime"`
	SourceCodeId   string `json:"sourceCodeId"`
	Url            string `json:"url"`
}

func (genericPlugin GenericPlugin) PersistFilePath(destPlugin plugins.AbstractPlugin) error {
	destPlugin.SetFilePath(genericPlugin.FilePath)
	return nil
}

func (genericPlugin GenericPlugin) PersistFileSha(destPlugin plugins.AbstractPlugin) error {
	sourceCodeSha, _ := plugins.CalculateFileSha(genericPlugin.FilePath)
	genericPlugin.FileSha = sourceCodeSha
	destPlugin.SetFileSha(genericPlugin.FileSha)
	return nil
}

func (genericPlugin GenericPlugin) CallGetPlugin(client *provider.Client) (plugins.AbstractPlugin, error) {
	return plugins.GetPlugin[GenericPlugin](genericPlugin.ID, GenericPluginDataType.Path, "", client)
}

func (genericPlugin GenericPlugin) CallReadProcess(client *provider.Client) ([]byte, error) {
	resp, err := http.Get(genericPlugin.SourceCodeLink.Url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, err
}
