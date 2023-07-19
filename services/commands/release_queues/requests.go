package release_queues

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
	"github.com/leanspace/terraform-provider-leanspace/provider"
)

type apiCreateOrUpdateGlobalTransmissionMetatadata struct {
	GlobalTransmissionMetadata []general_objects.KeyValue `json:"globalTransmissionMetadata"`
}

func (queue *ReleaseQueue) toAPIFormat() ([]byte, error) {
	createOrUpdateGlobalTransmissionMetadata := apiCreateOrUpdateGlobalTransmissionMetatadata{
		GlobalTransmissionMetadata: queue.GlobalTransmissionMetadata,
	}
	return json.Marshal(createOrUpdateGlobalTransmissionMetadata)
}

func createOrUpdateGlobalTransmissionMetadata(queue *ReleaseQueue, releaseQueueId string, client *provider.Client) error {
	data, err := queue.toAPIFormat()
	if err != nil {
		return err
	}
	path := fmt.Sprintf("%s/%s/%s/global-transmission-metadata", client.HostURL, ReleaseQueueDataType.Path, releaseQueueId)
	req, err := http.NewRequest("PUT", path, strings.NewReader(string(data)))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	_, err, _ = client.DoRequest(req, &(client).Token)
	if err != nil {
		return err
	}
	return nil
}

func deleteGlobalTransmissionMetadata(queue *ReleaseQueue, releaseQueueId string, key string, client *provider.Client) error {
	path := fmt.Sprintf("%s/%s/%s/global-transmission-metadata/%s", client.HostURL, ReleaseQueueDataType.Path, releaseQueueId, key)
	req, err := http.NewRequest("DELETE", path, nil)
	if err != nil {
		return err
	}
	_, err, _ = client.DoRequest(req, &(client).Token)
	if err != nil {
		return err
	}
	return nil
}

func (queue *ReleaseQueue) PostCreateProcess(client *provider.Client, created any) error {
	createdQueue := created.(*ReleaseQueue)
	if len(queue.GlobalTransmissionMetadata) == 0 {
		return nil
	}
	if err := createOrUpdateGlobalTransmissionMetadata(queue, createdQueue.ID, client); err != nil {
		return err
	}
	return nil
}

func (queue *ReleaseQueue) PostUpdateProcess(client *provider.Client, updated any) error {
	updatedQueue := updated.(*ReleaseQueue)
	for _, globalTransmissionMetadata := range updatedQueue.GlobalTransmissionMetadata { // Remove extra global transmission metadata
		if !helper.Contains(queue.GlobalTransmissionMetadata, globalTransmissionMetadata) {
			if err := deleteGlobalTransmissionMetadata(queue, updatedQueue.ID, globalTransmissionMetadata.Key, client); err != nil {
				return err
			}
		}
	}
	return queue.PostCreateProcess(client, updated) // Add needed global transmission metadata
}
