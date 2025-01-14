package streams_queue

import (
	"encoding/json"
	"fmt"
	"github.com/leanspace/terraform-provider-leanspace/helper"
	_ "github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/services/streams/streams"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/leanspace/terraform-provider-leanspace/provider"
)

type apiStreamQueueCreateInfo struct {
	Command streams.Stream `json:"command"`
}

type apiStreamQueueUpdateInfo struct {
	Command  streams.Stream `json:"command"`
	StreamId string         `json:"streamId"`
}

type apiStreamQueueResponse struct {
	Status   string `json:"status"`
	StreamId string `json:"streamId"`
}

type apiStreamQueueCreationResponse struct {
	ID string `json:"id"`
}

type StreamQueueInformation struct {
	Status        string
	StreamId      string
	StreamQueueId string
}

type searchStreamQueueResponse struct {
	Content []contentItem `json:"content"`
}

type contentItem struct {
	ID        string `json:"id"`
	StreamID  string `json:"streamId"`
	Status    string `json:"status"`
	RequestID string `json:"requestId"`
}

func toAPICreateFormat(stream *streamQueue) ([]byte, error) {
	streamQueue := apiStreamQueueCreateInfo{
		Command: stream.stream,
	}

	return json.Marshal(streamQueue)
}

func toAPIUpdateFormat(stream *streams.Stream) ([]byte, error) {
	streamQueue := apiStreamQueueUpdateInfo{
		Command:  *stream,
		StreamId: stream.ID,
	}
	return json.Marshal(streamQueue)
}

func CustomEncoding(streamQueue *streamQueue, data []byte, isUpdating bool) (io.Reader, string, error) {
	var streamQueueData []byte
	var err error

	streamQueue.stream.PreMarshallProcess()
	streamQueueData, err = toAPIUpdateFormat(&streamQueue.stream)
	if err != nil {
		return nil, "", err
	}
	return strings.NewReader(string(streamQueueData)), "application/json", nil
}

func CreateStream(stream *streamQueue, client *provider.Client) (*streamQueue, error) {
	streamQueueData, err := toAPICreateFormat(stream)
	if err != nil {
		return nil, err
	}

	requestContent, contentType, err := CustomEncoding(stream, streamQueueData, true)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/streams-repository/stream-queues", client.HostURL), requestContent)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	data, err, _ := client.DoRequest(req, &client.Token)
	if err != nil {
		return nil, err
	}

	var streamQueue apiStreamQueueCreationResponse
	if err := json.Unmarshal(data, &streamQueue); err != nil {
		return nil, err
	}
	streamId, err := waitForStreamQueueCompletion(streamQueue.ID, client)
	if err != nil {
		return nil, err
	}

	return fetchStreamInfo(nil, streamId, client, streamQueue.ID)
}

func UpdateStream(updatedStreamQueue *streamQueue, client *provider.Client, id string) (*streamQueue, error) {
	if err := performStreamUpdate(client, updatedStreamQueue); err != nil {
		return nil, err
	}
	streamId, err := waitForStreamQueueCompletion(updatedStreamQueue.streamQueueId, client)
	if err != nil {
		return nil, err
	}

	return fetchStreamInfo(&updatedStreamQueue.stream, streamId, client, updatedStreamQueue.streamQueueId)
}

func performStreamUpdate(client *provider.Client, streamQueue *streamQueue) error {
	updateData, err := json.Marshal(streamQueue.stream)
	if err != nil {
		return err
	}

	requestContent, contentType, err := CustomEncoding(streamQueue, updateData, true)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/streams-repository/stream-queues/%s", client.HostURL, streamQueue.streamQueueId), requestContent)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", contentType)
	_, err, _ = client.DoRequest(req, &client.Token)
	if err != nil {
		return err
	}

	return nil
}

func waitForStreamQueueCompletion(streamQueueId string, client *provider.Client) (string, error) {
	var streamQueueInfo *apiStreamQueueResponse
	var err error
	for {
		time.Sleep(1 * time.Second)
		streamQueueInfo, err = getStreamQueue(streamQueueId, client)
		if err != nil {
			return "", err
		}
		if streamQueueInfo.Status == "SUCCEEDED" {
			break
		} else if streamQueueInfo.Status == "FAILED" {
			return "", fmt.Errorf("stream queue processing failed for stream queue ID %s", streamQueueId)
		}
	}
	return streamQueueInfo.StreamId, nil
}

func fetchStreamInfo(stream *streams.Stream, streamId string, client *provider.Client, streamQueueId string) (*streamQueue, error) {
	if stream == nil {
		stream = &streams.Stream{}
	}
	streamInfo, err := getStream(streamId, client)
	if err != nil {
		return nil, err
	}
	updateStreamFields(stream, streamInfo)
	helper.Logger.Printf("fetchStreamInfo: %v", streamQueueId)
	helper.Logger.Printf("fetchStreamInfo: %v", stream)
	helper.Logger.Printf("fetchStreamInfo id: %v", stream.ID)
	stream.PostUnmarshallProcess()
	return &streamQueue{
		stream:        *stream,
		streamQueueId: streamQueueId,
	}, nil
}

func updateStreamFields(stream *streams.Stream, streamInfo *streams.Stream) {
	stream.ID = streamInfo.ID
	stream.Version = streamInfo.Version
	stream.Name = streamInfo.Name
	stream.Description = streamInfo.Description
	stream.AssetId = streamInfo.AssetId
	stream.Configuration = streamInfo.Configuration
	stream.Mappings = streamInfo.Mappings
	stream.CreatedAt = streamInfo.CreatedAt
	stream.CreatedBy = streamInfo.CreatedBy
	stream.LastModifiedAt = streamInfo.LastModifiedAt
	stream.LastModifiedBy = streamInfo.LastModifiedBy
}

func getStreamQueue(streamQueueId string, client *provider.Client) (*apiStreamQueueResponse, error) {
	path := fmt.Sprintf("%s/streams-repository/stream-queues/%s", client.HostURL, streamQueueId)
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	data, err, _ := client.DoRequest(req, &client.Token)
	if err != nil {
		return nil, err
	}

	var streamQueue apiStreamQueueResponse
	if err := json.Unmarshal(data, &streamQueue); err != nil {
		return nil, err
	}
	return &streamQueue, nil
}

func getStream(streamId string, client *provider.Client) (*streams.Stream, error) {
	path := fmt.Sprintf("%s/%s/%s", client.HostURL, path, streamId)
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	data, err, code := client.DoRequest(req, &client.Token)
	if code == http.StatusNotFound {
		return nil, fmt.Errorf("stream not found")
	}
	if err != nil {
		return nil, err
	}

	var stream streams.Stream
	if err := json.Unmarshal(data, &stream); err != nil {
		return nil, err
	}
	return &stream, nil

}
