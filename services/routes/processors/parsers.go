package processors

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/leanspace/terraform-provider-leanspace/helper"

	"github.com/leanspace/terraform-provider-leanspace/provider"
)

func (processor *Processor) ToMap() map[string]any {
	processorMap := processor.ToAuditMap()
	processorMap["name"] = helper.NilIfEmpty(processor.Name)
	processorMap["description"] = helper.NilIfEmpty(processor.Description)
	processorMap["version"] = helper.NilIfEmpty(processor.Version)
	processorMap["type"] = helper.NilIfEmpty(processor.Type)
	processorMap["file_path"] = helper.NilIfEmpty(processor.FilePath)
	processorMap["file_sha"] = helper.NilIfEmpty(processor.FileSha)
	return processorMap
}

func (processor *Processor) FromMap(processorMap map[string]any) error {
	processor.FromAuditMap(processorMap)
	processor.Name = helper.CastString(processorMap, "name")
	processor.Description = helper.CastString(processorMap, "description")
	processor.Version = helper.CastString(processorMap, "version")
	processor.Type = helper.CastString(processorMap, "type")
	processor.FilePath = helper.CastString(processorMap, "file_path")
	processor.FileSha = helper.CastString(processorMap, "file_sha")
	return nil
}

// Persist the file sha - this data is not returned from the backend, so when the resource
// is loaded (from create/read/update) the path is empty, and so terraform thinks the field was
// changed. This workaround prevents the value from changing - it's processed by terraform
// when reading the config and never changes again (except if the file changes).
func (processor *Processor) SetFileSha() error {
	fileData, err := os.ReadFile(processor.FilePath)
	if err != nil {
		processor.FileSha = "absent"
		return nil
	}
	hasher := sha256.New()
	hasher.Write(fileData)
	processor.FileSha = base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return nil
}

// Persist the file path - this data is not returned from the backend, so when the resource
// is loaded (from create/read/update) the path is empty, and so terraform thinks the field was
// changed. This workaround prevents the value from changing - it's loaded by terraform
// when reading the config and never changes again (except if the config changes).
func (processor *Processor) persistFilePath(destProcessor *Processor) error {
	destProcessor.FilePath = processor.FilePath
	return nil
}

func (processor *Processor) persistFileSha(destProcessor *Processor) error {
	err := processor.SetFileSha()
	if err != nil {
		return err
	}
	destProcessor.FileSha = processor.FileSha
	return nil
}

func (processor *Processor) PostCreateProcess(_ *provider.Client, destProcessorRaw any) error {
	createdProcessor := destProcessorRaw.(*Processor)
	err := processor.persistFilePath(createdProcessor)
	if err != nil {
		return err
	}
	return processor.persistFileSha(createdProcessor)
}

type apiValidProcessors struct {
	ProcessorIds []string `json:"processorIds"`
}

func (processor *Processor) PreDeleteProcess(client *provider.Client, destProcessorRaw any) error {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/routes-repository/processors/%s/routes", client.HostURL, processor.ID), nil)
	if err != nil {
		return err
	}

	body, err, _ := client.DoRequest(req, &(client).Token)
	if err != nil {
		return err
	}
	var element AttachedRoute
	err = json.Unmarshal(body, &element)
	if err != nil {
		return err
	}
	processorsIds, err := json.Marshal(apiValidProcessors{ProcessorIds: []string{processor.ID}})
	if err != nil {
		return err
	}

	for _, route := range element.Content {
		req, err = http.NewRequest("PUT", fmt.Sprintf("%s/routes-repository/routes/%s/processors", client.HostURL, route.ID), strings.NewReader(string(processorsIds)))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			return err
		}

		_, err, _ := client.DoRequest(req, &(client).Token)
		if err != nil {
			return err
		}
	}

	return nil
}
func (processor *Processor) PostUpdateProcess(_ *provider.Client, destProcessorRaw any) error {
	// No post update process needed, only delete and create and implemented
	return nil
}
func (processor *Processor) PostReadProcess(client *provider.Client, destProcessorRaw any) error {
	createdProcessor := destProcessorRaw.(*Processor)
	err := processor.persistFilePath(createdProcessor)
	if err != nil {
		return err
	}
	err = processor.persistFileSha(createdProcessor)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/routes-repository/processors/%s/generate-download-link", client.HostURL, createdProcessor.ID), nil)
	if err != nil {
		return err
	}

	body, err, _ := client.DoRequest(req, &(client).Token)
	if err != nil {
		return err
	}
	var element ProcessorUrl
	err = json.Unmarshal(body, &element)
	if err != nil {
		return err
	}

	req, err = http.NewRequest("GET", element.Url, nil)
	if err != nil {
		return err
	}

	body, err, _ = client.DoRequest(req, nil)
	if err != nil {
		return err
	}

	hasher := sha256.New()
	hasher.Write(body)
	createdProcessor.FileSha = base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	if createdProcessor.FileSha != processor.FileSha && processor.FileSha != "" {
		createdProcessor.FilePath = "file_changed" // this will cause the resource to be considered as changed
	}

	return nil
}
