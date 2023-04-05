package plugins

import (
	"io"
	"os"
	"net/http"

	"github.com/leanspace/terraform-provider-leanspace/provider"
)

func (plugin *Plugin) PostCreateProcess(client *provider.Client, created any) error {
	createdPlugin := created.(*Plugin)

	pluginFile, err := os.Open(plugin.FilePath)
    if err != nil {
        if os.IsNotExist(err) {
            return nil
        }
        return nil
    }

    info, err := pluginFile.Stat()
    if err != nil {
        panic(err)
    }

    _, err = uploadFile(createdPlugin.Url, info.Size(), pluginFile)
    if err != nil {
        panic(err)
    }

    return plugin.persistFilePath(created.(*Plugin))
}

func uploadFile(url string, contentLength int64, body io.Reader) (resp *http.Response, err error) {
	putRequest, err := http.NewRequest("PUT", url, body)
	if err != nil {
		return nil, err
	}
	putRequest.ContentLength = contentLength
	return http.DefaultClient.Do(putRequest)
}

// Persist the file path - this data is not returned from the backend, so when the resource
// is loaded (from create/read/update) the path is empty, and so terraform thinks the field was
// changed. This workaround prevents the value from changing - it's loaded by terraform
// when reading the config and never changes again (except if the config changes).
func (plugin *Plugin) persistFilePath(destPlugin *Plugin) error {
	destPlugin.FilePath = plugin.FilePath
	return nil
}

func (plugin *Plugin) PostUpdateProcess(_ *provider.Client, destPluginRaw any) error {
	return plugin.persistFilePath(destPluginRaw.(*Plugin))
}
func (plugin *Plugin) PostReadProcess(_ *provider.Client, destPluginRaw any) error {
	return plugin.persistFilePath(destPluginRaw.(*Plugin))
}
