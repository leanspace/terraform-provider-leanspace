package processors

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/leanspace/terraform-provider-leanspace/provider"
)

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
		createdProcessor.FilePath = "file_changed"
	}
	return nil
}
