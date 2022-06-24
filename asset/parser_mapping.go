package asset

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
	}
}

func (client *Client) forCommandDefinitions() GenericResourceType[CommandDefinition] {
	return GenericResourceType[CommandDefinition]{
		Client: client,
		Path:   "asset-repository/command-definitions",
	}
}
