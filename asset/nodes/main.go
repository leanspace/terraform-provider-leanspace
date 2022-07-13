package nodes

import "terraform-provider-asset/asset"

var NodeDataType = asset.DataSourceType[Node, *Node]{
	ResourceIdentifier: "leanspace_nodes",
	Name:               "node",
	Path:               "asset-repository/nodes",
	Schema:             rootNodeSchema,
}
