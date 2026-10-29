package units

import (
	"github.com/leanspace/terraform-provider-leanspace/provider"
)

var UnitDataType = provider.DataSourceType[Unit, *Unit]{
	ResourceIdentifier: "leanspace_units",
	SecurityTags: &provider.SecurityTagsDefinition{
		ResourceIdentifier: "leanspace_unit_security_tags",
		TargetIDAttribute:  "unit_id",
	},
	Path:   "asset-repository/units",
	Schema: unitSchema,
}
