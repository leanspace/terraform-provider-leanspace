package units

import (
	"leanspace-terraform-provider/provider"
)

var UnitDataType = provider.DataSourceType[Unit, *Unit]{
	ResourceIdentifier: "leanspace_units",
	Name:               "unit",
	Path:               "asset-repository/units",
	Schema:             unitSchema,
}
