package asset

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (client *Client) GetAllNodes() (*PaginatedList[Node], error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/asset-repository/nodes", client.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err, _ := client.doRequest(req, &client.Token)
	if err != nil {
		return nil, err
	}

	nodes := PaginatedList[Node]{}
	err = json.Unmarshal(body, &nodes)
	if err != nil {
		return nil, err
	}

	return &nodes, nil
}

func (client *Client) GetNode(nodeId string) (*Node, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/asset-repository/nodes/%s", client.HostURL, nodeId), nil)
	if err != nil {
		return nil, err
	}

	body, err, _ := client.doRequest(req, &client.Token)
	if err != nil {
		return nil, err
	}

	node := Node{}
	err = json.Unmarshal(body, &node)
	if err != nil {
		return nil, err
	}

	return &node, nil
}

func (client *Client) CreateNode(createNode Node) (*Node, error) {
	rb, err := json.Marshal(createNode)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/asset-repository/nodes", client.HostURL), strings.NewReader(string(rb)))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}

	body, err, _ := client.doRequest(req, &client.Token)
	if err != nil {
		return nil, err
	}

	node := Node{}
	err = json.Unmarshal(body, &node)
	if err != nil {
		return nil, err
	}

	return &node, nil
}

func (client *Client) UpdateNode(nodeId string, createNode Node) (*Node, error) {
	rb, err := json.Marshal(createNode)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/asset-repository/nodes/%s", client.HostURL, nodeId), strings.NewReader(string(rb)))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}

	body, err, _ := client.doRequest(req, &client.Token)
	if err != nil {
		return nil, err
	}

	node := Node{}
	err = json.Unmarshal(body, &node)
	if err != nil {
		return nil, err
	}

	return &node, nil
}

func (client *Client) DeleteNode(nodeId string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/asset-repository/nodes/%s", client.HostURL, nodeId), nil)
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
