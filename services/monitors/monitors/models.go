package monitors

import (
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
	"github.com/leanspace/terraform-provider-leanspace/services/monitors/action_templates"
)

type Monitor struct {
	ID                string                            `json:"id"`
	Name              string                            `json:"name"`
	Description       string                            `json:"description"`
	Status            string                            `json:"status"`
	MetricId          string                            `json:"metricId"`
	NodeId            string                            `json:"nodeId"`
	Rule              Rule                              `json:"rule"`
	ActionTemplates   []action_templates.ActionTemplate `json:"actionTemplates"`
	ActionTemplateIds []string                          `json:"actionTemplateIds"`
	Tags              []general_objects.KeyValue        `json:"tags"`
	CreatedAt         string                            `json:"createdAt"`
	CreatedBy         string                            `json:"createdBy"`
	LastModifiedAt    string                            `json:"lastModifiedAt"`
	LastModifiedBy    string                            `json:"lastModifiedBy"`
}

func (monitor *Monitor) GetID() string { return monitor.ID }

type Rule struct {
	ComparisonOperator string  `json:"comparisonOperator"`
	ComparisonValue    float64 `json:"comparisonValue"`
	Tolerance          float64 `json:"tolerance,omitempty"`
}
