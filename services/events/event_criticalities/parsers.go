package event_criticalities

func (eventCriticality *EventCriticalities) ToMap() map[string]any {
	eventCriticalityMap := make(map[string]any)
	eventCriticalityMap["id"] = eventCriticality.ID
	eventCriticalityMap["name"] = eventCriticality.Name
	eventCriticalityMap["created_at"] = eventCriticality.CreatedAt
	eventCriticalityMap["created_by"] = eventCriticality.CreatedBy
	eventCriticalityMap["last_modified_at"] = eventCriticality.LastModifiedAt
	eventCriticalityMap["last_modified_by"] = eventCriticality.LastModifiedBy
	eventCriticalityMap["read_only"] = eventCriticality.ReadOnly
	return eventCriticalityMap
}

func (eventCriticality *EventCriticalities) FromMap(eventCriticalityMap map[string]any) error {
	eventCriticality.ID = eventCriticalityMap["id"].(string)
	eventCriticality.Name = eventCriticalityMap["name"].(string)
	eventCriticality.CreatedAt = eventCriticalityMap["created_at"].(string)
	eventCriticality.CreatedBy = eventCriticalityMap["created_by"].(string)
	eventCriticality.LastModifiedAt = eventCriticalityMap["last_modified_at"].(string)
	eventCriticality.LastModifiedBy = eventCriticalityMap["last_modified_by"].(string)
	eventCriticality.ReadOnly = eventCriticalityMap["read_only"].(bool)
	return nil
}
