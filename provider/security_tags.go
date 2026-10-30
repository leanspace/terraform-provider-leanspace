package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type SecurityTagsModel interface {
	GetSecurityTags() []string
	SetSecurityTags([]string)
}

func WithSecurityTagsSchema(resourceSchema map[string]*schema.Schema) map[string]*schema.Schema {
	resourceSchema["security_tags"] = &schema.Schema{
		Type: schema.TypeSet, Optional: true, Elem: &schema.Schema{Type: schema.TypeString},
		Description: "Complete security-tag collection. Omit on creation to leave tags unchanged; an explicit empty set clears tags. Removing this attribute after managing tags also clears the collection. Refresh and import retrieve actual tags.",
	}
	return resourceSchema
}

func withoutSecurityTagsSchema(resourceSchema map[string]*schema.Schema) map[string]*schema.Schema {
	filteredSchema := make(map[string]*schema.Schema, len(resourceSchema))
	for key, attribute := range resourceSchema {
		if key != "security_tags" {
			filteredSchema[key] = attribute
		}
	}
	return filteredSchema
}

func (dataSource DataSourceType[T, PT]) securityTagsModel(value PT) (SecurityTagsModel, error) {
	model, supported := any(value).(SecurityTagsModel)
	if !supported {
		return nil, fmt.Errorf("resource %s has security_tags schema but its model does not implement SecurityTagsModel", dataSource.ResourceIdentifier)
	}
	return model, nil
}

func (dataSource DataSourceType[T, PT]) writeSecurityTags(ctx context.Context, client *Client, parentID string, value PT) error {
	model, err := dataSource.securityTagsModel(value)
	if err != nil {
		return err
	}
	_, err = sendSecurityTagRequest(ctx, client, dataSource.Path, parentID, http.MethodPut, model.GetSecurityTags())
	return err
}

func (dataSource DataSourceType[T, PT]) loadSecurityTagsAfterRead(ctx context.Context, client *Client, parentID string, value PT) error {
	model, err := dataSource.securityTagsModel(value)
	if err != nil {
		return err
	}
	responseBody, err := sendSecurityTagRequest(ctx, client, dataSource.Path, parentID, http.MethodGet, nil)
	if err != nil {
		return err
	}
	var tagValues []any
	if err := json.Unmarshal(responseBody, &tagValues); err != nil {
		return err
	}
	if tagValues == nil {
		return fmt.Errorf("expected a JSON string array for security tags")
	}
	securityTags := make([]string, len(tagValues))
	for index, tagValue := range tagValues {
		tag, valid := tagValue.(string)
		if !valid {
			return fmt.Errorf("expected a JSON string array for security tags")
		}
		securityTags[index] = tag
	}
	model.SetSecurityTags(securityTags)
	return nil
}

func sendSecurityTagRequest(ctx context.Context, client *Client, basePath, parentID, method string, securityTags []string) ([]byte, error) {
	var requestBody []byte
	if method == http.MethodPut {
		if securityTags == nil {
			securityTags = []string{}
		}
		var err error
		requestBody, err = json.Marshal(struct {
			SecurityTags []string `json:"securityTags"`
		}{securityTags})
		if err != nil {
			return nil, err
		}
	}
	request, err := http.NewRequestWithContext(ctx, method, fmt.Sprintf("%s/%s/%s/security-tags", client.HostURL, basePath, parentID), bytes.NewReader(requestBody))
	if err != nil {
		return nil, err
	}
	if method == http.MethodPut {
		request.Header.Set("Content-Type", "application/json")
	}
	responseBody, err, _ := client.DoRequest(request, &client.Token)
	return responseBody, err
}
