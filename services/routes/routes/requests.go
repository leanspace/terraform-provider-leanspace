package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/provider"
)

type apiValidProcessors struct {
	ProcessorIds []string `json:"processorIds"`
}

func (route *Route) attachDetach(action string, processors []string, client *provider.Client) error {
	processorData, err := json.Marshal(apiValidProcessors{ProcessorIds: processors})
	if err != nil {
		return err
	}

	path := fmt.Sprintf("%s/%s/%s/processors", client.HostURL, RouteDataType.Path, route.ID)
	req, err := http.NewRequest(action, path, strings.NewReader(string(processorData)))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	_, err, _ = client.DoRequest(req, &(client).Token)
	if err != nil {
		return err
	}
	return nil
}

func (route *Route) attachProcessors(processors []string, client *provider.Client) error {
	return route.attachDetach("POST", processors, client)
}

func (route *Route) detachProcessors(processors []string, client *provider.Client) error {
	return route.attachDetach("PUT", processors, client)
}

func (route *Route) PostCreateProcess(client *provider.Client, created any) error {
	createdRoute := created.(*Route)
	if len(createdRoute.ProcessorIds) > 0 {
		err := route.attachProcessors(createdRoute.ProcessorIds, client)
		if err != nil {
			return err
		}
	}
	return nil
}

func (route *Route) PostUpdateProcess(client *provider.Client, updated any) error {
	updatedRoute := updated.(*Route)
	currentProcessors := updatedRoute.ProcessorIds

	expectedProcessors := route.ProcessorIds

	processorsToRemove := []string{}
	processorsToAdd := []string{}

	// Diff processors to see what to add/remove
	for _, processor := range currentProcessors {
		if !helper.Contains(expectedProcessors, processor) {
			processorsToRemove = append(processorsToRemove, processor)
		}
	}
	for _, processor := range expectedProcessors {
		if !helper.Contains(currentProcessors, processor) {
			processorsToAdd = append(processorsToAdd, processor)
		}
	}

	// Apply diff
	if len(processorsToRemove) > 0 {
		err := route.detachProcessors(processorsToRemove, client)
		if err != nil {
			return err
		}
	}
	if len(processorsToAdd) > 0 {
		err := route.attachProcessors(processorsToAdd, client)
		if err != nil {
			return err
		}
	}

	return nil
}
