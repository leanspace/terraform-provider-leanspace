package remote_agents

type RemoteAgent struct {
	ID               string      `json:"id"`
	Name             string      `json:"name"`
	Description      string      `json:"description,omitempty"`
	ServiceAccountId string      `json:"serviceAccountId"`
	Connectors       []Connector `json:"connectors"`
	CreatedAt        string      `json:"createdAt"`
	CreatedBy        string      `json:"createdBy"`
	LastModifiedAt   string      `json:"lastModifiedAt"`
	LastModifiedBy   string      `json:"lastModifiedBy"`
}

func (agent *RemoteAgent) GetID() string { return agent.ID }

type Connector struct {
	ID             string `json:"id"`
	GatewayId      string `json:"gatewayId"`
	Type           string `json:"type"`
	Socket         Socket `json:"socket"`
	CreatedAt      string `json:"createdAt"`
	CreatedBy      string `json:"createdBy"`
	LastModifiedAt string `json:"lastModifiedAt"`
	LastModifiedBy string `json:"lastModifiedBy"`

	// inbound only
	StreamId    string     `json:"streamId"`
	Destination ConnTarget `json:"destination"`

	// outbound only
	CommandQueueId string     `json:"commandQueueId"`
	Source         ConnTarget `json:"source"`
}

type Socket struct {
	Type string `json:"type"`
	Host string `json:"host,omitempty"`
	Port int    `json:"port"`
}

type ConnTarget struct {
	Type    string `json:"type"`
	Binding string `json:"binding"`
}
