package nodes

import "github.com/leanspace/terraform-provider-leanspace/provider"

var NodeDataType = provider.DataSourceType[Node, *Node]{
	ResourceIdentifier: "leanspace_nodes",
	Path:               "asset-repository/nodes",
	Schema:             rootNodeSchema,
	FilterSchema:       dataSourceFilterSchema,
}
