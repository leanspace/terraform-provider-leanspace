package event_criticalities

import "github.com/leanspace/terraform-provider-leanspace/provider"

var EventsCriticalitiesDataType = provider.DataSourceType[EventsCriticalities, *EventsCriticalities]{
	ResourceIdentifier: "leanspace_events_criticalities",
	Path:               "events/event-criticalities",
	Schema:             EventsCriticalitiesSchema,
	FilterSchema:       dataSourceFilterSchema,
}
