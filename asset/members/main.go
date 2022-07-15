package members

import (
	"terraform-provider-asset/asset"
)

var MemberDataType = asset.DataSourceType[Member, *Member]{
	ResourceIdentifier: "leanspace_members",
	Name:               "member",
	Path:               "teams-repository/members",
	Schema:             memberSchema,
}
