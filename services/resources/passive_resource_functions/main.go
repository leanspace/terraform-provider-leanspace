package passive_resource_functions

import "github.com/leanspace/terraform-provider-leanspace/provider"

var PassiveResourceFunctionDataType = provider.DataSourceType[PassiveResourceFunction, *PassiveResourceFunction]{
	ResourceIdentifier: "leanspace_passive_resource_functions",
	Path:               "resources-repository/passive-resource-functions",
	Schema:             passiveResourceFunctionSchema,
	FilterSchema:       dataSourceFilterSchema,
}
