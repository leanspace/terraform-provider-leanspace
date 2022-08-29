package remote_agents

import "github.com/leanspace/terraform-provider-leanspace/provider"

var RemoteAgentDataType = provider.DataSourceType[RemoteAgent, *RemoteAgent]{
	ResourceIdentifier: "leanspace_remote_agents",
	Path:               "agents-repository/agents",
	Schema:             remoteAgentSchema,
	FilterSchema:       dataSourceFilterSchema,
}
