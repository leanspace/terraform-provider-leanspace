package remote_agents

import (
	"terraform-provider-asset/asset"
)

var RemoteAgentDataType = asset.DataSourceType[RemoteAgent, *RemoteAgent]{
	ResourceIdentifier: "leanspace_remote_agents",
	Name:               "remote_agent",
	Path:               "agents-repository/agents",
	Schema:             remoteAgentSchema,
}
