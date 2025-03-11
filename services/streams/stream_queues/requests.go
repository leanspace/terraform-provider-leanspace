package stream_queues

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/leanspace/terraform-provider-leanspace/services/streams/streams"

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

type streamQueueInformation struct {
	Status        string
	StreamId      string
	StreamQueueId string
}

type searchStreamQueueResponse struct {
	Content []streamQueueContentItem `json:"content"`
}

type streamQueueContentItem struct {
	ID        string `json:"id"`
	StreamID  string `json:"streamId"`
	Status    string `json:"status"`
	RequestID string `json:"requestId"`
}

func toAPICreateFormat(stream *streams.Stream) ([]byte, error) {
	streamQueue := apiStreamQueueCreateInfo{
		Command: *stream,
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

func CustomEncoding(stream *streams.Stream, data []byte, isUpdating bool) (io.Reader, string, error) {
	var streamQueueData []byte
	var err error

	stream.PreMarshallProcess()
	streamQueueData, err = toAPIUpdateFormat(stream)
	if err != nil {
		return nil, "", err
	}
	return strings.NewReader(string(streamQueueData)), "application/json", nil
}

func CreateStream(stream *streams.Stream, client *provider.Client) (*streams.Stream, error) {
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

	return fetchStreamInfo(nil, streamId, client)
}

func UpdateStream(updatedStream *streams.Stream, client *provider.Client, id string) (*streams.Stream, error) {
	streamQueueInfo, err := fetchStreamQueueInfo(updatedStream.ID, client)
	if err != nil {
		return nil, err
	}

	if err := performStreamUpdate(client, streamQueueInfo.StreamQueueId, updatedStream); err != nil {
		return nil, err
	}
	streamId, err := waitForStreamQueueCompletion(streamQueueInfo.StreamQueueId, client)
	if err != nil {
		return nil, err
	}

	return fetchStreamInfo(updatedStream, streamId, client)
}

func performStreamUpdate(client *provider.Client, streamQueueId string, updatedStream *streams.Stream) error {
	updateData, err := json.Marshal(updatedStream)
	if err != nil {
		return err
	}

	requestContent, contentType, err := CustomEncoding(updatedStream, updateData, true)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/streams-repository/stream-queues/%s", client.HostURL, streamQueueId), requestContent)
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
	var originalByteData *[]byte
	for {
		time.Sleep(1 * time.Second)
		originalByteData, streamQueueInfo, err = getStreamQueue(streamQueueId, client)
		if err != nil {
			return "", err
		}
		if streamQueueInfo.Status == "SUCCEEDED" {
			return streamQueueInfo.StreamId, nil
		} else if streamQueueInfo.Status == "FAILED" {
			var jsonValue bytes.Buffer
			json.Indent(&jsonValue, *originalByteData, "", "    ")
			return "", fmt.Errorf("stream queue processing failed for stream queue ID %s with errors %s", streamQueueId, jsonValue.String())
		}
	}

}

func fetchStreamInfo(stream *streams.Stream, streamId string, client *provider.Client) (*streams.Stream, error) {
	if stream == nil {
		stream = &streams.Stream{}
	}
	streamInfo, err := getStream(streamId, client)
	if err != nil {
		return nil, err
	}
	updateStreamFields(stream, streamInfo)
	return stream, nil
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

func getStreamQueue(streamQueueId string, client *provider.Client) (*[]byte, *apiStreamQueueResponse, error) {
	path := fmt.Sprintf("%s/streams-repository/stream-queues/%s", client.HostURL, streamQueueId)
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	data, err, _ := client.DoRequest(req, &client.Token)
	if err != nil {
		return nil, nil, err
	}

	var streamQueue apiStreamQueueResponse
	if err := json.Unmarshal(data, &streamQueue); err != nil {
		return nil, nil, err
	}
	return &data, &streamQueue, nil
}

func getStream(streamId string, client *provider.Client) (*streams.Stream, error) {
	path := fmt.Sprintf("%s/%s/%s", client.HostURL, Path, streamId)
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

func fetchStreamQueueInfo(streamId string, client *provider.Client) (*streamQueueInformation, error) {
	path := fmt.Sprintf("%s/streams-repository/stream-queues?streamIds=%s", client.HostURL, streamId)
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	data, err, code := client.DoRequest(req, &client.Token)
	if code == http.StatusNotFound {
		return nil, fmt.Errorf("stream queue not found")
	}
	if err != nil {
		return nil, err
	}

	var response searchStreamQueueResponse
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, err
	}

	if len(response.Content) != 1 {
		return nil, fmt.Errorf("invalid number of stream queues found for stream ID %s", streamId)
	}

	item := response.Content[0]
	return &streamQueueInformation{
		Status:        item.Status,
		StreamId:      item.StreamID,
		StreamQueueId: item.ID,
	}, nil
}
