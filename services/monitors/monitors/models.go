package monitors

import (
	"leanspace-terraform-provider/helper/general_objects"
	"leanspace-terraform-provider/services/monitors/action_templates"
)

type Monitor struct {
	ID                        string                            `json:"id" terra:"id"`
	Name                      string                            `json:"name" terra:"name"`
	Description               string                            `json:"description" terra:"description"`
	Status                    string                            `json:"status" terra:"status"`
	PollingFrequencyInMinutes int                               `json:"pollingFrequencyInMinutes" terra:"polling_frequency_in_minutes"`
	MetricId                  string                            `json:"metricId" terra:"metric_id"`
	NodeId                    string                            `json:"nodeId" terra:"node_id"`
	Statistics                Statistics                        `json:"statistics" terra:"statistics"`
	Expression                Expression                        `json:"expression" terra:"expression"`
	ActionTemplates           []action_templates.ActionTemplate `json:"actionTemplates" terra:"action_templates"`
	ActionTemplateIds         []string                          `json:"actionTemplateIds" terra:"action_template_ids"`
	Tags                      []general_objects.Tag             `json:"tags" terra:"tags"`
	CreatedAt                 string                            `json:"createdAt" terra:"created_at"`
	CreatedBy                 string                            `json:"createdBy" terra:"created_by"`
	LastModifiedAt            string                            `json:"lastModifiedAt" terra:"last_modified_at"`
	LastModifiedBy            string                            `json:"lastModifiedBy" terra:"last_modified_by"`
}

func (monitor *Monitor) GetID() string { return monitor.ID }

type Statistics struct {
	LastEvaluation Evaluation `json:"lastEvaluation" terra:"last_evaluation"`
}

type Evaluation struct {
	Timestamp string  `json:"timestamp" terra:"timestamp"`
	Value     float64 `json:"value" terra:"value"`
	Status    string  `json:"status" terra:"status"`
}

type Expression struct {
	ComparisonOperator  string  `json:"comparisonOperator" terra:"comparison_operator"`
	ComparisonValue     float64 `json:"comparisonValue" terra:"comparison_value"`
	AggregationFunction string  `json:"aggregationFunction" terra:"aggregation_function"`
	Tolerance           float64 `json:"tolerance,omitempty" terra:"tolerance,omitempty"`
}
