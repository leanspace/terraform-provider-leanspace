package orbit_resources

import "github.com/leanspace/terraform-provider-leanspace/provider"

var OrbitResourceDataType = provider.DataSourceType[OrbitResource, *OrbitResource]{
	ResourceIdentifier: "leanspace_orbit_resources",
	Path:               "orbits-repository/orbit-resources",
	Schema:             orbitResourceSchema,
	FilterSchema:       dataSourceFilterSchema,
}
