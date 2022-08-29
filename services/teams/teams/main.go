package teams

import (
	"leanspace-terraform-provider/provider"
)

var TeamDataType = provider.DataSourceType[Team, *Team]{
	ResourceIdentifier: "leanspace_teams",
	Path:               "teams-repository/teams",
	Schema:             teamSchema,
	FilterSchema:       dataSourceFilterSchema,
}
