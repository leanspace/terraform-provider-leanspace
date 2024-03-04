package resource_functions

import "github.com/leanspace/terraform-provider-leanspace/provider"

var ResourceFunctionDataType = provider.DataSourceType[ResourceFunction, *ResourceFunction]{
	ResourceIdentifier: "leanspace_resource_functions",
	Path:               "activities-repository/activity-definitions/resource-functions",
	Schema:             resourceFunctionSchema,
	FilterSchema:       dataSourceFilterSchema,
}
