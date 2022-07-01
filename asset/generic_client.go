package asset

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"terraform-provider-asset/asset/general_objects"
)

func (client GenericClient[T, PT]) GetAll() (*general_objects.PaginatedList[T], error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", client.Client.HostURL, client.Path), nil)
	if err != nil {
		return nil, err
	}

	body, err, _ := client.Client.doRequest(req, &(client.Client).Token)
	if err != nil {
		return nil, err
	}

	values := general_objects.PaginatedList[T]{}
	err = json.Unmarshal(body, &values)
	if err != nil {
		return nil, err
	}

	return &values, nil
}

func (client GenericClient[T, PT]) Get(id string) (PT, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/%s", client.Client.HostURL, client.Path, id), nil)
	if err != nil {
		return nil, err
	}

	body, err, _ := client.Client.doRequest(req, &(client.Client).Token)
	if err != nil {
		return nil, err
	}

	var value T
	err = json.Unmarshal(body, &value)
	if err != nil {
		return nil, err
	}

	return &value, nil
}

func (client GenericClient[T, PT]) Create(createElement PT) (PT, error) {
	rb, err := json.Marshal(createElement)
	if err != nil {
		return nil, err
	}

	path := client.Path
	if client.CreatePath != nil {
		path = client.CreatePath(createElement)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s", client.Client.HostURL, path), strings.NewReader(string(rb)))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}

	body, err, _ := client.Client.doRequest(req, &(client.Client).Token)
	if err != nil {
		return nil, err
	}

	var value T
	err = json.Unmarshal(body, &value)
	if err != nil {
		return nil, err
	}

	return &value, nil
}

func (client GenericClient[T, PT]) Update(nodeId string, createElement PT) (*T, error) {
	rb, err := json.Marshal(createElement)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/%s/%s", client.Client.HostURL, client.Path, nodeId), strings.NewReader(string(rb)))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}

	body, err, _ := client.Client.doRequest(req, &(client.Client).Token)
	if err != nil {
		return nil, err
	}

	var value T
	err = json.Unmarshal(body, &value)
	if err != nil {
		return nil, err
	}

	return &value, nil
}

func (client GenericClient[T, PT]) Delete(nodeId string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/%s/%s", client.Client.HostURL, client.Path, nodeId), nil)
	if err != nil {
		return err
	}

	_, err, statusCode := client.Client.doRequest(req, &(client.Client).Token)
	// If it has been deleted outside terraform, it should not fail here
	if statusCode != http.StatusNotFound && err != nil {
		return err
	}

	return nil
}
