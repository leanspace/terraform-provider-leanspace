package asset

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (client *Client) GetAllAssets() (*PaginatedList[Asset], error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/asset-repository/nodes", client.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err, _ := client.doRequest(req, &client.Token)
	if err != nil {
		return nil, err
	}

	assets := PaginatedList[Asset]{}
	err = json.Unmarshal(body, &assets)
	if err != nil {
		return nil, err
	}

	return &assets, nil
}

func (client *Client) GetAsset(assetId string) (*Asset, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/asset-repository/nodes/%s", client.HostURL, assetId), nil)
	if err != nil {
		return nil, err
	}

	body, err, _ := client.doRequest(req, &client.Token)
	if err != nil {
		return nil, err
	}

	asset := Asset{}
	err = json.Unmarshal(body, &asset)
	if err != nil {
		return nil, err
	}

	return &asset, nil
}

func (client *Client) CreateOrder(createAsset Asset) (*Asset, error) {
	rb, err := json.Marshal(createAsset)
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

	asset := Asset{}
	err = json.Unmarshal(body, &asset)
	if err != nil {
		return nil, err
	}

	return &asset, nil
}

func (client *Client) UpdateAsset(assetId string, createAsset Asset) (*Asset, error) {
	rb, err := json.Marshal(createAsset)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/asset-repository/nodes/%s", client.HostURL, assetId), strings.NewReader(string(rb)))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}

	body, err, _ := client.doRequest(req, &client.Token)
	if err != nil {
		return nil, err
	}

	asset := Asset{}
	err = json.Unmarshal(body, &asset)
	if err != nil {
		return nil, err
	}

	return &asset, nil
}

func (client *Client) DeleteAsset(assetId string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/asset-repository/nodes/%s", client.HostURL, assetId), nil)
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
