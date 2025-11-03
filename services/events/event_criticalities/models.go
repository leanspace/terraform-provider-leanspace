package event_criticalities

import "github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

type EventCriticalities struct {
	ID             string                     `json:"id"`
	Name           string                     `json:"name"`
	CreatedAt      string                     `json:"createdAt"`
	CreatedBy      string                     `json:"createdBy"`
	LastModifiedAt string                     `json:"lastModifiedAt"`
	LastModifiedBy string                     `json:"lastModifiedBy"`
	ReadOnly       bool                       `json:"readOnly"`
	Tags           []general_objects.KeyValue `json:"tags,omitempty"`
}

func (eventCriticalities *EventCriticalities) GetID() string { return eventCriticalities.ID }
