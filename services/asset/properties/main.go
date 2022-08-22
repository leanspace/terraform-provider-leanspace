package properties

import (
	"fmt"
	"leanspace-terraform-provider/provider"
)

var PropertyDataType = provider.DataSourceType[Property[any], *Property[any]]{
	ResourceIdentifier: "leanspace_properties",
	Path:               "asset-repository/properties",
	Schema:             propertySchema,
	CreatePath: func(p *Property[any]) string {
		return fmt.Sprintf("asset-repository/nodes/%s/properties", p.NodeId)
	},
}
