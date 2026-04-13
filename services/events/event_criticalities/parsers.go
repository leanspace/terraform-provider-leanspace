package event_criticalities

import (
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

func (eventCriticality *EventCriticalities) ToMap() map[string]any {
	eventCriticalityMap := eventCriticality.ToAuditMap()
	eventCriticalityMap["name"] = helper.NilIfEmpty(eventCriticality.Name)
	eventCriticalityMap["read_only"] = helper.NilIfEmpty(eventCriticality.ReadOnly)
	eventCriticalityMap["tags"] = helper.ParseToMaps(eventCriticality.Tags)
	return eventCriticalityMap
}

func (eventCriticality *EventCriticalities) FromMap(eventCriticalityMap map[string]any) error {
	eventCriticality.FromAuditMap(eventCriticalityMap)
	eventCriticality.Name = helper.CastString(eventCriticalityMap, "name")
	eventCriticality.ReadOnly = helper.CastBool(eventCriticalityMap, "read_only")
	if tags, err := helper.ParseFromMaps[general_objects.KeyValue](helper.CastSlice(eventCriticalityMap, "tags")); err != nil {
		return err
	} else {
		eventCriticality.Tags = tags
	}
	return nil
}
