package activity_definitions

import "github.com/leanspace/terraform-provider-leanspace/provider"

var ActivityDefinitionDataType = provider.DataSourceType[ActivityDefinition, *ActivityDefinition]{
	ResourceIdentifier: "leanspace_activity_definitions",
	Path:               "activities-repository/activity-definitions",
	Schema:             activityDefinitionSchema,
	FilterSchema:       dataSourceFilterSchema,
}
