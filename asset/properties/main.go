package properties

import (
	"fmt"
	"terraform-provider-asset/asset"
)

var PropertyDataType = asset.DataSourceType[Property[any]]{
	ResourceIdentifier: "leanspace_properties",
	Name:               "property",
	Path:               "asset-repository/properties",
	CreatePath: func(p Property[any]) string {
		return fmt.Sprintf("asset-repository/nodes/%s/properties", p.NodeId)
	},

	Schema: propertySchema,

	GetID:       func(p *Property[any]) string { return p.ID },
	MapToStruct: getPropertyData,
	StructToMap: propertyStructToInterface,
}
