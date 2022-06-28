package nodes

import "terraform-provider-asset/asset"

var NodeDataType = asset.DataSourceType[Node]{
	ResourceIdentifier: "leanspace_nodes",
	Name:               "node",
	Path:               "asset-repository/nodes",

	Schema:     nodeSchema,
	RootSchema: rootNodeSchema,

	GetID:       func(n *Node) string { return n.ID },
	MapToStruct: nodeInterfaceToStruct,
	StructToMap: nodeStructToInterfaceBase,
}
