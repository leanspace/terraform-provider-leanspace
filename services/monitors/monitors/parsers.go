package monitors

import (
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func (monitor *Monitor) ToMap() map[string]any {
	monitorMap := make(map[string]any)
	monitorMap["id"] = monitor.ID
	monitorMap["name"] = monitor.Name
	monitorMap["description"] = monitor.Description
	monitorMap["status"] = monitor.Status
	monitorMap["metric_id"] = monitor.MetricId
	monitorMap["node_id"] = monitor.NodeId
	monitorMap["rule"] = []any{monitor.Rule.ToMap()}
	monitorMap["action_templates"] = helper.ParseToMaps(monitor.ActionTemplates)
	if monitor.ActionTemplateLinks != nil {
		monitorMap["action_template_links"] = helper.ParseToMaps(monitor.ActionTemplateLinks)
	}
	monitorMap["tags"] = helper.ParseToMaps(monitor.Tags)
	monitorMap["created_at"] = monitor.CreatedAt
	monitorMap["created_by"] = monitor.CreatedBy
	monitorMap["last_modified_at"] = monitor.LastModifiedAt
	monitorMap["last_modified_by"] = monitor.LastModifiedBy
	monitorMap["type"] = "REALTIME"
	return monitorMap
}

func (rule *Rule) ToMap() map[string]any {
	ruleMap := make(map[string]any)
	ruleMap["comparison_operator"] = rule.ComparisonOperator
	ruleMap["comparison_value"] = rule.ComparisonValue
	ruleMap["tolerance"] = rule.Tolerance
	return ruleMap
}

func (actionTemplateLink *ActionTemplateLink) ToMap() map[string]any {
	actionTemplateLinkMap := make(map[string]any)
	actionTemplateLinkMap["id"] = actionTemplateLink.ID
	actionTemplateLinkMap["triggered_on"] = actionTemplateLink.TriggeredOn
	return actionTemplateLinkMap
}

func (monitor *Monitor) FromMap(monitorMap map[string]any) error {
	monitor.ID = monitorMap["id"].(string)
	monitor.Name = monitorMap["name"].(string)
	monitor.Description = monitorMap["description"].(string)
	monitor.Status = monitorMap["status"].(string)
	monitor.MetricId = monitorMap["metric_id"].(string)
	monitor.NodeId = monitorMap["node_id"].(string)
	if len(monitorMap["rule"].([]any)) > 0 {
		if err := monitor.Rule.FromMap(monitorMap["rule"].([]any)[0].(map[string]any)); err != nil {
			return err
		}
	}
	if actionTemplates, err := helper.ParseFromMaps[ActionTemplate](monitorMap["action_templates"].(*schema.Set).List()); err != nil {
		return err
	} else {
		monitor.ActionTemplates = actionTemplates
	}
	if actionTemplateLinks, err := helper.ParseFromMaps[ActionTemplateLink](monitorMap["action_template_links"].(*schema.Set).List()); err != nil {
		return err
	} else {
		monitor.ActionTemplateLinks = actionTemplateLinks
	}
	if tags, err := helper.ParseFromMaps[general_objects.KeyValue](monitorMap["tags"].(*schema.Set).List()); err != nil {
		return err
	} else {
		monitor.Tags = tags
	}
	monitor.CreatedAt = monitorMap["created_at"].(string)
	monitor.CreatedBy = monitorMap["created_by"].(string)
	monitor.LastModifiedAt = monitorMap["last_modified_at"].(string)
	monitor.LastModifiedBy = monitorMap["last_modified_by"].(string)
	return nil
}

func (rule *Rule) FromMap(ruleMap map[string]any) error {
	rule.ComparisonOperator = ruleMap["comparison_operator"].(string)
	rule.ComparisonValue = ruleMap["comparison_value"].(float64)
	rule.Tolerance = ruleMap["tolerance"].(float64)
	return nil
}

func (actionTemplateLink *ActionTemplateLink) FromMap(actionTemplateLinkMap map[string]any) error {
	actionTemplateLink.ID = actionTemplateLinkMap["id"].(string)
	actionTemplateLink.TriggeredOn = make([]string, actionTemplateLinkMap["triggered_on"].(*schema.Set).Len())
	for i, value := range actionTemplateLinkMap["triggered_on"].(*schema.Set).List() {
		actionTemplateLink.TriggeredOn[i] = value.(string)
	}
	return nil
}

// The API doesn't return "action template IDs", it's a field derived from
// "actionTemplates" to make it easier to configure.
// We thus need to persist the IDs in that separate array.
func (monitor *Monitor) PostUnmarshallProcess() error {
	monitor.ActionTemplateLinks = make([]ActionTemplateLink, len(monitor.ActionTemplates))
	for i, value := range monitor.ActionTemplates {
		monitor.ActionTemplateLinks[i].ID = value.ID
		monitor.ActionTemplateLinks[i].TriggeredOn = value.TriggeredOn
	}
	return nil
}

func (actionTemplate *ActionTemplate) ToMap() map[string]any {
	actionTemplateMap := make(map[string]any)
	actionTemplateMap["id"] = actionTemplate.ID
	actionTemplateMap["name"] = actionTemplate.Name
	actionTemplateMap["type"] = actionTemplate.Type
	actionTemplateMap["url"] = actionTemplate.URL
	actionTemplateMap["payload"] = actionTemplate.Payload
	actionTemplateMap["headers"] = actionTemplate.Headers
	actionTemplateMap["triggered_on"] = actionTemplate.TriggeredOn
	actionTemplateMap["created_at"] = actionTemplate.CreatedAt
	actionTemplateMap["created_by"] = actionTemplate.CreatedBy
	actionTemplateMap["last_modified_at"] = actionTemplate.LastModifiedAt
	actionTemplateMap["last_modified_by"] = actionTemplate.LastModifiedBy
	return actionTemplateMap
}

func (actionTemplate *ActionTemplate) FromMap(actionTemplateMap map[string]any) error {
	actionTemplate.ID = actionTemplateMap["id"].(string)
	actionTemplate.Name = actionTemplateMap["name"].(string)
	actionTemplate.Type = actionTemplateMap["type"].(string)
	actionTemplate.URL = actionTemplateMap["url"].(string)
	actionTemplate.Payload = actionTemplateMap["payload"].(string)
	actionTemplate.Headers = make(map[string]string, len(actionTemplateMap["headers"].(map[string]any)))
	for key, value := range actionTemplateMap["headers"].(map[string]any) {
		actionTemplate.Headers[key] = value.(string)
	}
	actionTemplate.TriggeredOn = make([]string, actionTemplateMap["triggered_on"].(*schema.Set).Len())
	for i, value := range actionTemplateMap["triggered_on"].(*schema.Set).List() {
		actionTemplate.TriggeredOn[i] = value.(string)
	}
	actionTemplate.CreatedAt = actionTemplateMap["created_at"].(string)
	actionTemplate.CreatedBy = actionTemplateMap["created_by"].(string)
	actionTemplate.LastModifiedAt = actionTemplateMap["last_modified_at"].(string)
	actionTemplate.LastModifiedBy = actionTemplateMap["last_modified_by"].(string)
	return nil
}
