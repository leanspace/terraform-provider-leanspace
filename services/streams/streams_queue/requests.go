package streams_queue

import (
	"encoding/json"
	"fmt"
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/leanspace/terraform-provider-leanspace/provider"
)

type apiStreamQueueCreateInfo struct {
	Command Stream `json:"command"`
}

type apiStreamQueueUpdateInfo struct {
	Command  Stream `json:"command"`
	StreamId string `json:"streamId"`
}

type apiStreamQueueResponse struct {
	Status   string `json:"status"`
	StreamId string `json:"streamId"`
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

func (stream *Stream) toAPICreateFormat() ([]byte, error) {
	streamQueue := apiStreamQueueCreateInfo{
		Command: *stream,
	}
	return json.Marshal(streamQueue)
}

func (stream *Stream) toAPIUpdateFormat() ([]byte, error) {
	streamQueue := apiStreamQueueUpdateInfo{
		Command:  *stream,
		StreamId: stream.ID,
	}
	return json.Marshal(streamQueue)
}

func (stream *Stream) CustomEncoding(data []byte, isUpdating bool) (io.Reader, string, error) {
	var streamQueueData []byte
	var err error
	if isUpdating {
		stream.PreMarshallProcess()
		streamQueueData, err = stream.toAPIUpdateFormat()
	} else {
		streamQueueData, err = stream.toAPICreateFormat()
	}
	if err != nil {
		return nil, "", err
	}
	helper.Logger.Printf("streamQueueData %s", streamQueueData)
	return strings.NewReader(string(streamQueueData)), "application/json", nil
}

func (stream *Stream) UpdateStream(client *provider.Client, id string, updatedStream *Stream) (*Stream, error) {
	streamQueueInfo, err := fetchStreamQueueInfo(stream.ID, client)
	helper.Logger.Printf("streamQueueInfo %v", streamQueueInfo)
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

	return fetchUpdatedStreamInfo(updatedStream, streamId, client)
}

func performStreamUpdate(client *provider.Client, streamQueueId string, updatedStream *Stream) error {
	updateData, err := json.Marshal(updatedStream)
	if err != nil {
		return err
	}

	requestContent, contentType, err := updatedStream.CustomEncoding(updateData, true)
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

func (stream *Stream) PostCreateProcess(client *provider.Client, destStreamRaw any) error {
	createdStream := destStreamRaw.(*Stream)
	streamId, err := waitForStreamQueueCompletion(createdStream.ID, client)
	if err != nil {
		return err
	}
	updateStreamInfo(createdStream, streamId, client)
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

func fetchUpdatedStreamInfo(stream *Stream, streamId string, client *provider.Client) (*Stream, error) {
	streamInfo, err := getStream(streamId, client)
	if err != nil {
		return nil, err
	}
	updateStreamFields(stream, streamInfo)
	return stream, nil
}

func updateStreamInfo(stream *Stream, streamId string, client *provider.Client) error {
	streamInfo, err := getStream(streamId, client)
	if err != nil {
		return err
	}
	updateStreamFields(stream, streamInfo)
	stream.PostUnmarshallProcess()
	return nil
}

func updateStreamFields(stream *Stream, streamInfo *Stream) {
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

func getStream(streamId string, client *provider.Client) (*Stream, error) {
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

	var stream Stream
	if err := json.Unmarshal(data, &stream); err != nil {
		return nil, err
	}
	return &stream, nil
}

func fetchStreamQueueInfo(streamId string, client *provider.Client) (*StreamQueueInformation, error) {
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
		helper.Logger.Printf("element content %v", response.Content)
		return nil, fmt.Errorf("invalid number of stream queues found for stream ID %s", streamId)
	}

	item := response.Content[0]
	return &StreamQueueInformation{
		Status:        item.Status,
		StreamId:      item.StreamID,
		StreamQueueId: item.ID,
	}, nil
}
