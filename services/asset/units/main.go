package units

import (
	"github.com/leanspace/terraform-provider-leanspace/provider"
)

var UnitDataType = provider.DataSourceType[Unit, *Unit]{
	ResourceIdentifier: "leanspace_units",
	Path:               "asset-repository/units",
	Schema:             unitSchema,
	TFModelFactory:     func() any { return &UnitTF{} },
}
