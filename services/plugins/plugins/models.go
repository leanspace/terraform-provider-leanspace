package plugins

type Plugin struct {
	ID                               string `json:"id" terra:"id"`
	Type                             string `json:"type" terra:"type"`
	ImplementationClassName          string `json:"implementationClassName" terra:"implementation_class_name"`
	Name                             string `json:"name" terra:"name"`
	Description                      string `json:"description" terra:"description"`
	SourceCodeFileDownloadAuthorized bool   `json:"sourceCodeFileDownloadAuthorized" terra:"source_code_file_download_authorized"`
	FilePath                         string `json:"filePath" terra:"file_path"`
	CreatedAt                        string `json:"createdAt" terra:"created_at"`
	CreatedBy                        string `json:"createdBy" terra:"created_by"`
	LastModifiedAt                   string `json:"lastModifiedAt" terra:"last_modified_at"`
	LastModifiedBy                   string `json:"lastModifiedBy" terra:"last_modified_by"`
}

func (plugin *Plugin) GetID() string { return plugin.ID }
