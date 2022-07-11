package asset

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"terraform-provider-asset/asset/general_objects"
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

func (client GenericClient[T, PT]) GetAll() (*general_objects.PaginatedList[T], error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", client.Client.HostURL, client.Path), nil)
	if err != nil {
		return nil, err
	}

	body, err, _ := client.Client.DoRequest(req, &(client.Client).Token)
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

// Retrieve a resource of this type with the given ID.
// Can return:
// - PT, nil, if the resource was fetched
// - nil, error, if there was an error when fetching the resource
// - nil, nil, if the resource was not found (and no other error occurred)
func (client GenericClient[T, PT]) Get(id string) (PT, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/%s", client.Client.HostURL, client.Path, id), nil)
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

	return value, nil
}

func (client GenericClient[T, PT]) Create(createElement PT) (PT, error) {
	rb, err := client.marshalElement(createElement)
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
	rb, err := client.marshalElement(updateElement)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/%s/%s", client.Client.HostURL, client.Path, elementId), strings.NewReader(string(rb)))
	req.Header.Set("Content-Type", "application/json")
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
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/%s/%s", client.Client.HostURL, client.Path, elementId), nil)
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
