package monitors

import (
	"leanspace-terraform-provider/helper/general_objects"
	"leanspace-terraform-provider/services/monitors/action_templates"
)

type Monitor struct {
	ID                        string                            `json:"id"`
	Name                      string                            `json:"name"`
	Description               string                            `json:"description"`
	Status                    string                            `json:"status"`
	PollingFrequencyInMinutes int                               `json:"pollingFrequencyInMinutes"`
	MetricId                  string                            `json:"metricId"`
	NodeId                    string                            `json:"nodeId"`
	Statistics                Statistics                        `json:"statistics"`
	Expression                Expression                        `json:"expression"`
	ActionTemplates           []action_templates.ActionTemplate `json:"actionTemplates"`
	ActionTemplateIds         []string                          `json:"actionTemplateIds"`
	Tags                      []general_objects.Tag             `json:"tags"`
	CreatedAt                 string                            `json:"createdAt"`
	CreatedBy                 string                            `json:"createdBy"`
	LastModifiedAt            string                            `json:"lastModifiedAt"`
	LastModifiedBy            string                            `json:"lastModifiedBy"`
}

func (monitor *Monitor) GetID() string { return monitor.ID }

type Statistics struct {
	LastEvaluation Evaluation `json:"lastEvaluation"`
}

type Evaluation struct {
	Timestamp string  `json:"timestamp"`
	Value     float64 `json:"value"`
	Status    string  `json:"status"`
}

type Expression struct {
	ComparisonOperator  string  `json:"comparisonOperator"`
	ComparisonValue     float64 `json:"comparisonValue"`
	AggregationFunction string  `json:"aggregationFunction"`
	Tolerance           float64 `json:"tolerance,omitempty"`
}
