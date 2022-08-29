package members

import "github.com/leanspace/terraform-provider-leanspace/provider"

var MemberDataType = provider.DataSourceType[Member, *Member]{
	ResourceIdentifier: "leanspace_members",
	Path:               "teams-repository/members",
	Schema:             memberSchema,
	FilterSchema:       dataSourceFilterSchema,
}
