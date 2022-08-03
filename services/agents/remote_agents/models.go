package remote_agents

type RemoteAgent struct {
	ID               string      `json:"id" terra:"id"`
	Name             string      `json:"name" terra:"name"`
	Description      string      `json:"description,omitempty" terra:"description,omitempty"`
	ServiceAccountId string      `json:"serviceAccountId" terra:"service_account_id"`
	Connectors       []Connector `json:"connectors" terra:"connectors"`
	CreatedAt        string      `json:"createdAt" terra:"created_at"`
	CreatedBy        string      `json:"createdBy" terra:"created_by"`
	LastModifiedAt   string      `json:"lastModifiedAt" terra:"last_modified_at"`
	LastModifiedBy   string      `json:"lastModifiedBy" terra:"last_modified_by"`
}

func (agent *RemoteAgent) GetID() string { return agent.ID }

type Connector struct {
	ID             string `json:"id" terra:"id"`
	GatewayId      string `json:"gatewayId" terra:"gateway_id"`
	Type           string `json:"type" terra:"type"`
	Socket         Socket `json:"socket" terra:"socket"`
	CreatedAt      string `json:"createdAt" terra:"created_at"`
	CreatedBy      string `json:"createdBy" terra:"created_by"`
	LastModifiedAt string `json:"lastModifiedAt" terra:"last_modified_at"`
	LastModifiedBy string `json:"lastModifiedBy" terra:"last_modified_by"`

	// inbound only
	StreamId    string     `json:"streamId" terra:"stream_id"`
	Destination ConnTarget `json:"destination" terra:"destination"`

	// outbound only
	CommandQueueId string     `json:"commandQueueId" terra:"command_queue_id"`
	Source         ConnTarget `json:"source" terra:"source"`
}

type Socket struct {
	Type string `json:"type" terra:"type"`
	Host string `json:"host,omitempty" terra:"host,omitempty"`
	Port int    `json:"port" terra:"port"`
}

type ConnTarget struct {
	Type    string `json:"type" terra:"type"`
	Binding string `json:"binding" terra:"binding"`
}
