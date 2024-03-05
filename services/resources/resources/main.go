package resources

import "github.com/leanspace/terraform-provider-leanspace/provider"

var ResourceDataType = provider.DataSourceType[Resource, *Resource]{
	ResourceIdentifier: "leanspace_resources",
	Path:               "resources-repository/resources",
	Schema:             resourceSchema,
	FilterSchema:       dataSourceFilterSchema,
}
