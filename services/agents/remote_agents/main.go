package remote_agents

import "leanspace-terraform-provider/provider"

var RemoteAgentDataType = provider.DataSourceType[RemoteAgent, *RemoteAgent]{
	ResourceIdentifier: "leanspace_remote_agents",
	Name:               "remote_agent",
	Path:               "agents-repository/agents",
	Schema:             remoteAgentSchema,
}
