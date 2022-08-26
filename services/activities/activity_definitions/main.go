package activity_definitions

import "leanspace-terraform-provider/provider"

var ActivityDefinitionDataType = provider.DataSourceType[ActivityDefinition, *ActivityDefinition]{
	ResourceIdentifier: "leanspace_activity_definitions",
	Path:               "activities-repository/activity-definitions",
	Schema:             activityDefinitionSchema,
	FilterSchema:       dataSourceFilterSchema,
}
