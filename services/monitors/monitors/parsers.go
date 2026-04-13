package monitors

import (
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

func (monitor *Monitor) ToMap() map[string]any {
	monitorMap := monitor.ToAuditMap()
	monitorMap["name"] = helper.NilIfEmpty(monitor.Name)
	monitorMap["description"] = helper.NilIfEmpty(monitor.Description)
	monitorMap["status"] = helper.NilIfEmpty(monitor.Status)
	monitorMap["metric_id"] = helper.NilIfEmpty(monitor.MetricId)
	monitorMap["node_id"] = helper.NilIfEmpty(monitor.NodeId)
	monitorMap["rule"] = monitor.Rule.ToMap()
	monitorMap["action_templates"] = helper.ParseToMaps(monitor.ActionTemplates)
	if monitor.ActionTemplateLinks != nil {
		monitorMap["action_template_links"] = helper.ParseToMaps(monitor.ActionTemplateLinks)
	}
	monitorMap["tags"] = helper.ParseToMaps(monitor.Tags)
	monitorMap["type"] = "REALTIME"
	return monitorMap
}

func (rule *Rule) ToMap() map[string]any {
	ruleMap := make(map[string]any)
	ruleMap["comparison_operator"] = helper.NilIfEmpty(rule.ComparisonOperator)
	ruleMap["comparison_value"] = helper.NilIfEmpty(rule.ComparisonValue)
	ruleMap["tolerance"] = helper.Float64PtrToAny(rule.Tolerance)
	return ruleMap
}

func (actionTemplateLink *ActionTemplateLink) ToMap() map[string]any {
	actionTemplateLinkMap := make(map[string]any)
	actionTemplateLinkMap["id"] = helper.NilIfEmpty(actionTemplateLink.ID)
	if len(actionTemplateLink.TriggeredOn) > 0 {
		actionTemplateLinkMap["triggered_on"] = actionTemplateLink.TriggeredOn
	}
	return actionTemplateLinkMap
}

func (monitor *Monitor) FromMap(monitorMap map[string]any) error {
	monitor.FromAuditMap(monitorMap)
	monitor.Name = helper.CastString(monitorMap, "name")
	monitor.Description = helper.CastString(monitorMap, "description")
	monitor.Status = helper.CastString(monitorMap, "status")
	monitor.MetricId = helper.CastString(monitorMap, "metric_id")
	monitor.NodeId = helper.CastString(monitorMap, "node_id")
	if err := monitor.Rule.FromMap(helper.CastMapAny(monitorMap, "rule")); err != nil {
		return err
	}
	if actionTemplates, err := helper.ParseFromMaps[ActionTemplate](helper.CastSlice(monitorMap, "action_templates")); err != nil {
		return err
	} else {
		monitor.ActionTemplates = actionTemplates
	}
	if actionTemplateLinks, err := helper.ParseFromMaps[ActionTemplateLink](helper.CastSlice(monitorMap, "action_template_links")); err != nil {
		return err
	} else {
		monitor.ActionTemplateLinks = actionTemplateLinks
	}
	if tags, err := helper.ParseFromMaps[general_objects.KeyValue](helper.CastSlice(monitorMap, "tags")); err != nil {
		return err
	} else {
		monitor.Tags = tags
	}
	return nil
}

func (rule *Rule) FromMap(ruleMap map[string]any) error {
	rule.ComparisonOperator = helper.CastString(ruleMap, "comparison_operator")
	rule.ComparisonValue = helper.CastFloat64(ruleMap, "comparison_value")
	rule.Tolerance = helper.CastFloat64Ptr(ruleMap, "tolerance")
	return nil
}

func (actionTemplateLink *ActionTemplateLink) FromMap(actionTemplateLinkMap map[string]any) error {
	actionTemplateLink.ID = helper.CastString(actionTemplateLinkMap, "id")
	actionTemplateLink.TriggeredOn = make([]string, len(helper.CastSlice(actionTemplateLinkMap, "triggered_on")))
	for i, value := range helper.CastSlice(actionTemplateLinkMap, "triggered_on") {
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
	actionTemplateMap := actionTemplate.ToAuditMap()
	actionTemplateMap["name"] = helper.NilIfEmpty(actionTemplate.Name)
	actionTemplateMap["type"] = helper.NilIfEmpty(actionTemplate.Type)
	actionTemplateMap["url"] = helper.NilIfEmpty(actionTemplate.URL)
	actionTemplateMap["payload"] = helper.NilIfEmpty(actionTemplate.Payload)
	actionTemplateMap["headers"] = helper.NilIfEmpty(actionTemplate.Headers)
	actionTemplateMap["triggered_on"] = helper.NilIfEmpty(actionTemplate.TriggeredOn)
	return actionTemplateMap
}

func (actionTemplate *ActionTemplate) FromMap(actionTemplateMap map[string]any) error {
	actionTemplate.FromAuditMap(actionTemplateMap)
	actionTemplate.Name = helper.CastString(actionTemplateMap, "name")
	actionTemplate.Type = helper.CastString(actionTemplateMap, "type")
	actionTemplate.URL = helper.CastString(actionTemplateMap, "url")
	actionTemplate.Payload = helper.CastString(actionTemplateMap, "payload")
	actionTemplate.Headers = make(map[string]string, len(helper.CastMapAny(actionTemplateMap, "headers")))
	for key, value := range helper.CastMapAny(actionTemplateMap, "headers") {
		actionTemplate.Headers[key] = value.(string)
	}
	actionTemplate.TriggeredOn = make([]string, len(helper.CastSlice(actionTemplateMap, "triggered_on")))
	for i, value := range helper.CastSlice(actionTemplateMap, "triggered_on") {
		actionTemplate.TriggeredOn[i] = value.(string)
	}
	return nil
}
