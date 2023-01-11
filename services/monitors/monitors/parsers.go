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
	monitorMap["polling_frequency_in_minutes"] = monitor.PollingFrequencyInMinutes
	monitorMap["metric_id"] = monitor.MetricId
	monitorMap["node_id"] = monitor.NodeId
	monitorMap["statistics"] = []any{monitor.Statistics.ToMap()}
	monitorMap["expression"] = []any{monitor.Expression.ToMap()}
	monitorMap["action_templates"] = helper.ParseToMaps(monitor.ActionTemplates)
	monitorMap["action_template_ids"] = monitor.ActionTemplateIds
	monitorMap["tags"] = helper.ParseToMaps(monitor.Tags)
	monitorMap["created_at"] = monitor.CreatedAt
	monitorMap["created_by"] = monitor.CreatedBy
	monitorMap["last_modified_at"] = monitor.LastModifiedAt
	monitorMap["last_modified_by"] = monitor.LastModifiedBy
	return monitorMap
}

func (statistics *Statistics) ToMap() map[string]any {
	statisticsMap := make(map[string]any)
	statisticsMap["last_evaluation"] = []any{statistics.LastEvaluation.ToMap()}
	return statisticsMap
}

func (evaluation *Evaluation) ToMap() map[string]any {
	evaluationMap := make(map[string]any)
	evaluationMap["timestamp"] = evaluation.Timestamp
	evaluationMap["value"] = evaluation.Value
	evaluationMap["status"] = evaluation.Status
	return evaluationMap
}

func (expression *Expression) ToMap() map[string]any {
	expressionMap := make(map[string]any)
	expressionMap["comparison_operator"] = expression.ComparisonOperator
	expressionMap["comparison_value"] = expression.ComparisonValue
	expressionMap["aggregation_function"] = expression.AggregationFunction
	expressionMap["tolerance"] = expression.Tolerance
	return expressionMap
}

func (monitor *Monitor) FromMap(monitorMap map[string]any) error {
	monitor.ID = monitorMap["id"].(string)
	monitor.Name = monitorMap["name"].(string)
	monitor.Description = monitorMap["description"].(string)
	monitor.Status = monitorMap["status"].(string)
	monitor.PollingFrequencyInMinutes = monitorMap["polling_frequency_in_minutes"].(int)
	monitor.MetricId = monitorMap["metric_id"].(string)
	monitor.NodeId = monitorMap["node_id"].(string)
	if len(monitorMap["statistics"].([]any)) > 0 {
		if err := monitor.Statistics.FromMap(monitorMap["statistics"].([]any)[0].(map[string]any)); err != nil {
			return err
		}
	}
	if len(monitorMap["expression"].([]any)) > 0 {
		if err := monitor.Expression.FromMap(monitorMap["expression"].([]any)[0].(map[string]any)); err != nil {
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
	if tags, err := helper.ParseFromMaps[general_objects.Tag](monitorMap["tags"].(*schema.Set).List()); err != nil {
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

func (statistics *Statistics) FromMap(statisticsMap map[string]any) error {
	if len(statisticsMap["last_evaluation"].([]any)) > 0 {
		if err := statistics.LastEvaluation.FromMap(statisticsMap["last_evaluation"].([]any)[0].(map[string]any)); err != nil {
			return err
		}
	}
	return nil
}

func (evaluation *Evaluation) FromMap(evaluationMap map[string]any) error {
	evaluation.Timestamp = evaluationMap["timestamp"].(string)
	evaluation.Value = evaluationMap["value"].(float64)
	evaluation.Status = evaluationMap["status"].(string)
	return nil
}

func (expression *Expression) FromMap(expressionMap map[string]any) error {
	expression.ComparisonOperator = expressionMap["comparison_operator"].(string)
	expression.ComparisonValue = expressionMap["comparison_value"].(float64)
	expression.AggregationFunction = expressionMap["aggregation_function"].(string)
	expression.Tolerance = expressionMap["tolerance"].(float64)
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
