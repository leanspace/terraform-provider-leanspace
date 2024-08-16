package monitors

import (
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

type Monitor struct {
	ID                  string                     `json:"id"`
	Name                string                     `json:"name"`
	Description         string                     `json:"description"`
	Status              string                     `json:"status"`
	MetricId            string                     `json:"metricId"`
	NodeId              string                     `json:"nodeId"`
	Rule                Rule                       `json:"rule"`
	ActionTemplates     []ActionTemplate           `json:"actionTemplates"`
	ActionTemplateLinks []ActionTemplateLink       `json:"actionTemplateLinks"`
	Tags                []general_objects.KeyValue `json:"tags"`
	CreatedAt           string                     `json:"createdAt"`
	CreatedBy           string                     `json:"createdBy"`
	LastModifiedAt      string                     `json:"lastModifiedAt"`
	LastModifiedBy      string                     `json:"lastModifiedBy"`
}

func (monitor *Monitor) GetID() string { return monitor.ID }

type Rule struct {
	ComparisonOperator string  `json:"comparisonOperator"`
	ComparisonValue    float64 `json:"comparisonValue"`
	Tolerance          float64 `json:"tolerance,omitempty"`
}

type ActionTemplateLink struct {
	ID          string   `json:"id"`
	TriggeredOn []string `json:"triggeredOn"`
}

type ActionTemplate struct {
	ID             string            `json:"id"`
	Name           string            `json:"name"`
	Type           string            `json:"type"`
	URL            string            `json:"url"`
	Payload        string            `json:"payload"`
	Headers        map[string]string `json:"headers"`
	TriggeredOn    []string          `json:"triggeredOn"`
	CreatedAt      string            `json:"createdAt"`
	CreatedBy      string            `json:"createdBy"`
	LastModifiedAt string            `json:"lastModifiedAt"`
	LastModifiedBy string            `json:"lastModifiedBy"`
}
