package activity_definitions

import (
	"terraform-provider-asset/asset"
)

var ActivityDefinitionDataType = asset.DataSourceType[ActivityDefinition, *ActivityDefinition]{
	ResourceIdentifier: "leanspace_activity_definitions",
	Name:               "activity_definition",
	Path:               "activities-repository/activity-definitions",
	Schema:             activityDefinitionSchema,
}
