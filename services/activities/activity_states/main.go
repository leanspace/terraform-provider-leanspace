package activity_states

import "github.com/leanspace/terraform-provider-leanspace/provider"

var ActivityStateDataType = provider.DataSourceType[ActivityState, *ActivityState]{
	ResourceIdentifier: "leanspace_activity_states",
	Path:               "activities-repository/activities/states",
	Schema:             activityStateSchema,
	FilterSchema:       nil,
}
