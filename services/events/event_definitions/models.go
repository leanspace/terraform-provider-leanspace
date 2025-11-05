package event_definitions

import "github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

type EventsDefinition struct {
	ID             string                     `json:"id"`
	Name           string                     `json:"name"`
	Source         string                     `json:"source"`
	State          string                     `json:"state"`
	Description    string                     `json:"description,omitempty"`
	Criticality    string                     `json:"criticality,omitempty"`
	Rules          []Rules[any]               `json:"rules,omitempty"`
	CreatedAt      string                     `json:"createdAt"`
	CreatedBy      string                     `json:"createdBy"`
	LastModifiedAt string                     `json:"lastModifiedAt"`
	LastModifiedBy string                     `json:"lastModifiedBy"`
	Tags           []general_objects.KeyValue `json:"tags,omitempty"`
}

func (eventDefinition *EventsDefinition) GetID() string { return eventDefinition.ID }

type Rules[T any] struct {
	Operator        string              `json:"operator"`
	Path            string              `json:"path"`
	ComparisonValue *ComparisonValue[T] `json:"comparisonValue"`
}

type ComparisonValue[T any] struct {
	Value T      `json:"value,omitempty"`
	Type  string `json:"type"`
}
