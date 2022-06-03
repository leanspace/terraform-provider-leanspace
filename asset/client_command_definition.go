package asset

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (client *Client) GetAllCommandDefinitions() (*PaginatedList[CommandDefinition], error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/asset-repository/command-definitions", client.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err, _ := client.doRequest(req, &client.Token)
	if err != nil {
		return nil, err
	}

	commandDefinitions := PaginatedList[CommandDefinition]{}
	err = json.Unmarshal(body, &commandDefinitions)
	if err != nil {
		return nil, err
	}

	return &commandDefinitions, nil
}

func (client *Client) GetCommandDefinition(commandDefinitionId string) (*CommandDefinition, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/asset-repository/command-definitions/%s", client.HostURL, commandDefinitionId), nil)
	if err != nil {
		return nil, err
	}

	body, err, _ := client.doRequest(req, &client.Token)
	if err != nil {
		return nil, err
	}

	commandDefinition := CommandDefinition{}
	err = json.Unmarshal(body, &commandDefinition)
	if err != nil {
		return nil, err
	}

	return &commandDefinition, nil
}

func (client *Client) CreateCommandDefinition(assetId string, createCommandDefinition CommandDefinition) (*CommandDefinition, error) {
	rb, err := json.Marshal(createCommandDefinition)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/asset-repository/nodes/%s/command-definitions", client.HostURL, assetId), strings.NewReader(string(rb)))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}

	body, err, _ := client.doRequest(req, &client.Token)
	if err != nil {
		return nil, err
	}

	commandDefinition := CommandDefinition{}
	err = json.Unmarshal(body, &commandDefinition)
	if err != nil {
		return nil, err
	}

	return &commandDefinition, nil
}

func (client *Client) UpdateCommandDefinition(commandDefinitionId string, createCommandDefinition CommandDefinition) (*CommandDefinition, error) {
	rb, err := json.Marshal(createCommandDefinition)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/asset-repository/command-definitions/%s", client.HostURL, commandDefinitionId), strings.NewReader(string(rb)))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}

	body, err, _ := client.doRequest(req, &client.Token)
	if err != nil {
		return nil, err
	}

	commandDefinition := CommandDefinition{}
	err = json.Unmarshal(body, &commandDefinition)
	if err != nil {
		return nil, err
	}

	return &commandDefinition, nil
}

func (client *Client) DeleteCommandDefinition(commandDefinitionId string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/asset-repository/command-definitions/%s", client.HostURL, commandDefinitionId), nil)
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
