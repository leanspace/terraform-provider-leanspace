package leaf_space_connections

import (
	"fmt"
	"github.com/leanspace/terraform-provider-leanspace/provider"
)

var LeafSpaceConnectionDataType = provider.DataSourceType[LeafSpaceConnection, *LeafSpaceConnection]{
	ResourceIdentifier: "leanspace_leaf_space_integrations",
	Path:               "integration-leafspace/connections",
	Schema:             leafSpaceConnectionSchema,
	FilterSchema:       leafSpaceConnectionFilterSchema,
	ReadPath: func(id string) string {
		return fmt.Sprintf("integration-leafspace/connections")
	},
	DeletePath: func(id string) string {
		return fmt.Sprintf("integration-leafspace/connections")
	},
	UpdatePath: func(id string) string {
		return fmt.Sprintf("integration-leafspace/connections")
	},
	IsUnique: true,
}
