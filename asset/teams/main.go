package teams

import (
	"terraform-provider-asset/asset"
)

var TeamDataType = asset.DataSourceType[Team, *Team]{
	ResourceIdentifier: "leanspace_teams",
	Name:               "team",
	Path:               "teams-repository/teams",
	Schema:             teamSchema,
}
