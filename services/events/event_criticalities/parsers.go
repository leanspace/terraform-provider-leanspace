package event_criticalities

func (eventCriticality *EventsCriticalities) ToMap() map[string]any {
	eventCriticalityMap := make(map[string]any)
	eventCriticalityMap["id"] = eventCriticality.ID
	eventCriticalityMap["name"] = eventCriticality.Name

	return eventCriticalityMap
}

func (eventCriticality *EventsCriticalities) FromMap(eventCriticalityMap map[string]any) error {
	eventCriticality.ID = eventCriticalityMap["id"].(string)
	eventCriticality.Name = eventCriticalityMap["name"].(string)

	return nil
}
