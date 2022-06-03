package asset

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (client *Client) GetAllProperties() (*PaginatedList[Property[interface{}]], error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/asset-repository/properties", client.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err, _ := client.doRequest(req, &client.Token)
	if err != nil {
		return nil, err
	}

	properties := PaginatedList[Property[interface{}]]{}
	err = json.Unmarshal(body, &properties)
	if err != nil {
		return nil, err
	}

	return &properties, nil
}

func (client *Client) GetProperty(propertyId string) (*Property[interface{}], error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/asset-repository/properties/%s", client.HostURL, propertyId), nil)
	if err != nil {
		return nil, err
	}

	body, err, _ := client.doRequest(req, &client.Token)
	if err != nil {
		return nil, err
	}

	property := Property[interface{}]{}
	err = json.Unmarshal(body, &property)
	if err != nil {
		return nil, err
	}

	return &property, nil
}

func (client *Client) CreateProperty(assetId string, createProperty Property[interface{}]) (*Property[interface{}], error) {
	rb, err := json.Marshal(createProperty)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/asset-repository/nodes/%s/properties", client.HostURL, assetId), strings.NewReader(string(rb)))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}

	body, err, _ := client.doRequest(req, &client.Token)
	if err != nil {
		return nil, err
	}

	property := Property[interface{}]{}
	err = json.Unmarshal(body, &property)
	if err != nil {
		return nil, err
	}

	return &property, nil
}

func (client *Client) UpdateProperty(propertyId string, createProperty Property[interface{}]) (*Property[interface{}], error) {
	rb, err := json.Marshal(createProperty)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/asset-repository/properties/%s", client.HostURL, propertyId), strings.NewReader(string(rb)))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}

	body, err, _ := client.doRequest(req, &client.Token)
	if err != nil {
		return nil, err
	}

	property := Property[interface{}]{}
	err = json.Unmarshal(body, &property)
	if err != nil {
		return nil, err
	}

	return &property, nil
}

func (client *Client) DeleteProperty(propertyId string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/asset-repository/properties/%s", client.HostURL, propertyId), nil)
	if err != nil {
		return err
	}

	_, err, statusCode := client.doRequest(req, &client.Token)
	// If it as been deleted outside terraform, it should not fail here
	if statusCode != http.StatusNotFound && err != nil {
		return err
	}

	return nil
}
