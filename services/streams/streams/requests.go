package streams

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/leanspace/terraform-provider-leanspace/provider"
)

type apiStreamQueueInfo struct {
	Command Stream `json:"command"`
}

type apiStreamQueueResponse struct {
	Status   string `json:"status"`
	StreamId string `json:"streamId"`
}

func (stream *Stream) toAPIFormat() ([]byte, error) {
	streamQueue := apiStreamQueueInfo{
		Command: *stream,
	}
	return json.Marshal(streamQueue)
}

func (stream *Stream) CustomEncoding(data []byte, isUpdating bool) (io.Reader, string, error) {
	if isUpdating {
		return strings.NewReader(string(data)), "application/json", nil
	}
	streamQueueData, err := stream.toAPIFormat()
	if err != nil {
		return nil, "", err
	}
	return strings.NewReader(string(streamQueueData)), "application/json", nil
}

func (stream *Stream) PostCreateProcess(client *provider.Client, destStreamRaw any) error {
	createdStream := destStreamRaw.(*Stream)

	var streamQueue apiStreamQueueResponse
	currentStatus := "UNKNOWN"
	var byteData []byte

	// do ... while loop
	for ok := true; ok; ok = currentStatus != "SUCCEEDED" && currentStatus != "FAILED" {
		time.Sleep(1 * time.Second)
		originalByteData, streamQueuePointer, err := GetStreamQueue(createdStream.ID, client)
		if err != nil {
			return err
		}
		streamQueue = *streamQueuePointer
		currentStatus = streamQueue.Status
		byteData = *originalByteData
	}

	if currentStatus == "FAILED" {
		var jsonValue bytes.Buffer
		json.Indent(&jsonValue, byteData, "", "    ")
		return fmt.Errorf("Stream creation failed with errors: %s", jsonValue.String())
	}

	stream.ID = streamQueue.StreamId
	streamInfo, err := GetStream(stream.ID, client)
	if err != nil {
		return err
	}

	createdStream.ID = streamInfo.ID
	createdStream.Version = streamInfo.Version
	createdStream.Name = streamInfo.Name
	createdStream.Description = streamInfo.Description
	createdStream.AssetId = streamInfo.AssetId
	createdStream.Configuration = streamInfo.Configuration
	createdStream.Mappings = streamInfo.Mappings
	createdStream.CreatedAt = streamInfo.CreatedAt
	createdStream.CreatedBy = streamInfo.CreatedBy
	createdStream.LastModifiedAt = streamInfo.LastModifiedAt
	createdStream.LastModifiedBy = streamInfo.LastModifiedBy

	return nil
}

func GetStreamQueue(streamQueueId string, client *provider.Client) (*[]byte, *apiStreamQueueResponse, error) {
	path := fmt.Sprintf("%s/streams-repository/stream-queues/%s", client.HostURL, streamQueueId)
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}
	data, err, code := client.DoRequest(req, &(client).Token)
	if code == http.StatusNotFound {
		return nil, nil, nil
	}
	if err != nil {
		return nil, nil, err
	}
	var element apiStreamQueueResponse
	err = json.Unmarshal(data, &element)
	if err != nil {
		return nil, nil, err
	}
	return &data, &element, nil
}

func GetStream(streamId string, client *provider.Client) (*Stream, error) {
	path := fmt.Sprintf("%s/%s/%s", client.HostURL, StreamDataType.Path, streamId)
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
	var element Stream
	err = json.Unmarshal(data, &element)
	if err != nil {
		return nil, err
	}
	return &element, nil
}
