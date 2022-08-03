package monitors

import (
	"fmt"
	"leanspace-terraform-provider/helper"
	"leanspace-terraform-provider/provider"
	"net/http"
)

func (monitor *Monitor) actionTemplateChange(action string, actionTemplateId string, client *provider.Client) error {
	path := fmt.Sprintf("%s/%s/%s/action-templates/%s", client.HostURL, MonitorDataType.Path, monitor.ID, actionTemplateId)
	req, err := http.NewRequest(action, path, nil)
	if err != nil {
		return err
	}
	_, err, _ = client.DoRequest(req, &(client).Token)
	if err != nil {
		return err
	}
	return nil
}

func (monitor *Monitor) addActionTemplate(actionTemplateId string, client *provider.Client) error {
	return monitor.actionTemplateChange("POST", actionTemplateId, client)
}

func (monitor *Monitor) removeActionTemplate(actionTemplateId string, client *provider.Client) error {
	return monitor.actionTemplateChange("DELETE", actionTemplateId, client)
}

func (monitor *Monitor) PostUpdateProcess(client *provider.Client, monitorRaw any) error {
	monitorCurrent := monitorRaw.(*Monitor)
	currentActionTemplates := monitorCurrent.ActionTemplateIds
	expectedActionTemplates := monitor.ActionTemplateIds

	actionTemplatesToRemove := []string{}
	actionTemplatesToAdd := []string{}

	// Diff action templates to see what to add/remove
	for _, actionTemplate := range currentActionTemplates {
		if !helper.Contains(expectedActionTemplates, actionTemplate) {
			actionTemplatesToRemove = append(actionTemplatesToRemove, actionTemplate)
		}
	}
	for _, actionTemplate := range expectedActionTemplates {
		if !helper.Contains(currentActionTemplates, actionTemplate) {
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
