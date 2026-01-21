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

// Persist the file sha - this data is not returned from the backend, so when the resource
// is loaded (from create/read/update) the path is empty, and so terraform thinks the field was
// changed. This workaround prevents the value from changing - it's processed by terraform
// when reading the config and never changes again (except if the file changes).
func CalculateFileSha(sourceCodePath string) (string, error) {
	if sourceCodePath == "" {
		return "", nil
	}
	fileData, err := os.ReadFile(sourceCodePath)
	if err != nil {
		return "absent", nil
	}
	hasher := sha256.New()
	hasher.Write(fileData)
	sourceCodeSha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	return sourceCodeSha, nil
}

func DoPostCreateProcess[P AbstractPlugin](client *provider.Client, currentPlugin AbstractPlugin, getRequestPath string, destPluginRaw any) error {
	createdPlugin := destPluginRaw.(AbstractPlugin)

	pluginRetrieved, err := createdPlugin.CallGetPlugin(client)
	if err != nil {
		return err
	}
	err = currentPlugin.PersistFileSha(createdPlugin)
	if err != nil {
		return err
	}
	err = currentPlugin.PersistFilePath(createdPlugin)
	if err != nil {
		return err
	}

	startTime := time.Now()
	for pluginRetrieved.GetStatus() != "ACTIVE" && pluginRetrieved.GetStatus() != "FAILED" && time.Since(startTime).Seconds() < client.RetryTimeout.Seconds() {
		pluginRetrieved, err = createdPlugin.CallGetPlugin(client)
		if err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}

	if pluginRetrieved.GetStatus() == "FAILED" {
		return fmt.Errorf("Plugin creation failed")
	}

	currentPlugin.SetStatus(pluginRetrieved.GetStatus())
	return nil
}

func GetPlugin[P any](id string, requestPath string, extraPath string, client *provider.Client) (*P, error) {
	path := fmt.Sprintf("%s/%s/%s%s", client.HostURL, requestPath, id, extraPath)
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
	var element P
	err = json.Unmarshal(data, &element)
	if err != nil {
		return nil, err
	}
	return &element, nil
}

func DoPostReadProcess[P AbstractPlugin](client *provider.Client, currentPlugin AbstractPlugin, destPluginRaw any) error {
	createdPlugin := destPluginRaw.(AbstractPlugin)
	err := currentPlugin.PersistFilePath(createdPlugin)
	if err != nil {
		return err
	}
	sourceCodeSha, _ := CalculateFileSha(currentPlugin.GetFilePath())
	currentPlugin.SetFileSha(sourceCodeSha)

	body, err := currentPlugin.CallReadProcess(client)
	if err != nil {
		return err
	}

	hasher := sha256.New()
	hasher.Write(body)
	createdPlugin.SetFileSha(base64.URLEncoding.EncodeToString(hasher.Sum(nil)))
	if createdPlugin.GetFileSha() != currentPlugin.GetFileSha() && currentPlugin.GetFilePath() != "" {
		createdPlugin.SetFilePath("file_changed") // this will cause the resource to be considered as changed
	}

	return nil
}
