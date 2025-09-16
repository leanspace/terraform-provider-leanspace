package event_criticalities

type EventsCriticalities struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	CreatedAt      string `json:"createdAt"`
	CreatedBy      string `json:"createdBy"`
	LastModifiedAt string `json:"lastModifiedAt"`
	LastModifiedBy string `json:"lastModifiedBy"`
}

func (eventCriticalities *EventsCriticalities) GetID() string { return eventCriticalities.ID }
