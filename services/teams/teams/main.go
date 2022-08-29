package teams

import (
	"github.com/leanspace/terraform-provider-leanspace/provider"
)

var TeamDataType = provider.DataSourceType[Team, *Team]{
	ResourceIdentifier: "leanspace_teams",
	Path:               "teams-repository/teams",
	Schema:             teamSchema,
	FilterSchema:       dataSourceFilterSchema,
}
