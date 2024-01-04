package orbits

import "github.com/leanspace/terraform-provider-leanspace/provider"

var OrbitDataType = provider.DataSourceType[Orbit, *Orbit]{
	ResourceIdentifier: "leanspace_orbits",
	Path:               "orbits-repository/orbits",
	Schema:             orbitSchema,
	FilterSchema:       dataSourceFilterSchema,
}
