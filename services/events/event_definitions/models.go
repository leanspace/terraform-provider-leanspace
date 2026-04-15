package event_definitions

import "github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

type EventsDefinition struct {
	general_objects.AuditModel
	Name        string                     `json:"name"`
	Source      string                     `json:"source"`
	State       string                     `json:"state"`
	Description *string                    `json:"description,omitempty"`
	Criticality *string                    `json:"criticality,omitempty"`
	Rules       []Rules[any]               `json:"rules,omitempty"`
	Tags        []general_objects.KeyValue `json:"tags,omitempty"`
}

type Rules[T any] struct {
	Operator        string              `json:"operator"`
	Path            string              `json:"path"`
	ComparisonValue *ComparisonValue[T] `json:"comparisonValue"`
}

type ComparisonValue[T any] struct {
	Value T      `json:"value,omitempty"`
	Type  string `json:"type"`
}
