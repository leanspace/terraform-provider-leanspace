package event_criticalities

type EventCriticalities struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	CreatedAt      string `json:"createdAt"`
	CreatedBy      string `json:"createdBy"`
	LastModifiedAt string `json:"lastModifiedAt"`
	LastModifiedBy string `json:"lastModifiedBy"`
	ReadOnly       bool   `json:"readOnly"`
}

func (eventCriticalities *EventCriticalities) GetID() string { return eventCriticalities.ID }
