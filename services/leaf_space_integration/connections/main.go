package connections

import (
	"github.com/leanspace/terraform-provider-leanspace/provider"
)

var LeafSpaceConnectionDataType = provider.DataSourceType[LeafSpaceConnection, *LeafSpaceConnection]{
	ResourceIdentifier: "leanspace_leaf_space_integrations",
	Path:               path,
	Schema:             leafSpaceConnectionSchema,
	FilterSchema:       leafSpaceConnectionFilterSchema,
	ReadPath: func(id string) string {
		return path
	},
	DeletePath: func(id string) string {
		return path
	},
	UpdatePath: func(id string) string {
		return path
	},
	IsUnique: true,
}

var path = "integration-leafspace/connections"
