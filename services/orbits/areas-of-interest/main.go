package areas_of_interest

import "github.com/leanspace/terraform-provider-leanspace/provider"

var AreaOfInterestDataType = provider.DataSourceType[AreaOfInterest, *AreaOfInterest]{
	ResourceIdentifier: "leanspace_areas_of_interest",
	Path:               "orbits-repository/area-of-interests",
	Schema:             areaOfInterestSchema,
	FilterSchema:       dataSourceFilterSchema,
}
