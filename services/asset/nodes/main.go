package nodes

import "github.com/leanspace/terraform-provider-leanspace/provider"

var NodeDataType = provider.DataSourceType[Node, *Node]{
	ResourceIdentifier: "leanspace_nodes",
	SecurityTags: &provider.SecurityTagsDefinition{
		ResourceIdentifier: "leanspace_node_security_tags",
		TargetIDAttribute:  "node_id",
	},
	Path:         "asset-repository/nodes",
	Schema:       rootNodeSchema,
	FilterSchema: dataSourceFilterSchema,
}
