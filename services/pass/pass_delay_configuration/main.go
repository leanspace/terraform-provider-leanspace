package pass_delay_configuration

import (
	"github.com/leanspace/terraform-provider-leanspace/provider"
)

var PassDelayConfigurationType = provider.DataSourceType[PassDelayConfiguration, *PassDelayConfiguration]{
	ResourceIdentifier: "leanspace_pass_delay_configuration",
	Path:               path,
	Schema:             passDelayConfigurationSchema,
	FilterSchema:       nil,
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

var path = "passes-repository/passes/delay/configurations"
