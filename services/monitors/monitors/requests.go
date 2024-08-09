package monitors

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/leanspace/terraform-provider-leanspace/provider"
)

type apiLinkActionTemplate struct {
	TriggeredOn []string `json:"triggeredOn"`
}

func (actionTemplateLink *ActionTemplateLink) toAPIFormat() ([]byte, error) {
	linkActionTemplate := apiLinkActionTemplate{
		TriggeredOn: actionTemplateLink.TriggeredOn,
	}
	return json.Marshal(linkActionTemplate)
}

func (monitor *Monitor) actionTemplateChange(action string, actionTemplateLink ActionTemplateLink, client *provider.Client) error {
	data, err := actionTemplateLink.toAPIFormat() // ignored when deleting
	if err != nil {
		return err
	}
	path := fmt.Sprintf("%s/%s/%s/action-templates/%s", client.HostURL, MonitorDataType.Path, monitor.ID, actionTemplateLink.ID)
	req, err := http.NewRequest(action, path, strings.NewReader(string(data)))
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

func (monitor *Monitor) addActionTemplate(actionTemplateLink ActionTemplateLink, client *provider.Client) error {
	return monitor.actionTemplateChange("POST", actionTemplateLink, client)
}

func (monitor *Monitor) removeActionTemplate(actionTemplateLink ActionTemplateLink, client *provider.Client) error {
	return monitor.actionTemplateChange("DELETE", actionTemplateLink, client)
}

func (monitor *Monitor) PostCreateProcess(client *provider.Client, monitorRaw any) error {
	createdMonitor := monitorRaw.(*Monitor)
	expectedActionTemplates := monitor.ActionTemplateLinks

	// Add all members directly
	for _, actionTemplate := range expectedActionTemplates {
		err := createdMonitor.addActionTemplate(actionTemplate, client)
		if err != nil {
			return err
		}
	}

	return nil
}

func (monitor *Monitor) PostUpdateProcess(client *provider.Client, monitorRaw any) error {
	monitorCurrent := monitorRaw.(*Monitor)
	currentActionTemplates := monitorCurrent.ActionTemplateLinks
	expectedActionTemplates := monitor.ActionTemplateLinks

	actionTemplatesToRemove := []ActionTemplateLink{}
	actionTemplatesToAdd := []ActionTemplateLink{}

	// Diff action templates to see what to add/remove
	for _, actionTemplate := range currentActionTemplates {
		if !contains(expectedActionTemplates, actionTemplate) {
			actionTemplatesToRemove = append(actionTemplatesToRemove, actionTemplate)
		}
	}
	for _, actionTemplate := range expectedActionTemplates {
		if !contains(currentActionTemplates, actionTemplate) {
			actionTemplatesToAdd = append(actionTemplatesToAdd, actionTemplate)
		}
	}

	// Apply diff
	if len(actionTemplatesToRemove) > 0 {
		for _, actionTemplate := range actionTemplatesToRemove {
			err := monitor.removeActionTemplate(actionTemplate, client)
			if err != nil {
				return err
			}
		}
	}
	if len(actionTemplatesToAdd) > 0 {
		for _, actionTemplate := range actionTemplatesToAdd {
			err := monitor.addActionTemplate(actionTemplate, client)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func contains(slice []ActionTemplateLink, value ActionTemplateLink) bool {
	for _, v := range slice {
		if v.ID == value.ID && stringSliceAreEqual(v.TriggeredOn, value.TriggeredOn) {
			return true
		}
	}
	return false
}

func stringSliceAreEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
