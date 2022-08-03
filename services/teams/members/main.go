package members

import "leanspace-terraform-provider/provider"

var MemberDataType = provider.DataSourceType[Member, *Member]{
	ResourceIdentifier: "leanspace_members",
	Name:               "member",
	Path:               "teams-repository/members",
	Schema:             memberSchema,
}
