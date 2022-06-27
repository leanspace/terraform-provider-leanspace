package asset

import "fmt"

type GenericResourceType[T any] struct {
	Client     *Client
	Path       string
	CreatePath func(T) string
}

func (client *Client) forNodes() GenericResourceType[Node] {
	return GenericResourceType[Node]{
		Client: client,
		Path:   "asset-repository/nodes",
	}
}

func (client *Client) forProperties() GenericResourceType[Property[any]] {
	return GenericResourceType[Property[any]]{
		Client: client,
		Path:   "asset-repository/properties",
		CreatePath: func(p Property[any]) string {
			return fmt.Sprintf("asset-repository/nodes/%s/properties", p.NodeId)
		},
	}
}

func (client *Client) forCommandDefinitions() GenericResourceType[CommandDefinition] {
	return GenericResourceType[CommandDefinition]{
		Client: client,
		Path:   "asset-repository/command-definitions",
		CreatePath: func(c CommandDefinition) string {
			return fmt.Sprintf("asset-repository/nodes/%s/command-definitions", c.NodeId)
		},
	}
}
