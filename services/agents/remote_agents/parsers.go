package remote_agents

import (
	"github.com/leanspace/terraform-provider-leanspace/helper"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func (agent *RemoteAgent) ToMap() map[string]any {
	agentMap := make(map[string]any)
	agentMap["id"] = agent.ID
	agentMap["name"] = agent.Name
	agentMap["description"] = agent.Description
	agentMap["service_account_id"] = agent.ServiceAccountId
	agentMap["connectors"] = helper.ParseToMaps(agent.Connectors)
	agentMap["created_at"] = agent.CreatedAt
	agentMap["created_by"] = agent.CreatedBy
	agentMap["last_modified_at"] = agent.LastModifiedAt
	agentMap["last_modified_by"] = agent.LastModifiedBy
	return agentMap
}

func (connector *Connector) ToMap() map[string]any {
	connectorMap := make(map[string]any)
	connectorMap["id"] = connector.ID
	connectorMap["gateway_id"] = connector.GatewayId
	connectorMap["type"] = connector.Type
	connectorMap["socket"] = []any{connector.Socket.ToMap()}
	if connector.Type == "INBOUND" {
		connectorMap["stream_id"] = connector.StreamId
		connectorMap["destination"] = []any{connector.Destination.ToMap()}
	} else if connector.Type == "OUTBOUND" {
		connectorMap["command_queue_id"] = connector.CommandQueueId
		connectorMap["source"] = []any{connector.Source.ToMap()}
	}
	connectorMap["created_at"] = connector.CreatedAt
	connectorMap["created_by"] = connector.CreatedBy
	connectorMap["last_modified_at"] = connector.LastModifiedAt
	connectorMap["last_modified_by"] = connector.LastModifiedBy
	return connectorMap
}

func (socket *Socket) ToMap() map[string]any {
	socketMap := make(map[string]any)
	socketMap["type"] = socket.Type
	socketMap["host"] = socket.Host
	socketMap["port"] = socket.Port
	return socketMap
}

func (connTarget *ConnTarget) ToMap() map[string]any {
	connTargetMap := make(map[string]any)
	connTargetMap["type"] = connTarget.Type
	connTargetMap["binding"] = connTarget.Binding
	return connTargetMap
}

func (agent *RemoteAgent) FromMap(agentMap map[string]any) error {
	agent.ID = agentMap["id"].(string)
	agent.Name = agentMap["name"].(string)
	agent.Description = agentMap["description"].(string)
	agent.ServiceAccountId = agentMap["service_account_id"].(string)
	if connectors, err := helper.ParseFromMaps[Connector](agentMap["connectors"].(*schema.Set).List()); err != nil {
		return err
	} else {
		agent.Connectors = connectors
	}
	agent.CreatedAt = agentMap["created_at"].(string)
	agent.CreatedBy = agentMap["created_by"].(string)
	agent.LastModifiedAt = agentMap["last_modified_at"].(string)
	agent.LastModifiedBy = agentMap["last_modified_by"].(string)

	return nil
}

func (connector *Connector) FromMap(connectorMap map[string]any) error {
	connector.ID = connectorMap["id"].(string)
	connector.GatewayId = connectorMap["gateway_id"].(string)
	connector.Type = connectorMap["type"].(string)
	if len(connectorMap["socket"].([]any)) > 0 {
		if err := connector.Socket.FromMap(connectorMap["socket"].([]any)[0].(map[string]any)); err != nil {
			return err
		}
	}
	if connector.Type == "INBOUND" {
		connector.StreamId = connectorMap["stream_id"].(string)
		if len(connectorMap["destination"].([]any)) > 0 {
			if err := connector.Destination.FromMap(connectorMap["destination"].([]any)[0].(map[string]any)); err != nil {
				return err
			}
		}
	} else if connector.Type == "OUTBOUND" {
		connector.CommandQueueId = connectorMap["command_queue_id"].(string)
		if len(connectorMap["source"].([]any)) > 0 {
			if err := connector.Source.FromMap(connectorMap["source"].([]any)[0].(map[string]any)); err != nil {
				return err
			}
		}
	}
	connector.CreatedAt = connectorMap["created_at"].(string)
	connector.CreatedBy = connectorMap["created_by"].(string)
	connector.LastModifiedAt = connectorMap["last_modified_at"].(string)
	connector.LastModifiedBy = connectorMap["last_modified_by"].(string)
	return nil
}

func (socket *Socket) FromMap(socketMap map[string]any) error {
	socket.Type = socketMap["type"].(string)
	socket.Host = socketMap["host"].(string)
	socket.Port = socketMap["port"].(int)
	return nil
}

func (connTarget *ConnTarget) FromMap(connTargetMap map[string]any) error {
	connTarget.Type = connTargetMap["type"].(string)
	connTarget.Binding = connTargetMap["binding"].(string)
	return nil
}
