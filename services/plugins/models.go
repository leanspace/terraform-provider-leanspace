package plugins

import "github.com/leanspace/terraform-provider-leanspace/provider"

type AbstractPlugin interface {
	GetID() string
	GetStatus() string
	SetStatus(status string)
	GetFilePath() string
	SetFilePath(filePath string)
	GetFileSha() string
	SetFileSha(fileSha string)

	// Persist the file path - this data is not returned from the backend, so when the resource
	// is loaded (from create/read/update) the path is empty, and so terraform thinks the field was
	// changed. This workaround prevents the value from changing - it's loaded by terraform
	// when reading the config and never changes again (except if the config changes).
	PersistFilePath(destPlugin AbstractPlugin) error
	PersistFileSha(destPlugin AbstractPlugin) error

	CallGetPlugin(client *provider.Client) (AbstractPlugin, error)
	CallReadProcess(client *provider.Client) ([]byte, error)
}
