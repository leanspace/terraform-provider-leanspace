package events_definitions

import "github.com/leanspace/terraform-provider-leanspace/provider"

var EventsDefinitionDataType = provider.DataSourceType[EventsDefinition, *EventsDefinition]{
	ResourceIdentifier: "leanspace_events_definitions",
	Path:               "events/event-definitions",
	Schema:             eventsDefinitions,
	FilterSchema:       dataSourceFilterSchema,
}
