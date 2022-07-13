package properties

import (
	"fmt"
	"terraform-provider-asset/asset"
)

var PropertyDataType = asset.DataSourceType[Property[any], *Property[any]]{
	ResourceIdentifier: "leanspace_properties",
	Name:               "property",
	Path:               "asset-repository/properties",
	Schema:             propertySchema,
	CreatePath: func(p *Property[any]) string {
		return fmt.Sprintf("asset-repository/nodes/%s/properties", p.NodeId)
	},
}
