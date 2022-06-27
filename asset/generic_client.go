package asset

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (resource GenericResourceType[T]) GetAll() (*PaginatedList[T], error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", resource.Client.HostURL, resource.Path), nil)
	if err != nil {
		return nil, err
	}

	body, err, _ := resource.Client.doRequest(req, &(resource.Client).Token)
	if err != nil {
		return nil, err
	}

	elements := PaginatedList[T]{}
	err = json.Unmarshal(body, &elements)
	if err != nil {
		return nil, err
	}

	return &elements, nil
}

func (resource GenericResourceType[T]) Get(id string) (*T, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/%s", resource.Client.HostURL, resource.Path, id), nil)
	if err != nil {
		return nil, err
	}

	body, err, _ := resource.Client.doRequest(req, &(resource.Client).Token)
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

func (resource GenericResourceType[T]) Create(createElement T) (*T, error) {
	rb, err := json.Marshal(createElement)
	if err != nil {
		return nil, err
	}

	path := resource.Path
	if resource.CreatePath != nil {
		path = resource.CreatePath(createElement)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s", resource.Client.HostURL, path), strings.NewReader(string(rb)))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}

	body, err, _ := resource.Client.doRequest(req, &(resource.Client).Token)
	if err != nil {
		return nil, err
	}

	var element T
	err = json.Unmarshal(body, &element)
	if err != nil {
		return nil, err
	}

	return &element, nil
}

func (resource GenericResourceType[T]) Update(nodeId string, createElement T) (*T, error) {
	rb, err := json.Marshal(createElement)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/%s/%s", resource.Client.HostURL, resource.Path, nodeId), strings.NewReader(string(rb)))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}

	body, err, _ := resource.Client.doRequest(req, &(resource.Client).Token)
	if err != nil {
		return nil, err
	}

	var element T
	err = json.Unmarshal(body, &element)
	if err != nil {
		return nil, err
	}

	return &element, nil
}

func (resource GenericResourceType[T]) Delete(nodeId string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/%s/%s", resource.Client.HostURL, resource.Path, nodeId), nil)
	if err != nil {
		return err
	}

	_, err, statusCode := resource.Client.doRequest(req, &(resource.Client).Token)
	// If it has been deleted outside terraform, it should not fail here
	if statusCode != http.StatusNotFound && err != nil {
		return err
	}

	return nil
}
