package plugins

type Plugin struct {
	ID                               string `json:"id"`
	Type                             string `json:"type"`
	ImplementationClassName          string `json:"implementationClassName"`
	Name                             string `json:"name"`
	Description                      string `json:"description"`
	SourceCodeFileDownloadAuthorized bool   `json:"sourceCodeFileDownloadAuthorized"`
	FilePath                         string `json:"filePath"`
	CreatedAt                        string `json:"createdAt"`
	CreatedBy                        string `json:"createdBy"`
	LastModifiedAt                   string `json:"lastModifiedAt"`
	LastModifiedBy                   string `json:"lastModifiedBy"`
}

func (plugin *Plugin) GetID() string { return plugin.ID }
