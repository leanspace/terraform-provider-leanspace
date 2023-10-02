package routes

import "github.com/leanspace/terraform-provider-leanspace/provider"

var RouteDataType = provider.DataSourceType[Route, *Route]{
	ResourceIdentifier: "leanspace_routes",
	Path:               "routes-repository/routes",
	Schema:             routeSchema,
	FilterSchema:       dataSourceFilterSchema,
}
