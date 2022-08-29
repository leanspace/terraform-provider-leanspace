package members

import "leanspace-terraform-provider/provider"

var MemberDataType = provider.DataSourceType[Member, *Member]{
	ResourceIdentifier: "leanspace_members",
	Path:               "teams-repository/members",
	Schema:             memberSchema,
	FilterSchema:       dataSourceFilterSchema,
}
