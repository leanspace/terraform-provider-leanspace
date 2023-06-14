package monitors

import (
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
	"github.com/leanspace/terraform-provider-leanspace/services/monitors/action_templates"

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
	monitorMap["action_template_ids"] = monitor.ActionTemplateIds
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
	if actionTemplates, err := helper.ParseFromMaps[action_templates.ActionTemplate](monitorMap["action_templates"].(*schema.Set).List()); err != nil {
		return err
	} else {
		monitor.ActionTemplates = actionTemplates
	}
	monitor.ActionTemplateIds = make([]string, monitorMap["action_template_ids"].(*schema.Set).Len())
	for i, value := range monitorMap["action_template_ids"].(*schema.Set).List() {
		monitor.ActionTemplateIds[i] = value.(string)
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

// The API doesn't return "action template IDs", it's a field derived from
// "actionTemplates" to make it easier to configure.
// We thus need to persist the IDs in that separate array.
func (monitor *Monitor) PostUnmarshallProcess() error {
	monitor.ActionTemplateIds = make([]string, len(monitor.ActionTemplates))
	for i, value := range monitor.ActionTemplates {
		monitor.ActionTemplateIds[i] = value.ID
	}
	return nil
}
