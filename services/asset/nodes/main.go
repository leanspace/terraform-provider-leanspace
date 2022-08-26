package nodes

import "leanspace-terraform-provider/provider"

var NodeDataType = provider.DataSourceType[Node, *Node]{
	ResourceIdentifier: "leanspace_nodes",
	Path:               "asset-repository/nodes",
	Schema:             rootNodeSchema,
	FilterSchema:       dataSourceFilterSchema,
}
