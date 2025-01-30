package pass_delay_configuration

import (
	"fmt"
	"github.com/leanspace/terraform-provider-leanspace/provider"
)

var PassDelayConfigurationType = provider.DataSourceType[PassDelayConfiguration, *PassDelayConfiguration]{
	ResourceIdentifier: "leanspace_pass_delay_configuration",
	Path:               "passes-repository/passes/delay/configurations",
	Schema:             passDelayConfigurationSchema,
	FilterSchema:       nil,
	ReadPath: func(id string) string {
		return fmt.Sprintf("passes-repository/passes/delay/configurations")
	},
	DeletePath: func(id string) string {
		return fmt.Sprintf("passes-repository/passes/delay/configurations")
	},
	UpdatePath: func(id string) string {
		return fmt.Sprintf("passes-repository/passes/delay/configurations")
	},
	IsUnique: true,
}
