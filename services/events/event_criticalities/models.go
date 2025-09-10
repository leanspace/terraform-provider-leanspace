package event_criticalities

type EventsCriticalities struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (eventCriticalities *EventsCriticalities) GetID() string { return eventCriticalities.ID }
