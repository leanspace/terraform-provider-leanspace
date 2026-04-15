package monitors

import (
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

type Monitor struct {
	general_objects.AuditModel
	Name                string                     `json:"name"`
	Description         *string                    `json:"description"`
	Status              string                     `json:"status"`
	MetricId            string                     `json:"metricId"`
	NodeId              string                     `json:"nodeId"`
	Rule                Rule                       `json:"rule"`
	ActionTemplates     []ActionTemplate           `json:"actionTemplates"`
	ActionTemplateLinks []ActionTemplateLink       `json:"actionTemplateLinks"`
	Tags                []general_objects.KeyValue `json:"tags"`
}

type Rule struct {
	ComparisonOperator string   `json:"comparisonOperator"`
	ComparisonValue    float64  `json:"comparisonValue"`
	Tolerance          *float64 `json:"tolerance,omitempty"`
}

type ActionTemplateLink struct {
	ID          string   `json:"id"`
	TriggeredOn []string `json:"triggeredOn"`
}

type ActionTemplate struct {
	general_objects.AuditModel
	Name        string            `json:"name"`
	Type        string            `json:"type"`
	URL         string            `json:"url"`
	Payload     string            `json:"payload"`
	Headers     map[string]string `json:"headers"`
	TriggeredOn []string          `json:"triggeredOn"`
}
