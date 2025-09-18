package event_criticalities

import "github.com/leanspace/terraform-provider-leanspace/provider"

var EventCriticalitiesDataType = provider.DataSourceType[EventCriticalities, *EventCriticalities]{
	ResourceIdentifier: "leanspace_event_criticalities",
	Path:               "events/event-criticalities",
	Schema:             EventCriticalitiesSchema,
	FilterSchema:       dataSourceFilterSchema,
}
