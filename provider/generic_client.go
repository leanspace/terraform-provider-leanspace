package provider

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

func (client GenericClient[T, PT]) marshalElement(element PT) ([]byte, error) {
	if extraElement, ok := any(element).(ExtraMarshallModel); ok {
		if err := extraElement.PreMarshallProcess(); err != nil {
			return nil, err
		}
	}
	return json.Marshal(element)
}

func (client GenericClient[T, PT]) unmarshalElement(data []byte) (PT, error) {
	var element PT = new(T)
	err := json.Unmarshal(data, element)
	if err != nil {
		return nil, err
	}
	if extraElement, ok := any(element).(ExtraUnmarshallModel); ok {
		if err := extraElement.PostUnmarshallProcess(); err != nil {
			return nil, err
		}
	}
	return element, nil
}

func (client GenericClient[T, PT]) encodeElement(element PT, isUpdating bool) (io.Reader, string, error) {
	data, err := client.marshalElement(element)
	if err != nil {
		return nil, "", err
	}
	if customEncoding, ok := any(element).(CustomEncodingModel); ok {
		return customEncoding.CustomEncoding(data, isUpdating)
	} else {
		return strings.NewReader(string(data)), "application/json", nil
	}
}

func (client GenericClient[T, PT]) encodeQueryParams(filters map[string]any) string {
	queryParams := url.Values{}

	addValue := func(key string, value any) {
		if value == "" {
			return
		}
		if str, isString := value.(string); isString {
			queryParams.Add(key, str)
		} else {
			queryParams.Add(key, fmt.Sprint(value))
		}
	}

	for key, value := range filters {
		parsedKey := helper.SnakeToCamelCase(key)
		if list, isList := value.([]any); isList {
			for _, subValue := range list {
				addValue(parsedKey, subValue)
			}
		} else {
			addValue(parsedKey, value)
		}
	}
	return queryParams.Encode()
}

func (client GenericClient[T, PT]) GetAll(filters map[string]any) (*general_objects.PaginatedList[T, PT], error) {
	path := fmt.Sprintf("%s/%s", client.Client.HostURL, client.Path)
	if filters != nil {
		path += "?" + client.encodeQueryParams(filters)
	}
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	body, err, _ := client.Client.DoRequest(req, &(client.Client).Token)
	if err != nil {
		return nil, err
	}

	values := general_objects.PaginatedList[T, PT]{}
	err = json.Unmarshal(body, &values)
	if err != nil {
		return nil, err
	}

	return &values, nil
}

func (client GenericClient[T, PT]) GetUnique() (PT, error) {
	path := fmt.Sprintf("%s/%s", client.Client.HostURL, client.Path)
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	body, err, _ := client.Client.DoRequest(req, &(client.Client).Token)
	if err != nil {
		return nil, err
	}

	value, err := client.unmarshalElement(body)
	if err != nil {
		return nil, err
	}

	return value, nil
}

// Retrieve a resource of this type with the given ID.
// Can return:
// - PT, nil, if the resource was fetched
// - nil, error, if there was an error when fetching the resource
// - nil, nil, if the resource was not found (and no other error occurred)
func (client GenericClient[T, PT]) Get(id string, readElement PT) (PT, error) {
	path := fmt.Sprintf("%s/%s", client.Path, id)
	if client.ReadPath != nil {
		path = client.ReadPath(id)
	}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", client.Client.HostURL, path), nil)
	if err != nil {
		return nil, err
	}

	body, err, statusCode := client.Client.DoRequest(req, &(client.Client).Token)
	if statusCode == http.StatusNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	value, err := client.unmarshalElement(body)
	if err != nil {
		return nil, err
	}

	if postRead, ok := any(readElement).(PostReadModel); ok {
		if err := postRead.PostReadProcess(client.Client, value); err != nil {
			return nil, err
		}
	}

	return value, nil
}

func (client GenericClient[T, PT]) Create(createElement PT) (PT, error) {
	path := client.Path
	if client.CreatePath != nil {
		path = client.CreatePath(createElement)
	}

	requestContent, contentType, err := client.encodeElement(createElement, false)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s", client.Client.HostURL, path), requestContent)
	req.Header.Set("Content-Type", contentType)
	if err != nil {
		return nil, err
	}

	body, err, _ := client.Client.DoRequest(req, &(client.Client).Token)

	// Here, maybe check if the error is because of a 409 (conflict), in which case
	// we update the object and continue as if it was created.
	if err != nil {
		return nil, err
	}

	value, err := client.unmarshalElement(body)
	if err != nil {
		return nil, err
	}

	if postCreate, ok := any(createElement).(PostCreateModel); ok {
		if err := postCreate.PostCreateProcess(client.Client, value); err != nil {
			return nil, err
		}
	}

	return value, nil
}

func (client GenericClient[T, PT]) Update(elementId string, updateElement PT) (PT, error) {
	requestContent, contentType, err := client.encodeElement(updateElement, true)
	path := fmt.Sprintf("%s/%s", client.Path, elementId)
	if client.UpdatePath != nil {
		path = client.UpdatePath(elementId)
	}
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/%s", client.Client.HostURL, path), requestContent)
	req.Header.Set("Content-Type", contentType)
	if err != nil {
		return nil, err
	}

	body, err, _ := client.Client.DoRequest(req, &(client.Client).Token)
	if err != nil {
		return nil, err
	}

	value, err := client.unmarshalElement(body)
	if err != nil {
		return nil, err
	}

	if postUpdate, ok := any(updateElement).(PostUpdateModel); ok {
		if err := postUpdate.PostUpdateProcess(client.Client, value); err != nil {
			return nil, err
		}
	}

	return value, nil
}

func (client GenericClient[T, PT]) Delete(elementId string, element PT) error {
	path := fmt.Sprintf("%s/%s", client.Path, elementId)
	if client.DeletePath != nil {
		path = client.DeletePath(elementId)
	}
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/%s", client.Client.HostURL, path), nil)
	if err != nil {
		return err
	}

	_, err, statusCode := client.Client.DoRequest(req, &(client.Client).Token)
	// If it has been deleted outside terraform, it should not fail here
	if statusCode != http.StatusNotFound && err != nil {
		return err
	}

	if postDelete, ok := any(element).(PostDeleteModel); element != nil && ok {
		if err := postDelete.PostDeleteProcess(client.Client); err != nil {
			return err
		}
	}

	return nil
}
