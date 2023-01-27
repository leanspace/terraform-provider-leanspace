package properties

import (
	"fmt"

	"github.com/leanspace/terraform-provider-leanspace/provider"
)

var PropertyDataType = provider.DataSourceType[Property[any], *Property[any]]{
	ResourceIdentifier: "leanspace_properties",
	Path:               "asset-repository/properties/v2",
	Schema:             PropertySchema,
	FilterSchema:       dataSourceFilterSchema,
	CreatePath: func(p *Property[any]) string {
		return fmt.Sprintf("asset-repository/nodes/%s/properties/v2", p.NodeId)
	},
}
